package main

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {

	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	j, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(j)
	if err != nil {
		return err
	}
	return nil
}
