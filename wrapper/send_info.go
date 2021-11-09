package wrapper

import (
	"bytes"
	"fmt"
	dbt "neocheckin_cache/database"
	dbm "neocheckin_cache/database/models"
	"neocheckin_cache/utils"
	em "neocheckin_cache/wrapper/models/exported_models"
	rqm "neocheckin_cache/wrapper/models/request_models"
	rsm "neocheckin_cache/wrapper/models/response_models"
	"net/http"
)

// FIXME jeg ved ikke om koden virker
// men tak for endelig at extrace conversion
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

// FIXME jeg ved ikke om koden virker
func sendQueuedTasks(db dbt.AbstractDatabase, pk string) {
	t, err := db.GetAllTasks()
	if err != nil {
		// TODO: add to task logs
		return
	}
	for i := 0; i < len(t); i++ {
		status, err := SendTask(em.Task{
			TaskId:       t[i].TaskId,
			Name:         t[i].Name,
			EmployeeRfid: t[i].EmployeeRfid,
			PostKey:      pk,
			SystemId:     t[i].SystemId,
			Timestamp:    t[i].Timestamp,
		}, db, true)
		if status == http.StatusOK && err == nil {
			err := db.DeleteTaskWithDatabaseId(t[i].DatabaseId)
			if err != nil {
				//TODO: add to task logs
				continue
			}
		}
	}
}

// FIXME jeg ved ikke om koden virker, også for lang
func SendTask(t em.Task, db dbt.AbstractDatabase, queued bool) (int, error) {

	// TODO: add to task logs
	req, err := createRequestWithBody(t)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		addTaskToQueue(db, t)
		// TODO: add to task logs
		return http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		message, parseError := getResponseErrorMessage(resp, queued, db, t)
		if parseError != nil {
			return http.StatusInternalServerError, parseError
		} else {
			// TODO: add error message to task logs
			fmt.Println(message)
		}

		if resp.StatusCode == http.StatusInternalServerError && !queued {

			addTaskToQueue(db, t)

		}

		return resp.StatusCode, nil
	}

	if resp.StatusCode == http.StatusOK && err == nil && !queued {
		sendQueuedTasks(db, t.PostKey)
	}

	return http.StatusOK, nil
}

func addTaskToQueue(db dbt.AbstractDatabase, t em.Task) {
	// TODO: add db errors to logs
	db.AddTask(dbm.Task{
		TaskId:       t.TaskId,
		Name:         t.Name,
		EmployeeRfid: t.EmployeeRfid,
		SystemId:     t.SystemId,
		Timestamp:    t.Timestamp,
	})
}

func getResponseErrorMessage(resp *http.Response, queued bool, db dbt.AbstractDatabase, t em.Task) (string, error) {
	// TODO: add to task logs
	parsedError, errorDuringParsing := parseResponseError(resp)
	if errorDuringParsing != nil {
		return "", errorDuringParsing
	}

	return parsedError.Error, nil
}

func parseResponseError(resp *http.Response) (rsm.Error, error) {
	rErr := rsm.Error{}
	pErr := utils.ParseBody(utils.ParseableBody{Body: resp.Body, Header: resp.Header}, &rErr)
	if pErr != nil {
		return rsm.Error{}, pErr
	}
	return rErr, nil
}

func createRequestWithBody(t em.Task) (*http.Request, error) {
	enc, err := utils.JsonEncode(convertTaskToRequest(t))
	if err != nil {
		return nil, err
	}

	req, err := utils.CreatePostRequest("/tasks/add", t.PostKey, bytes.NewBuffer(enc))
	if err != nil {
		return nil, err
	}
	return req, nil
}
