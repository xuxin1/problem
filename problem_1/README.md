# problem

## 文件

* go_server.go `gin 框架的server`
* py_server.py `flask 框架的server`
* http.go   `模拟并发客户端`


## 运行

`cd ~/problem_1`

```bash
go mod tidy
go run go_server.go
go run http.go 500  # 500 并发
```





