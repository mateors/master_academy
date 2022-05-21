package main

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

func main() {

	
	// open an existing file
	wb, err := xlsx.OpenFile("cod.xlsx")
	if err != nil {
		panic(err)
	}
	// wb now contains a reference to the workbook
	// show all the sheets in the workbook
	fmt.Println("Sheets in this file:")
	for i, sh := range wb.Sheets {
		fmt.Println(i, sh.Name)

	}

	xls := wb.Sheet["Sheet1"]
	//xls.Row()
	fmt.Println(xls.MaxRow, xls.MaxCol)
	xls.ForEachRow(rowVisitor)
	//first, err := xls.Row(1)
	//fmt.Println(err, first.ForEachCell())

	fmt.Println("----")
}

func rowVisitor(r *xlsx.Row) error {
	//fmt.Println(r)
	return r.ForEachCell(cellVisitor)
}

func cellVisitor(c *xlsx.Cell) error {
	value, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Cell value:", value)
	}
	return err
}
