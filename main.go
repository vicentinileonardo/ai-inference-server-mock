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
	CloudProvider   string       `json:"cloud_provider"`
	EligibleRegions []RegionInfo `json:"eligible_regions"`
}

type SchedulingInfo struct {
	SchedulingTime    string `json:"schedulingTime"`
	CloudProvider     string `json:"cloudProvider"`
	SchedulingRegion  string `json:"schedulingRegion"`
	SchedulingCountry string `json:"schedulingCountry"`
}

var (
	providers = []string{"AWS", "GCP", "Azure"}
	regions   = []string{"us-east-1", "us-west-2", "eu-west-1", "ap-southeast-1"}
)

func getRandomFutureTime() time.Time {
	now := time.Now()

	// Define the time window
	minOffset := 30 * time.Minute
	maxOffset := 48 * time.Hour // 2 days

	// Calculate random duration between minOffset and maxOffset
	randomDuration := time.Duration(rand.Int63n(int64(maxOffset-minOffset))) + minOffset

	return now.Add(randomDuration)
}

func getSchedulingInfo(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %s for %s\n", r.RemoteAddr, r.URL.Path)

	info := SchedulingInfo{
		SchedulingTime:   getRandomFutureTime().Format(time.RFC3339),
		CloudProvider:    providers[rand.Intn(len(providers))],
		SchedulingRegion: regions[rand.Intn(len(regions))],
	}

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
