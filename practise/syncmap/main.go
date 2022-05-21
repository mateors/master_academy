package main

import (
	"log"
	"strings"
	"sync"
	"time"

	"fmt"

	. "github.com/ahmetb/go-linq/v3"
	"github.com/go-gota/gota/dataframe"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/puzpuzpuz/xsync"
	"github.com/tidwall/shardmap"
	"github.com/xuri/excelize/v2"
)

//https://forum.golangbridge.org/t/safe-concurrent-read-and-write-in-maps/15419
type MyMap struct {
	//sync.RWMutex
	sync.Mutex
	//data map[string]interface{}
	data map[string][]map[string]interface{}
}

func (m *MyMap) read(key string) ([]map[string]interface{}, bool) {
	//m.RLock()
	m.Lock()
	result, ok := m.data[key]
	//m.RUnlock()
	m.Unlock()
	return result, ok
}

func (m *MyMap) write(key string, value []map[string]interface{}) {
	m.Lock()
	m.data[key] = value
	m.Unlock()
}

func main() {

	start := time.Now()

	// wg := sync.WaitGroup{}
	// wg.Add(5)

	// go func() {
	// 	dataFrame() //
	// 	wg.Done()
	// }()

	// go func() {
	// 	conMap() //
	// 	wg.Done()
	// }()

	// go func() {
	// 	myMap() //
	// 	wg.Done()
	// }()

	// go func() {
	// 	sharedMap()
	// 	wg.Done()
	// }()

	// go func() {
	// 	xsyncMap() //
	// 	wg.Done()
	// }()

	// go func() {
	// 	syncMap() //
	// 	wg.Done()
	// }()

	// go func() {
	// 	manualGroupVals() //
	// 	wg.Done()
	// }()

	//wg.Wait()
	goLinq()
	fmt.Println("TotalTimeTaken:", time.Since(start).Seconds(), "s")

}

func goLinq() {

	rows := []map[string]interface{}{
		{"name": "Mostain", "salary": 30000, "ShipmentTag": "b001"},
		{"name": "Sanzida", "salary": 20000, "ShipmentTag": "b001"},
		{"name": "Riaz", "salary": 21000, "ShipmentTag": "b002"},
		{"name": "Pallobi", "salary": 30000, "ShipmentTag": "b001"},
		{"name": "Moaz", "salary": 20000, "ShipmentTag": "b001"},
		{"name": "Tareq", "salary": 21000, "ShipmentTag": "b002"},
	}

	//var frows = make([]map[string]interface{}, 0)
	//src := reflect.ValueOf(rows)
	//keys := src.MapKeys()
	//item := src.Index(0).Interface()

	//fmt.Println(src.Kind(), item)
	//fmt.Println(keys)

	//var result interface{} //= make([]map[string]interface{},0)

	// result := make(map[string]interface{})

	// From(rows).WhereT(func(x map[string]interface{}) bool {
	// 	return len(x["name"].(string)) > 0

	// }).GroupByT(
	// 	func(x map[string]interface{}) interface{} {

	// 		return x["ShipmentTag"]
	// 	},
	// 	func(x interface{}) interface{} {
	// 		return x

	// 	}).ToMapBy(&result,
	// 	func(i interface{}) interface{} {
	// 		return i.(Group).Key

	// 	}, func(i interface{}) interface{} {
	// 		return i.(Group).Group

	// 	})

	// for key, val := range result {
	// 	fmt.Println(key, len(val.([]interface{})))
	// }

	result := make(map[string]interface{})

	From(rows).Where(func(c interface{}) bool {
		cmap := c.(map[string]interface{})
		return len(cmap["name"].(string)) > 0

	}).GroupBy(func(c interface{}) interface{} {
		cmap := c.(map[string]interface{})
		return cmap["ShipmentTag"]

	}, func(ShipmentTag interface{}) interface{} {
		return ShipmentTag

	}).ToMapBy(&result,
		func(i interface{}) interface{} {
			return i.(Group).Key
		}, func(i interface{}) interface{} {
			return i.(Group).Group
		})

	fmt.Println(len(result))
	// for i, val := range data.Results() {
	// 	fmt.Printf("%d %T\n",i, val)
	// }

}

func keySel(i interface{}) interface{} {
	return i.(Group).Key
}

func valSel(x interface{}) interface{} {

	return x

}

func manualGroupVals() {

	start := time.Now()
	var guvs = make(map[string][]map[string]interface{})
	//groupDataFrame := make(map[string]interface{})
	grpcols := []string{"ShipmentTag"}
	// rows := []map[string]interface{}{
	// 	{"name": "Mostain", "salary": 30000, "ShipmentTag": "b001"},
	// 	{"name": "Sanzida", "salary": 20000, "ShipmentTag": "b001"},
	// 	{"name": "Riaz", "salary": 21000, "ShipmentTag": "b002"},
	// 	{"name": "Pallobi", "salary": 30000, "ShipmentTag": "b001"},
	// 	{"name": "Moaz", "salary": 20000, "ShipmentTag": "b001"},
	// 	{"name": "Tareq", "salary": 21000, "ShipmentTag": "b002"},
	// }

	rows, err := excelReader("cod2.xlsx", "Sheet1", nil)
	if err != nil {
		log.Println(err)
		return
	}

	for _, row := range rows {

		key := ""
		for i, col := range grpcols {

			format := ""
			if i == 0 {
				format = "%s%v"
			} else {
				format = "%s_%v"
			}
			//keyname := fmt.Sprint("%v", row[col])
			key = fmt.Sprintf(format, key, row[col])
			//guvs[keyname] = append(guvs[keyname], col)
			//fmt.Println(">>", key, row[col])
		}
		guvs[key] = append(guvs[key], row)
	}

	fmt.Println("manualGroupVals:", len(guvs), time.Since(start).Seconds(), "s")
	// for key, val := range guvs {
	// 	fmt.Printf("%v %d %T %v\n", key, len(val), val, val[0]["ShipmentTag"])
	// }

}

func dataFrame() {

	// rows := []map[string]interface{}{
	// 	{"name": "Mostain", "salary": 30000, "ShipmentTag": "b001"},
	// 	{"name": "Sanzida", "salary": 20000, "ShipmentTag": "b001"},
	// 	{"name": "Riaz", "salary": 21000, "ShipmentTag": "b002"},
	// 	{"name": "Pallobi", "salary": 15000, "ShipmentTag": "b001"},
	// 	{"name": "Moaz", "salary": 20000, "ShipmentTag": "b001"},
	// 	{"name": "Tareq", "salary": 22000, "ShipmentTag": "b002"},
	// }

	sTart := time.Now()
	rows, err := excelReader("cod2.xlsx", "Sheet1", nil)
	if err != nil {
		log.Println(err)
		return
	}

	df := dataframe.LoadMaps(rows)

	dfg := df.GroupBy("ShipmentTag")
	rmap := dfg.GetGroups()
	//fmt.Println(rmap)
	// for key, val := range rmap {
	// 	fmt.Println(key, "**", len(val.Maps()))
	// }
	take := time.Since(sTart).Seconds()
	fmt.Println("dataFrame:", len(rmap), take, "s")

}

//GODEBUG=gctrace=1 ./sy
//https://github.com/cornelk/hashmap
//https://stackoverflow.com/questions/55389749/is-writing-to-a-mutex-map-with-multiple-goroutines-faster-than-one-and-why

func conMap() {

	sTart := time.Now()
	rows, err := excelReader("cod2.xlsx", "Sheet1", nil)
	if err != nil {
		log.Println(err)
		return
	}

	var smap = cmap.New()
	var wg sync.WaitGroup
	totalRows := len(rows)
	limit := 1000
	se := calcMaster(totalRows, limit)
	wg.Add(len(se))

	for _, s := range se {
		go adpro5(&smap, &wg, rows[s.Start:s.End])
	}

	wg.Wait()

	rmap := []map[string]interface{}{}
	for mval := range smap.IterBuffered() {

		//fmt.Println(mval.Key, "==", mval.Val)
		vals, isOk := mval.Val.([]interface{})
		if isOk {
			for _, val := range vals {
				kv := make(map[string]interface{})
				kv[fmt.Sprint(val)] = mval.Key
				//fmt.Println(val, "==>", key)
				rmap = append(rmap, kv)
			}
		}
	}

	fmt.Println("conMap:", len(rmap), time.Since(sTart).Seconds(), "s")

}

func myMap() {

	sTart := time.Now()
	rows, err := excelReader("cod2.xlsx", "Sheet1", nil)
	if err != nil {
		log.Println(err)
		return
	}
	// rows := []map[string]interface{}{
	// 	{"name": "Mostain", "salary": 30000, "ShipmentTag": "b001"},
	// 	{"name": "Sanzida", "salary": 20000, "ShipmentTag": "b001"},
	// 	{"name": "Riaz", "salary": 21000, "ShipmentTag": "b002"},
	// 	{"name": "Pallobi", "salary": 30000, "ShipmentTag": "b001"},
	// 	{"name": "Moaz", "salary": 20000, "ShipmentTag": "b001"},
	// 	{"name": "Tareq", "salary": 21000, "ShipmentTag": "b002"},
	// }

	//var shmap = shardmap.New(1) //xsync.NewMap() //sync.Map
	dmap := make(map[string][]map[string]interface{})
	var smap MyMap
	smap.data = dmap

	var wg sync.WaitGroup
	totalRows := len(rows)
	limit := 500
	se := calcMaster(totalRows, limit)
	wg.Add(len(se))

	for _, s := range se {
		go adpro4(&smap, &wg, rows[s.Start:s.End])
	}

	wg.Wait()

	// rmap := []map[string]interface{}{}
	// //fmt.Printf("%v %T\n", smap.data, smap.data) //map[string][]map[string]interface{}
	// for key, mval := range smap.data {

	// 	vals, isOk := mval.([]interface{})
	// 	if isOk {
	// 		for _, val := range vals {
	// 			kv := make(map[string]interface{})
	// 			kv[fmt.Sprint(val)] = key
	// 			//fmt.Println(val, "==>", key)
	// 			rmap = append(rmap, kv)
	// 		}
	// 	}
	// }

	fmt.Println("myMap:", len(smap.data), time.Since(sTart).Seconds(), "s")

}

func sharedMap() {

	sTart := time.Now()
	rows, err := excelReader("cod2.xlsx", "Sheet1", nil)
	if err != nil {
		log.Println(err)
		return
	}

	//var shmap = shardmap.New(1) //xsync.NewMap() //sync.Map
	var smap shardmap.Map

	var wg sync.WaitGroup
	totalRows := len(rows)
	limit := 1000
	se := calcMaster(totalRows, limit)
	wg.Add(len(se))

	for _, s := range se {
		go adpro3(&smap, &wg, rows[s.Start:s.End])
	}

	wg.Wait()

	rmap := []map[string]interface{}{}
	smap.Range(func(key string, val interface{}) bool {
		vals := val.([]interface{})
		for _, val := range vals {
			kv := make(map[string]interface{})
			kv[fmt.Sprint(val)] = key
			rmap = append(rmap, kv)
		}
		return true
	})
	fmt.Println("sharedMap:", len(rmap), time.Since(sTart).Seconds(), "s")

}

func xsyncMap() {

	sTart := time.Now()
	rows, err := excelReader("cod2.xlsx", "Sheet1", nil)
	if err != nil {
		log.Println(err)
		return
	}

	var smap = xsync.NewMap() //sync.Map
	var wg sync.WaitGroup
	totalRows := len(rows)
	limit := 1000
	se := calcMaster(totalRows, limit)
	wg.Add(len(se))

	for _, s := range se {
		go adpro2(smap, &wg, rows[s.Start:s.End])
	}

	wg.Wait()

	rmap := []map[string]interface{}{}
	smap.Range(func(key string, val interface{}) bool {

		//vals := val.([]map[string]interface{})
		// for k, val := range vals {
		// 	kv := make(map[string]interface{})
		// 	kv[fmt.Sprint(k)] = key
		// 	rmap = append(rmap, kv)
		// }
		//rmap = append(rmap, vals...)
		kv := make(map[string]interface{})
		kv[fmt.Sprint(key)] = key
		rmap = append(rmap, kv)
		return true
	})
	fmt.Println("xsyncMap:", len(rmap), time.Since(sTart).Seconds(), "s")

}

func syncMap() {

	sTart := time.Now()
	// rows := []map[string]interface{}{
	// 	{"name": "Mostain", "salary": 30000, "batch": "b001"},
	// 	{"name": "Sanzida", "salary": 20000, "batch": "b001"},
	// 	{"name": "Riaz", "salary": 21000, "batch": "b002"},
	// 	{"name": "Pallobi", "salary": 30000, "batch": "b001"},
	// 	{"name": "Moaz", "salary": 20000, "batch": "b001"},
	// 	{"name": "Tareq", "salary": 21000, "batch": "b002"},
	// }
	rows, err := excelReader("cod2.xlsx", "Sheet1", nil)
	if err != nil {
		log.Println(err)
		return
	}

	var smap sync.Map
	var wg sync.WaitGroup
	totalRows := len(rows)
	limit := 1000
	se := calcMaster(totalRows, limit)
	wg.Add(len(se))

	for _, s := range se {
		go adpro(&smap, &wg, rows[s.Start:s.End])
	}

	//go adpro(&smap, &wg, rows[0:3])
	//go adpro(&smap, &wg, rows[3:])
	wg.Wait()

	rmap := []map[string]interface{}{}
	//fmt.Println(smap)
	smap.Range(func(key, val interface{}) bool {

		//vals := val.([]map[string]interface{})
		//fmt.Printf("%s %v %T\n", key, vals, vals)
		//for _, val := range vals {
		//kv := make(map[string]interface{})
		//kv[fmt.Sprint(val)] = key
		//rmap = append(rmap, val)
		//}

		kv := make(map[string]interface{})
		kv[fmt.Sprint(key)] = key
		rmap = append(rmap, kv)
		return true
	})

	fmt.Println("syncMap:", len(rmap), time.Since(sTart).Seconds(), "s")

}

type startEnd struct {
	Start int
	End   int
}

func excelReader(filePath, sheetName string, dateTimeFormat map[string]string) ([]map[string]interface{}, error) {

	sRows := make([]map[string]interface{}, 0)
	keys := make([]string, 0) //first row columns

	xlf, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	//smap := xlf.GetSheetMap()
	//sheetName := smap[1] //xlf.GetSheetName(1)
	//fmt.Println("sheetName:", sheetName, smap, smap[1])

	if dateTimeFormat != nil {

		if dateCols, isFound := dateTimeFormat["date"]; isFound {
			dates := strings.Split(dateCols, ",")
			//https://xuri.me/excelize/en/style.html#number_format
			styleID, _ := xlf.NewStyle(`{"number_format":15}`) //14
			for _, column := range dates {
				//xlf.SetCellStyle(sheetName, "K", "K", styleID)
				err = xlf.SetColStyle(sheetName, column, styleID) //column must be in capital
				if err != nil {
					log.Println("ErrDateFormat", err)
				}
			}
		}

		if dateTimeCols, isFound := dateTimeFormat["datetime"]; isFound {
			dates := strings.Split(dateTimeCols, ",")
			styleID, _ := xlf.NewStyle(`{"number_format":22}`)
			for _, column := range dates {
				err = xlf.SetColStyle(sheetName, column, styleID) //column must be in capital
				if err != nil {
					log.Println("ErrDateFormat", err)
				}
			}
		}

	}

	var i int
	erows, err := xlf.Rows(sheetName)
	if err != nil {
		return nil, err
	}

	for erows.Next() {

		cols, err := erows.Columns()
		if err != nil {
			continue
		}

		if len(cols) == 0 {
			continue
		}

		i++
		if i == 1 {
			keys = cols
			continue
		}

		vmap := make(map[string]interface{})
		for j := 0; j < len(keys); j++ {
			key := strings.TrimSpace(keys[j]) //trime
			if len(key) > 0 {
				var val string
				if j < len(cols) {
					val = strings.TrimSpace(cols[j]) //trime
				}
				vmap[key] = val
				//fmt.Println(key, "==>", val)
			}
		}
		if len(vmap) > 0 {
			sRows = append(sRows, vmap)
		}
	}
	return sRows, nil
}

func calcMaster(totalRows, limit int) []*startEnd {

	//var totalRows uint32 = 196542
	//var limit uint32 = 1000
	var se = []*startEnd{}
	var i int = 0

	for {
		if totalRows <= i { //totalRows <= i
			se = append(se, &startEnd{Start: i, End: totalRows + 1})
			break
		}
		start := i
		end := i + limit
		if end > totalRows {
			diff := totalRows - start
			end = start + diff
		}
		//fmt.Println(start, end)
		se = append(se, &startEnd{Start: start, End: end})
		i = end
		if end == totalRows {
			break
		}
	}
	return se
}

func adpro(smap *sync.Map, wg *sync.WaitGroup, rows []map[string]interface{}) {

	defer wg.Done()
	grpFields := []string{"ShipmentTag"}

	for _, row := range rows {

		// for _, grp := range grpFields {

		// 	if val, isOk := smap.Load(grp); !isOk {
		// 		slc := make([]interface{}, 0)
		// 		vals := append(slc, row[grp])
		// 		smap.Store(grp, vals)
		// 		fmt.Println(">", grp, row[grp])

		// 	} else {
		// 		//fmt.Printf("%v %T\n", val, val)
		// 		pvals, isOk := val.([]interface{})
		// 		if isOk {

		// 			isFound := FindExist(pvals, row[grp])
		// 			//fmt.Println(row[grp], isFound, pvals)
		// 			if isFound == false {
		// 				uvals := append(pvals, row[grp])
		// 				smap.Store(grp, uvals)
		// 			}
		// 		}
		// 	}
		// }

		key := ""
		for i, grp := range grpFields {
			format := ""
			if i == 0 {
				format = "%s%v"
			} else {
				format = "%s_%v"
			}
			key = fmt.Sprintf(format, key, row[grp])
		}
		// eany, _ := smap.Load(key)
		// evals, isOk := eany.([]map[string]interface{})
		// if isOk {
		// 	evals = append(evals, row)
		// 	smap.Store(key, evals)

		// } else {
		// 	evals := make([]map[string]interface{}, 0)
		// 	evals = append(evals, row)
		// 	smap.Store(key, evals)
		// }
		smap.Store(key, "ShipmentTag")

	}

}

//xsync
func adpro2(smap *xsync.Map, wg *sync.WaitGroup, rows []map[string]interface{}) {

	defer wg.Done()
	grpFields := []string{"ShipmentTag"}

	for _, row := range rows {

		// for _, grp := range grpFields {

		// 	if val, isOk := smap.Load(grp); !isOk {
		// 		slc := make([]interface{}, 0)
		// 		vals := append(slc, row[grp])
		// 		smap.Store(grp, vals)
		// 		fmt.Println(">", grp, row[grp])

		// 	} else {
		// 		//fmt.Printf("%v %T\n", val, val)
		// 		pvals, isOk := val.([]interface{})
		// 		if isOk {

		// 			isFound := FindExist(pvals, row[grp])
		// 			//fmt.Println(row[grp], isFound, pvals)
		// 			if isFound == false {
		// 				uvals := append(pvals, row[grp])
		// 				smap.Store(grp, uvals)
		// 			}

		// 		}
		// 	}
		// }

		key := ""
		for i, grp := range grpFields {
			format := ""
			if i == 0 {
				format = "%s%v"
			} else {
				format = "%s_%v"
			}
			key = fmt.Sprintf(format, key, row[grp])
		}

		smap.Store(key, "ShipmentTag")
		// evals, isOk := smap.Load(key)
		// if !isOk {
		// 	evals := make([]map[string]interface{}, 0)
		// 	smap.Store(key, append(evals, row))
		// } else {
		// 	vals := evals.([]map[string]interface{})
		// 	smap.Store(key, append(vals, row))
		// }
	}

}

func adpro3(smap *shardmap.Map, wg *sync.WaitGroup, rows []map[string]interface{}) {

	defer wg.Done()
	grpFields := []string{"ShipmentTag"}

	for _, row := range rows {

		for _, grp := range grpFields {

			if val, isOk := smap.Get(grp); !isOk {
				slc := make([]interface{}, 0)
				vals := append(slc, row[grp])
				smap.Set(grp, vals)
				fmt.Println(">", grp, row[grp])

			} else {
				//fmt.Printf("%v %T\n", val, val)
				pvals, isOk := val.([]interface{})
				if isOk {

					isFound := FindExist(pvals, row[grp])
					//fmt.Println(row[grp], isFound, pvals)
					if isFound == false {
						uvals := append(pvals, row[grp])
						smap.Set(grp, uvals)
					}

				}
			}
		}
	}

}

//myMap
func adpro4(smap *MyMap, wg *sync.WaitGroup, rows []map[string]interface{}) {

	defer wg.Done()
	grpFields := []string{"ShipmentTag"}

	for _, row := range rows {

		key := ""
		for i, grp := range grpFields {
			format := ""
			if i == 0 {
				format = "%s%v"
			} else {
				format = "%s_%v"
			}
			key = fmt.Sprintf(format, key, row[grp])
		}

		evals, _ := smap.read(key)
		evals = append(evals, row)
		smap.write(key, evals)
		//fmt.Println(key, isOk, len(evals))
	}

}

func adpro5(smap *cmap.ConcurrentMap, wg *sync.WaitGroup, rows []map[string]interface{}) {

	defer wg.Done()
	grpFields := []string{"ShipmentTag"}

	for _, row := range rows {

		for _, grp := range grpFields {

			if val, isOk := smap.Get(grp); !isOk {
				slc := make([]interface{}, 0)
				vals := append(slc, row[grp])
				//smap.write(grp, vals)
				smap.Set(grp, vals)
				fmt.Println(">", grp, row[grp])

			} else {
				//fmt.Printf("%v %T\n", val, val)
				pvals, isOk := val.([]interface{})
				if isOk {

					isFound := FindExist(pvals, row[grp])
					//fmt.Println(row[grp], isFound, pvals)
					if isFound == false {
						uvals := append(pvals, row[grp])
						smap.Set(grp, uvals)
					}

				}
			}
		}
	}

}

func FindExist(vals []interface{}, val interface{}) bool {

	for _, sval := range vals {
		if sval == val {
			return true
		}
	}
	return false
}
