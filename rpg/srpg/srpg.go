package srpg

import "strconv"

type Carte struct {
	Valeur  string
	Couleur string
	Famille string
	URL     string
}

// terrain
type Damier struct {
	Case [10][10]Case
}

// chacun des pions de chacuns des joueurs
type Case struct {
	Joueur  bool
	Vide    bool
	Is_dame bool
}

// je n'ai besoin que de ca pour lancer la partie
type Partie struct {
	Damier Damier
}

func (carte Carte) Affichage() string {
	var ret string

	ret += "la valeur de ma carte est: " + carte.Valeur + "\nCa couleur est: " + carte.Couleur + "\nCa famille est: " + carte.Famille

	return ret
}

func (damier Damier) Affichage_Damier() string {
	var ret string
	ret += "   0 1 2 3 4 5 6 7 8 9\n"

	for i, v := range damier.Case {
		ret += strconv.Itoa(i) + " "
		for _, w := range v {

			if w.Vide {
				ret += " _"
			} else {
				if w.Joueur {
					ret += " O"
				} else {
					ret += " X"
				}
			}
		}
		ret += "\n"
	}

	return ret
}
