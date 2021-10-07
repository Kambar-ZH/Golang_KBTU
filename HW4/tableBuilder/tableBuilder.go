package tablebuilder

import (
	"fmt"
	"hw4/models"
	"hw4/helper"
	"sort"
	"time"
)

func showCurrentTime() {
	fmt.Printf("Time: %v\n", time.Now().Format("02-Jan-2006 15:04:05"))
}

func sortParticipantsByScore(participants []models.Participant) []models.Participant {
	sort.Slice(participants, func(i, j int) bool {
		return participants[i].CurrentScore < participants[j].CurrentScore
	})
	return participants
}

func ShowTable(participants []models.Participant) {
	showCurrentTime()
	sortedParticipants := sortParticipantsByScore(participants)
	for i, p := range sortedParticipants {
		fmt.Printf("| %d | User: %16s | Score: %5d | New rating prediction: %5d |\n", i+1, p.User.Handle, p.CurrentScore, helper.EvaluateNewRating(p))
	}
}