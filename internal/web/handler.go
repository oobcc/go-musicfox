package web

import (
	"fmt"
	"net/http"
)

var Status *musicStatus

type musicStatus struct {
	Lyrics     string
	IsPlay     bool
	SongName   string
	ArtistName string
}

func getMusicMessage(w http.ResponseWriter, r *http.Request) {

	icon := ""

	if Status.IsPlay {
		icon = "▶"
	} else {
		icon = "||"
	}

	fmt.Fprintf(w, "%s %s %s %s", icon, Status.SongName, Status.ArtistName, Status.Lyrics)

}

func Headler() {
	Status = &musicStatus{}
	Status.IsPlay = false
	Status.Lyrics = ""

	http.HandleFunc("/", getMusicMessage)

	http.ListenAndServe(":9800", nil)
}
