package commands

import (
	"discord_starbot/permissions"
	"discord_starbot/utils"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var SetupCommand = &discordgo.ApplicationCommand{
	Name:        "setup",
	Description: "sets up the starboard",
}

func HandleSetupCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channel, err := s.GuildChannelCreateComplex(os.Getenv("GUILD_ID"), discordgo.GuildChannelCreateData{
		Name: "starboard",
		Type: discordgo.ChannelTypeGuildNews,
		PermissionOverwrites: []*discordgo.PermissionOverwrite{
			{
				ID:   os.Getenv("GUILD_ID"),
				Type: discordgo.PermissionOverwriteTypeRole,
				Deny: permissions.SendMessages | permissions.CreatePublicThreads,
			},
		},
	})

	if err != nil {
		log.Println("commands/setup.go:", err)
	}

	err = utils.Setenv(".env", "POST_TARGET", channel.ID)
	if err != nil {
		log.Println("commands/setup.go:", err)
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: ":+1:",
		},
	})

	if err != nil {
		log.Println("commands/setup.go:", err)
	}
}
