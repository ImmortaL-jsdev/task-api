package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task-api/internal/database"
	"task-api/internal/models"
)

type Handlers struct {
	store *database.TaskStore
}

func NewHandler(store *database.TaskStore) *Handlers {
	return &Handlers{
		store: store,
	}
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, map[string]string{"error": message})
}
func parseTaskID(r *http.Request) (int, error) {
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/tasks/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		return 0, http.ErrNoLocation
	}
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (h *Handlers) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.store.GetAll()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Ошибка получения задач!")
		return
	}

	respondWithJSON(w, http.StatusOK, tasks)
}
func (h *Handlers) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := parseTaskID(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Некорректный id задачи!")
		return
	}
	task, err := h.store.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusBadRequest, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, task)
}

func (h *Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	var input models.CreateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Некорректно отправленные данные!")
		return
	}
	if strings.TrimSpace(input.Title) == "" {
		respondWithError(w, http.StatusBadRequest, "Заголовок задачи должен присутствовать!")
		return
	}
	task, err := h.store.Create(&input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, task)
}

func (h *Handlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := parseTaskID(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Некорректный id задачи!")
		return
	}

	var input models.UpdateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Некорректные данные!")
		return
	}
	if input.Title != nil && strings.TrimSpace(*input.Title) == "" {
		respondWithError(w, http.StatusBadRequest, "Заголовок обязателен!")
		return
	}

	task, err := h.store.Update(id, input)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, task)
}

func (h *Handlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := parseTaskID(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Некорректный id задачи!")
		return
	}

	if err := h.store.Delete(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "success"})
}
