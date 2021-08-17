package main

import "fmt"

func getIllegalKey(kfuin uint64, userId string) string {
	return fmt.Sprintf("%d_%s", kfuin, userId)
}

func main() {
	fmt.Printf("%s\n", getIllegalKey(123454321, "56778"))

	fmt.Println(uint32(1))
}
