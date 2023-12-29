package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

// WrappedError represents an error with additional context information.
type WrappedError struct {
	Context string
	Err     error
}

// Error returns a string representation of the error.
func (w *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", w.Context, w.Err)
}

// Wrap creates a new WrappedError with the given error and context information.
func Wrap(err error, info string) *WrappedError {
	return &WrappedError{Context: info, Err: err}
}

func main() {
	banner()
	var (
		dirNameFlag string
		fileExtFlag string
	)

	// Parse command line options
	createDir := flag.Bool("d", false, "Create a directory")
	flag.StringVar(&dirNameFlag, "dirName", "", "Directory name")
	flag.StringVar(&fileExtFlag, "ext", "", "File extension")
	flag.Parse()
	initProject(*createDir, dirNameFlag, fileExtFlag)

}

func banner() {
	fmt.Println("gowrite - The simple writing project initializer")
	fmt.Println()
	fmt.Println("usage: git gowrite [options]")
	fmt.Println()
}

func checkIfExists() {}

func initProject(createDir bool, dirName, fileExt string) {
	if createDir {
		// Check if the directory name is provided
		if dirName == "" {
			fmt.Println("Error: Directory name is required when using -d flag.")
			return
		}

		// Create the directory
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			// Wrap the error with context information
			wrappedDirErr := Wrap(err, "Error creating directory")
			fmt.Println(wrappedDirErr)
			return
		}

		// Use WaitGroup to wait for all goroutines to finish
		var wg sync.WaitGroup

		fileNames := []string{
			"characters",
			"settings",
			"locations",
			"plot_outline",
			"creatures",
			"magic_system",
			"world_map",
			"inspiration",
			"quotes",
			"extra",
		}

		// Loop to create files concurrently
		for _, fileName := range fileNames {
			wg.Add(1)
			go func(fileName string) {
				defer wg.Done()

				fullFileName := fileName
				if fileExt != "" {
					fullFileName += "." + fileExt
				}

				// Create the file
				file, err := os.Create(fullFileName)
				if err != nil {
					fmt.Println("Error creating file:", err)
					return
				}
				defer file.Close()

				fmt.Printf("File %s created.\n", fullFileName)
			}(fileName)
		}

		// Wait for all goroutines to finish before exiting
		wg.Wait()
	}
}
