package main

import (
	"fmt"
	"os"
	"regexp"
)

type rule struct {
	name         string
	shortReason  string
	regexPattern string
}

func main() {
	fileContents := getLogFileContents()

	regexRules := getRules()

	exitStatus := 0
	for _, rule := range regexRules {
		pattern := regexp.MustCompilePOSIX(rule.regexPattern)
		match := pattern.Match([]byte(fileContents))
		if match {
			fmt.Println(fmt.Sprintf("%s: %s", rule.name, string(pattern.Find([]byte(fileContents)))))
			exitStatus = 1
		}
	}
	os.Exit(exitStatus)
}

func getRules() []rule {
	return []rule{
		{"Generic warning", "Warnings are bad.", ".*warning:.*"},
	}
}

func getLogFileContents() string {
	if len(os.Args) < 2 {
		fmt.Println("No file specified")
		os.Exit(1)
	}
	arg := os.Args[1]
	dat, err := os.ReadFile(arg)
	check(err)
	fileContents := string(dat)
	return fileContents
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
