package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Match struct {
	ID              int    `json:"id"`
	EquipoLocal     string `json:"equipoLocal"`
	EquipoVisitante string `json:"equipoVisitante"`
	Fecha           string `json:"fecha"`
	Goals           int    `json:"goals,omitempty"`
	YellowCards     int    `json:"yellowCards,omitempty"`
	RedCards        int    `json:"redCards,omitempty"`
	ExtraTime       bool   `json:"extraTime,omitempty"`
}

var matches []Match

func getAllMatches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}

func getMatchByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	for _, match := range matches {
		if match.ID == id {
			json.NewEncoder(w).Encode(match)
			return
		}
	}
	http.Error(w, "Partido no encontrado", http.StatusNotFound)
}

func createMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMatch Match
	if err := json.NewDecoder(r.Body).Decode(&newMatch); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	newMatch.ID = rand.Intn(1000000)
	matches = append(matches, newMatch)
	json.NewEncoder(w).Encode(newMatch)
}

func updateMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var updatedData Match
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	for i, match := range matches {
		if match.ID == id {
			match.EquipoLocal = updatedData.EquipoLocal
			match.EquipoVisitante = updatedData.EquipoVisitante
			match.Fecha = updatedData.Fecha

			if updatedData.Goals != 0 {
				match.Goals = updatedData.Goals
			}
			if updatedData.YellowCards != 0 {
				match.YellowCards = updatedData.YellowCards
			}
			if updatedData.RedCards != 0 {
				match.RedCards = updatedData.RedCards
			}
			match.ExtraTime = updatedData.ExtraTime

			matches[i] = match
			json.NewEncoder(w).Encode(match)
			return
		}
	}
	http.Error(w, "Partido no encontrado", http.StatusNotFound)
}

func deleteMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	for i, match := range matches {
		if match.ID == id {
			matches = append(matches[:i], matches[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Partido no encontrado", http.StatusNotFound)
}

func patchRegisterGoal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	for i, match := range matches {
		if match.ID == id {
			matches[i].Goals++
			json.NewEncoder(w).Encode(matches[i])
			return
		}
	}
	http.Error(w, "Partido no encontrado", http.StatusNotFound)
}

func patchRegisterYellowCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	for i, match := range matches {
		if match.ID == id {
			matches[i].YellowCards++
			json.NewEncoder(w).Encode(matches[i])
			return
		}
	}
	http.Error(w, "Partido no encontrado", http.StatusNotFound)
}

func patchRegisterRedCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	for i, match := range matches {
		if match.ID == id {
			matches[i].RedCards++
			json.NewEncoder(w).Encode(matches[i])
			return
		}
	}
	http.Error(w, "Partido no encontrado", http.StatusNotFound)
}

func patchSetExtraTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	for i, match := range matches {
		if match.ID == id {
			matches[i].ExtraTime = true
			json.NewEncoder(w).Encode(matches[i])
			return
		}
	}
	http.Error(w, "Partido no encontrado", http.StatusNotFound)
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {

	matches = append(matches, Match{
		ID:              1,
		EquipoLocal:     "Real Madrid",
		EquipoVisitante: "Barcelona",
		Fecha:           "03/27/2025",
	})
	matches = append(matches, Match{
		ID:              2,
		EquipoLocal:     "Atlético Madrid",
		EquipoVisitante: "Sevilla",
		Fecha:           "03/28/2025",
	})

	router := mux.NewRouter()
	router.Use(enableCORS)

	router.HandleFunc("/api/matches", getAllMatches).Methods("GET")
	router.HandleFunc("/api/matches/{id}", getMatchByID).Methods("GET")
	router.HandleFunc("/api/matches", createMatch).Methods("POST")
	router.HandleFunc("/api/matches/{id}", updateMatch).Methods("PUT")
	router.HandleFunc("/api/matches/{id}", deleteMatch).Methods("DELETE")

	router.HandleFunc("/api/matches/{id}/goal", patchRegisterGoal).Methods("PATCH")
	router.HandleFunc("/api/matches/{id}/yellowcard", patchRegisterYellowCard).Methods("PATCH")
	router.HandleFunc("/api/matches/{id}/redcard", patchRegisterRedCard).Methods("PATCH")
	router.HandleFunc("/api/matches/{id}/extratime", patchSetExtraTime).Methods("PATCH")

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
