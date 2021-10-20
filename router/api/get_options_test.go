package api

import (
	m "neocheckin_cache/database/models"
	em "neocheckin_cache/router/api/models/exported_models"
	"testing"
)

func convertedOptionIsEqual(o1 []em.Option, o2 []em.Option) bool {
	if len(o1) != len(o2) {
		return false
	}
	for i := 0; i < len(o1); i++ {
		if o1[i] != o2[i] {
			return false
		}
	}

	return true
}

func TestConvertOptionsToExportedModels(t *testing.T) {
	{ // arrays should be same length
		o := make([]m.Option, 24)
		cO := ConvertOptionsToExportedModels(o)
		if len(o) != len(cO) {
			t.Error("should be same length")
		}
	}
	{ // should convert to exported_models.Option
		o := []m.Option{
			{
				WrapperId: 1,
				Name:      "Option 1",
				Schedule: m.Schedule{
					From: m.ScheduleTime{
						Second: 0,
						Minute: 0,
						Hour:   0,
					},
					To: m.ScheduleTime{
						Second: 0,
						Minute: 0,
						Hour:   24,
					},
					Days: m.ScheduleDays{
						Monday:    true,
						Tuesday:   true,
						Wednesday: true,
						Thursday:  true,
						Friday:    true,
						Saturday:  true,
						Sunday:    true,
					},
				},
			},
			{
				WrapperId: 2,
				Name:      "Option 2",
				Schedule: m.Schedule{
					From: m.ScheduleTime{},
					To:   m.ScheduleTime{},
					Days: m.ScheduleDays{},
				},
			},
		}
		expected := []em.Option{
			{
				Id:        1,
				Name:      "Option 1",
				Available: em.OptionAvailable(em.AVAILABLE),
			},
			{
				Id:        2,
				Name:      "Option 2",
				Available: em.OptionAvailable(em.NOT_AVAILABLE),
			},
		}
		cO := ConvertOptionsToExportedModels(o)
		if len(o) != len(cO) {
			t.Error("should be same length")
		}
		if !convertedOptionIsEqual(cO, expected) {
			t.Error("should equal expected")
		}
	}
}
