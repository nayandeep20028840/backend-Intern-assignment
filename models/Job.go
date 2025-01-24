package models

import (
	"context"
	"sync"
)

// *************************

// JobError represents an error encountered while processing a specific store.
type JobError struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}

// Job represents a task that processes multiple stores.
type Job struct {
	ID       int
	Results  []JobResult
	Status   string
	Error    []string             // Slice to store error messages, if any
	mux      sync.Mutex           // Mutex for thread-safe access to the job's fields
	Ctx      context.Context      // Context to manage cancellation of the job
	CancelFn context.CancelFunc   // Function to cancel the job's context
}

// JobResult represents the result of processing a single store.
type JobResult struct {
	StoreID   string `json:"store_id"`
	ImageURL  string `json:"image_url"`
	Perimeter int    `json:"perimeter"`         // Calculated perimeter of the store image
	Error     string `json:"error,omitempty"`   // Optional error message if processing failed
}

// MarkFailed marks the job as failed and appends an error message.
func (j *Job) MarkFailed(err string) {
	j.mux.Lock()
	defer j.mux.Unlock()
	j.Status = "failed"
	j.Error = append(j.Error, err)
	j.CancelFn()
}

func (j *Job) MarkCompleted() {
	j.mux.Lock()
	defer j.mux.Unlock()
	if j.Status != "failed" {
		j.Status = "completed"
	}
}

// AddResult adds a processing result for a specific store to the job's results slice.
func (j *Job) AddResult(result JobResult) {
	j.mux.Lock()
	defer j.mux.Unlock()
	j.Results = append(j.Results, result)
}

// GetStatus returns the current status of the job as a map, including results or errors.
func (j *Job) GetStatus() map[string]interface{} {
	j.mux.Lock()
	defer j.mux.Unlock()   // Unlock the mutex once the function is done

	response := map[string]interface{}{
		"status": j.Status,
		"job_id": j.ID,
	}
	if j.Status == "failed" {
		response["error"] = j.Error
	} else if j.Status == "completed" {
		response["results"] = j.Results
	}
	return response
}
