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
		log.Printf("Cannot open the session: %v", openErr)
	}

	// When everything is over, make sure we close the session
	defer session.Close()

	testGuild := os.Getenv("GUILD_ID")

	_, err := session.ApplicationCommandCreate(session.State.User.ID, testGuild, &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Ping, Pong",
	})
	if err != nil {
		log.Fatalf("Failed to create 'ping' command: %v", err)
	}
	_, err = session.ApplicationCommandCreate(session.State.User.ID, testGuild, &discordgo.ApplicationCommand{
		Name:        "roll",
		Description: "Roll some dice! (e.g. /roll 3d4 + d12)",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:         "dice",
				Type:         discordgo.ApplicationCommandOptionString,
				Description:  "Dice to roll (e.g. \"d4\" or \"3d12 + d20\")",
				Required:     true,
				Autocomplete: false,
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to create 'roll' command: %v", err)
	}

	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			switch i.ApplicationCommandData().Name {
			case "ping":
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "pong",
					},
				})
				if err != nil {
					log.Printf("Failed to respond to command: %v", err)
				}
			case "roll":
				RollInteraction(s, i)
			}
		}
	})

	testChannel := "559936001305214999"

	mes, messageErr := session.ChannelMessageSend(testChannel, "What's up gamers?")

	if messageErr != nil {
		log.Printf("Cannot send message: %v", messageErr)
	} else {
		log.Printf("Sent message: %v", mes)

		session.MessageReactionAdd(mes.ChannelID, mes.ID, "1")
	}

	// Create a channel for waiting on Interrupt
	stop := make(chan os.Signal)
	// Use stop channel for os.Interrupt
	signal.Notify(stop, os.Interrupt)
	// Wait on a signal from stop channel
	<-stop
	// Cleanup tasks now that we recieved os.Interrupt
	log.Println("Gracefully shutdowning")

	session.ChannelMessageSend(testChannel, "I'm going to bed!")
}
