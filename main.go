package main

import (
	"golang-discord-bot/bot"
	"golang-discord-bot/config"
	"log"
)

// lien d'invitation du bot dans le serveur
// https://discord.com/oauth2/authorize?client_id=993483152762876014&scope=bot

func main() {

	//Read the config file
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	//Launch Run in bot.go
	bot.Run()
	<-make(chan struct{}) // Keep the bot running
	return
}
