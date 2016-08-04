package main

import (
  "fmt"
  "bufio"
  "os"
)

func init() {

}

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Press G and enter to Start: ")
  text, _ := reader.ReadString('\n')
  if(text == "G" || text == "g") {
    //Start the GAME
  


  }
}
