package api

import (
	"fmt"
	db "neocheckin_cache/database"
	dbm "neocheckin_cache/database/models"
	em "neocheckin_cache/router/api/models/exported_models"
	rsm "neocheckin_cache/router/api/models/response_models"
	"neocheckin_cache/utils"
	"net/http"
	"time"
)

func GetOptionsEndpoint(rw http.ResponseWriter, rq http.Request, db db.AbstractDatabase) {
	rw.Header().Add("Content-Type", "application/json")

	dbO, err := db.GetAllOptions()
	if err != nil {
		utils.WriteServerError(rw, err)
		return
	}
	converted := ConvertOptionsToExportedModels(dbO)

	encoded, err := utils.JsonEncode(rsm.Options{
		Options: converted,
	})

	if err == nil {
		fmt.Fprintf(rw, "%s", encoded)
		return
	} else {
		utils.WriteServerError(rw, err)
		return
	}
}

func optionIsDuringWeekday(d dbm.ScheduleDays, w time.Weekday) bool {
	// could not find an easier way to do this sadly.
	if w == time.Monday && d.Monday {
		return true
	} else if w == time.Tuesday && d.Tuesday {
		return true
	} else if w == time.Wednesday && d.Wednesday {
		return true
	} else if w == time.Thursday && d.Thursday {
		return true
	} else if w == time.Friday && d.Friday {
		return true
	} else if w == time.Saturday && d.Saturday {
		return true
	} else if w == time.Sunday && d.Sunday {
		return true
	}
	return false
}

func scheduleTimeToSeconds(s dbm.ScheduleTime) int {
	return s.Second + s.Minute*60 + s.Hour*60*60
}

func OptionIsAvailable(o dbm.Option) bool {
	t := time.Now()
	w := t.Weekday()
	if optionIsDuringWeekday(o.Schedule.Days, w) {
		h, m, s := t.Clock()
		tS := s + m*60 + h*60*60
		frS := scheduleTimeToSeconds(o.Schedule.From)
		toS := scheduleTimeToSeconds(o.Schedule.To)

		return frS <= tS && tS <= toS
	}
	return false
}

func ConvertOptionsToExportedModels(d []dbm.Option) []em.Option {
	r := make([]em.Option, len(d))
	for i := 0; i < len(d); i++ {
		a := OptionIsAvailable(d[i])
		oa := em.OptionAvailable(em.INVALID)
		if a {
			oa = em.OptionAvailable(em.AVAILABLE)
		} else {
			oa = em.OptionAvailable(em.NOT_AVAILABLE)
		}

		r[i] = em.Option{
			Id:        d[i].WrapperId,
			Name:      d[i].Name,
			Available: oa,
		}
	}

	return r
}
