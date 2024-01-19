package main

import (
	"fmt"
	"os"
	"time"

	clockface "github.com/bparolini/learn-go-with-tests/16-mathematics"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Wrong number of arguments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "clockface":
		clockfaceMain()
	}
}

// region clockface

func clockfaceMain() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}

// endregion
