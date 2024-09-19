package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Quote struct {
    Text   string `json:"content"`
    Author string `json:"author"`
}

func fetchQuote() (Quote, error) {
    resp, err := http.Get("https://api.quotable.io/random")
    if err != nil {
        return Quote{}, err
    }
    defer resp.Body.Close()

    var quote Quote
    if err := json.NewDecoder(resp.Body).Decode(&quote); err != nil {
        return Quote{}, err
    }

    return quote, nil
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
    quote, err := fetchQuote()
    if err != nil {
        http.Error(w, "无法获取名言", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "<h1>\"%s\"</h1><p>- %s</p>", quote.Text, quote.Author)
}

func main() {
    http.HandleFunc("/", quoteHandler)
    fmt.Println("服务器正在 http://localhost:8080 运行...")
    http.ListenAndServe(":8080", nil)
}
