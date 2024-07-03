package utils

import "os"

func ReadInputFile(inputFile string) (string, error) {
	inputContent, err := os.ReadFile(inputFile)
	if err != nil {
		return "", err
	}
	return string(inputContent), nil
}

func WriteOutputFile(outputFile, content string) error {
	err := os.WriteFile(outputFile, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}
