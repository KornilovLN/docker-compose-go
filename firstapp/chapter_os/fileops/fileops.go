package fileops

import (
	logging "firstapp/chapter_os/log"
	"os"
	"os/exec"
)

var debugMode = true

func CreateFile(filename string) {
	logging.LogInfo("Creating file: '%s'", filename)

	filePtr, err := os.Create(filename)
	if err != nil {
		logging.LogError(err)
	}
	defer filePtr.Close() // закрываем файл
	logging.LogInfo("File '%s' created successfully", filename)
}

func OpenFile(filename string) {
	logging.LogInfo("Opening file '%s'", filename)

	filePtr, err := os.Open(filename)
	if err != nil {
		logging.LogError(err)
	}
	defer filePtr.Close() // закрываем файл

	logging.LogInfo("File '%s' opened successfully", filename)
}

func ReadFileContents(filename string) {
	logging.LogInfo("Reading file '%s' contents", filename)

	bytes, err := os.ReadFile(filename)
	if err != nil {
		logging.LogError(err)
	}
	var fileText = string(bytes[:])

	logging.LogInfo("File '%s' contents: '%s'", filename, fileText)
}

func WriteFileContents(filename string, content string) {
	logging.LogInfo("Writing to file '%s'", filename)

	err := os.WriteFile(filename, []byte(content), 0666)
	if err != nil {
		logging.LogFatal(err)
	}

	logging.LogInfo("File '%s' written successfully", filename)
}

func GetTreeOutput(path string) (string, error) {
	cmd := exec.Command("tree", path)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
