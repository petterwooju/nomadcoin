package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/petterwoooju/nomadcoin/explorer"
	"github.com/petterwoooju/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome\n\n")
	fmt.Printf("Please use following flags: \n\n")
	fmt.Printf("-port=4000: start the HTML Explorer\n")
	fmt.Printf("-mode=rest: Choose between 'html' and 'rest'\n")
	runtime.Goexit()
}
func Start() {

	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")

	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest' ")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode)

}
