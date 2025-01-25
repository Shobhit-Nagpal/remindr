package db

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shobhit-Nagpal/remindr/internal/jobs"
	"github.com/Shobhit-Nagpal/remindr/internal/utils"
)

type DB struct {
	Jobs []jobs.Job `json:"jobs"`
}

func InitDB() error {
	dbPath, err := utils.GetDBPath()
	if err != nil {
		return err
	}

	if _, err = os.Stat(dbPath); os.IsNotExist(err) {
		err = os.Mkdir(dbPath, 0700)
		if err != nil {
			return err
		}

		err = createDBFile(dbPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func createDBFile(dir string) error {
	fileName := fmt.Sprintf("%s/db.json", dir)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	db := DB{
		Jobs: []jobs.Job{},
	}

	data, err := json.Marshal(db)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func GetAllJobs() ([]*jobs.Job, error) {
	jobs := []*jobs.Job{}
	filePath, err := utils.GetDBFile()
	if err != nil {
		return jobs, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return jobs, err
	}

	err = json.Unmarshal(data, &jobs)
	if err != nil {
		return jobs, err
	}

	return jobs, nil
}
