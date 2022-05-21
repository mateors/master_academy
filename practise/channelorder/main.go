package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	//wg.Add(1)
	//c1 := make(chan map[string]string, 1) //no wait required
	//go sendMap(c1)
	//go receiveMap(c1)               //finish this
	//fmt.Println("Welcome", len(c1)) //then execute this line

	//fmt.Println(<-c1)
	// for v := range c1 {
	// 	fmt.Println(v)
	// }
	//wg.Wait()

	//wg.Add(5)
	s := time.Now()

	rows := []map[string]interface{}{
		{"name": "Mostain", "age": 38},
		{"name": "Moaz", "age": 25},
		{"name": "Nahid", "age": 26},
		{"name": "Anonnya", "age": 27},
		{"name": "Tareq", "age": 29},
		{"name": "Eshita", "age": 27},
		{"name": "Riyaz", "age": 21},
		{"name": "Pallabi", "age": 21},
		{"name": "Zubair", "age": 22},
		{"name": "Sumi", "age": 30},
		{"name": "Unknown", "age": 0},
		{"name": "Unknown2", "age": 0},
		{"name": "Unknown3", "age": 0},
		{"name": "Unknown4", "age": 0},
		{"name": "Unknown5", "age": 0},
	}

	limit := CalculateLimit(len(rows))
	fmt.Println("CalculateLimit:", limit)
	c1 := make(chan []map[string]interface{}, 5)

	se := calcMaster(len(rows), limit)
	for _, val := range se {
		go rowProcessor(c1, val, rows)

	}

	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	//time.Sleep(time.Millisecond * 100)
	var nmap = make([]map[string]interface{}, 0)
	for range se {
		vals := <-c1
		for _, row := range vals {
			nmap = append(nmap, row)
		}
		fmt.Println("receiver:", len(vals))
	}

	//close(c1)
	fmt.Println("Done!", len(nmap))
	timeTaken := time.Since(s).Milliseconds()
	fmt.Println("TimeTaken:", timeTaken, "ms")

}

func CalculateLimit(totalRow int) int {
	ncpu := runtime.NumCPU()
	res := float64(totalRow) / float64(ncpu)
	return int(math.Ceil(res))
}

//sender
func rowProcessor(c chan<- []map[string]interface{}, sval *startEnd, rows []map[string]interface{}) {

	fmt.Println("sender:", sval.Start, sval.End)
	time.Sleep(time.Millisecond * 2000)
	c <- rows[sval.Start:sval.End]
}

type startEnd struct {
	Start int
	End   int
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

//send only or incoming
func sendMap(cstr chan<- map[string]string) {

	//defer wg.Done()
	time.Sleep(time.Millisecond * 4000)
	nmap := map[string]string{
		"name":      "Mostain",
		"education": "CIS",
	}
	cstr <- nmap
	close(cstr)

}

//receive only
func receiveMap(cstr <-chan map[string]string) {

	defer wg.Done()
	nmap := <-cstr
	for key, val := range nmap {
		fmt.Println(key, val)
	}

}

//send only
func send(cstr chan<- string) {

	time.Sleep(time.Millisecond * 100)
	cstr <- "Hello"

}

//receive only
func receive(cstr <-chan string) {

	msg := <-cstr
	fmt.Println(msg)

}
