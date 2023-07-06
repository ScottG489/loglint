package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type rule struct {
	code         string
	name         string
	labels       []string
	shortReason  string
	regexPattern string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file specified")
		os.Exit(1)
	}
	filename := os.Args[1]

	file, err := os.Open(filename)
	check(err)

	exitStatus := validateRules(getRules(), file)
	err = file.Close()
	check(err)
	os.Exit(exitStatus)
}

func validateRules(regexRules []rule, file *os.File) int {
	exitStatus := 0
	lineNumber := 1
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lineNumber++
		text := scanner.Text()
		for _, rule := range regexRules {
			pattern := regexp.MustCompile(rule.regexPattern)
			match := pattern.FindString(text)
			if match != "" {
				fmt.Println(fmt.Sprintf("[%s] %s:%d: '%s'",
					rule.code,
					rule.name,
					lineNumber,
					match))
				exitStatus = 1
			}
		}
	}
	return exitStatus
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
		{
			"LL105",
			"NPM warning",
			[]string{"npm"},
			"Generic NPM warning.",
			"^npm WARN.*",
		},
	}
}
