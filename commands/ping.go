package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	start := time.Now()

	msg, err := s.ChannelMessageSend(m.ChannelID, "Pong")
	if err != nil {
		fmt.Println("failed to send message", err)
	}

	t := time.Now()
	elapsed := t.Sub(start)

	s.ChannelMessageEdit(m.ChannelID, msg.ID, fmt.Sprintf("Pong! `%dms`", elapsed))
}
