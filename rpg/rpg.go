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

func C_Case(Joueur bool, Vide bool, Is_dame bool) srpg.Case {
	var ret srpg.Case

	ret.Joueur = Joueur
	ret.Vide = Vide
	ret.Is_dame = Is_dame

	return ret
}

func C_Damier() srpg.Damier {
	var ret srpg.Damier

	for y, v := range ret.Case {
		for x, _ := range v {
			if (x+y)%2 == 0 {
				ret.Case[y][x].Vide = true
				ret.Case[y][x].Joueur = false
				ret.Case[y][x].Is_dame = false
			} else {
				if y < 4 {
					ret.Case[y][x].Vide = false
					ret.Case[y][x].Joueur = false
					ret.Case[y][x].Is_dame = false
				} else if y > 5 {
					ret.Case[y][x].Vide = false
					ret.Case[y][x].Joueur = true
					ret.Case[y][x].Is_dame = false
				} else {
					ret.Case[y][x].Vide = true
					ret.Case[y][x].Joueur = false
					ret.Case[y][x].Is_dame = false
				}
			}
		}
	}

	// for y := 0; y < len(ret.Case); y++ {
	// 	for x := 0; x < len(ret.Case[y]); x++ {
	// 		if (x+y)%2 == 0 {
	// 			ret.Case[y][x].Vide = x + y
	// 			ret.Case[y][x].Joueur = 0
	// 			ret.Case[y][x].Is_dame = false
	// 		} else {
	// 			ret.Case[y][x].Vide = -1
	// 			ret.Case[y][x].Joueur = -1
	// 			ret.Case[y][x].Is_dame = false
	// 		}

	// 	}

	// }

	return ret
}
