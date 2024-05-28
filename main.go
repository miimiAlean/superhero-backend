package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Superhero struct {
	Name      string `json:"name"`
	Biography struct {
		FullName string `json:"fullName"`
	} `json:"biography"`
	Powerstats struct {
		Intelligence int `json:"intelligence"`
		Strength     int `json:"strength"`
		Speed        int `json:"speed"`
		Durability   int `json:"durability"`
		Power        int `json:"power"`
		Combat       int `json:"combat"`
	} `json:"powerstats"`
	Images struct {
		XS string `json:"xs"`
		SM string `json:"sm"`
		MD string `json:"md"`
		LG string `json:"lg"`
	} `json:"images"`
}

var superheroes = map[string]Superhero{
	"wolverine": {
		Name: "Wolverine",
		Biography: struct {
			FullName string `json:"fullName"`
		}{FullName: "John Logan"},
		Powerstats: struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		}{
			Intelligence: 63, Strength: 32, Speed: 50, Durability: 100, Power: 89, Combat: 100,
		},
		Images: struct {
			XS string `json:"xs"`
			SM string `json:"sm"`
			MD string `json:"md"`
			LG string `json:"lg"`
		}{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/717-wolverine.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/717-wolverine.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/717-wolverine.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/717-wolverine.jpg",
		},
	},
	"spiderman": {
		Name: "Spider-Man",
		Biography: struct {
			FullName string `json:"fullName"`
		}{FullName: "Peter Parker"},
		Powerstats: struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		}{
			Intelligence: 90, Strength: 55, Speed: 67, Durability: 75, Power: 74, Combat: 85,
		},
		Images: struct {
			XS string `json:"xs"`
			SM string `json:"sm"`
			MD string `json:"md"`
			LG string `json:"lg"`
		}{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/620-spider-man.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/620-spider-man.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/620-spider-man.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/620-spider-man.jpg",
		},
	},
	// Agrega los demás superhéroes aquí
	"ironman": {
		Name: "Iron Man",
		Biography: struct {
			FullName string `json:"fullName"`
		}{FullName: "Tony Stark"},
		Powerstats: struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		}{
			Intelligence: 100, Strength: 85, Speed: 58, Durability: 85, Power: 100, Combat: 64,
		},
		Images: struct {
			XS string `json:"xs"`
			SM string `json:"sm"`
			MD string `json:"md"`
			LG string `json:"lg"`
		}{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/346-iron-man.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/346-iron-man.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/346-iron-man.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/346-iron-man.jpg",
		},
	},
	"blackwidow": {
		Name: "Black Widow",
		Biography: struct {
			FullName string `json:"fullName"`
		}{FullName: "Natasha Romanoff"},
		Powerstats: struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		}{
			Intelligence: 75, Strength: 13, Speed: 33, Durability: 30, Power: 36, Combat: 100,
		},
		Images: struct {
			XS string `json:"xs"`
			SM string `json:"sm"`
			MD string `json:"md"`
			LG string `json:"lg"`
		}{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/107-black-widow.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/107-black-widow.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/107-black-widow.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/107-black-widow.jpg",
		},
	},
	"thor": {
		Name: "Thor",
		Biography: struct {
			FullName string `json:"fullName"`
		}{FullName: "Thor Odinson"},
		Powerstats: struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		}{
			Intelligence: 69, Strength: 100, Speed: 83, Durability: 100, Power: 100, Combat: 100,
		},
		Images: struct {
			XS string `json:"xs"`
			SM string `json:"sm"`
			MD string `json:"md"`
			LG string `json:"lg"`
		}{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/659-thor.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/659-thor.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/659-thor.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/659-thor.jpg",
		},
	},
}

func getSuperhero(w http.ResponseWriter, r *http.Request) {
	hero := r.URL.Query().Get("hero")
	if hero == "" {
		http.Error(w, "Missing hero query parameter", http.StatusBadRequest)
		return
	}

	superhero, ok := superheroes[hero]
	if !ok {
		http.Error(w, "Superhero not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(superhero)
}

func main() {
	http.HandleFunc("/api/superhero", getSuperhero)
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
