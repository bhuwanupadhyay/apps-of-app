package main

import (
	"fmt"
	"io.github.bhuwanupadhyay/apps-of-app/cmd"
	"os"
)

func main() {
	workingDir, err := extractTemporarily()
	if err != nil {
		panic(err)
	}
	cmd.Execute(workingDir)
	fmt.Println("")
	cleanWorkingDir(workingDir)
}

func cleanWorkingDir(workingDir string) {
	// cleanup
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			panic(err)
		}
	}(workingDir)
}
