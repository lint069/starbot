package commands

import (
	"discord_starbot/config"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if !strings.HasPrefix(m.Content, config.BotPrefix) {
		return
	}

	args := strings.Fields(m.Content[len(config.BotPrefix):])
	if len(args) == 0 {
		return
	}

	command := strings.ToLower(args[0])
	latency := s.HeartbeatLatency()

	switch command {
	case "ping":
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Pong! `%dms`", latency.Milliseconds()))
	}
}
