# curl-to-k6

k6 script generator from curl command 

## Description

simple cli for k6 script generator from curl command

you need to create `.feature` file first using gherkin syntex

## Getting Started

### Dependencies


* node js >=12
* go 1.18


### Installing

* Install prettier
```bash
    npm i -g prettier
```

* Install cli apps

```bash
 go install github.com/dionaditya/curl-to-k6@v0.1.1
```

### Executing program

* create file with .feature extension using gherkin syntax something like this

```
upload-post.feature
Feature: upload image endpoint
  Scenario: Generate k6 script for create nwe user
    Given curl command:
    """
    curl -v -F filename=image.mp4 -F upload=@image.ppt http://localhost:8080/api/upload
    """
    And k6 options:
      """
        {
          stages: [
            {
                 duration: "5m", target: 60
            },
          ],
        }
      """
    And The file name is "generated_test.js"
```

The feature file should have Scenario with 3 steps definitions

1. Given step followed by curl command inside step argument
1. And step followed by k6 script options inside step argument
1. And step followed by generated file name inside double quote

You can modify any wording except the rules above.

Support multilple scenario inside feature file. 

* Run command 

```bash
curl-to-k6 generate -s 'source_dir_of_feature_files' -o 'output_dir_of_generated_k6_script'
```

Example

This command below will find any feature files inside current directory and put generated k6 script inside current directory
```bash
curl-to-k6 generate -s '.' -o '.'
```

This command below will find any feature files inside `feature`s directory and put generated k6 script inside `scripts` directory

```bash
curl-to-k6 generate -s './features' -o './scripts'
```
