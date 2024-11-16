package suffler

import (
	"math/rand"
	data "quiz/data"
)

func Shuffle(questions []data.Question) {
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

}
