package spinner

import (
	"fmt"
	"time"
)

func Spinner() {
	// Create a spinner.
	spinner := make([]rune, 5)
	spinner[0] = '|'
	spinner[1] = '/'
	spinner[2] = '-'
	spinner[3] = '\\'
	spinner[4] = '|'

	// Start the spinner.
	for i := 0; i < 5; i++ {
		fmt.Print(spinner[i])
		time.Sleep(100 * time.Millisecond)
	}
	// Finish the spinner.
	fmt.Println()
}
