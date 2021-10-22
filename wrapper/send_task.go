package wrapper

import (
	"bytes"
	"fmt"
	db "neocheckin_cache/database"
	dbm "neocheckin_cache/database/models"
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

func sendQueuedTasks(db db.AbstractDatabase, pk string) {
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

func SendTask(t em.Task, db db.AbstractDatabase, queued bool) (int, error) {

	enc, err := utils.JsonEncode(convertTaskToRequest(t))
	if err != nil {
		// TODO: add to task logs
		return http.StatusInternalServerError, err
	}

	resp, err := http.Post("http://localhost:7000", "application/json", bytes.NewBuffer(enc))
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
		pErr := utils.ParseBody(utils.ParseableBody{Body: resp.Body, Header: resp.Header}, rErr)
		if pErr != nil {
			return http.StatusInternalServerError, pErr
		}
		// TODO: add to task logs
		return resp.StatusCode, fmt.Errorf(rErr.Error)
	}

	if resp.StatusCode == http.StatusOK && err == nil && !queued {
		sendQueuedTasks(db, t.PostKey)
	}

	return resp.StatusCode, nil
}
