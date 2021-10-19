package api

import (
	"neocheckin_cache/database"
	m "neocheckin_cache/database/models"
	"testing"
)

func TestGetEmployeeFromRfid(t *testing.T) {

	{ // 0 employees in database
		db := database.MockMemoryDatabase{}

		_, err := GetEmployeeFromRfid(&db, "")

		if db.GetEmployeeWithRfidCallAmount != 1 {
			t.Errorf("should call db.GetEmployeeWithRfid once")
		}

		if err.Error() != "not found" {
			t.Errorf("should return error \"not found\"")
		}
	}

	{ // 1 employees in database
		db := database.MockMemoryDatabase{}

		mockEmployee := m.Employee{
			Rfid:       "12345678",
			Name:       "Ole Soelberg",
			Flex:       3600,
			Working:    true,
			Department: "Programmør",
			Photo:      "base64:soelberg",
		}

		db.InsertEmployee(mockEmployee)

		employee, err := GetEmployeeFromRfid(&db, "12345678")

		if err != nil {
			t.Errorf("should not return error")
		}

		if employee.Name != mockEmployee.Name {
			t.Errorf("should match, expected %q got %q", employee.Name, mockEmployee.Name)
		}

	}

	{ // multiple employees
		db := database.MockMemoryDatabase{}

		employees := []m.Employee{
			{
				Rfid:       "87654321",
				Name:       "Ole Helledie",
				Flex:       -3600,
				Working:    false,
				Department: "It-support",
				Photo:      "base64:helledie",
			},
			{
				Rfid:       "12345678",
				Name:       "Ole Soelberg",
				Flex:       3600,
				Working:    true,
				Department: "Programmør",
				Photo:      "base64:soelberg",
			},
			{
				Rfid:       "12345678",
				Name:       "Soelberg Ole",
				Flex:       0,
				Working:    true,
				Department: "Programmør",
				Photo:      "base64:ole",
			},
		}

		db.InsertEmployee(employees[0])
		db.InsertEmployee(employees[1])
		db.InsertEmployee(employees[2])

		employee, err := GetEmployeeFromRfid(&db, employees[1].Rfid)

		if err != nil {
			t.Errorf("should not return error")
		}

		if employees[1].Name != employee.Name {
			t.Errorf("should match, expected %q got %q", employee.Name, employees[1].Name)
		}

		employee, err = GetEmployeeFromRfid(&db, "01010101")

		if err == nil {
			t.Errorf("should return error")
		}

		for i := range employees {
			if employees[i].Name == employee.Name {
				t.Errorf("should not match")
			}
		}

	}

}

func TestConvertOptionsToExportedModels(t *testing.T) {
	{ // arrays should be same length
		o := make([]m.Option, 24)
		cO := ConvertOptionsToExportedModels(o)
		if len(o) != len(cO) {
			t.Error("should be same length")
		}
	}
}
