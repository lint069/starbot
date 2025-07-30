package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	token := os.Getenv("BOT_TOKEN")
	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal("Error creating Discord session", err)
	}

	_ = discord

	err = discord.Open()
	if err != nil {
		log.Fatal("Error opening connection", err)
	}

	log.Println("bot online")

	select {}
}
