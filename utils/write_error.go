package utils

import (
	"fmt"
	rsm "neocheckin_cache/router/api/models/response_models"
	"net/http"
)

// FIXME jeg ved ikke om koden virker
func WriteError(rw http.ResponseWriter, err error) {
	encoded, err := JsonEncode(rsm.Error{
		Error: err.Error(),
	})

	if err == nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, "%s", encoded)
		return
	} else {
		WriteServerError(rw, err)
		return
	}
}

// FIXME jeg ved ikke om koden virker
func WriteServerError(rw http.ResponseWriter, err error) {
	encoded, err := JsonEncode(rsm.Error{
		Error: err.Error(),
	})

	if err == nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "%s", encoded)
		return
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
