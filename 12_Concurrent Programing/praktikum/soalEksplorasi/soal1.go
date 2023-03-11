// saya memodifikasi soal menjadi 4 sub url, sehingga ada 4 go routine yang di jalankan
// karena jika hanya 1 url, sehingga hanya 1 go routine yang dijalankan, maka keuntungan dari penggunaan go routine kurang didapatkan

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type FakeStore struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      struct {
		Rate  float64 `json:"rate"`
		Count int     `json:"count"`
	} `json:"rating"`
}

func main() {
	ch := make(chan FakeStore)
	urls := []string{
		"https://fakestoreapi.com/products/1",
		"https://fakestoreapi.com/products/2",
		"https://fakestoreapi.com/products/3",
		"https://fakestoreapi.com/products/4",
	}
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go getResponse(ch, &wg, url)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	
	fmt.Println("PRODUCT DATA")
	for i:=0; i<len(urls);i++{
		result:=<-ch
		fmt.Println("===")
		fmt.Printf("title: %s\n",result.Title)
		fmt.Printf("price: %f\n",result.Price)
		fmt.Printf("category: %s\n",result.Category)
		fmt.Println("===")
	}
}

func getResponse(channel chan FakeStore, wg *sync.WaitGroup, url string) {
	defer wg.Done()
	respone, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	responseData, err := ioutil.ReadAll(respone.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject FakeStore
	json.Unmarshal(responseData, &responseObject)
	channel <- responseObject
}
