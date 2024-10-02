package demo

import (
	"firstapp/chapter_os/fileops"
	logging "firstapp/chapter_os/log"
)

const (
	FILEname = "files/demo.txt"
	PATHname = "../"
)

func DemonstrateFileOperations() {
	logging.LogInfo("Starting file operations demonstration")

	fileops.CreateFile(FILEname)
	fileops.OpenFile(FILEname)
	fileops.ReadFileContents(FILEname)

	// Запись в файл результатов выполнения команды tree
	treeOutput, err := fileops.GetTreeOutput(PATHname)
	if err != nil {
		logging.LogError(err)
	} else {
		fileops.WriteFileContents(FILEname, treeOutput)
	}

	fileops.ReadFileContents(FILEname)

	logging.LogInfo("File operations demonstration completed")
}
