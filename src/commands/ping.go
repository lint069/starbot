package commands

import (
	"discord_starbot/permissions"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {

	userPermissions, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil {
		return
	}

	// Example of using permissions limitation
	hasPermissions := userPermissions&permissions.Administrator != 0
	if !hasPermissions {
		s.ChannelMessageSend(m.ChannelID, "You can't use this command, required permission: `Administrator`")
		return
	}

	latency := s.HeartbeatLatency()
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Pong! `%dms`", latency.Milliseconds()))
}
