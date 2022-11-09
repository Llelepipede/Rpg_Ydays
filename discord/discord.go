package discord

import (
	"Prog_Lud/config"
	"Prog_Lud/rpg"
	"Prog_Lud/rpg/srpg"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session
var damier srpg.Damier

func Run() {
	damier = rpg.C_Damier()
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
	// var ret discordgo.MessageEmbed
	// var image discordgo.MessageEmbedImage
	// Nouvelle_carte := rpg.C_Carte("Valet", "Rouge", "Carreaux", "https://upload.wikimedia.org/wikipedia/commons/thumb/2/27/Jack_of_diamonds_fr.svg/640px-Jack_of_diamonds_fr.svg.png")
	// image.URL = Nouvelle_carte.URL
	// image.ProxyURL = "./img/valet_carreau.png"
	// image.Height = 100
	// image.Width = 100
	// ret.Image = &image
	// ret.Title = Nouvelle_carte.Affichage()
	// adresse_m := &ret
	// _, _ = s.ChannelMessageSendEmbed(m.ChannelID, adresse_m)
	fmt.Printf("m.Content: %v\n", m.Content)
	if StartWith(m.Content, "damier") {
		new_text := config.Split(m.Content)

		s.ChannelMessageSend(m.ChannelID, "```"+damier.Affichage_Damier()+"```")
		if len(new_text) == 3 {
			list, _ := s.GuildMembersSearch(m.GuildID, new_text[1], 10)
			for _, v := range list {
				if v.User.Username == new_text[1] {
					s.ChannelMessageSend(m.ChannelID, v.Mention())
				}
			}
			list, _ = s.GuildMembersSearch(m.GuildID, new_text[2], 10)
			for _, v := range list {
				if v.User.Username == new_text[2] {
					s.ChannelMessageSend(m.ChannelID, v.Mention())
				}
			}

		}

	}
	if StartWith(m.Content, "move") {

	}

	// s.ChannelMessageSend(m.ChannelID, m.Author.Mention())

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
