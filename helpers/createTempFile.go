package helpers

import (
	"os"
	"path/filepath"
	"xhinobi-go/constants"
)


func CreateTempFile(text string) (string, error) {
	var tempFilePath string
	if constants.IsCloudEnvironment {
		tempFilePath = filepath.Join(os.Getenv("HOME"), constants.TEMP_FILE_NAME)
	} else {
		tempFilePath = filepath.Join(os.TempDir(), constants.TEMP_FILE_NAME)
	}

	_, err := os.Stat(tempFilePath)
	if err == nil {
		err = os.Remove(tempFilePath)
		if err != nil {
			return "", err
		}
	}

	err = os.WriteFile(tempFilePath, []byte(text), 0644)
	if err != nil {
		return "", err
	}

	return tempFilePath, nil
}