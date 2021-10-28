package api

import (
	"fmt"
	c "neocheckin_cache/config"
	dbt "neocheckin_cache/database"
	dbm "neocheckin_cache/database/models"
	em "neocheckin_cache/router/api/models/exported_models"
	rsm "neocheckin_cache/router/api/models/response_models"
	"neocheckin_cache/utils"
	"net/http"
	"time"
)

func GetOptionsEndpoint(rw http.ResponseWriter, rq http.Request, db dbt.AbstractDatabase) {
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
	switch w {
	case time.Monday:
		return d.Monday
	case time.Tuesday:
		return d.Tuesday
	case time.Wednesday:
		return d.Wednesday
	case time.Thursday:
		return d.Thursday
	case time.Friday:
		return d.Friday
	case time.Saturday:
		return d.Saturday
	case time.Sunday:
		return d.Sunday
	default:
		return false
	}
}

func scheduleTimeToSeconds(s dbm.ScheduleTime) int {
	return s.Second + s.Minute*60 + s.Hour*60*60
}

func OptionIsAvailable(o dbm.Option) bool {
	if len(o.Locations) != 0 {
		conf := c.Read()
		loc := conf["LOCATION"]
		for i := 0; i < len(o.Locations); i++ {
			if loc == o.Locations[i] {
				return true
			}
		}
		return false
	}

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

		oa := em.OptionAvailable(em.NOT_AVAILABLE)
		if a && d[i].Priority {
			oa = em.OptionAvailable(em.PRIORITY)
		} else if a {
			oa = em.OptionAvailable(em.AVAILABLE)
		}

		r[i] = em.Option{
			Id:        d[i].WrapperId,
			Name:      d[i].Name,
			Available: oa,
		}
	}

	return r
}
