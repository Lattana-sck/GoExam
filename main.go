package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
    startPort = 1024
    endPort   = 65535
    timeout   = 2 * time.Second
)

func scanPort(ip string, port int, wg *sync.WaitGroup, openPorts chan int) {
    defer wg.Done()

    address := fmt.Sprintf("%s:%d", ip, port)
    conn, err := net.DialTimeout("tcp", address, timeout)

    if err == nil {
        conn.Close()
        openPorts <- port

        // Appeler getPing pour effectuer une requête HTTP GET
        go getPing(ip, port)
    }
}

func getPing(ip string, port int) {
    url := fmt.Sprintf("http://%s:%d/ping", ip, port)

    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Erreur lors de la requête HTTP GET : %v\n", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return
        }
        fmt.Printf("Réponse de la requête HTTP GET pour le port %d : %s\n", port, string(body))
    }
}

func main() {
    ip := "10.49.122.144"
    var wg sync.WaitGroup

    openPorts := make(chan int)

    for port := startPort; port <= endPort; port++ {
        wg.Add(1)
        go scanPort(ip, port, &wg, openPorts)
    }

    go func() {
        wg.Wait()
        close(openPorts)
    }()

    for openPort := range openPorts {
        fmt.Printf("Port ouvert : %d\n", openPort)
    }

    app := fiber.New()
    app.Listen(":3000") // Vous pouvez utiliser un port différent si nécessaire
}