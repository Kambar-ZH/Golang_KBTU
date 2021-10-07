package helper

import (
	"hw4/models"
	"math"
)

func EvaluateNewRating(participant models.Participant) int {
	change := math.Sqrt(math.Abs(float64(participant.CurrentScore) - float64(participant.Penalty) - float64(participant.User.Rating)))
	if participant.CurrentScore-participant.Penalty > participant.User.Rating {
		return participant.User.Rating + int(change)
	}
	return participant.User.Rating - int(change)
}