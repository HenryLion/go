package main
import (
	"fmt"
	"encoding/json"
	"log"
)

type Movie struct {
	Title string
	year int `json:"released"`
	Color bool `json:"color,omitempty"` // no json output if the field has the zero value for its type
	Actors []string
}

func main () {
	var movie = []Movie {
		{Title: "Casablanca",
		 year: 1942,
		 Color: false,
		 Actors: []string {"Humphrey Bogart", "Ingrid Bergman"}},
	}

	fmt.Println(movie)

	fmt.Println("-----------------\n")

	// only exported fields are marshaled
	// so year not marshaled
	data, err := json.Marshal(movie)
	if err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// ummarshal
	var titles []struct { Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}