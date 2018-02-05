package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/tbertrand25/gitlab"
)

func main() {
	// Ensure we have a valid command

	if os.Args[1] == "search-repo" {
		if len(os.Args) == 3 {
			client := gitlab.MakeGlClient()
			name := os.Args[2]
			fmt.Println(client.SearchProjects(name))
		}
	} else if os.Args[1] == "search-groups" {
		if len(os.Args) == 3 {
			client := gitlab.MakeGlClient()
			name := os.Args[2]
			fmt.Println(client.SearchGroups(name))
		}
	} else if os.Args[1] == "create-repo" {
		if len(os.Args) == 3 {
			client := gitlab.MakeGlClient()
			path := os.Args[2]
			client.CreateProject(path)
		}
	} else if os.Args[1] == "create-branch" {
		if len(os.Args) == 5 {
			client := gitlab.MakeGlClient()
			path := os.Args[2]
			branchName := os.Args[3]
			ref := os.Args[4]
			client.CreateBranch(path, branchName, ref)
		}
	} else if os.Args[1] == "protect-branch" {
		if len(os.Args) == 5 {
			client := gitlab.MakeGlClient()
			path := os.Args[2]
			push, _ := strconv.Atoi(os.Args[3])
			merge, _ := strconv.Atoi(os.Args[4])
			client.ProtectBranch(path, push, merge)
		}
	} else if os.Args[1] == "create-gms-repo" {
		if len(os.Args) == 3 {
			client := gitlab.MakeGlClient()
			path := os.Args[2]
			client.CreateProject(path)
			client.CreateBranch(path, "develop", "master")
			client.ProtectBranch(path+"#master", 0, 40)
			client.ProtectBranch(path+"#develop", 0, 30)
			client.SetDefaultBranch(path + "#develop")
		}
	} else if os.Args[1] == "set-default-branch" {
		if len(os.Args) == 3 {
			client := gitlab.MakeGlClient()
			path := os.Args[2]
			client.SetDefaultBranch(path)
		}
	}
}
