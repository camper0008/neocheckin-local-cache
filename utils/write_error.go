package utils

import (
	"fmt"
	rsm "neocheckin_cache/router/api/models/response_models"
	"net/http"
)

func WriteError(rw http.ResponseWriter, err error) {
	encoded, err := JsonEncode(rsm.Error{
		Error: err.Error(),
	})

	if err == nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, "%s", encoded)
		return
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
