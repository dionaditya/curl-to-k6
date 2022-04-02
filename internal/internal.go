package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	curlParser "github.com/dionaditya/curl-to-k6/curl_parser"
	"github.com/dionaditya/curl-to-k6/templates"

	helpers "github.com/dionaditya/curl-to-k6/helpers"

	gherkin "github.com/cucumber/gherkin/go"
)

func generatek6Code(fileName string, directoryName string, outDir string) {
	content, err := ioutil.ReadFile(directoryName + "/" + fileName)

	if err != nil {
		fmt.Println("test")
		fmt.Println(err)
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

		if _, err := os.Stat(outDir); os.IsNotExist(err) {
			err = os.Mkdir(outDir, 0755)
			// TODO: handle error
		}

		f, err := os.Create(outDir + "/" + fileOutput)

		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = f.WriteString(stringFile)

		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(fileOutput, "written successfully")
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}

		cmd := exec.Command("prettier", "--write", outDir)
		_, err = cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

	}
}

func Run(sourceDir string, outDir string) {

	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !strings.Contains(file.Name(), ".feature") {
			continue
		}

		generatek6Code(file.Name(), sourceDir, outDir)
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
