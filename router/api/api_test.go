package api

import (
	"neocheckin_cache/database"
	"neocheckin_cache/database/models"
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

		mockEmployee := models.Employee{
			Rfid:       "12345678",
			Name:       "Ole Soelberg",
			Flex:       3600,
			Working:    true,
			Department: "Programmør",
			Photo:      "base64:iguess",
		}

		db.InsertEmployee(mockEmployee)

		employee, err := GetEmployeeFromRfid(&db, "12345678")

		if err != nil {
			t.Errorf("should not return error")
		}

		if employee != mockEmployee {
			t.Errorf("should match")
		}

		if employee.Name != mockEmployee.Name {
			t.Errorf("should match")
		}

	}

	{ // multiple employees
		db := database.MockMemoryDatabase{}

		employees := []models.Employee{
			{
				DatabaseModel: models.DatabaseModel{
					DatabaseId: "0",
				},
				Rfid:       "87654321",
				Name:       "Ole Helledie",
				Flex:       -3600,
				Working:    false,
				Department: "It-support",
				Photo:      "base64:iguess",
			},
			{
				DatabaseModel: models.DatabaseModel{
					DatabaseId: "1",
				},
				Rfid:       "12345678",
				Name:       "Ole Soelberg",
				Flex:       3600,
				Working:    true,
				Department: "Programmør",
				Photo:      "base64:iguess",
			},
			{
				DatabaseModel: models.DatabaseModel{
					DatabaseId: "2",
				},
				Rfid:       "12345678",
				Name:       "Praktikplads Taber",
				Flex:       6969,
				Working:    true,
				Department: "Programmør",
				Photo:      "base64:iguess",
			},
		}

		db.InsertEmployee(employees[0])
		db.InsertEmployee(employees[1])
		db.InsertEmployee(employees[2])

		employee, err := GetEmployeeFromRfid(&db, employees[1].Rfid)

		if err != nil {
			t.Errorf("should not return error")
		}

		if employees[1] != employee {
			t.Errorf("should match")
		}

		if employees[1].Name != employee.Name {
			t.Errorf("should match")
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
