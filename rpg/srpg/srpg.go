package srpg

type Carte struct {
	Valeur  string
	Couleur string
	Famille string
	URL     string
}

// terrain
type Damier struct {
}

// chacun des pions de chacuns des joueurs
type Pion struct {
}

// je n'ai besoin que de ca pour lancer la partie
type Partie struct {
}

func (carte Carte) Affichage() string {
	var ret string

	ret += "la valeur de ma carte est: " + carte.Valeur + "\nCa couleur est: " + carte.Couleur + "\nCa famille est: " + carte.Famille

	return ret
}
