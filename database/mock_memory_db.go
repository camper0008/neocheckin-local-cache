package database

import (
	"fmt"
	"neocheckin_cache/database/models"
	"neocheckin_cache/shared"
)

type MockMemoryDatabase struct {
	MemoryDatabase
	GetEmployeeWithRfidCallAmount int
}

func (db *MockMemoryDatabase) GetEmployeeWithRfid(rfid string) (models.Employee, error) {
	_, empl, err := findEmployee(db.employees, func(e models.Employee) bool {
		return e.Rfid == rfid
	})

	if err == nil {
		return *empl, nil
	}

	return models.Employee{}, fmt.Errorf("could not find Employee with rfid '%s'", rfid)
}

func (db *MockMemoryDatabase) GetEmployeeWithDatabaseId(id string) (models.Employee, error) {
	db.GetEmployeeWithRfidCallAmount++

	_, empl, err := findEmployee(db.employees, func(e models.Employee) bool {
		return e.DatabaseId == id
	})

	if err == nil {
		return *empl, nil
	}

	return models.Employee{}, fmt.Errorf("could not find Employee with database id '%s'", id)
}

func (db *MockMemoryDatabase) InsertEmployee(empl models.Employee) error {
	_, oldEmpl, err := findEmployee(db.employees, func(e models.Employee) bool {
		return e.DatabaseId == empl.DatabaseId
	})

	if err == nil && oldEmpl.DatabaseId == empl.DatabaseId {
		return fmt.Errorf("employee with database id '%s' already exists", oldEmpl.DatabaseId)
	}

	db.employees = append(db.employees, empl)
	return nil
}

func (db *MockMemoryDatabase) UpdateEmployeeWithDatabaseId(id string, props models.Employee) error {
	_, empl, err := findEmployee(db.employees, func(e models.Employee) bool {
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
	i, _, err := findEmployee(db.employees, func(e models.Employee) bool {
		return e.DatabaseId == id
	})

	if err == nil {
		db.employees[i] = db.employees[len(db.employees)-1]
		db.employees = db.employees[:len(db.employees)-1]
		return nil
	}

	return fmt.Errorf("could not find Employee with database id '%s'", id)
}

func (db *MockMemoryDatabase) GetOptionWithWrapperId(id shared.WrapperEnum) (models.Option, error) {
	_, opt, err := findOption(db.options, func(o models.Option) bool {
		return o.WrapperId == id
	})

	if err == nil {
		return *opt, nil
	}

	return models.Option{}, fmt.Errorf("could not find Option with wrapper id '%d'", id)
}

func (db *MockMemoryDatabase) GetOptionWithDatabaseId(id string) (models.Option, error) {
	_, opt, err := findOption(db.options, func(o models.Option) bool {
		return o.DatabaseId == id
	})

	if err == nil {
		return *opt, nil
	}

	return models.Option{}, fmt.Errorf("could not find Option with database id '%s'", id)
}

func (db *MockMemoryDatabase) InsertOption(opt models.Option) error {
	_, oldOpt, err := findOption(db.options, func(o models.Option) bool {
		return o.DatabaseId == opt.DatabaseId
	})

	if err == nil && opt.DatabaseId == oldOpt.DatabaseId {
		return fmt.Errorf("option with database id '%s' already exists", oldOpt.DatabaseId)
	}

	db.options = append(db.options, opt)
	return nil
}

func (db *MockMemoryDatabase) UpdateOptionWithDatabaseId(id string, props models.Option) error {
	_, opt, err := findOption(db.options, func(o models.Option) bool {
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
	i, _, err := findOption(db.options, func(e models.Option) bool {
		return e.DatabaseId == id
	})

	if err == nil {
		db.options[i] = db.options[len(db.options)-1]
		db.options = db.options[:len(db.options)-1]
		return nil
	}

	return fmt.Errorf("could not find Option with database id '%s'", id)
}

func (db *MockMemoryDatabase) AddAction(action models.Action) error {
	_, oldAction, err := findAction(db.actions, func(e models.Action) bool {
		return e.DatabaseId == action.DatabaseId
	})

	if err == nil && oldAction.DatabaseId == action.DatabaseId {
		return fmt.Errorf("action with database id '%s' already exists", oldAction.DatabaseId)
	}

	db.actions = append(db.actions, action)
	return nil
}

func (db *MockMemoryDatabase) DeleteActionWithDatabaseId(id string, action models.Action) error {
	i, _, err := findAction(db.actions, func(e models.Action) bool {
		return e.DatabaseId == action.DatabaseId
	})

	if err == nil {
		db.options[i] = db.options[len(db.options)-1]
		db.options = db.options[:len(db.options)-1]
		return nil
	}

	return fmt.Errorf("could not find Action with database id '%s'", id)
}
