package discord

import (
	"Prog_Lud/config"
	"Prog_Lud/rpg"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Run() {
	// create bot session
	goBot, err := discordgo.New("Bot " + config.Config.Token)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	// make the bot a user
	user, err := goBot.User("@me")
	if err != nil {
		fmt.Printf("user: %v\n", user)
		log.Fatal(err.Error())
		return
	}
	BotID = user.ID
	goBot.AddHandler(messageHandler)
	err = goBot.Open()

	if err != nil {
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if BotID == m.Author.ID || m.ChannelID != "1032264204289323068" {
		return
	}

	var ret discordgo.MessageEmbed
	var image discordgo.MessageEmbedImage
	Nouvelle_carte := rpg.C_Carte("Valet", "Rouge", "Carreaux", "https://upload.wikimedia.org/wikipedia/commons/thumb/2/27/Jack_of_diamonds_fr.svg/640px-Jack_of_diamonds_fr.svg.png")
	image.URL = Nouvelle_carte.URL
	image.ProxyURL = "./img/valet_carreau.png"
	image.Height = 100
	image.Width = 100
	ret.Image = &image
	ret.Title = Nouvelle_carte.Affichage()
	adresse_m := &ret
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)

}

func StartWith(content string, patern string) bool {
	if len(patern) > len(content) {
		return false
	}
	for i, v := range patern {
		if rune(content[i]) != v {
			return false
		}
	}
	return true
}
