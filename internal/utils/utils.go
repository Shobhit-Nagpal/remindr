package utils

import (
	"fmt"
	"os"
)

func GetHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir, nil
}

func GetDBPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.remindr", homeDir), nil
}


func GetDBFile() (string, error) {
	dbPath, err := GetDBPath()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/db.json", dbPath), nil
}
