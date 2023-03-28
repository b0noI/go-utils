package fs

import (
	"os"
	"os/user"
	"path/filepath"
)

func GetProgramFolderPath(programName string) (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	folderName := "." + programName
	return filepath.Join(currentUser.HomeDir, folderName), nil
}

func MaybeCreateProgramFolder(programName string) (string, error) {
	folderPath, err := GetProgramFolderPath(programName)
	if err != nil {
		return "", err
	}

	// Create the folder if it doesn't exist
	err = os.MkdirAll(folderPath, 0755)
	if err != nil {
		return "", err
	}
	return folderPath, nil
}
