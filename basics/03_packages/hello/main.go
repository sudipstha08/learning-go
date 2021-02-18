package hello

import (
	"fmt"

	"learning-go/basics/03_packages/grettings"
)

func main() {
	// Get a greeting message and print it.
	message := grettings.Hello("Dec")
	fmt.Println(message)
}
