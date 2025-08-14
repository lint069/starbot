package commands

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var CommandMap = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate){}

func InitCommands(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		SetupCommand,
	}

	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(os.Getenv("APP_ID"), "1369013231615737866", cmd)
		if err != nil {
			log.Printf("Cannot create slash command %q: %v", cmd.Name, err)
		}
	}

	CommandMap["ping"] = Ping
}
