package main

import (
	"fmt"
	"os"

	"goreloaded/utils"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("We need: prog_name input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	inputContent, err := utils.ReadInputFile(inputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	editedContent, _ := utils.ContentEdit(inputContent)

	err = utils.WriteOutputFile(outputFile, editedContent)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("It's OK!")
}
