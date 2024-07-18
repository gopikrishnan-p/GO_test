// // // package main

// // // import "fmt"

// // // func main() {
// // // 	fmt.Println("Hello, World!")
// // // }
// // package main

// // import (
// //     "fmt"
// //     "io/ioutil"
// //     "net/http"
// // )

// // func main() {
// //     // Create a new HTTP client
// //     client := &http.Client{}

// //     // Create a new HTTP request
// //     req, err := http.NewRequest("GET", "https://example.com", nil)
// //     if err != nil {
// //         fmt.Println(err)
// //         return
// //     }

// //     // Add headers if needed
// //     req.Header.Add("Accept", "application/json")

// //     // Send the request
// //     resp, err := client.Do(req)
// //     if err != nil {
// //         fmt.Println(err)
// //         return
// //     }
// //     defer resp.Body.Close()

// //     // Read the response body
// //     body, err := ioutil.ReadAll(resp.Body)
// //     if err != nil {
// //         fmt.Println(err)
// //         return
// //     }

// //     // Print the response body
// //     fmt.Println(string(body))
// // }
// package main

// import (
//     "encoding/json"
//     "net/http"
//     "time"
// )

// type TimeRequest struct {
//     Time     string `json:"time"`
//     Duration string `json:"duration"`
// }

// type TimeResponse struct {
//     Result time.Time `json:"result"`
// }

// func timeHandler(w http.ResponseWriter, r *http.Request) {
//     var req TimeRequest
//     err := json.NewDecoder(r.Body).Decode(&req)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }

//     t, err := time.Parse(time.RFC3339, req.Time)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }

//     d, err := time.ParseDuration(req.Duration)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }

//     result := t.Add(d)

//     resp := TimeResponse{Result: result}
//     json.NewEncoder(w).Encode(resp)
// }

// func main() {
//     http.HandleFunc("/time", timeHandler)
//     http.ListenAndServe(":8080", nil)
// }
// //
package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	// Define the struct for the response
	Message string `json:"message"`
}

// handleRequest handles an HTTP request and sends a JSON response with the current time.
//
// Parameters:
// - w: http.ResponseWriter - the response writer used to write the response.
// - r: *http.Request - the request object containing information about the HTTP request.
//
// Return type: None.
func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	currentTime := time.Now()
	// Format the current time
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	println(formattedTime)
	// Prepare the response
	response := Response{
		Message: "Hello! " + formattedTime,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Start the server
	http.HandleFunc("/", handleRequest)
	// Listen on port 8080
	http.ListenAndServe(":8080", nil)
}
