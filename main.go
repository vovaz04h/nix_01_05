package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Post ...
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

const (
	POSTS_COUNT  = 1
	BASE_URL     = "https://jsonplaceholder.typicode.com/posts/"
	STORAGE_PATH = "./storage/posts/"
)

var (
	wg sync.WaitGroup
)

func getPost(id int) {

	defer wg.Done()

	resp, err := http.Get(BASE_URL + strconv.Itoa(id))
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Println(err)
		return
	}

	data, err := json.MarshalIndent(post, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(STORAGE_PATH+strconv.Itoa(id), data, 0644)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {

	for postID := 1; postID <= POSTS_COUNT; postID++ {
		wg.Add(1)
		go getPost(postID)
	}

	wg.Wait()
}
