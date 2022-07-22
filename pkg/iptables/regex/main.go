package main

import (
	"fmt"
	"regexp"
	"strings"
)

var commentRegex *regexp.Regexp

func init() {
	var err error
	commentRegex, err = regexp.Compile(`\/\*(?:[^\*]|\*+[^\/\*])*\*+\/`)
	if err != nil {
		panic(err)
	}
}

func main() {
	action := "xahahah"
	// comments := commentRegex.FindStringSubmatch(action)
	a, comments := parseCommnet(action)
	for _, v := range comments {
		fmt.Println(string(v))
	}
	fmt.Println(a)
}

func parseCommnet(action string) (string, []string) {
	comments := commentRegex.FindAllString(action, -1)
	return strings.TrimSpace(commentRegex.ReplaceAllString(action, "")), comments
}
