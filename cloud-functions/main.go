package function

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("EntryPoint", EntryPoint)
}

func EntryPoint(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	resp, err := process(idStr)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

type responseSchema struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	Age  int    `json:"age"`
}

func process(idStr string) ([]byte, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s as int; %w", idStr, err)
	}
	resp := responseSchema{ID: id}
	switch id % 3 {
	case 0:
		resp.Name = "Alice"
		resp.City = "Sapporo"
		resp.Age = 10
	case 1:
		resp.Name = "Bravo"
		resp.City = "Tokyo"
		resp.Age = 20
	default:
		resp.Name = "Charlie"
		resp.City = "Naha"
		resp.Age = 30
	}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON; %w", err)
	}
	return respBytes, nil
}