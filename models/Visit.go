package models


// Visit represents information about a store visit.
type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURL  []string `json:"image_url"`    // List of image URLs captured during the visit.
	VisitTime string   `json:"visit_time"`
}

type JobRequest struct {
	Count  int     `json:"count"`
	Visits []Visit `json:"visits"`  
}
