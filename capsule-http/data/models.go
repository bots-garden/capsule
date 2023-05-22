// Package data -> this package is responsible for the data layer
package data

import (
	"errors"
	"os/exec"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Status (of a task)
type Status int64

// Enums of Status
const (
	Waiting   Status = 0
	Started   Status = 1
	Finished  Status = 2
	Failed    Status = 3
	Cancelled Status = 4
)

// CapsuleTask is a struct with the parameters to start a Capsule process
type CapsuleTask struct {
	//Group            string  `json:"group"`
	FunctionName     string   `json:"name"`
	FunctionRevision string   `json:"revision"`
	Description      string   `json:"description"`
	Path             string   `json:"path"`
	Args             []string `json:"args"`
	Env              []string `json:"env"`
}


//! a same function for a same revision can exist multiple times
//? how to calculate the ID (and be sure that it is unique)?

// CapsuleProcess is a struct to describe a running Capsule process
type CapsuleProcess struct {
	//Group             string        `json:"group"`
	ID                string    `json:"id"`
	FunctionName      string    `json:"name"`
	FunctionRevision  string    `json:"revision"`
	Description       string    `json:"description"`
	CurrentStatus     Status    `json:"currentStatus"`
	StatusDescription string    `json:"statusDescription"`
	CreatedAt         time.Time `json:"createdAt"` // record the time the event was requested
	StartedAt         time.Time `json:"startedAt"`
	FinishedAt        time.Time `json:"finishedAt"`
	CancelledAt       time.Time `json:"cancelledAt"`
	FailedAt          time.Time `json:"failedAt"`
	CheckedAt         time.Time `json:"checkedAt"`
	Pid               int       `json:"pid"`
	Path              string    `json:"path"`
	Args              []string  `json:"args"`
	Env               []string  `json:"env"`
	Cmd               *exec.Cmd `json:"-"`
}

var runningCapsules sync.Map // Map of Capsule processes

// SetCapsuleProcessRecord stores a CapsuleProcess for a given key in the runningCapsules map.
//
// key: a string representing the key to store the CapsuleProcess.
// process: a CapsuleProcess to be stored in the runningCapsules map.
func SetCapsuleProcessRecord(key string, process CapsuleProcess) {
	process.ID = key
	runningCapsules.Store(key, process)
}

// CreateCapsuleProcessRecord creates a new CapsuleProcess record and returns its key.
//
// process: a CapsuleProcess struct containing information about the process.
// string: the key of the newly created CapsuleProcess record.
func CreateCapsuleProcessRecord(process CapsuleProcess) string {
	key := uuid.New().String()
	SetCapsuleProcessRecord(key, process)
	return key
}

// GetCapsuleProcessRecord retrieves the CapsuleProcess associated with a given key.
//
// key string: the key to look up in the runningCapsules map.
// CapsuleProcess: the CapsuleProcess associated with the given key, or an empty CapsuleProcess if the key is not found.
func GetCapsuleProcessRecord(key string) (CapsuleProcess, error) {
	process, ok :=runningCapsules.Load(key)
	if !ok {
		return CapsuleProcess{}, errors.New("Capsule process not found")
	}
	return process.(CapsuleProcess), nil
}

// DeleteCapsuleProcessRecord deletes a process record from the runningCapsules map given a key.
//
// key: a string representing the key to delete from the map.
// returns: a string representing the deleted key.
func DeleteCapsuleProcessRecord(key string) string {
	runningCapsules.Delete(key)
	return key
}

// GetAllCapsuleProcessRecords returns all the capsule process records.
//
// None.
// []CapsuleProcess.
func GetAllCapsuleProcessRecords() []CapsuleProcess {
	var processes []CapsuleProcess
	runningCapsules.Range(func(key, value interface{}) bool {
		processes = append(processes, value.(CapsuleProcess))
		return true
	})
	return processes
}

// GetCapsuleProcessRecordsRelatedTo returns an array of CapsuleProcess 
// records that are related to the given function name.
//
// functionName: A string representing the name of the function to match.
// []CapsuleProcess: An array of CapsuleProcess records related to the 
// given function name.
func GetCapsuleProcessRecordsRelatedTo(functionName string) []CapsuleProcess {
	var processes []CapsuleProcess
	runningCapsules.Range(func(key, value interface{}) bool {
		if value.(CapsuleProcess).FunctionName == functionName {
			processes = append(processes, value.(CapsuleProcess))
		}
		return true
	})
	return processes
}

// TODO: method(s) to delete process(es)
