package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	curlParser "github.com/curl-k6/curl_parser"
	"github.com/curl-k6/templates"

	helpers "github.com/curl-k6/helpers"

	gherkin "github.com/cucumber/gherkin/go"
)

func generatek6Code(fileName string) {
	content, err := ioutil.ReadFile("./features/" + fileName)

	if err != nil {
		log.Fatal(err)
	}

	reader := strings.NewReader(string(content))

	gherkinDocument, err := gherkin.ParseGherkinDocument(reader)

	if err != nil {
		log.Fatal(err)
	}

	var scenario *gherkin.Scenario

	for i := range gherkinDocument.Feature.Children {
		scenario = gherkinDocument.Feature.Children[i].(*gherkin.Scenario)

		curlCommand := scenario.Steps[0].Argument.(*gherkin.DocString).Content

		k6Script := scenario.Steps[1].Argument.(*gherkin.DocString).Content

		fileOutput := getQuotedStrings2(strings.Split(scenario.Steps[2].Text, "And")[0])[0]

		request, _ := curlParser.Parse(curlCommand)

		headers, _ := json.Marshal(request.Header)

		stringFile := templates.GenerateStarterCode(request.Url, strings.ToLower(request.Method), k6Script, helpers.GetBody(request.Body), string(headers), helpers.ProduceFormData(request.Files))

		f, err := os.Create("./features/" + fileOutput)

		if err != nil {
			fmt.Println(err)
			return
		}
		l, err := f.WriteString(stringFile)

		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(l, "bytes written successfully")
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}

		cmd := exec.Command("prettier", "--write", ".")
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Print the output
		fmt.Println(string(stdout))
	}
}

func main() {

	files, err := ioutil.ReadDir("./features")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		generatek6Code(file.Name())
	}

}

var re2 = regexp.MustCompile(`".*?"`)

func getQuotedStrings2(s string) []string {
	ms := re2.FindAllString(s, -1)
	ss := make([]string, len(ms))
	for i, m := range ms {
		ss[i] = m[1 : len(m)-1] // Note the substring of the match.
	}
	return ss

}
