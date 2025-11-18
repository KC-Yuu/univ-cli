package custom

import (
	"math/rand"
	"time"
)

var quotes = []string{
	"La seule façon de faire du bon travail est d'aimer ce que vous faites. - Steve Jobs",
	"L'innovation distingue un leader d'un suiveur. - Steve Jobs",
	"Le code est comme l'humour. Quand vous devez l'expliquer, c'est mauvais. - Cory House",
	"D'abord, résolvez le problème. Ensuite, écrivez le code. - John Johnson",
	"Le meilleur code est celui qui n'a pas besoin d'être écrit. - Jeff Atwood",
	"La simplicité est la sophistication suprême. - Leonardo da Vinci",
	"Tout le monde devrait apprendre à programmer car cela vous apprend à penser. - Steve Jobs",
	"Les programmes doivent être écrits pour que les gens les lisent, et seulement accessoirement pour que les machines les exécutent. - Harold Abelson",
	"N'ayez pas peur de la perfection, vous ne l'atteindrez jamais. - Salvador Dali",
	"Le progrès n'est pas un accident, c'est une nécessité. - Herbert Spencer",
}

func GetRandomQuote() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return quotes[r.Intn(len(quotes))]
}
