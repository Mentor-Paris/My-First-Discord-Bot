package bot

import (
	"fmt"
	"golang-discord-bot/config"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Run() {

	// Create bot session
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		panic(err)
	}

	// Make the bot a user
	user, err := goBot.User("@me")
	if err != nil {
		fmt.Println("Your Token is wrong")
		panic(err)
	}
	//Allows all events
	goBot.Identify.Intents = discordgo.IntentsAll

	//Get the bot user ID
	BotID = user.ID

	//Create an event Handler, when condition of messageHandler trigger launch the function
	goBot.AddHandler(messageHandler)

	//Launch the bot
	err = goBot.Open()
	if err != nil {
		panic(err)
	}
	fmt.Println("Bot is running !")
}

// Trigger on message send
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}
	//Struct of message : https://pkg.go.dev/github.com/bwmarrin/discordgo@v0.25.0?utm_source=gopls#Message

	// If the message is "Beep" reply with "Boop"
	if m.Content == "Beep" {
		s.ChannelMessageSend(m.ChannelID, "Boop")
	}

	// If the message is "Boop" reply with "Beep"
	if m.Content == "Boop" {
		s.ChannelMessageSend(m.ChannelID, "Beep")
	}
}
