package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var session *discordgo.Session

func init() {
	// Load .env file
	godotenv.Load(".env")

	botSecret := os.Getenv("BOT_SECRET")
	var err error
	session, err = discordgo.New("Bot " + botSecret) // TODO Store this secret differently
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func main() {
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		// Notify us when the Bot is online
		log.Println("Bot is up!")
	})
	// Open our session
	openErr := session.Open()
	// If we fail to open the session, log and exit
	if openErr != nil {
		log.Fatalf("Cannot open the session: %v", openErr)
	}

	// When everything is over, make sure we close the session
	defer session.Close()

	// _, err := s.ApplicationCommandCreate(s.State.User.ID, "559936001305214997", "ping") // TODO Store Guild differently
	// if err != nil {
	// 	log.Panicf("Cannot create 'ping' command: %v", err)
	// }

	session.ChannelMessageSend("559936001305214999", "We got beef :cut_of_meat:")

	// Create a channel for waiting on Interrupt
	stop := make(chan os.Signal)
	// Use stop channel for os.Interrupt
	signal.Notify(stop, os.Interrupt)
	// Wait on a signal from stop channel
	<-stop
	// Cleanup tasks now that we recieved os.Interrupt
	log.Println("Gracefully shutdowning")

	session.ChannelMessageSend("559936001305214999", "Goodbye cruel world!")
}
