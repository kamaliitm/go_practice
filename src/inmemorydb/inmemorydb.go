package inmemorydb

type COLUMN_TYPE int

const (
	INTEGER COLUMN_TYPE = iota
	STRING
)

type ColumnData struct {
	ColumnObj Column
	Val       any
}

func (c *ColumnData) getVal() any {
	if c.ColumnObj.ColType == INTEGER {
		return c.Val.(int)
	}

	return c.Val.(string)
}

type Column struct {
	Name    string
	ColType COLUMN_TYPE
	// Val     string
}

type Row struct {
	ID   int
	Data map[string]ColumnData
}

type Filter interface {
	where()
}

type Table interface {
	InsertInto(data []ColumnData) error
	SelectFrom(columns []string, filter Filter) ([][]ColumnData, error)
	Update(data []ColumnData, filter Filter) error
	Remove(filter Filter) error
}

type DB interface {
	CreateTable(tableName string, columns []Column) error
	RenameTable(tableName string, newName string) error
	DropTable(tableName string) error

	GetTable(tableName string) error
}
