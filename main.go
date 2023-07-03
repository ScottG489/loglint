package main

import (
	"fmt"
	"os"
	"regexp"
)

type rule struct {
	code         string
	name         string
	labels     []string
	shortReason  string
	regexPattern string
}

func main() {
	fileContents := getLogFileContents()

	regexRules := getRules()

	exitStatus := 0
	for _, rule := range regexRules {
		pattern := regexp.MustCompile(rule.regexPattern)
		match := pattern.Match([]byte(fileContents))
		if match {
			fmt.Println(fmt.Sprintf("%s: '%s'", rule.name, string(pattern.Find([]byte(fileContents)))))
			exitStatus = 1
		}
	}
	os.Exit(exitStatus)
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

func getRules() []rule {
	return []rule{
		{
			"LL100",
			"Generic error",
			[]string{"generic"},
			"Errors are bad.",
			"(?i).*error:.*",
		},
		{
			"LL102",
			"Generic warning",
			[]string{"generic"},
			"Warnings are bad.",
			"(?i).*warning:.*",
		},
		{
			"LL103",
			"Terraform lock file changed",
			[]string{"terraform"},
			"Terraform lock file should be checked in.",
			"^Terraform has made some changes to the provider dependency selections recorded",
		},
		{
			"LL104",
			"Terraform lock file created",
			[]string{"terraform"},
			"Terraform lock file should be checked in.",
			"^Terraform has created a lock file \\S* to record the provider",
		},
	}
}
