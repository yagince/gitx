package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
)

type Branch struct {
	BranchType  string
	IssueNumber string
	Description string
}

func (b Branch) Print() {
	log.Printf("%s:\t%s", color.GreenString("BranchType"), color.CyanString(b.BranchType))
	log.Printf("%s:\t%s", color.GreenString("IssueNumber"), color.CyanString(b.IssueNumber))
	log.Printf("%s:\t%s", color.GreenString("Description"), color.CyanString(b.Description))
}

func (b Branch) Start() {
	out, err := exec.Command("git", "checkout", "-b", b.Name()).CombinedOutput()
	if err != nil {
		log.Println(color.RedString("Error has occured"))
		log.Println(string(out))
		os.Exit(1)
	}
	log.Print(string(out))
}

func (b Branch) Name() string {
	return fmt.Sprintf("%s/%s-%s", b.BranchType, b.IssueNumber, b.Description)
}

func main() {
	log.SetFlags(0)
	flag.Parse()
	args := flag.Args()

	if len(args) < 3 {
		log.Println(color.RedString("arguments was not enough"))
		log.Println("ex) git start f 1234 start-branch")
		os.Exit(1)
	}

	branchType := detectBranchType(args[0])
	issueNumber := args[1]
	description := args[2]

	branch := Branch{
		BranchType:  branchType,
		IssueNumber: issueNumber,
		Description: description,
	}
	branch.Print()
	branch.Start()
}

func detectBranchType(str string) string {
	switch str {
	case "f":
		return "feature"
	case "c":
		return "camp"
	case "h":
		return "hotfix"
	case "s":
		return "spark"
	default:
		return str
	}
}
