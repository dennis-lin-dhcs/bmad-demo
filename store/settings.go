package store

import (
    "database/sql"

    "app-template/model"
)

func GetSettings(db *sql.DB, provider string) (model.LLMSettings, error) {
    row := db.QueryRow(`
        SELECT provider, endpoint, model, api_key, temperature, system_prompt
        FROM llm_settings
        WHERE provider = ?
    `, provider)

    var s model.LLMSettings
    err := row.Scan(&s.Provider, &s.Endpoint, &s.Model, &s.APIKey, &s.Temperature, &s.SystemPrompt)
    if err == sql.ErrNoRows {
        return model.DefaultSettings(provider), nil
    }
    if err != nil {
        return model.LLMSettings{}, err
    }
    return s, nil
}

func ListSettings(db *sql.DB) ([]model.LLMSettings, error) {
    providers := model.Providers()
    out := make([]model.LLMSettings, 0, len(providers))
    for _, provider := range providers {
        s, err := GetSettings(db, provider)
        if err != nil {
            return nil, err
        }
        out = append(out, s)
    }
    return out, nil
}

func UpsertSettings(db *sql.DB, s model.LLMSettings) error {
    _, err := db.Exec(`
        INSERT INTO llm_settings(provider, endpoint, model, api_key, temperature, system_prompt, updated_at)
        VALUES(?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
        ON CONFLICT(provider) DO UPDATE SET
            endpoint = excluded.endpoint,
            model = excluded.model,
            api_key = excluded.api_key,
            temperature = excluded.temperature,
            system_prompt = excluded.system_prompt,
            updated_at = CURRENT_TIMESTAMP
    `, s.Provider, s.Endpoint, s.Model, s.APIKey, s.Temperature, s.SystemPrompt)
    return err
}

func DeleteSettings(db *sql.DB, provider string) error {
    _, err := db.Exec(`DELETE FROM llm_settings WHERE provider = ?`, provider)
    return err
}
