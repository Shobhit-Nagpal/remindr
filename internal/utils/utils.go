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

func CreateDBFile(dir string) error {
	fileName := fmt.Sprintf("%s/db.json", dir)
	_, err := os.Create(fileName)
	if err != nil {
		return err
	}

	return nil
}

func GetDBFile() (string, error) {
	dbPath, err := GetDBPath()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/db.json", dbPath), nil
}
