package handlers

import (
	"discord_starbot/commands"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := os.Getenv("PREFIX")

	if m.Author.Bot {
		return
	}

	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	args := strings.Fields(m.Content[len(prefix):])
	if len(args) == 0 {
		return
	}

	command := strings.ToLower(args[0])

	if cmdFunc, ok := commands.CommandMap[command]; ok {
		cmdFunc(s, m)
	}
}
