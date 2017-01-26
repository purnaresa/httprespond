package httprespond

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data     interface{}   `json:"data,omitempty"`
	Errors   []interface{} `json:"errors,omitempty"`
	Meta     interface{}   `json:"meta,omitempty"`
	Links    Links         `json:"links,omitempty"`
	Included interface{}   `json:"included,omitempty"`
}

type Links struct {
	Self       string      `json:"self,omitempty"`
	Related    interface{} `json:"related,omitempty"`
	Pagination Pagination  `json:"pagination,omitempty"`
}

type Pagination struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
}

func Success(data interface{}, meta interface{}, included interface{}, pagination Pagination, status int, w http.ResponseWriter, r *http.Request) {
	response := Response{
		Data:     data,
		Meta:     meta,
		Included: included,
		Links: Links{
			Self:       r.URL.Path,
			Pagination: pagination,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
