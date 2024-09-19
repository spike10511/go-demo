package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "log"
    "strings"
)

// fetchTitle 函数用于获取指定URL的网页标题
func fetchTitle(url string, results chan string) {
    resp, err := http.Get(url)
    if err != nil {
        results <- fmt.Sprintf("error: could not fetch %s", url)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        results <- fmt.Sprintf("error: %s returned status %d", url, resp.StatusCode)
        return
    }

    doc, err := html.Parse(resp.Body)
    if err != nil {
        results <- fmt.Sprintf("error: could not parse %s", url)
        return
    }

    var title string
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
            title = n.FirstChild.Data
            return
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)

    if title == "" {
        results <- fmt.Sprintf("error: no title found for %s", url)
    } else {
        results <- fmt.Sprintf("title: %s (from %s)", strings.TrimSpace(title), url)
    }
}

func main() {
    urls := []string{
        "https://golang.org",
        "https://www.google.com",
        "https://www.github.com",
        "https://news.ycombinator.com",
        "https://www.reddit.com",
    }

    results := make(chan string, len(urls))

    // 启动多个goroutine来并行获取网页标题
    for _, url := range urls {
        go fetchTitle(url, results)
    }

    // 收集结果并输出
    for i := 0; i < len(urls); i++ {
        fmt.Println(<-results)
    }
}
