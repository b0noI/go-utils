package fs

import (
	"os"
	"os/user"
	"path/filepath"
)

func MaybeCreateProgramFolder(programName string) error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	folderName := "." + programName
	folderPath := filepath.Join(currentUser.HomeDir, folderName)

	// Create the folder if it doesn't exist
	err = os.MkdirAll(folderPath, 0755)
	if err != nil {
		return err
	}
	return nil
}
