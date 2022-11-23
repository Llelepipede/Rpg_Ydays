package discord

import (
	"Prog_Lud/config"
	"Prog_Lud/discord/sdiscord"
	"Prog_Lud/rpg"
	"Prog_Lud/rpg/srpg"
	"fmt"
	"log"
	"strconv"

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
		// separe notre message en mot pour pouvoir les compter
		new_text := config.Split(m.Content)

		// cree mon tableau total de sauvegarde
		var totalsauvegarde []sdiscord.SauvPartie
		// je le remplie
		totalsauvegarde, err := sdiscord.ReadSauvegarde()
		// je l'affiche au cas ou
		fmt.Printf("err: %v\n", err)
		fmt.Printf("totalsauvegarde: %v\n", totalsauvegarde)

		// je crée la sauvegarde de la partie actuelle
		var sauvegarde sdiscord.SauvPartie
		// je lui remplie les donner du damier actuelle, et un id de 1
		sauvegarde.Damier = damier
		sauvegarde.Id = 1

		// je verifie si il y a exactement 3 mots dans mon message
		if len(new_text) == 3 {
			// je recupere la list de tout les utilisateurs connecter
			// ou leur pseudo comment par le 2eme mot de mon message
			list, _ := s.GuildMembersSearch(m.GuildID, new_text[1], 10)
			// je parcours la list des utilisateur trouver
			for _, v := range list {
				// je vérifie si il sont parfaitement egale ou pas
				// a mon 2eme
				if v.User.Username == new_text[1] {
					// je sauvegarde l'id de l'utilisateur dans ma sauvegarde
					sauvegarde.Joueur1 = v.User.ID
					// je mentionne l'utiliateur
					s.ChannelMessageSend(m.ChannelID, v.Mention())
				}
			}
			// je refait tout pareil , mais cette fois si avec le 3eme mot
			list, _ = s.GuildMembersSearch(m.GuildID, new_text[2], 10)
			for _, v := range list {
				if v.User.Username == new_text[2] {
					sauvegarde.Joueur2 = v.User.ID
					s.ChannelMessageSend(m.ChannelID, v.Mention())
				}
			}

		}
		// j'ajoute ma nouvelle sauvegarde au tableau total de sauvegarde
		totalsauvegarde = append(totalsauvegarde, sauvegarde)
		// j'ecris dans le fichier sauvegarde.json les données de toute
		// mes sauvegarde
		sdiscord.SauvegardePartie(totalsauvegarde)

	}
	// move 1.5  ["move","1.5"]
	// move 1.5 1.7
	if StartWith(m.Content, "move") {
		new_text := config.Split(m.Content)
		// s.ChannelMessageSend(m.ChannelID, "```"+damier.Affichage_Damier()+"```")

		x, _ := strconv.Atoi(string(new_text[1][0]))
		y, _ := strconv.Atoi(string(new_text[1][2]))

		if damier.Case[y][x].Vide {
			s.ChannelMessageSend(m.ChannelID, "vide")

		} else {
			s.ChannelMessageSend(m.ChannelID, "plein")
			x2, _ := strconv.Atoi(string(new_text[2][0]))
			y2, _ := strconv.Atoi(string(new_text[2][2]))
			damier.Case[y][x].Vide = true
			damier.Case[y2][x2].Vide = false
			damier.Case[y2][x2].Joueur = damier.Case[y][x].Joueur

			s.ChannelMessageSend(m.ChannelID, "```"+damier.Affichage_Damier()+"```")
		}

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
