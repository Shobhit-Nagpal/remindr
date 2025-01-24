package db

import (
	"os"

	"github.com/Shobhit-Nagpal/remindr/internal/jobs"
	"github.com/Shobhit-Nagpal/remindr/internal/utils"
)

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
	}

	return nil
}

func GetAllJobs() []jobs.Job {
	return []jobs.Job{}
}
