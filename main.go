package main

import (
	"fmt"
	"io.github.bhuwanupadhyay/apps-of-app/cmd"
	"os"
)

func main() {

	fmt.Println("------------------ <BEGIN> ----------------------------")
	workingDir, err := extractTemporarily()
	fmt.Printf("Working directory: %s\n", workingDir)
	//err = listDir(workingDir)
	fmt.Println("------------------ <EXECUTE> --------------------------")
	fmt.Println("")

	if err != nil {
		panic(err)
	}

	cmd.Execute(workingDir)

	fmt.Println("")
	fmt.Println("----------------- <FINISHED> -------------------------")
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
