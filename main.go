package main

import (
	"fmt"
	"flag"
	"os"
	"encoding/json"
	"github.com/junhaideng/portScanner/scanner"
)

var filename string 
var protocol string 
var hostname string 


func init(){
	flag.StringVar(&filename, "f", "port.json", "filename used to save the port opened information")
	flag.StringVar(&protocol, "p", "tcp", "protocol used to scan port")
	flag.StringVar(&hostname, "h", "", "the target host")
}

func main(){
	flag.Parse()
	if hostname == "" {
		fmt.Println("please specify host !")
		return 
	}
	results := scanner.ScanHostPortOpenAsync(hostname, protocol)
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