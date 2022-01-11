package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	NEWS_URL       = "https://saurav.tech/NewsAPI/everything/cnn.json"
	QUOTE_URL      = "https://www.affirmations.dev"
	JOKE_URL       = "https://icanhazdadjoke.com"
	MAX_NEWS_ITEMS = 3
)

type content struct {
	body string
	err  error
}

func getNews(newsChan chan<- content) {
	var parsedBody struct {
		Articles []struct {
			Title  string `json:"title"`
			Author string `json:"author"`
			Desc   string `json:"description"`
			Url    string `json:"url"`
		} `json:"articles"`
	}

	err := fetchContent(NEWS_URL, &parsedBody)
	if err == nil && len(parsedBody.Articles) == 0 {
		err = fmt.Errorf("empty news response")
	}

	if err != nil {
		newsChan <- content{err: fmt.Errorf("error in retrieving news: %s", err.Error())}
		return
	}

	newsItems := make([]string, 0, MAX_NEWS_ITEMS)
	for i := 0; i < MAX_NEWS_ITEMS; i++ {
		article := parsedBody.Articles[i+1]
		newsItem := fmt.Sprintf("%s\n- %s\n\n%s\n\n%s\n------", article.Title, article.Author, article.Desc, article.Url)
		newsItems = append(newsItems, newsItem)
	}

	newsBody := fmt.Sprintf("*** News of the day ***\n\n%s\n", strings.Join(newsItems, "\n\n"))
	newsChan <- content{body: newsBody}
}

func getQuote(quoteChan chan<- content) {
	var parsedBody struct {
		Affirmation string `json:"affirmation"`
	}

	err := fetchContent(QUOTE_URL, &parsedBody)
	if err == nil && parsedBody.Affirmation == "" {
		err = fmt.Errorf("empty quote response")
	}

	if err != nil {
		quoteChan <- content{err: fmt.Errorf("error in retrieving quote: %s", err.Error())}
		return
	}

	quote := fmt.Sprintf("*** Quote of the day ***\n%q\n", parsedBody.Affirmation)
	quoteChan <- content{body: quote}
}

func getJoke(jokeChan chan<- content) {
	var parsedBody struct {
		Joke string `json:"joke"`
	}

	err := fetchContent(JOKE_URL, &parsedBody)
	if err == nil && parsedBody.Joke == "" {
		err = fmt.Errorf("empty joke response")
	}

	if err != nil {
		jokeChan <- content{err: fmt.Errorf("error in retrieving joke: %s", err.Error())}
		return
	}

	joke := fmt.Sprintf("*** Joke of the day ***\n%q\n", parsedBody.Joke)
	jokeChan <- content{body: joke}
}

func fetchContent(url string, parsedBody interface{}) error {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error in fetching data: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error in reading response body: %s", err)
	}

	err = json.Unmarshal(body, parsedBody)
	if err != nil {
		return fmt.Errorf("Error in parsing body: %s", err)
	}

	return nil
}

func main() {
	newsChan := make(chan content)
	quoteChan := make(chan content)
	jokeChan := make(chan content)

	go getNews(newsChan)
	go getQuote(quoteChan)
	go getJoke(jokeChan)

	newsContent := <-newsChan
	quoteContent := <-quoteChan
	jokeContent := <-jokeChan

	fmt.Printf("Hello! Here's your briefing:\n\n")
	for _, content := range []content{newsContent, quoteContent, jokeContent} {
		if content.err != nil {
			fmt.Println(content.err)
		} else {
			fmt.Println(content.body)
		}
		fmt.Println()
	}
}
