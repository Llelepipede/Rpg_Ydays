package rpg

import "Prog_Lud/rpg/srpg"

func C_Carte(Valeur string, Couleur string, Famille string, URL string) srpg.Carte {
	var ret srpg.Carte

	ret.Valeur = Valeur
	ret.Couleur = Couleur
	ret.Famille = Famille
	ret.URL = URL

	return ret
}
