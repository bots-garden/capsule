// Package data -> this package is responsible for the data layer
package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/bots-garden/capsule/capsule-http/tools"
)

var runningCapsules sync.Map // Map of Capsule processes
var storage = tools.GetEnv("CAPSULE_PROCESSES_STORAGE", "processes.json")

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
	Index             int    `json:"index"`
	FunctionName      string    `json:"name"`
	FunctionRevision  string    `json:"revision"`
	HTTPPort          string    `json:"httpPort"`
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

// SetCapsuleProcessRecord stores a CapsuleProcess for a given key in the runningCapsules map.
//
// key: a string representing the key to store the CapsuleProcess.
// process: a CapsuleProcess to be stored in the runningCapsules map.
func SetCapsuleProcessRecord(process CapsuleProcess) string {
	key:= process.FunctionName+"/"+process.FunctionRevision+"/"+strconv.Itoa(process.Index)
	//fmt.Println("🔑", key)
	runningCapsules.Store(key, process)
	return key
}

// SearchLastIndexOfProcessRecord returns the last order number of the processes
// with the given function name and revision.
//
// functionName: the name of the function to search for.
// functionRevision: the revision of the function to search for.
// int: the last order number of the processes found.
func SearchLastIndexOfProcessRecord(functionName, functionRevision string) int {
	var processes []CapsuleProcess
	runningCapsules.Range(func(key, value interface{}) bool {
		if value.(CapsuleProcess).FunctionName == functionName && value.(CapsuleProcess).FunctionRevision == functionRevision {
			processes = append(processes, value.(CapsuleProcess))
		}
		return true
	})
	orderNum := -1

	for _, process := range processes {
		if process.Index > orderNum {
			orderNum = process.Index
		}
	}
	return orderNum
}


// CreateCapsuleProcessRecord creates a new CapsuleProcess record and returns its key.
//
// process: a CapsuleProcess struct containing information about the process.
// string: the key of the newly created CapsuleProcess record.
func CreateCapsuleProcessRecord(process CapsuleProcess) string {
	process.Index = SearchLastIndexOfProcessRecord(process.FunctionName, process.FunctionRevision) + 1
	return SetCapsuleProcessRecord(process)
}

// GetCapsuleProcessRecord retrieves the CapsuleProcess associated with a given key.
//
// key string: the key to look up in the runningCapsules map.
// CapsuleProcess: the CapsuleProcess associated with the given key, or an empty CapsuleProcess if the key is not found.
func GetCapsuleProcessRecord(key string) (CapsuleProcess, error) {
	process, ok := runningCapsules.Load(key)
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

// === Persistence ===

// SaveCapsuleProcessRecords saves the runningCapsules map to a json file
// at the specified storage location. It returns an error if one occurs.
//
// Return: error
func SaveCapsuleProcessRecords() error {
	// save map to json file
	jsonString, err := json.MarshalIndent(runningCapsules, "", "  ")
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(storage, jsonString, 0644)
	if err != nil {
		log.Println(err)
	}
	return err
}
