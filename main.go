package main

import (
	"fmt"
	"log"
	"os"
	"time"

	clockface "github.com/bparolini/learn-go-with-tests/16-mathematics"
	blogposts "github.com/bparolini/learn-go-with-tests/17-reading_files"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Wrong number of arguments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "clockface":
		clockfaceMain()
	case "reading_files":
		readingFilesMain()
	}
}

// region clockface

func clockfaceMain() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}

// endregion

// region reading files

func readingFilesMain() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("17-reading_files/posts"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(posts)
}

// endregion
