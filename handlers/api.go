package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strings"

    "app-template/model"
    "app-template/store"
)

func RegisterAPI(mux *http.ServeMux, db *sql.DB) {
    mux.HandleFunc("/api/settings", withCORS(func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            settings, err := store.ListSettings(db)
            if err != nil {
                writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
                return
            }
            writeJSON(w, http.StatusOK, settings)
        case http.MethodOptions:
            w.WriteHeader(http.StatusNoContent)
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
        }
    }))

    mux.HandleFunc("/api/settings/", withCORS(func(w http.ResponseWriter, r *http.Request) {
        provider := strings.TrimPrefix(r.URL.Path, "/api/settings/")
        provider = strings.Trim(provider, "/")
        if provider == "" {
            writeJSON(w, http.StatusBadRequest, map[string]string{"error": "missing provider"})
            return
        }

        switch r.Method {
        case http.MethodGet:
            settings, err := store.GetSettings(db, provider)
            if err != nil {
                writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
                return
            }
            writeJSON(w, http.StatusOK, settings)
        case http.MethodPut, http.MethodPost:
            var s model.LLMSettings
            if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
                writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
                return
            }
            s.Provider = provider
            if s.Endpoint == "" || s.Model == "" {
                writeJSON(w, http.StatusBadRequest, map[string]string{"error": "endpoint and model are required"})
                return
            }
            if err := store.UpsertSettings(db, s); err != nil {
                writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
                return
            }
            writeJSON(w, http.StatusOK, s)
        case http.MethodDelete:
            if err := store.DeleteSettings(db, provider); err != nil {
                writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
                return
            }
            writeJSON(w, http.StatusOK, map[string]any{"provider": provider, "deleted": true, "defaults": model.DefaultSettings(provider)})
        case http.MethodOptions:
            w.WriteHeader(http.StatusNoContent)
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
        }
    }))
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        next(w, r)
    }
}

func writeJSON(w http.ResponseWriter, status int, value any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    _ = json.NewEncoder(w).Encode(value)
}
