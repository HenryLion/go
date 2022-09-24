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

func getids () (ids []uint32) {
	idlist := []uint32 {1,2,3}
	ids = idlist
	return
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

	ids := []uint64 {123456852,45874525}
	id2 := ids
	id2[0] = 3
	fmt.Printf("ids is: %v\n",ids)
	fmt.Printf("id2 is: %v\n",id2)
	ids = append(ids, 0)
	jsonStr, err := json.Marshal(ids)

	fmt.Printf("%s\n", jsonStr)

	fmt.Println(fmt.Sprintf("(id in %s)",jsonStr))
	fmt.Println(getids())

	fmt.Println(len(ids))
	for i,v := range ids {
		fmt.Println(i, v)
	}
}