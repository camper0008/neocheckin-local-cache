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

	enc, err := utils.JsonEncode(convertTaskToRequest(t))
	if err != nil {
		// TODO: add to task logs
		return http.StatusInternalServerError, err
	}

	req, err := utils.CreatePostRequest("/tasks/add", t.PostKey, bytes.NewBuffer(enc))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		db.AddTask(dbm.Task{
			TaskId:       t.TaskId,
			Name:         t.Name,
			EmployeeRfid: t.EmployeeRfid,
			SystemId:     t.SystemId,
			Timestamp:    t.Timestamp,
		})
		// TODO: add to task logs
		return http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError {
		if resp.StatusCode == http.StatusInternalServerError && !queued {
			// TODO: add db errors to logs
			db.AddTask(dbm.Task{
				TaskId:       t.TaskId,
				Name:         t.Name,
				EmployeeRfid: t.EmployeeRfid,
				SystemId:     t.SystemId,
				Timestamp:    t.Timestamp,
			})
		}
		rErr := rsm.Error{}
		pErr := utils.ParseBody(utils.ParseableBody{Body: resp.Body, Header: resp.Header}, &rErr)
		if pErr != nil {
			return http.StatusInternalServerError, pErr
		}
		// TODO: add to task logs
		return resp.StatusCode, fmt.Errorf(rErr.Error)
	}

	if resp.StatusCode >= 400 {
		return resp.StatusCode, fmt.Errorf("recieved unexpected status code")
	}

	if resp.StatusCode == http.StatusOK && err == nil && !queued {
		sendQueuedTasks(db, t.PostKey)
	}

	return http.StatusOK, nil
}
