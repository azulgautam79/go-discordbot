🤖 Go Discord Bot

A simple Discord bot written in Go using the discordgo
 library.

The bot provides basic utility and information commands such as ping, help, dice rolling, message echoing, user information, and server information.

✨ Features

🏓 Ping — Check whether the bot is online

❓ Help — Display all available commands

ℹ️ Info — Display information about the bot

📢 Echo — Repeat a message

🎲 Roll — Roll a dice with a configurable number of sides

👤 User Info — Display information about the current user

🏠 Server Info — Display information about the current server

🛠️ Built With

Go

DiscordGo

Discord Gateway API

📋 Requirements

Before running the bot, make sure you have:

Go installed

A Discord account

A Discord application/bot created in the Discord Developer Portal

Your Discord bot token

The bot invited to your Discord server

🚀 Getting Started
1. Clone the Repository
git clone https://github.com/your-username/your-repository.git
cd your-repository


Replace the repository URL with your actual GitHub repository.

2. Install Dependencies

If your project already has a go.mod file:

go mod tidy


Or install DiscordGo directly:

go get github.com/bwmarrin/discordgo

3. Configure Your Bot Token

Your bot needs a Discord token to connect to Discord.

Never commit your bot token to GitHub.

A common approach is to store it in an environment variable.

Linux / macOS
export DISCORD_TOKEN="your-bot-token"

Windows PowerShell
$env:DISCORD_TOKEN="your-bot-token"


Then read the token in Go:

token := os.Getenv("DISCORD_TOKEN")

4. Configure Discord Intents

If your bot reads message content, make sure Message Content Intent is enabled.

Go to:

Discord Developer Portal → Your Application → Bot → Privileged Gateway Intents

Enable:

Message Content Intent

Only enable other privileged intents if your bot actually needs them.

For example:

dg.Identify.Intents =
    discordgo.IntentsGuilds |
    discordgo.IntentsGuildMessages |
    discordgo.IntentsMessageContent


If the required intent is not enabled in the Developer Portal, Discord may close the Gateway connection with:

websocket: close 4014: Disallowed intent(s)

📜 Commands

Assuming your bot prefix is !:

Command	Description
!ping	Check if the bot is alive
!help	Show the list of available commands
!info	Show information about the bot
!echo <text>	Repeat the provided text
!roll [sides]	Roll a dice
!userinfo	Show information about your Discord account
!serverinfo	Show information about the current server
Examples
Ping
!ping


Response:

pong!

Roll a Six-Sided Dice
!roll

Roll a 20-Sided Dice
!roll 20

Echo a Message
!echo Hello, Discord!

📁 Project Structure

A simple project structure could look like this:

your-bot/
├── bot/
│   ├── commands.go
│   └── ...
├── main.go
├── go.mod
├── go.sum
└── README.md

▶️ Running the Bot

Run the project with:

go run .


Or build it first:

go build -o bot


Then run the compiled application:

Linux / macOS
./bot

Windows
.\bot.exe

🔐 Security

Never upload your Discord bot token to GitHub or share it publicly.

Avoid code like:

discordgo.New("Bot YOUR_TOKEN_HERE")


Instead, use an environment variable or another secure configuration method.

If your token is accidentally exposed, immediately regenerate it from the Discord Developer Portal.

🐛 Troubleshooting
websocket: close 4014: Disallowed intent(s)

This means your bot requested a privileged Gateway intent that Discord has not allowed.

Check:

Discord Developer Portal → Application → Bot → Privileged Gateway Intents

Make sure the intents requested by your Go code are enabled.

For this bot, if you process message content, Message Content Intent is typically required.

exit status 0xc000013a

On Windows, this usually means the application was interrupted, commonly by pressing Ctrl+C.

It is generally not a Discord connection error.

📄 License

This project is open source. Add your preferred license here, for example:

MIT License


Made with ❤️ and Go.# go-discordbot
