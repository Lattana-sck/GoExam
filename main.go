package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	startPort = 1024
	endPort   = 8192
	timeout   = 2 * time.Second
)

var (
	rightPort  int
	userSecret string
	mutex      sync.Mutex
)

func scanPort(ip string, port int, wg *sync.WaitGroup, openPorts chan int) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, timeout)

	if err == nil {
		conn.Close()
		openPorts <- port
		go getPing(ip, port)
	}
}

func getPing(ip string, port int) {
	url := fmt.Sprintf("http://%s:%d/ping", ip, port)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		fmt.Printf("http://10.49.122.144:%d/ping : %s\n", port, string(body))

		mutex.Lock()
		rightPort = port
		mutex.Unlock()
	}
}

func signUp(ip string, port int) {
	url := fmt.Sprintf("http://%s:%d/signup", ip, port)

	data := map[string]string{"User": "Lattana"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erreur lors de la requête POST:", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse POST:", err)
		return
	}

	fmt.Printf("http://10.49.122.144:%d/signup :%s\n", port, string(responseBody))
}

func check(ip string, port int) {
	url := fmt.Sprintf("http://%s:%d/check", ip, port)

	data := map[string]string{"User": "Lattana"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erreur lors de la requête POST:", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse POST:", err)
		return
	}

	fmt.Printf("http://10.49.122.144:%d/check : %s\n", port, string(responseBody))
}

func getUserSecret(ip string, port int) {
	for {
		url := fmt.Sprintf("http://%s:%d/getUserSecret", ip, port)

		data := map[string]string{"User": "Lattana"}
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Erreur lors de la conversion en JSON:", err)
			return
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Erreur lors de la requête POST:", err)
			return
		}
		defer resp.Body.Close()

		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de la réponse POST:", err)
			return
		}

		if string(responseBody) != "Really don't feel like working today huh..." {
			fmt.Printf("http://10.49.122.144:%d/getUserSecret : %s\n", port, string(responseBody))
			mutex.Lock()
			userSecret = string(responseBody)[13:]
			userSecret = strings.TrimRight(userSecret, "\n")
			mutex.Unlock()
			break
		}
	}
}

func getUserLevel(ip string, port int) {

	url := fmt.Sprintf("http://%s:%d/getUserLevel", ip, port)

	data := map[string]string{
		"User":   "Lattana",
		"Secret": userSecret,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erreur lors de la requête POST:", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse POST:", err)
		return
	}

	fmt.Printf("http://10.49.122.144:%d/getUserLevel : %s\n", port, string(responseBody))
}

func getUserPoints(ip string, port int) {

	url := fmt.Sprintf("http://%s:%d/getUserPoints", ip, port)

	data := map[string]string{
		"User":   "Lattana",
		"Secret": userSecret,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erreur lors de la requête POST:", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse POST:", err)
		return
	}

	fmt.Printf("http://10.49.122.144:%d/getUserPoints : %s\n", port, string(responseBody))
}

func getHint(ip string, port int) {

	url := fmt.Sprintf("http://%s:%d/iNeedAHint", ip, port)

	data := map[string]string{
		"User":   "Lattana",
		"Secret": userSecret,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erreur lors de la requête POST:", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse POST:", err)
		return
	}

	fmt.Printf("http://10.49.122.144:%d/iNeedAHint : %s\n", port, string(responseBody))
}

func getChallenge(ip string, port int) {
    url := fmt.Sprintf("http://%s:%d/enterChallenge", ip, port)

	data := map[string]string{
		"User":   "Lattana",
		"Secret": userSecret,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erreur lors de la requête POST:", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse POST:", err)
		return
	}

	fmt.Printf("http://10.49.122.144:%d/enterChallenge : %s\n", port, string(responseBody))
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

	go signUp(ip, rightPort)
	go check(ip, rightPort)

	var secretWg sync.WaitGroup

	secretWg.Add(1)
	go func() {
		defer secretWg.Done()
		getUserSecret(ip, rightPort)
	}()

	secretWg.Wait()

	var levelWg sync.WaitGroup

	levelWg.Add(1)
	go func() {
		defer levelWg.Done()
		getUserLevel(ip, rightPort)
	}()

	levelWg.Wait()

	var pointsWg sync.WaitGroup

	pointsWg.Add(1)
	go func() {
		defer pointsWg.Done()
		getUserPoints(ip, rightPort)
	}()
	pointsWg.Wait()
    
    var hintWg sync.WaitGroup

	hintWg.Add(1)
	go func() {
        defer hintWg.Done()
        getHint(ip, rightPort)
    }()
    hintWg.Wait()

    var GetChallengeWg sync.WaitGroup
    GetChallengeWg.Add(1)
    go func() {
        defer GetChallengeWg.Done()
        getChallenge(ip, rightPort)
    }()
    GetChallengeWg.Wait()

	app := fiber.New()
	app.Listen(":3000")
}
