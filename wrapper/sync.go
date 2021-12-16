package wrapper

import (
	"fmt"
	"neocheckin_cache/database"
	"neocheckin_cache/utils"
	"time"
)

func syncTaskTypes(db database.AbstractDatabase, l *utils.Logger) {
	t, err := GetTaskTypes(l)
	if err != nil {
		fmt.Printf("Error synchronizing task types: %v\n", err)
		l.FormatAndAppendToLogFile(fmt.Sprintf("Error synchronizing task types: %v", err))
	} else {
		err = updateDbFromTaskTypesResponse(db, t)
		if err != nil {
			fmt.Printf("Error synchronizing task types: %v\n", err)
			l.FormatAndAppendToLogFile(fmt.Sprintf("Error synchronizing task types: %v", err))
		}
	}
}

func syncEmployeesWithoutPhoto(db database.AbstractDatabase, l *utils.Logger) {
	e, err := getEmployeesWithoutPhoto(l)
	if err != nil {
		fmt.Printf("Error synchronizing employees: %v\n", err)
		l.FormatAndAppendToLogFile(fmt.Sprintf("Error synchronizing employees: %v", err))
	} else {
		errors := updateDbFromSyncEmployeesResponse(db, e)
		for i := 0; i < len(errors); i++ {
			if errors[i] != nil {
				fmt.Printf("Error synchronizing employees without photo: %v\n", errors[i])
				l.FormatAndAppendToLogFile(fmt.Sprintf("Error synchronizing employees without photo: %v", errors[i]))
			}
		}
	}
}
func syncEmployeesWithPhoto(db database.AbstractDatabase, l *utils.Logger) {
	e, err := getEmployeesWithPhoto(l)
	if err != nil {
		fmt.Printf("Error synchronizing employees: %v\n", err)
		l.FormatAndAppendToLogFile(fmt.Sprintf("Error synchronizing employees: %v", err))
	} else {
		err = updateDbFromAllEmployeesResponse(db, e)
		if err != nil {
			fmt.Printf("Error synchronizing employees: %v\n", err)
			l.FormatAndAppendToLogFile(fmt.Sprintf("Error synchronizing employees: %v", err))
		}
	}
}

func InitialSync(db database.AbstractDatabase, l *utils.Logger) {
	fmt.Println("Attempting to synchronize...")
	syncTaskTypes(db, l)
	syncEmployeesWithPhoto(db, l)
	fmt.Println("Done")
}

func ScheduleSync(db database.AbstractDatabase, l *utils.Logger) {
	for range time.Tick(time.Minute * 1) {
		fmt.Println("Attempting to synchronize...")
		syncTaskTypes(db, l)
		syncEmployeesWithoutPhoto(db, l)
		fmt.Println("Done")
	}
}
