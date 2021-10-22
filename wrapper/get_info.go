// TODO: log all errors
package wrapper

import (
	"neocheckin_cache/utils"
	rm "neocheckin_cache/wrapper/models/response_models"
	"net/http"
)

func GetTaskTypes() (rm.GetTaskTypes, error) {
	// TODO: use config url
	resp, err := http.Get("http://localhost:7000/api/tasks/types")
	if err != nil {
		return rm.GetTaskTypes{}, err
	}

	parsed := rm.GetTaskTypes{}
	err = utils.ParseBody(utils.ParseableBody{
		Body:   resp.Body,
		Header: resp.Header,
	}, &parsed)
	if err != nil {
		return rm.GetTaskTypes{}, err
	}

	return parsed, nil
}

func GetEmployees() (rm.GetEmployees, error) {
	// TODO: use config url
	resp, err := http.Get("http://localhost:7000/api/employees/all")
	if err != nil {
		return rm.GetEmployees{}, err
	}

	parsed := rm.GetEmployees{}
	err = utils.ParseBody(utils.ParseableBody{
		Body:   resp.Body,
		Header: resp.Header,
	}, &parsed)
	if err != nil {
		return rm.GetEmployees{}, err
	}

	return parsed, nil
}
