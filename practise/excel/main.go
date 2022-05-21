package main

import (
	"fmt"
	"log"

	"github.com/pbnjay/grate"
	_ "github.com/pbnjay/grate/simple" // tsv and csv support
	_ "github.com/pbnjay/grate/xls"
	_ "github.com/pbnjay/grate/xlsx"
)

func main() {
	// wb, _ := grate.Open(os.Args[1]) // open the file
	// sheets, _ := wb.List()          // list available sheets
	// for _, s := range sheets {      // enumerate each sheet name
	// 	sheet, _ := wb.Get(s) // open the sheet
	// 	for sheet.Next() {    // enumerate each row of data
	// 		row := sheet.Strings() // get the row's content as []string
	// 		fmt.Println(strings.Join(row, "\t"))
	// 	}
	// }
	// wb.Close()
	src, err := grate.Open("cod.xlsx")
	if err != nil {
		log.Println("ERR1", err)
		return
	}
	sheets, err := src.List()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(sheets)
}
