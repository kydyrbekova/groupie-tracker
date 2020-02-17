package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

var art Artist
var api API

type API struct {
	Ar Artist
}
type Artist []struct {
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	// Locations    string   `json:"locations"`
	// ConcertDates string   `json:"concertDates"`
	// Relations    string   `json:"relations"`
}

type Locations struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	ID    int64    `json:"id"`
	Dates []string `json:"dates"`
}

type Index struct {
	ID             int64               `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// var API struct {
// 	Artist    Artist
// 	Data      Dates
// 	Locations Locations
// }

func TreatingData() {
	artists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	// locations, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	// dates, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	// relations, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatal(err)
	}

	fileArtists, err := ioutil.ReadAll(artists.Body)
	// fileLocations, err := ioutil.ReadAll(locations.Body)
	// fileDates, err := ioutil.ReadAll(dates.Body)
	// fileRelations, err := ioutil.ReadAll(relations.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(fileArtists, &art)
	api.Ar = art

	// json.Unmarshal(fileLocations, &art)

	// json.Unmarshal(fileDates, &art)

	// json.Unmarshal(fileRelations, &art)

}

func main() {

	TreatingData()
	http.HandleFunc("/", handlefunc)
	fmt.Println("Listening to port :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handlefunc(w http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.Error(w, "Error 404\n Page not Found", http.StatusNotFound)
		return
	}
	temp := template.Must(template.ParseFiles("html/index.html"))

	fmt.Println(art)

	temp.ExecuteTemplate(w, "index.html", art)
}