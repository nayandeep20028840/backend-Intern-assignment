package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"server/models"
	"server/services"
)

func SubmitJob(w http.ResponseWriter, r *http.Request) {
	var jobReq models.JobRequest
	if err := json.NewDecoder(r.Body).Decode(&jobReq); err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	// Check for count mismatch
	if jobReq.Count != len(jobReq.Visits) {
		http.Error(w, `{"error": "Count mismatch"}`, http.StatusBadRequest)
		return
	}

	// Validate each visit object
	for _, visit := range jobReq.Visits {
		if visit.StoreID == "" {
			http.Error(w, `{"error": "Missing store_id in visits/Field missing"}`, http.StatusBadRequest)
			return
		}

		if len(visit.ImageURL) == 0 {
			http.Error(w, `{"error": "Missing image_url in visits/Field missing"}`, http.StatusBadRequest)
			return
		}
	}

	// Submit the job if validation passes
	jobID := services.SubmitJob(jobReq)

	// Respond with job ID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"job_id": jobID})
}

func GetJobStatus(w http.ResponseWriter, r *http.Request) {
	jobIDStr := r.URL.Query().Get("jobid")
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		http.Error(w, `{}`, http.StatusBadRequest)
		return
	}

	response, err := services.GetJobStatus(jobID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
