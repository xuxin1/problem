package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var NUM = 200

func main() {

	NUM, _ = strconv.Atoi(os.Args[1])
	waitgroup := sync.WaitGroup{}
	start := time.Now()

	for i := 0; i < NUM; i++ {
		go ask(i, &waitgroup)
	}
	cost := time.Since(start)
	waitgroup.Wait()
	fmt.Printf("%s cost :%s\t", NUM, cost)

}
func ask(i int, waitgroup *sync.WaitGroup) {
	waitgroup.Add(1)
	defer func() {
		fmt.Printf("%d 号连接关闭！\n", i)
		waitgroup.Done()
	}()
	targetUrl := "http://localhost:8081/ping" // gin 接口
	// targetUrl := "http://localhost:8082/"     // flask 接口
	_, err := http.Get(targetUrl)
	if err != nil {
		fmt.Printf("%d  请求失败！，%v\n", i, err)
	}
}
