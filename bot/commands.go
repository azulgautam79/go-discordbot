package bot

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func pingCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "pong!")
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

func helpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	help := strings.Join([]string{
		"**Available commands:**",
		BotPrefix + "ping - Check if the bot is alive",
		BotPrefix + "help - Show this help message",
		BotPrefix + "info - Show bot info",
		BotPrefix + "echo <text> - Repeat your message",
		BotPrefix + "roll [sides] - Roll a dice (default 6 sides)",
		BotPrefix + "userinfo - Info about your account",
		BotPrefix + "serverinfo - Info about this server",
	}, "\n")

	_, err := s.ChannelMessageSend(m.ChannelID, help)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

func infoCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "I'm a Discord bot written in Go using discordgo!")
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

func echoCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 0 {
		s.ChannelMessageSend(m.ChannelID, "Usage: "+BotPrefix+"echo <text>")
		return
	}
	_, err := s.ChannelMessageSend(m.ChannelID, strings.Join(args, " "))
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

func rollCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	sides := 6
	if len(args) > 0 {
		n, err := strconv.Atoi(args[0])
		if err != nil || n < 1 {
			s.ChannelMessageSend(m.ChannelID, "Please provide a valid number of sides.")
			return
		}
		sides = n
	}

	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(sides) + 1
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You rolled a **%d** (d%d)", result, sides))
}

func userInfoCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	u := m.Author
	info := fmt.Sprintf("**User Info**\nName: %s#%s\nID: %s\nBot: %t", u.Username, u.Discriminator, u.ID, u.Bot)
	s.ChannelMessageSend(m.ChannelID, info)
}

func serverInfoCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	g, err := s.Guild(m.GuildID)
	if err != nil {
		fmt.Println("Error getting guild:", err)
		return
	}
	info := fmt.Sprintf("**Server Info**\nName: %s\nID: %s\nMembers: %d", g.Name, g.ID, g.MemberCount)
	s.ChannelMessageSend(m.ChannelID, info)
}
