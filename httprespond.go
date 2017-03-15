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
	Related    string      `json:"related,omitempty"`
	First      string      `json:"first,omitempty"`
	Last       string      `json:"last,omitempty"`
	Prev       string      `json:"prev,omitempty"`
	Next       string      `json:"next,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	HasNextPage bool `json:"has_next_page"`
	HasPrevPage bool `json:"has_prev_page"`
}

func Success(data interface{}, meta interface{}, included interface{}, pagination *Pagination, status int, w http.ResponseWriter, r *http.Request) {
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
