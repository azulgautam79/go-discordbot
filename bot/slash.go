package bot

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

// slashCommands defines the slash commands the bot registers.
var slashCommands = []*discordgo.ApplicationCommand{
	{Name: "ping", Description: "Check if the bot is alive"},
	{Name: "help", Description: "Show help message"},
	{Name: "info", Description: "Show bot info"},
	{
		Name:        "echo",
		Description: "Repeat your message",
		Options: []*discordgo.ApplicationCommandOption{
			{Type: discordgo.ApplicationCommandOptionString, Name: "text", Description: "Text to repeat", Required: true},
		},
	},
	{
		Name:        "roll",
		Description: "Roll a dice",
		Options: []*discordgo.ApplicationCommandOption{
			{Type: discordgo.ApplicationCommandOptionInteger, Name: "sides", Description: "Number of sides (default 6)", Required: false},
		},
	},
	{Name: "userinfo", Description: "Info about your account"},
	{Name: "serverinfo", Description: "Info about this server"},
}

func registerSlashCommands(s *discordgo.Session) {
	for _, cmd := range slashCommands {
		_, err := s.ApplicationCommandCreate(BotID, "", cmd)
		if err != nil {
			fmt.Println("Error creating slash command", cmd.Name, ":", err)
		}
	}
	fmt.Println("Slash commands registered")
}

func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()

	var content string

	switch data.Name {
	case "ping":
		content = "pong!"
	case "help":
		content = "**Available slash commands:**\n/ping\n/help\n/info\n/echo\n/roll\n/userinfo\n/serverinfo"
	case "info":
		content = "I'm a Discord bot written in Go using discordgo!"
	case "echo":
		options := data.Options
		if len(options) > 0 {
			content = options[0].StringValue()
		} else {
			content = "Nothing to echo."
		}
	case "roll":
		sides := 6
		if len(data.Options) > 0 {
			sides = int(data.Options[0].IntValue())
			if sides < 1 {
				sides = 6
			}
		}
		rand.Seed(time.Now().UnixNano())
		result := rand.Intn(sides) + 1
		content = fmt.Sprintf("You rolled a **%d** (d%d)", result, sides)
	case "userinfo":
		u := i.Member.User
		if u == nil {
			u = i.User
		}
		content = fmt.Sprintf("**User Info**\nName: %s#%s\nID: %s\nBot: %t", u.Username, u.Discriminator, u.ID, u.Bot)
	case "serverinfo":
		g, err := s.Guild(i.GuildID)
		if err != nil {
			content = "Could not get server info."
		} else {
			content = fmt.Sprintf("**Server Info**\nName: %s\nID: %s\nMembers: %d", g.Name, g.ID, g.MemberCount)
		}
	default:
		content = "Unknown command."
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: content},
	})
	if err != nil {
		fmt.Println("Error responding to interaction:", err)
	}
}

// Helper to convert DataOption to easier access if needed
func getOption(options []*discordgo.ApplicationCommandInteractionDataOption, name string) *discordgo.ApplicationCommandInteractionDataOption {
	for _, o := range options {
		if o.Name == name {
			return o
		}
	}
	return nil
}
