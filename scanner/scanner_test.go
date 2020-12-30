package scanner

import (
	"fmt"
	"os"
	"testing"
	"encoding/json"
)

func TestScanPort(t *testing.T){
	fmt.Println(ScanPort(TCP, "127.0.0.1",443))
	fmt.Println(ScanPort(UDP, "127.0.0.1", 443))
}

func TestScanHost(t *testing.T){
	data := ScanHostPortOpen("127.0.0.1", "tcp")
	file, err := os.Create("port.json")
	if err != nil{
		t.Fatal(err)
		return 
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent(" ", "  ")
	encoder.Encode(data)
}

func TestScanHostAsync(t *testing.T){
	results := ScanHostPortOpenAsync("182.61.200.7", TCP)
	fmt.Println(results)
	fmt.Println(len(results))
}

func TestRun(t *testing.T){
	Run("127.0.0.1", "tcp")
}