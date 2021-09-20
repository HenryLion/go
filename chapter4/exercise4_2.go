package main
import (
	"fmt"
	"os"
	"crypto/sha256"
	"crypto/sha512"
)

func rst (str string) {
	if len(os.Args) == 1 {
		fmt.Printf("%x\n",sha256.Sum256([]byte(str)) )
	} else if (len(os.Args) == 2) {
		if (os.Args[1] == "384"){
			fmt.Printf("%x\n",sha512.Sum384([]byte(str)) )
		} else if (os.Args[1] == "512") {
			fmt.Printf("%x\n", sha512.Sum512([]byte(str)))
		} else {
			fmt.Println("param error.")
		}
	} else {
		fmt.Println("param error.")
	}
}


func main () {
	rst("x")
}