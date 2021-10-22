package database

import (
	"fmt"
	m "neocheckin_cache/database/models"
	"neocheckin_cache/utils"
)

type MemoryDatabase struct {
	employees []m.Employee
	options   []m.Option
	tasks     []m.Task
}

func findEmployee(a []m.Employee, f func(m.Employee) bool) (int, *m.Employee, error) {
	for i, v := range a {
		if f(v) {
			return i, &v, nil
		}
	}
	return -1, nil, fmt.Errorf("employee not found")
}

func findOption(a []m.Option, f func(m.Option) bool) (int, *m.Option, error) {
	for i, v := range a {
		if f(v) {
			return i, &v, nil
		}
	}
	return -1, nil, fmt.Errorf("option not found")
}

func findTask(a []m.Task, f func(m.Task) bool) (int, *m.Task, error) {
	for i, v := range a {
		if f(v) {
			return i, &v, nil
		}
	}
	return -1, nil, fmt.Errorf("task not found")
}

func (db *MemoryDatabase) GetEmployeeWithRfid(rfid string) (m.Employee, error) {
	_, empl, err := findEmployee(db.employees, func(e m.Employee) bool {
		return e.Rfid == rfid
	})

	if err == nil {
		return *empl, nil
	}

	return m.Employee{}, fmt.Errorf("could not find Employee with rfid '%s'", rfid)
}

func (db *MemoryDatabase) GetEmployeeWithDatabaseId(id string) (m.Employee, error) {
	_, empl, err := findEmployee(db.employees, func(e m.Employee) bool {
		return e.DatabaseId == id
	})

	if err == nil {
		return *empl, nil
	}

	return m.Employee{}, fmt.Errorf("could not find Employee with database id '%s'", id)
}

func (db *MemoryDatabase) GetAllEmployees() ([]m.Employee, error) {
	return db.employees, nil
}

func (db *MemoryDatabase) ReplaceEmployees(e []m.Employee) error {
	db.employees = e
	return nil
}

func (db *MemoryDatabase) InsertEmployee(empl m.Employee) error {
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

func (db *MemoryDatabase) UpdateEmployeeWithDatabaseId(id string, props m.Employee) error {
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

func (db *MemoryDatabase) DeleteEmployeeWithDatabaseId(id string) error {
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

func (db *MemoryDatabase) GetOptionWithWrapperId(id int) (m.Option, error) {
	_, opt, err := findOption(db.options, func(o m.Option) bool {
		return o.WrapperId == id
	})

	if err == nil {
		return *opt, nil
	}

	return m.Option{}, fmt.Errorf("could not find Option with wrapper id '%d'", id)
}

func (db *MemoryDatabase) GetOptionWithDatabaseId(id string) (m.Option, error) {
	_, opt, err := findOption(db.options, func(o m.Option) bool {
		return o.DatabaseId == id
	})

	if err == nil {
		return *opt, nil
	}

	return m.Option{}, fmt.Errorf("could not find Option with database id '%s'", id)
}

func (db *MemoryDatabase) GetAllOptions() ([]m.Option, error) {
	return db.options, nil
}

func (db *MemoryDatabase) ReplaceOptions(o []m.Option) error {
	db.options = o
	return nil
}

func (db *MemoryDatabase) InsertOption(opt m.Option) error {
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

func (db *MemoryDatabase) UpdateOptionWithDatabaseId(id string, props m.Option) error {
	_, opt, err := findOption(db.options, func(o m.Option) bool {
		return o.DatabaseId == id
	})

	if err == nil {
		opt.Name = props.Name
		opt.WrapperId = props.WrapperId
		opt.Schedule = props.Schedule
	}

	return fmt.Errorf("could not find Option with database id '%s'", id)
}

func (db *MemoryDatabase) DeleteOptionWithDatabaseId(id string) error {
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

func (db *MemoryDatabase) AddTask(task m.Task) error {
	_, oldTask, err := findTask(db.tasks, func(e m.Task) bool {
		return e.DatabaseId == task.DatabaseId
	})

	if err == nil && oldTask.DatabaseId == task.DatabaseId {
		return fmt.Errorf("task with database id '%s' already exists", oldTask.DatabaseId)
	}

	db.tasks = append(db.tasks, task)
	return nil
}

func (db *MemoryDatabase) GetAllTasks() ([]m.Task, error) {
	return db.tasks, nil
}

func (db *MemoryDatabase) DeleteTaskWithDatabaseId(id string) error {
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
