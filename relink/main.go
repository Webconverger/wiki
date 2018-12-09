package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`\[\[(.*?)\]\]`)
var reTitle = regexp.MustCompile(`\[\[\!meta title="(.*)"]]`)

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	output := reTitle.ReplaceAllStringFunc(string(b), frontmatter)
	output = re.ReplaceAllStringFunc(output, relink)
	fmt.Print(output)
}

func relink(input string) string {
	parts := strings.Split(re.FindStringSubmatch(input)[1], "|")
	var name, target string
	if len(parts) == 2 {
		name, target = parts[0], parts[1]
	} else {
		name, target = parts[0], parts[0]
	}
	return fmt.Sprintf("[%s](/%s/)", strings.Replace(name, "_", " ", -1), strings.ToLower(target))
}

func frontmatter(input string) (output string) {
	output = "---\n"
	js, err := json.MarshalIndent(struct {
		Title     string `json:"title"`
		Permalink string `json:"permalink"`
	}{reTitle.FindStringSubmatch(input)[1],
		"/" + strings.TrimSuffix(os.Args[1], filepath.Ext(os.Args[1])) + "/",
	}, "", "    ")
	if err != nil {
		panic(err)
	}
	output += string(js)
	output += "\n---"
	return
}
