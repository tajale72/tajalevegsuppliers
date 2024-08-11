package loadbalancer

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// Server represents a backend server with its address and port.
type Server struct {
	Address string
	Port    int
}

// LoadBalancer represents a simple load balancer.
type LoadBalancer struct {
	servers []*Server
	mutex   sync.Mutex
}

// NewLoadBalancer creates a new LoadBalancer instance with the given backend servers.
func NewLoadBalancer(servers []*Server) *LoadBalancer {
	return &LoadBalancer{
		servers: servers,
	}
}

// ChooseServer selects a backend server using a simple round-robin algorithm.
func (lb *LoadBalancer) ChooseServer() *Server {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	// Implement a round-robin algorithm
	server := lb.servers[0]
	lb.servers = append(lb.servers[1:], server)
	return server
}

// handleRequest is the Gin handler for incoming requests.
func (lb *LoadBalancer) handleRequest(c *gin.Context) {
	server := lb.ChooseServer()

	// Forward the request to the selected backend server
	targetURL := fmt.Sprintf("http://%s:%d%s", server.Address, server.Port, c.Request.RequestURI)
	resp, err := http.Get(targetURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to forward the request"})
		return
	}
	defer resp.Body.Close()

	// Copy the response from the backend server to the client
	c.Status(resp.StatusCode)
	c.Writer.Write([]byte(fmt.Sprintf("Response from %s: %s", server.Address, resp.Status)))
}

func main() {
	// Define backend servers
	servers := []*Server{
		{Address: "localhost", Port: 8080},
		{Address: "localhost", Port: 8081},
		// Add more servers as needed
	}

	// Create a load balancer
	lb := NewLoadBalancer(servers)

	// Create a Gin router
	router := gin.Default()

	// Handle all incoming requests using the load balancer
	router.Any("/name", lb.handleRequest)

	// Run the Gin server on port 8088
	err := router.Run(":8088")
	if err != nil {
		fmt.Println("Failed to start the server:", err)
	}
}
