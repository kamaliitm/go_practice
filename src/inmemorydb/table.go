package inmemorydb

import "errors"

type TableImpl struct {
	Name       string
	Rows       map[int]Row
	Columns    []Column
	LastRowNum int
}

func NewTable(name string, cols []Column) Table {
	return &TableImpl{
		Name:    name,
		Rows:    map[int]Row{},
		Columns: cols,
	}
}

func (t *TableImpl) InsertInto(data []ColumnData) error {
	t.LastRowNum += 1
	row := Row{
		ID: t.LastRowNum,
	}
	rowData := map[string]ColumnData{}
	for _, colData := range data {
		rowData[colData.ColumnObj.Name] = colData
	}
	row.Data = rowData
	t.Rows[t.LastRowNum] = row
	return nil
}

func (t *TableImpl) SelectFrom(columns []string, filter Filter) ([][]ColumnData, error) {
	returnData := [][]ColumnData{}
	for _, row := range t.Rows {
		returnRow := []ColumnData{}
		for _, colData := range row.Data {
			if stringExists(colData.ColumnObj.Name, columns) {
				colType := colData.ColumnObj.ColType
				if colType == INTEGER {
					colData.Val = colData.Val.(int)
				}
				returnRow = append(returnRow, colData)
			}
		}
		returnData = append(returnData, returnRow)
	}

	return returnData, nil
}

func (t *TableImpl) Update(data []ColumnData, filter Filter) error {
	err := t.validateSchema(data)
	if err != nil {
		return err
	}

	rows := t.applyFilter(filter)
	for id, row := range rows {
		for _, colData := range data {
			row.Data[colData.ColumnObj.Name] = colData
		}
		t.Rows[id] = row
	}

	return nil
}

func (t *TableImpl) Remove(filter Filter) error {
	rows := t.applyFilter(filter)
	for id := range rows {
		delete(t.Rows, id)
	}

	return nil
}

func (t *TableImpl) applyFilter(filter Filter) map[int]Row {
	return t.Rows
}

func (t *TableImpl) validateSchema(data []ColumnData) error {
	for _, col := range data {
		if !columnExists(col.ColumnObj.Name, t.Columns) {
			return errors.New("column does not exist")
		}
	}

	return nil
}

func columnExists(colName string, arr []Column) bool {
	for _, column := range arr {
		if column.Name == colName {
			return true
		}
	}

	return false
}

func stringExists(colName string, arr []string) bool {
	for _, s := range arr {
		if s == colName {
			return true
		}
	}

	return false
}
