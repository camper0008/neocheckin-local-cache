package database

import (
	"fmt"
	m "neocheckin_cache/database/models"
	"neocheckin_cache/utils"
)

type MockMemoryDatabase struct {
	MemoryDatabase
	GetEmployeeWithRfidCallAmount int
}

func (db *MockMemoryDatabase) GetEmployeeWithRfid(rfid string) (m.Employee, error) {
	db.GetEmployeeWithRfidCallAmount++
	_, empl, err := findEmployee(db.employees, func(e m.Employee) bool {
		return e.Rfid == rfid
	})

	if err == nil {
		return *empl, nil
	}

	return m.Employee{}, fmt.Errorf("could not find Employee with rfid '%s'", rfid)
}

func (db *MockMemoryDatabase) GetEmployeeWithDatabaseId(id string) (m.Employee, error) {
	db.GetEmployeeWithRfidCallAmount++
	_, empl, err := findEmployee(db.employees, func(e m.Employee) bool {
		return e.DatabaseId == id
	})

	if err == nil {
		return *empl, nil
	}

	return m.Employee{}, fmt.Errorf("could not find Employee with database id '%s'", id)
}

func (db *MockMemoryDatabase) GetAllEmployees() ([]m.Employee, error) {
	return db.employees, nil
}

func (db *MockMemoryDatabase) ReplaceEmployees(e []m.Employee) error {
	for i := range e {
		if e[i].DatabaseId == "" {
			id := utils.GenerateUUID()
			e[i].DatabaseId = id
		}
	}
	db.employees = e
	return nil
}

func (db *MockMemoryDatabase) InsertEmployee(empl m.Employee) error {
	if empl.DatabaseId == "" {
		empl.DatabaseId = utils.GenerateUUID()
	}
	_, oldEmpl, err := findEmployee(db.employees, func(e m.Employee) bool {
		return e.DatabaseId == empl.DatabaseId
	})

	if err == nil && oldEmpl.DatabaseId == empl.DatabaseId {
		return fmt.Errorf("employee with database id '%s' already exists", oldEmpl.DatabaseId)
	}

	db.employees = append(db.employees, empl)
	return nil
}

func (db *MockMemoryDatabase) UpdateEmployeeWithDatabaseId(id string, props m.Employee) error {
	i, _, err := findEmployee(db.employees, func(e m.Employee) bool {
		return e.DatabaseId == id
	})

	if err == nil {
		db.employees[i].Rfid = props.Rfid
		db.employees[i].Name = props.Name
		db.employees[i].Flex = props.Flex
		db.employees[i].Working = props.Working
		db.employees[i].Department = props.Department
		db.employees[i].Photo = props.Photo

		return nil
	}

	return fmt.Errorf("could not find Employee with database id '%s'", id)
}

func (db *MockMemoryDatabase) DeleteEmployeeWithDatabaseId(id string) error {
	i, _, err := findEmployee(db.employees, func(e m.Employee) bool {
		return e.DatabaseId == id
	})

	if err == nil {
		db.employees[i] = db.employees[len(db.employees)-1]
		db.employees = db.employees[:len(db.employees)-1]
		return nil
	}

	return fmt.Errorf("could not find Employee with database id '%s'", id)
}

func (db *MockMemoryDatabase) GetOptionWithWrapperId(id int) (m.Option, error) {
	_, opt, err := findOption(db.options, func(o m.Option) bool {
		return o.WrapperId == id
	})

	if err == nil {
		return *opt, nil
	}

	return m.Option{}, fmt.Errorf("could not find Option with wrapper id '%d'", id)
}

func (db *MockMemoryDatabase) GetOptionWithDatabaseId(id string) (m.Option, error) {
	_, opt, err := findOption(db.options, func(o m.Option) bool {
		return o.DatabaseId == id
	})

	if err == nil {
		return *opt, nil
	}

	return m.Option{}, fmt.Errorf("could not find Option with database id '%s'", id)
}

func (db *MockMemoryDatabase) GetAllOptions() ([]m.Option, error) {
	return db.options, nil
}

func (db *MockMemoryDatabase) ReplaceOptions(o []m.Option) error {
	for i := range o {
		if o[i].DatabaseId == "" {
			id := utils.GenerateUUID()
			o[i].DatabaseId = id
		}
	}
	db.options = o
	return nil
}

func (db *MockMemoryDatabase) InsertOption(opt m.Option) error {
	if opt.DatabaseId == "" {
		opt.DatabaseId = utils.GenerateUUID()
	}

	_, oldOpt, err := findOption(db.options, func(o m.Option) bool {
		return o.DatabaseId == opt.DatabaseId
	})

	if err == nil && opt.DatabaseId == oldOpt.DatabaseId {
		return fmt.Errorf("option with database id '%s' already exists", oldOpt.DatabaseId)
	}

	db.options = append(db.options, opt)
	return nil
}

func (db *MockMemoryDatabase) UpdateOptionWithDatabaseId(id string, props m.Option) error {
	i, _, err := findOption(db.options, func(o m.Option) bool {
		return o.DatabaseId == id
	})

	if err == nil {
		db.options[i].Name = props.Name
		db.options[i].WrapperId = props.WrapperId
		db.options[i].Schedule = props.Schedule
	}

	return fmt.Errorf("could not find Option with database id '%s'", id)
}

func (db *MockMemoryDatabase) DeleteOptionWithDatabaseId(id string) error {
	i, _, err := findOption(db.options, func(e m.Option) bool {
		return e.DatabaseId == id
	})

	if err == nil {
		db.options[i] = db.options[len(db.options)-1]
		db.options = db.options[:len(db.options)-1]
		return nil
	}

	return fmt.Errorf("could not find Option with database id '%s'", id)
}

func (db *MockMemoryDatabase) AddTask(task m.Task) error {
	_, oldTask, err := findTask(db.tasks, func(e m.Task) bool {
		return e.DatabaseId == task.DatabaseId
	})

	if err == nil && oldTask.DatabaseId == task.DatabaseId {
		return fmt.Errorf("task with database id '%s' already exists", oldTask.DatabaseId)
	}

	db.tasks = append(db.tasks, task)
	return nil
}

func (db *MockMemoryDatabase) GetAllTasks() ([]m.Task, error) {
	return db.tasks, nil
}

func (db *MockMemoryDatabase) DeleteTaskWithDatabaseId(id string) error {
	i, _, err := findTask(db.tasks, func(e m.Task) bool {
		return e.DatabaseId == id
	})

	if err == nil {
		db.options[i] = db.options[len(db.options)-1]
		db.options = db.options[:len(db.options)-1]
		return nil
	}

	return fmt.Errorf("could not find Task with database id '%s'", id)
}
