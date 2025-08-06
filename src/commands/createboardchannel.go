package commands

import (
	"discord_starbot/config"
	"discord_starbot/permissions"
	"log"

	"github.com/bwmarrin/discordgo"
)

var CreateChannelCommand = &discordgo.ApplicationCommand{
	Name:        "create-starboard-channel",
	Description: "Creates the channel that the bot will post to.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "custom-name",
			Description: "lets you choose a name for the channel, if empty channel name will be 'starboard'",
			Required:    false,
		},
	},
}

func CreateChannelCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userPermissions, err := s.UserChannelPermissions(i.Member.User.ID, i.ChannelID)
	if err != nil {
		return
	}

	hasPermissions := userPermissions&permissions.ManageChannels != 0

	if !hasPermissions {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You need the `Manage Channels` permissions to use this command.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	options := i.ApplicationCommandData().Options

	channelname := "starboard"
	if len(options) > 0 {
		channelname = options[0].StringValue()
	}

	channel, err := s.GuildChannelCreateComplex(config.GuildID, discordgo.GuildChannelCreateData{
		Name: channelname,
		Type: discordgo.ChannelTypeGuildText,
		PermissionOverwrites: []*discordgo.PermissionOverwrite{
			{
				ID:   config.GuildID,
				Type: discordgo.PermissionOverwriteTypeRole,
				Deny: permissions.SendMessages | permissions.CreatePublicThreads | permissions.CreatePrivateThreads,
			},
		},
	})

	if err != nil {
		log.Println("Couldn't create channel: ", err)
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Channel <#" + channel.ID + "> created.",
		},
	})
}
