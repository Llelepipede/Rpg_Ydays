package sdiscord

import (
	"Prog_Lud/rpg/srpg"
	"encoding/json"
	"io/ioutil"
)

// je crée ma structure qui stoquera mes donnée a sauvegarder
type SauvPartie struct {
	Id      int         `json:"Id"`
	Joueur1 string      `json:"Joueur1"`
	Joueur2 string      `json:"Joueur2"`
	Damier  srpg.Damier `json:"Damier"`
}

// je recupère mes donner pour les ecrirent dans le json
func SauvegardePartie(Sauvegarde []SauvPartie) error {
	// je transforme mes donnée qui sont dans sauvegarde pour les
	// rendre compatible avec du json
	file, err := json.Marshal(Sauvegarde)
	// je vérifie si il y a une erreur ou pas
	if err != nil {
		return err
	}
	// j'ecris dans le fichier
	ioutil.WriteFile("sauvegarde.json", file, 0777)
	return err
}

// je recupere les données du json
func ReadSauvegarde() ([]SauvPartie, error) {
	// je prepare ma variable qui stoquera mes données du json
	var ret []SauvPartie
	// je lis le json, et stoque les donnée dans file
	file, err := ioutil.ReadFile("sauvegarde.json")

	// je vérifie si il y a une erreur
	if err != nil {
		return ret, err
	}
	// je transforme mes donnée json en donnée structure que je stoque
	// dans ret
	err = json.Unmarshal(file, &ret)

	// je vérifie si il y a une erreur
	if err != nil {
		return ret, err
	}
	return ret, nil
}
