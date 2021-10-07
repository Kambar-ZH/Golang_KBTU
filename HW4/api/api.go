package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"hw4/models"
	"hw4/tablebuilder"
)

var Participants []models.Participant

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", returnAllParticipants)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func returnAllParticipants(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Participants)
}

func showStandings() {
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var participantsTable []models.Participant
	json.Unmarshal(body, &participantsTable)
	tablebuilder.ShowTable(participantsTable)
}

func updatePenalties() {
	for i := range Participants {
		Participants[i].Penalty += 20
	}
}

func Run() {
	Participants = []models.Participant{
		{User: models.User{Handle: "jeeraffo", Rating: 1925, ProfileLink: "http://localhost:8080/user/jeeraffo"}, CurrentScore: 2500},
		{User: models.User{Handle: "khiro", Rating: 2050, ProfileLink: "http://localhost:8080/user/khiro"}, CurrentScore: 4500},
	}
	go handleRequests()
	fmt.Println("Welcome to the ICPC competition!")
	show := time.NewTicker(5 * time.Second)
	tick := time.NewTicker(1 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for {
			select {
			case <-show.C:
				showStandings()
			case <- tick.C:
				updatePenalties()
			case <-ctx.Done():
				return
			}
		}
	}()
	fmt.Scanln()
}