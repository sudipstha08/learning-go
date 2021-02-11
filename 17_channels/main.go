 // Goroutines require a mechanism to communicate within themselves
 // The communication channel between the goroutines is referred to as channels. 
 // Channels ensure that the goroutines and the main thread can communicate with each other.

package main
import "fmt"

func main() {
  dataChannel := make(chan string, 3)  // BUFFER VALUE IS MARKED AS 3
	dataChannel <- "Some Sample Data"
	dataChannel <- "Some Other Sample Data"
  dataChannel <- "Buffered Channel"
  fmt.Println(<-dataChannel)
  fmt.Println(<-dataChannel)
  fmt.Println(<-dataChannel)
}
