package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

func main() {

	r := gin.Default()
	r.GET("/hello", Server)
	r.Run(":8088")

}

func Server(c *gin.Context) {
	res := LoadBalancer()
	c.JSON(http.StatusAccepted, res)
}

func LoadBalancer() string {
	wg.Add(2)
	ch1 := make(chan string)
	ch2 := make(chan string)
	go Server1(&wg, ch1)

	go Server2(&wg, ch2)
	// Use select to wait for either channel to receive a message
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
		return msg1
	case msg2 := <-ch2:
		fmt.Println(msg2)
		return msg2
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout: No message received in 3 seconds.")
		return "no message recieved"
	}
}

func Server1(wg *sync.WaitGroup, ch1 chan string) {
	fmt.Println("Calling Server 1")
	defer wg.Done()
	// Use a tool like "http.Client" to send requests to the load balancer
	// Adjust the request URL and method based on your actual API endpoints
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8081", nil)
	if err != nil {
		fmt.Println("get Request failed:", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		val := string(body)
		ch1 <- val
	}
}

func Server2(wg *sync.WaitGroup, ch2 chan string) {
	fmt.Println("Calling Server 2")
	defer wg.Done()
	// Use a tool like "http.Client" to send requests to the load balancer
	// Adjust the request URL and method based on your actual API endpoints
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://54.234.134.122:8080  ", nil)
	if err != nil {
		fmt.Println("get Request failed:", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		val := string(body)
		ch2 <- val
	}
}
