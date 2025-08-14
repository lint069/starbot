package main

import (
	"discord_starbot/commands"
	"discord_starbot/handlers"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func setStatus(s *discordgo.Session) {
	err := s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: "the stars",
				Type: discordgo.ActivityTypeWatching,
			},
		},
		Status: "idle",
		AFK:    true,
	})
	if err != nil {
		log.Println("error updating status:", err)
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	token := os.Getenv("APP_TOKEN")
	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal("Error creating Discord session", err)
	}

	commands.InitCommands(discord)
	discord.AddHandler(handlers.MessageHandler)
	discord.AddHandler(handlers.SlashCommandHandler)

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		setStatus(s)
	})

	err = discord.Open()
	if err != nil {
		log.Fatal("Error opening connection", err)
	}

	log.Println("bot online")

	select {}
}
