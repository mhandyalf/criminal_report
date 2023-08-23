package handlers

import (
	"criminal_report/entity"
	"criminal_report/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CriminalHandler struct {
	repo *repository.CriminalRepository
}

func NewCriminalHandler(repo *repository.CriminalRepository) *CriminalHandler {
	return &CriminalHandler{repo}
}

func (handler *CriminalHandler) GetCriminals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	criminals, err := handler.repo.GetAllCriminals()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(criminals)
}

func (handler *CriminalHandler) GetCriminal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	criminal, err := handler.repo.GetCriminalByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(criminal)
}

func (handler *CriminalHandler) CreateCriminal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var criminal entity.CriminalReport
	err := json.NewDecoder(r.Body).Decode(&criminal)
	if err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	id, err := handler.repo.CreateCriminal(criminal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"Crime Created Succesfully! id": id})
}

func (handler *CriminalHandler) UpdateCriminal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var criminal entity.CriminalReport
	err = json.NewDecoder(r.Body).Decode(&criminal)
	if err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	err = handler.repo.UpdateCriminal(id, criminal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Criminal updated successfully"})
}

func (handler *CriminalHandler) DeleteCriminal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = handler.repo.DeleteCriminal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Criminal deleted successfully"})
}
