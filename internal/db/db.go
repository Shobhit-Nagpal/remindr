package db

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Shobhit-Nagpal/remindr/internal/jobs"
)

type DB struct {
	dir  string
	file string
	mu   *sync.RWMutex
}

type Data struct {
	Jobs map[string]jobs.Job `json:"jobs"`
}

func NewDB(dir, file string) (*DB, error) {
	db := &DB{
		dir:  dir,
		file: file,
		mu:   &sync.RWMutex{},
	}

	err := db.ensureDB()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) createDBFile() error {
	fileName := fmt.Sprintf("%s/%s", db.dir, db.file)
	_, err := os.Create(fileName)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, []byte("{}"), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetAllJobs() (map[string]jobs.Job, error) {
	data := &Data{}
	db.mu.Lock()
	defer db.mu.Unlock()

	data, err := db.loadDB()
	if err != nil {
		return nil, err
	}

	return data.Jobs, nil
}

func (db *DB) CreateJob(message string, interval int, level string) (*jobs.Job, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	job := jobs.CreateJob(message, interval, jobs.Level(level))

	data, err := db.loadDB()
	if err != nil {
		return nil, err
	}

	data.Jobs[job.ID] = *job

	err = db.writeDB(data)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (db *DB) ensureDB() error {
	if _, err := os.Stat(db.dir); os.IsNotExist(err) {
		err = os.Mkdir(db.dir, 0700)
		if err != nil {
			return err
		}

		err = db.createDBFile()
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) DeleteJob(id string) (*jobs.Job, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, err := db.loadDB()
	if err != nil {
		return nil, err
	}

	deleteId := ""
	job := jobs.Job{}

	for jobId, j := range data.Jobs {
		if id == strings.Split(jobId, "-")[0] {
			deleteId = jobId
			job = j
		}
	}

	delete(data.Jobs, deleteId)

	err = db.writeDB(data)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (db *DB) UpdateJob(id string, active bool) (*jobs.Job, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, err := db.loadDB()
	if err != nil {
		return nil, err
	}

  updatedId := ""
	job := jobs.Job{}

	for jobId, j := range data.Jobs {
		if id == strings.Split(jobId, "-")[0] {
      updatedId = jobId
			job = j
		}
	}

	job.Active = active
	data.Jobs[updatedId] = job

	err = db.writeDB(data)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (db *DB) loadDB() (*Data, error) {
	data := &Data{
		Jobs: map[string]jobs.Job{},
	}

	filePath := fmt.Sprintf("%s/%s", db.dir, db.file)

	dbData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(dbData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (db *DB) writeDB(data *Data) error {
	dat, err := json.Marshal(data)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", db.dir, db.file)

	err = os.WriteFile(filePath, []byte(dat), 0644)
	if err != nil {
		return err
	}

	return nil
}
