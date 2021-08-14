package main

import (
	"encoding/json"
	"net/http"
)

type Statistic struct {
	Name     string `json:"name"`
	Manufact string `json:"manufactures"`
	Id       string `json:"id"`
}

type StatHandlers struct {
	store map[string]Statistic
}

func (h *StatHandlers) api(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
}

func (h *StatHandlers) get(w http.ResponseWriter, r *http.Request) {
	statistics := make([]Statistic, len(h.store))
	i := 0
	for _, statistic := range h.store {
		statistics[i] = statistic
		i++

	}
	JsonBytes, err := json.Marshal(statistics)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(JsonBytes)
}

func StatHanlder(w http.ResponseWriter, r *http.Request) {

}
func newStatHandlers() *StatHandlers {
	return &StatHandlers{
		store: map[string]Statistic{
			"id1": Statistic{
				Id:       "120002",
				Manufact: "standart",
				Name:     "bitcoin",
			},
		},
	}
}

func main() {
	StatHandlers := newStatHandlers()
	http.HandleFunc("/api/show_stat/", StatHandlers.api)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
