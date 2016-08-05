package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Press enter to Start: ")
	reader.Scan()
	text := reader.Text()
	if text == "" {
    //todo: start round until either player quits or len(deck)==0

	}
	if err := reader.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
