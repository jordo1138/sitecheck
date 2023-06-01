package main

import (
    "fmt"
    "time"
    "net/http"
    "bufio"
    "os"
	"github.com/montanaflynn/stats"


)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a website URL: ")
    url, _ := reader.ReadString('\n')
    url = url[:len(url)-1]
    
    fmt.Print("How often do you want to check the site (in seconds)? ")
    interval, _ := reader.ReadString('\n')
    interval = interval[:len(interval)-1]
    intervalSec, _ := time.ParseDuration(interval + "s")
    
    fmt.Print("For how long do you want to check the site (in minutes)? ")
    duration, _ := reader.ReadString('\n')
    duration = duration[:len(duration)-1]
    durationMin, _ := time.ParseDuration(duration + "m")
  
	var results []time.Duration
	var results2 []float64 
	var tmpResult float64
	start := time.Now()
    end := start.Add(durationMin)
    
    fmt.Println("Timestamp\tDuration\tStatus Code")
    
    for time.Now().Before(end) {
        start := time.Now()
        resp, err := http.Get(url)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        elapsed := time.Since(start)
        fmt.Printf("%s\t%s\t%d\n", start.Format(time.RFC3339), elapsed, resp.StatusCode)
        time.Sleep(intervalSec)
		results = append(results, elapsed)
		tmpResult = float64(elapsed)
		results2 = append (results2, tmpResult)
    }
	
/*	p50, _ := stats.Percentile(results2, 50)
	p95, _ := stats.Percentile(results2, 95) 
	p99, _ := stats.Percentile(results2, 99)
	
	fmt.Println("\nPercentiles:")
	fmt.Printf("p50: \n", p50)
	fmt.Printf("p95: \n", p95)
	fmt.Printf("p99: \n", p99)
*/
}