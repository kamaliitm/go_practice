package inmemorydb

import "errors"

type DBImpl struct {
	Name   string
	Tables map[string]Table
}

func NewDB(name string) DB {
	return &DBImpl{
		Name:   name,
		Tables: map[string]Table{},
	}
}

func (d *DBImpl) CreateTable(tableName string, columns []Column) error {
	if _, ok := d.Tables[tableName]; ok {
		return errors.New("table already exists")
	}

	table := NewTable(tableName, columns)
	d.Tables[tableName] = table

	return nil
}

func (d *DBImpl) RenameTable(tableName string, newName string) error {
	if _, ok := d.Tables[tableName]; ok {
		return errors.New("table already exists")
	}

	table := d.Tables[tableName]
	d.Tables[newName] = table
	delete(d.Tables, tableName)

	return nil
}

func (d *DBImpl) DropTable(tableName string) error {
	return nil
}

func (d *DBImpl) GetTable(tableName string) error {
	return nil
}
