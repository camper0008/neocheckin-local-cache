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
		return http.StatusInternalServerError, err
	}

	resp, err := http.Post("http://localhost:4000", "application/json", bytes.NewBuffer(enc))
	if err != nil {
		return resp.StatusCode, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError {
		err := rsm.Error{}
		utils.ParseBody(utils.ParseableBody{Body: resp.Body, Header: resp.Header}, err)
		return resp.StatusCode, fmt.Errorf(err.Error)
	}

	return http.StatusOK, nil
}
