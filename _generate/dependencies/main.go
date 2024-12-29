package main

import (
	"bufio"
	"fmt"
	. "github.com/Bookshelf-Writer/SimpleGenerator"
	"os"
	"strings"
)

// //////////////////////////////////////////////////////////////

func main() {
	file, err := os.Open("go.mod")
	if err != nil {
		fmt.Println("An error occurred while opening the file:", err)
		return
	}
	defer file.Close()

	dependencies := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "\t") {
			fields := strings.Fields(line)
			if len(fields) > 0 {
				dependencies[fields[0]] = fields[1]
			}
		}
	}

	obj := NewFilePathName("", "main")
	depMap := make(map[GeneratorValueObj]GeneratorValueObj)

	for dep, ver := range dependencies {
		depMap[GeneratorValueObj{Val: dep, Format: ""}] = GeneratorValueObj{Val: ver, Format: ""}
	}

	obj.AddMap("GlobalDependencies", obj.TypeString(), obj.TypeString(), depMap)
	obj.Save("value_dependencies.go")
}
