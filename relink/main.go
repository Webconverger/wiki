package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// KISS
var re = regexp.MustCompile(`\[\[(.*?)\]\]`)

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", relinkAll(string(b)))
}

func relink(input string) string {
	parts := strings.Split(re.FindStringSubmatch(input)[1], "|")
	log.Printf("parts: %#v", parts)
	var name, target string
	if len(parts) == 2 {
		name, target = parts[0], parts[1]
	} else {
		name, target = parts[0], parts[0]
	}
	return fmt.Sprintf("[%s](/%s/)", strings.Replace(name, "_", " ", -1), strings.ToLower(target))
}

func relinkAll(input string) string {
	return re.ReplaceAllStringFunc(input, relink)
}
