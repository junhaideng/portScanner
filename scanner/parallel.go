// 按照顺序进行扫描

package scanner

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func ScanPort(protocol string, hostname string, port int) (Result, bool) {
	// fmt.Printf("scanning port %d \n", port)
	p := strconv.Itoa(port)
	result := Result{Protocol: protocol, Port: p}

	addr := net.JoinHostPort(hostname, p)
	conn, err := net.DialTimeout(protocol, addr, 3*time.Second)
	if err != nil {
		result.State = Closed
		return result, false
	}
	defer conn.Close()
	result.State = Open
	return result, true
}

func ScanHostPortOpen(hostname string, protocol string) []Result {
	var results []Result
	for i := 0; i < PortNum; i++ {
		result, opened := ScanPort(protocol, hostname, i)
		if opened {
			fmt.Printf("%s is opened\n", net.JoinHostPort(hostname, strconv.Itoa(i)))
			results = append(results, result)
		}
	}
	return results
}
