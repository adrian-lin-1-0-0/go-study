package synchronouscommunication

import (
	"fmt"
	"net/http"
	"testing"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	c <- fmt.Sprintf("%s is up!", link)
}

func TestBestPractice(t *testing.T) {

	links := []string{
		"https://pkg.go.dev/bufio@go1.22.4",
		"https://pkg.go.dev/bytes@go1.22.4",
		"https://pkg.go.dev/crypto@go1.22.4",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}
