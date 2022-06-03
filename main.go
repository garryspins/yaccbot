package main

import (
	"log"
	"os"
	"os/signal"
	
	// "github.com/Moonlington/harmonia"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var botState *discordgo.Session

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
		return
	}

	botState, err = discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Error initializing BotState (%v) (%v)", err, botState)
		return
	}

	handler.session = botState
	botState.AddHandler(handler.BotHandler)
	
	err = botState.Open()
	if err != nil {
		log.Fatalf("Cannot open session: %v", err)
		return
	}
	defer botState.Close()
	
	handler.Create()
	handler.InitializeCommands()
	

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	handler.CleanupCommands()

	log.Println("Gracefully shutting down.")
}