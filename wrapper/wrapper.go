package wrapper

import (
	"bytes"
	"fmt"
	"neocheckin_cache/utils"
	em "neocheckin_cache/wrapper/models/exported_models"
	rqm "neocheckin_cache/wrapper/models/request_models"
	rsm "neocheckin_cache/wrapper/models/response_models"
	"net/http"
)

func convertTaskToRequest(t em.Task) rqm.AddTask {
	return rqm.AddTask{
		Name:         t.Name,
		TaskId:       t.TaskId,
		EmployeeRfid: t.EmployeeRfid,
		PostKey:      t.PostKey,
		SystemId:     t.SystemId,
		Timestamp:    t.Timestamp,
	}
}

func SendTask(t em.Task) (int, error) {

	enc, err := utils.JsonEncode(convertTaskToRequest(t))
	if err != nil {
		// TODO: add to task logs
		return http.StatusInternalServerError, err
	}

	resp, err := http.Post("http://localhost:4000", "application/json", bytes.NewBuffer(enc))
	if err != nil {
		if resp.StatusCode == http.StatusInternalServerError {
			// TODO: add to task queue
			return resp.StatusCode, err
		}
		// TODO: add to task logs
		return resp.StatusCode, err
	}

	status, err := handleSendTaskResponse(*resp)
	return status, err
}

func handleSendTaskResponse(r http.Response) (int, error) {
	defer r.Body.Close()

	if r.StatusCode == http.StatusBadRequest || r.StatusCode == http.StatusInternalServerError {
		if r.StatusCode == http.StatusInternalServerError {
			// TODO: add to task queue
			fmt.Println(":-(")
		}
		rErr := rsm.Error{}
		pErr := utils.ParseBody(utils.ParseableBody{Body: r.Body, Header: r.Header}, rErr)
		if pErr != nil {
			return http.StatusInternalServerError, pErr
		}
		return r.StatusCode, fmt.Errorf(rErr.Error)
		// TODO: add to task logs
	}

	return http.StatusOK, nil
}
