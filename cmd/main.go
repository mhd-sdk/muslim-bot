package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/mhd-sdk/muslim-bot/config"
)

func main() {
	config.LoadConfig()
	dg, err := discordgo.New("Bot " + config.EnvConfigs.DiscordToken)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// when the bot receive /ping command
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.ApplicationCommandData().Name == "ping" {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong!",
				},
			})
		}
	})

	// register /ping command
	var commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Type:        discordgo.ChatApplicationCommand,
			Description: "Replies with Pong!",
		},
	}

	// register the commands
	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("MuslimBot is now running.  Press CTRL + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}
