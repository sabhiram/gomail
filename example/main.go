package main

import (
	"fmt"
	"os"

	"github.com/sabhiram/gomail"
	"github.com/sabhiram/gomail/types"
)

func fatalOnError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	em, err := types.NewEmailAddress("foo@bar.com")
	fatalOnError(err)

	to := types.EmailAddresses{em}
	mp, err := gomail.NewMultipartMessage(to, "Hello There!", []*gomail.MultipartSection{
		gomail.NewMultipartTextSection([]byte("This is text!")),
		gomail.NewMultipartHTMLSection([]byte("<html><body><h1>This is html!</h1></body></html>")),
	})
	fatalOnError(err)

	fmt.Printf(mp.String())
}
