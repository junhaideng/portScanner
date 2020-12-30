// 使用并发扫描，不保证结果端口的顺序性

package scanner

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
	"encoding/json"
	"os"
)

var m sync.Mutex

// 扫描端口号，这里只返回打开的端口
func ScanPortAsync(protocol string, hostname string, port int, results *[]Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("scanning port %d \n", port)
	p := strconv.Itoa(port)
	result := Result{Protocol: protocol, Port: p}

	addr := net.JoinHostPort(hostname, p)
	// 继续扫描，如果端口号打开，那么不会返回err
	conn, err := net.DialTimeout(protocol, addr, 3*time.Second)
	if err != nil {
		result.State = Closed
	} else {
		defer conn.Close()
		result.State = Open

		// 加入开放的端口
		m.Lock()
		defer m.Unlock()
		*results = append(*results, result)
	}

}

func ScanHostPortOpenAsync(hostname string, protocol string) []Result {
	var results []Result
	var wg sync.WaitGroup

	for port := 0; port < PortNum; port++ {
		wg.Add(1)
		go ScanPortAsync(protocol, hostname, port, &results, &wg)
	}

	wg.Wait()

	return results
}

func Run(hostname string, protocol string){
	results := ScanHostPortOpenAsync(hostname, protocol)
	fmt.Println("Scan port finished, try to save information")
	file, err := os.Create("port.json")
	if err != nil{
		fmt.Println("create file error: ", err)
		return 
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent(" ", "  ")
	if err := encoder.Encode(results); err != nil{
		fmt.Println("save file error: ", err)
		return
	}
	fmt.Println("save file successfully")
}
