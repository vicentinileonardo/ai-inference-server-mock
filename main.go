package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type RegionInfo struct {
	Name             string `json:"name"`
	ISOCountryCodeA2 string `json:"iso_country_code_a2"`
	PhysicalLocation string `json:"physical_location"`
}

type SchedulingRequest struct {
	EligibleRegions []string `json:"eligible_regions"`
	Deadline        string   `json:"deadline"`
	Duration        string   `json:"duration"`
}

type SchedulingInfo struct {
	SchedulingTime    string `json:"schedulingTime"`
	CloudProvider     string `json:"cloudProvider"`
	SchedulingRegion  string `json:"schedulingRegion"`
	SchedulingCountry string `json:"schedulingCountry"`
}

var (
	providers = []string{"azure"}
)

func getRandomTimeBetween(start, end time.Time) time.Time {
	duration := end.Sub(start)
	randomDuration := time.Duration(rand.Int63n(int64(duration)))
	return start.Add(randomDuration)
}

func getSchedulingInfo(w http.ResponseWriter, r *http.Request) {
	// Ensure only POST requests are accepted
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming request
	var request SchedulingRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if len(request.EligibleRegions) == 0 {
		http.Error(w, "No eligible regions provided", http.StatusBadRequest)
		return
	}

	// Parse deadline
	deadline, err := time.Parse(time.RFC3339, request.Deadline)
	if err != nil {
		http.Error(w, "Invalid deadline format", http.StatusBadRequest)
		return
	}

	// Parse duration
	duration, err := time.ParseDuration(request.Duration)
	if err != nil {
		http.Error(w, "Invalid duration format", http.StatusBadRequest)
		return
	}

	// Calculate start time
	startTime := time.Now().UTC()
	endTime := deadline.Add(-duration)

	// Validate time range
	if endTime.Before(startTime) {
		http.Error(w, "Invalid time range", http.StatusBadRequest)
		return
	}

	// Select random scheduling time within the valid range
	schedulingTime := getRandomTimeBetween(startTime, endTime)

	// Select a random region from eligible regions
	schedulingRegion := request.EligibleRegions[rand.Intn(len(request.EligibleRegions))]

	// Prepare response with UTC time in ISO 8601 format
	info := SchedulingInfo{
		SchedulingTime:   schedulingTime.UTC().Format(time.RFC3339),
		CloudProvider:    providers[0],
		SchedulingRegion: schedulingRegion,
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(info); err != nil {
		log.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Successfully sent response: %+v\n", info)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func main() {
	rand.NewSource(time.Now().UnixNano())

	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/scheduling", getSchedulingInfo)

	log.Printf("Server starting on :8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
