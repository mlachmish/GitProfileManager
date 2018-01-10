package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"os"
	"bufio"
	"strings"
)

func main() {
	fmt.Println("Welcome to git profile manager")
	fmt.Println("------------------------------\n")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter workspace path: ")
	workspace, _ := reader.ReadString('\n')
	fmt.Print("Enter git user email: ")
	userEmail, _ := reader.ReadString('\n')

	successfulUpdates := updateGitUserEmail(strings.TrimSpace(workspace), strings.TrimSpace(userEmail))
	fmt.Printf("Done, successfully updated %d repos\n", successfulUpdates)
}


func updateGitUserEmail(workspace string, userEmail string) int {
	var successfulUpdates = 0

	files, err := ioutil.ReadDir(workspace)
	if err != nil {
		log.Println(err)
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Printf("Updating project: %s\n", f.Name())
			dirPath := workspace + "/" + f.Name()

			cmd:= exec.Command("git", "config", "user.email", userEmail)
			cmd.Dir = dirPath
			err := cmd.Run()
			if err != nil {
				log.Println(err)
				log.Printf("Trying recursively to find git repo with path %s\n", dirPath)
				successfulUpdates += updateGitUserEmail(dirPath, userEmail)
			} else {
				successfulUpdates++
			}
		}
	}

	return successfulUpdates
}
