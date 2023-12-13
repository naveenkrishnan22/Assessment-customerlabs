package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	Ev     string `json:"ev,omitempty"`
	Et     string `json:"et,omitempty"`
	ID     string `json:"id,omitempty"`
	UID    string `json:"uid,omitempty"`
	MID    string `json:"mid,omitempty"`
	T      string `json:"t,omitempty"`
	P      string `json:"p,omitempty"`
	L      string `json:"l,omitempty"`
	SC     string `json:"sc,omitempty"`
	ATRK1  string `json:"atrk1,omitempty"`
	ATRV1  string `json:"atrv1,omitempty"`
	ATRT1  string `json:"atrt1,omitempty"`
	ATRK2  string `json:"atrk2,omitempty"`
	ATRV2  string `json:"atrv2,omitempty"`
	ATRT2  string `json:"atrt2,omitempty"`
	ATRK3  string `json:"atrk3,omitempty"`
	ATRV3  string `json:"atrv3,omitempty"`
	ATRT3  string `json:"atrt3,omitempty"`
	ATRK4  string `json:"atrk4,omitempty"`
	ATRV4  string `json:"atrv4,omitempty"`
	ATRT4  string `json:"atrt4,omitempty"`
	UATRK1 string `json:"uatrk1,omitempty"`
	UATRV1 string `json:"uatrv1,omitempty"`
	UATRT1 string `json:"uatrt1,omitempty"`
	UATRK2 string `json:"uatrk2,omitempty"`
	UATRV2 string `json:"uatrv2,omitempty"`
	UATRT2 string `json:"uatrt2,omitempty"`
	UATRK3 string `json:"uatrk3,omitempty"`
	UATRV3 string `json:"uatrv3,omitempty"`
	UATRT3 string `json:"uatrt3,omitempty"`
	UATRK4 string `json:"uatrk4,omitempty"`
	UATRV4 string `json:"uatrv4,omitempty"`
	UATRT4 string `json:"uatrt4,omitempty"`
	UATRK5 string `json:"uatrk5,omitempty"`
	UATRV5 string `json:"uatrv5,omitempty"`
	UATRT5 string `json:"uatrt5,omitempty"`
	UATRK6 string `json:"uatrk6,omitempty"`
	UATRV6 string `json:"uatrv6,omitempty"`
	UATRT6 string `json:"uatrt6,omitempty"`
}

type Attribute struct {
	Value string `json:"value,omitempty"`
	Type  string `json:"type,omitempty"`
}

type ProcessedData struct {
	Event           string               `json:"event,omitempty"`
	EventType       string               `json:"event_type,omitempty"`
	AppID           string               `json:"app_id,omitempty"`
	UserID          string               `json:"user_id,omitempty"`
	MessageID       string               `json:"message_id,omitempty"`
	PageTitle       string               `json:"page_title,omitempty"`
	PageURL         string               `json:"page_url,omitempty"`
	BrowserLanguage string               `json:"browser_language,omitempty"`
	ScreenSize      string               `json:"screen_size,omitempty"`
	Attributes      map[string]Attribute `json:"attributes,omitempty"`
	Traits          map[string]Attribute `json:"traits,omitempty"`
}

func worker(requests <-chan Request) {
	for req := range requests {
		// Process the request and convert it to the desired format
		processedData := ProcessedData{
			Event:           req.Ev,
			EventType:       req.Et,
			AppID:           req.ID,
			UserID:          req.UID,
			MessageID:       req.MID,
			PageTitle:       req.T,
			PageURL:         req.P,
			BrowserLanguage: req.L,
			ScreenSize:      req.SC,
			Attributes:      make(map[string]Attribute),
			Traits:          make(map[string]Attribute),
		}

		// Map attributes
		processedData.Attributes["form_varient"] = Attribute{
			Value: req.ATRV1,
			Type:  req.ATRT1,
		}
		processedData.Attributes["ref"] = Attribute{
			Value: req.ATRV2,
			Type:  req.ATRT2,
		}
		processedData.Attributes["button_text"] = Attribute{
			Value: req.ATRV1,
			Type:  req.ATRT1,
		}
		processedData.Attributes["color_variation"] = Attribute{
			Value: req.ATRV2,
			Type:  req.ATRT2,
		}
		processedData.Attributes["page_path"] = Attribute{
			Value: req.ATRV3,
			Type:  req.ATRT3,
		}
		processedData.Attributes["source"] = Attribute{
			Value: req.ATRV4,
			Type:  req.ATRT4,
		}

		// Map traits
		processedData.Traits["name"] = Attribute{
			Value: req.UATRV1,
			Type:  req.UATRT1,
		}
		processedData.Traits["email"] = Attribute{
			Value: req.UATRV2,
			Type:  req.UATRT2,
		}
		processedData.Traits["age"] = Attribute{
			Value: req.UATRV3,
			Type:  req.UATRT3,
		}
		processedData.Traits["user_score"] = Attribute{
			Value: req.UATRV1,
			Type:  req.UATRT1,
		}
		processedData.Traits["gender"] = Attribute{
			Value: req.UATRV2,
			Type:  req.UATRT2,
		}
		processedData.Traits["tracking_code"] = Attribute{
			Value: req.UATRV3,
			Type:  req.UATRT3,
		}
		processedData.Traits["phone"] = Attribute{
			Value: req.UATRV4,
			Type:  req.UATRT4,
		}
		processedData.Traits["coupon_clicked"] = Attribute{
			Value: req.UATRV5,
			Type:  req.UATRT5,
		}
		processedData.Traits["opt_out"] = Attribute{
			Value: req.UATRV6,
			Type:  req.UATRT6,
		}

		// Convert processed data to JSON
		jsonData, err := json.Marshal(processedData)
		if err != nil {
			fmt.Println("Failed to marshal processed data:", err)
			continue
		}

		// Send the processed data to the webhook URL
		webhookURL := "https://webhook.site/5fd31bee-d2cb-4e31-a807-cf756a3545f7"
		resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Failed to send processed data to the webhook:", err)
			continue
		}

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Failed to read webhook response body:", err)
		}

		// Close the response body
		resp.Body.Close()

		// Print the webhook response
		fmt.Println("Webhook response:", string(body))
	}
}

func main() {
	// Create a channel to receive requests
	requests := make(chan Request)

	// Start the worker goroutine
	go worker(requests)

	// Handle requests
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		var request Request
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Send the request to the worker
		requests <- request

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Request received successfully"))

		// Display success message in the terminal
		fmt.Println("Request received and processed successfully")
	})

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
