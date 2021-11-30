package wrapper

import (
	"fmt"
	dbt "neocheckin_cache/database"
	dbm "neocheckin_cache/database/models"
	"neocheckin_cache/utils"
	em "neocheckin_cache/wrapper/models/exported_models"
	rqm "neocheckin_cache/wrapper/models/request_models"
	rsm "neocheckin_cache/wrapper/models/response_models"
	"net/http"
	"strings"
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
func sendQueuedTasks(db dbt.AbstractDatabase, pk string, l *utils.Logger) {
	t, err := db.GetAllTasks()
	if err != nil {
		l.FormatAndAppendToLogFile(fmt.Sprintf("error getting all tasks: %q", err.Error()))
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
		}, db, l, true)
		if status == http.StatusOK && err == nil {
			err := db.DeleteTaskWithDatabaseId(t[i].DatabaseId)
			if err != nil {
				l.FormatAndAppendToLogFile(fmt.Sprintf("error deleting queued task: %q", err.Error()))
				continue
			}
		}
	}
}

// FIXME jeg ved ikke om koden virker, også for lang
func SendTask(t em.Task, db dbt.AbstractDatabase, l *utils.Logger, queued bool) (int, error) {

	req, err := createRequestWithBody(t)
	if err != nil {
		l.FormatAndAppendToLogFile(fmt.Sprintf("unable to create request with body: %q", err.Error()))
		return http.StatusInternalServerError, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		addTaskToQueue(db, t, l)
		l.FormatAndAppendToLogFile(fmt.Sprintf("unable to send request: %q", err.Error()))
		return http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		message, parseError := getResponseErrorMessage(resp, queued, db, t, l)
		if parseError != nil {
			return http.StatusInternalServerError, parseError
		} else {
			l.FormatAndAppendToLogFile(fmt.Sprintf("got a status code >400 from wrapper with message %q", message))
		}

		if resp.StatusCode == http.StatusInternalServerError && !queued {
			addTaskToQueue(db, t, l)
		}

		return resp.StatusCode, nil
	}

	if resp.StatusCode == http.StatusOK && err == nil && !queued {
		sendQueuedTasks(db, t.PostKey, l)
	}

	return http.StatusOK, nil
}

func addTaskToQueue(db dbt.AbstractDatabase, t em.Task, l *utils.Logger) {
	err := db.InsertTask(dbm.Task{
		TaskId:       t.TaskId,
		Name:         t.Name,
		EmployeeRfid: t.EmployeeRfid,
		SystemId:     t.SystemId,
		Timestamp:    t.Timestamp,
	})
	if err != nil {
		l.FormatAndAppendToLogFile(fmt.Sprintf("error adding task to queue: %q", err.Error()))
	}
}

func getResponseErrorMessage(resp *http.Response, queued bool, db dbt.AbstractDatabase, t em.Task, l *utils.Logger) (string, error) {
	parsedError, errorDuringParsing := parseResponseError(resp)
	if errorDuringParsing != nil {
		l.FormatAndAppendToLogFile(fmt.Sprintf("error occurred during parsing response: %q", errorDuringParsing.Error()))
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

	req, err := utils.CreatePostRequest("/tasks/add", strings.NewReader(string(enc)))
	if err != nil {
		return nil, err
	}
	return req, nil
}
