package main

import (
	"fmt"
	"os"
)

//".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"

func main() {
	args := []string(os.Args[1:])
	fmt.Println(args)
}
