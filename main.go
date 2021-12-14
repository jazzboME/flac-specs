package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	pathPtr := flag.String("path", "./*.flac", "path to examine")
	boolPtr := flag.Bool("detail", false, "Determine if process should detail each file or produce summary")
	flag.Parse()

	files, _ := filepath.Glob(caseInsensitive(*pathPtr))
	if len(files) == 0 {
		fmt.Println("No matching files found.")
		os.Exit(0)
	}

	result, err := examine(files, *boolPtr)
	
	fmt.Print(result)
	if err!= nil {
		fmt.Printf("Error: %v", err)
	}
}