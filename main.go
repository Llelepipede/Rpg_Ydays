package main

import (
	"Prog_Lud/config"
	"Prog_Lud/discord"
	"log"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		log.Fatal(err)
		return
	} else {
		discord.Run()
		<-make(chan struct{})
		return
	}

	// fmt.Printf("%v\n", rpg.C_Damier().Affichage_Damier())
}
