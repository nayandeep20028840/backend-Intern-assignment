package services

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
	"server/models"
	"server/utils"
)

var (
	jobs       map[int]*models.Job
	jobCounter int
	jobMutex   sync.Mutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
	jobs = make(map[int]*models.Job)
	jobCounter = 1
}

func SubmitJob(jobReq models.JobRequest) int {
	job := &models.Job{
		ID:      jobCounter,
		Status:  "ongoing",
		Results: make([]models.JobResult, 0),
		Ctx:     context.Background(),
	}
	job.Ctx, job.CancelFn = context.WithCancel(job.Ctx)

	jobMutex.Lock()
	jobs[job.ID] = job
	jobCounter++
	jobMutex.Unlock()

	go processJob(job, jobReq)
	return job.ID
}

func processJob(job *models.Job, jobReq models.JobRequest) {
	var wg sync.WaitGroup

	for _, visit := range jobReq.Visits {
		if _, ok := stores[visit.StoreID]; !ok {
			job.MarkFailed(fmt.Sprintf("Invalid store_id: %s", visit.StoreID))
			return
		}

		wg.Add(1)
		go func(visit models.Visit) {
			defer wg.Done()
			for _, imgURL := range visit.ImageURL {
				select {
				case <-job.Ctx.Done():
					return
				default:
					img, err := utils.DownloadImage(imgURL)
					if err != nil {
						job.MarkFailed(fmt.Sprintf("Failed to download image: %v", err))
						return
					}

					perimeter := utils.CalculatePerimeter(img)
					time.Sleep(time.Duration(rand.Float64()*0.3+0.1) * time.Second)

					job.AddResult(models.JobResult{StoreID: visit.StoreID, ImageURL: imgURL, Perimeter: perimeter})
				}
			}
		}(visit)
	}

	wg.Wait()
	job.MarkCompleted()
}

func GetJobStatus(jobID int) (map[string]interface{}, error) {
	jobMutex.Lock()
	defer jobMutex.Unlock()

	if job, ok := jobs[jobID]; ok {
		return job.GetStatus(), nil
	}

	return nil, fmt.Errorf(`{}`)
}
