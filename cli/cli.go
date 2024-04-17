package cli

import (
	"flag"
	"fmt"
	"os"
)

func useage() {
	fmt.Printf("Welome to resistercoin\n\n")
	fmt.Printf("Please sue the following flags:\n\n")
	fmt.Printf("-port:\t\t Set the PORT of the server\n")
	fmt.Printf("-mode:\t\t Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		useage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest")

	fmt.Println(port)
	flag.Parse()

	switch *mode {
	case "rest":
		// rest.Start(*port)
	case "html":
		// explorer.Start(*port)
	default:
		useage()
	}
}
