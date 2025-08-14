package handlers

import (
	"discord_starbot/commands"

	"github.com/bwmarrin/discordgo"
)

func SlashCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "setup":
		commands.HandleSetupCommand(s, i)
	}
}
