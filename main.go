package main

import (
    "fmt"
    "net"
    "sync"
    "time"
)

const (
    startPort = 1
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
}
