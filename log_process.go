package main

import (
	"time"
	"fmt"
)

type LogProcess struct {
	path string // read file path
	influxDBDsn string // influxDB data source
	rc chan string // from read to process
	wc chan string // from process to write
}

// use pointer to avoid copy & save memory
func (l * LogProcess) ReadFromFile() {

}


func (l * LogProcess) Process() {
	
}

func (l * LogProcess) WriteToInfluxDB() {
	fmt.Println("Write To InfluxDB")
}

func main() {
	lp := &LogProcess{
		rc: make(chan string),
		wc: make(chan string),
		path: "/tmp/access.log",
		influxDBDsn: "username&password..",
	}
	// 併發執行
	// use `lp` instead of `*lp` to call a member function is available in Go
	go lp.ReadFromFile()
	go lp.Process()
	go lp.WriteToInfluxDB()

	time.Sleep(1 * time.Second)
}
