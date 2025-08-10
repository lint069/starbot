package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	latency := s.HeartbeatLatency()
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Pong! `%dms`", latency.Milliseconds()))
}
