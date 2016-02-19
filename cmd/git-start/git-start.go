package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
)

type Brunch struct {
	BrunchType  string
	IssueNumber string
	Description string
}

func (b Brunch) Print() {
	log.Printf("%s:\t%s", color.GreenString("BrunchType"), color.CyanString(b.BrunchType))
	log.Printf("%s:\t%s", color.GreenString("IssueNumber"), color.CyanString(b.IssueNumber))
	log.Printf("%s:\t%s", color.GreenString("Description"), color.CyanString(b.Description))
}

func (b Brunch) Start() {
	out, err := exec.Command("git", "checkout", "-b", b.Name()).CombinedOutput()
	if err != nil {
		log.Println(color.RedString("Error has occured"))
		log.Println(string(out))
		os.Exit(1)
	}
	log.Print(string(out))
}

func (b Brunch) Name() string {
	return fmt.Sprintf("%s/%s-%s", b.BrunchType, b.IssueNumber, b.Description)
}

func main() {
	log.SetFlags(0)
	flag.Parse()
	args := flag.Args()

	if len(args) < 3 {
		log.Println(color.RedString("arguments was not enough"))
		log.Println("ex) git start f 1234 start-brunch")
		os.Exit(1)
	}

	brunchType := detectBrunchType(args[0])
	issueNumber := args[1]
	description := args[2]

	brunch := Brunch{
		BrunchType:  brunchType,
		IssueNumber: issueNumber,
		Description: description,
	}
	brunch.Print()
	brunch.Start()
}

func detectBrunchType(str string) string {
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
