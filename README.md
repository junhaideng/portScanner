### 端口扫描工具

````bash
go get github.com/junhaideng/portScanner
````

使用方式
```go
go run main.go -h "127.0.0.1"
```
命令行参数，可以使用 `go run main.go --help` 进行查看
```
-f string
      filename used to save the port opened information (default "port.json")
-h string
      the target host
-p string
      protocol used to scan port (default "tcp")
```