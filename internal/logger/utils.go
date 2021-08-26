package logger

import (
	"os"
	"strings"
)

// Configure the file to output logs
// If the file exists, append the log to the file
// If the file does not exist, create the file and output the log
func openFile(fileName string) (*os.File, error) {
	if exists(fileName) {
		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0777)
		return f, err
	}
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0777)
	return f, err
}

// Extraction of file names to be listed in the log
func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]

}

// Check if the file exists
func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
