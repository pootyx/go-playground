package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func fetchYoutube(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	data, error := http.Get("https://www.youtube.com/")
	if error != nil {
		log.Fatal()
	}
	html, error := ioutil.ReadAll(data.Body)
	if error != nil {
		log.Fatal()
	}
	c <- string(html)
	return
}

func isThereAnyThe(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case data := <-c:
			fmt.Println(data[0])
			return
		}
	}
}

func main() {
	for i := 1;  i<=5; i++ {
		var wg = &sync.WaitGroup{}
		wg.Add(2)
		c := make(chan string)
		go fetchYoutube(c, wg)
		go isThereAnyThe(c, wg)
		wg.Wait()
	}

}
