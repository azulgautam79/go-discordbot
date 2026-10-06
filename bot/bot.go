package bot

import (
	"fmt"
	"strings"

	"github.com/azulgautam79/go-discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var BotPrefix string
var goBot *discordgo.Session

func Start(cfg config.Config) {
	var err error

	BotPrefix = cfg.BotPrefix

	goBot, err = discordgo.New("Bot " + cfg.Token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// Read messages from channels
	goBot.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println("Error getting bot user:", err)
		return
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)
	goBot.AddHandler(interactionHandler)
	goBot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		registerSlashCommands(s)
	})

	err = goBot.Open()
	if err != nil {
		fmt.Println("Error opening Discord connection:", err)
		return
	}

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself and other bots.
	if m.Author.ID == BotID || m.Author.Bot {
		return
	}

	// Only respond to messages that start with the prefix.
	if !strings.HasPrefix(m.Content, BotPrefix) {
		return
	}

	args := strings.Fields(m.Content[len(BotPrefix):])
	if len(args) == 0 {
		return
	}

	command := strings.ToLower(args[0])

	switch command {
	case "ping":
		pingCommand(s, m)
	case "help":
		helpCommand(s, m)
	case "info", "about":
		infoCommand(s, m)
	case "echo":
		echoCommand(s, m, args[1:])
	case "roll":
		rollCommand(s, m, args[1:])
	case "userinfo":
		userInfoCommand(s, m)
	case "serverinfo":
		serverInfoCommand(s, m)
	default:
		s.ChannelMessageSend(m.ChannelID, "Unknown command. Try `"+BotPrefix+"help`.")
	}
}
