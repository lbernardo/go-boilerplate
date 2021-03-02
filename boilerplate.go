package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var files = []string{
	"go.mod",
	"cmd/service/routes.go",
}

var pkg string

func init() {
	pkg = os.Args[1]
}

func main() {

	if pkg == "" {
		panic("Package is required")
	}
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		contentStr := string(content)
		contentStr = strings.ReplaceAll(contentStr, "github.com/lbernardo/go-boilerplate", pkg)
		if err := ioutil.WriteFile(file, []byte(contentStr), 07555) ; err != nil {
			fmt.Println("Error in write file", file)
		}
	}
}
