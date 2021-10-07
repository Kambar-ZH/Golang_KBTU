package models

type User struct {
	Handle      string `json:"handle"`
	Rating      int    `json:"rating"`
	ProfileLink string `json:"link"`
}

type Participant struct {
	User         User `json:"user"`
	CurrentScore int  `json:"score"`
	Penalty		 int  `json:"penalty"`
}