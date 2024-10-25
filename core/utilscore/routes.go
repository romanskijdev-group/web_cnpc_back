package utilscore

import (
	"encoding/json"
	"net/http"
)

// HandleRouterBodyRequest is a utility function that reads a request body into a JSON,
// performs an update function on the body and writes the result back as a JSON response.
//
// If the request body is invalid, it returns a 400 Bad Request response.
// If the update function fails, it returns a 500 Internal Server Error response.
// If the response cannot be encoded as JSON, it returns a 500 Internal Server Error response.
func HandleRouterBodyRequest(w http.ResponseWriter, r *http.Request, requestBody interface{}, updateFunc func(interface{}) (interface{}, error)) {
	err := json.NewDecoder(r.Body).Decode(requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res, err := updateFunc(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// HandleRouterGetResponse is a utility function that writes a JSON response.
//
// If an error is passed, it returns a 500 Internal Server Error response.
// If the response cannot be encoded as JSON, it returns a 500 Internal Server Error response.
func HandleRouterGetResponse(w http.ResponseWriter, r *http.Request, res interface{}, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
