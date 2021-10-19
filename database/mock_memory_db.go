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
	_, empl, err := findEmployee(db.employees, func(e m.Employee) bool {
		return e.DatabaseId == id
	})

	if err == nil {
		empl.Rfid = props.Rfid
		empl.Name = props.Name
		empl.Flex = props.Flex
		empl.Working = props.Working
		empl.Department = props.Department
		empl.Photo = props.Photo

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
	_, opt, err := findOption(db.options, func(o m.Option) bool {
		return o.DatabaseId == id
	})

	if err == nil {
		opt.Name = props.Name
		opt.WrapperId = props.WrapperId
		opt.Available = props.Available
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

func (db *MockMemoryDatabase) AddAction(action m.Action) error {
	_, oldAction, err := findAction(db.actions, func(e m.Action) bool {
		return e.DatabaseId == action.DatabaseId
	})

	if err == nil && oldAction.DatabaseId == action.DatabaseId {
		return fmt.Errorf("action with database id '%s' already exists", oldAction.DatabaseId)
	}

	db.actions = append(db.actions, action)
	return nil
}

func (db *MockMemoryDatabase) GetAllActions() ([]m.Action, error) {
	return db.actions, nil
}

func (db *MockMemoryDatabase) DeleteActionWithDatabaseId(id string, action m.Action) error {
	i, _, err := findAction(db.actions, func(e m.Action) bool {
		return e.DatabaseId == action.DatabaseId
	})

	if err == nil {
		db.options[i] = db.options[len(db.options)-1]
		db.options = db.options[:len(db.options)-1]
		return nil
	}

	return fmt.Errorf("could not find Action with database id '%s'", id)
}
