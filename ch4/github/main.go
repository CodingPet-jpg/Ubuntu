package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuseSearchResult struct {
	TotalCount int `json:"total_count"` // 只有名称带下划线的需要元数据标识
	Items      []*Issue
}

type Issue struct {
	Number    int
	Title     string
	State     string
	HtmlUrl   string `json:"html_url"`
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HtmlUrl string `json:"html_url"`
}

func main() {
	var format string = "2006-01-02"
	t := time.Now()
	var oneYearBefore string = t.AddDate(-1, 0, 0).Format(format)
	var halfYearBefore string = t.AddDate(0, -6, 0).Format(format)
	var oneMonthBefore string = t.AddDate(0, -1, 0).Format(format)
	var flag string = os.Args[1:][0]
	var compareF string
	switch flag {
	case "y":
		compareF = oneYearBefore
	case "h":
		compareF = halfYearBefore
	case "m":
		compareF = oneMonthBefore
	default:
		fmt.Println("Invaild argument")
		os.Exit(1)
	}
	var use string = "created:>" + compareF
	var q string = url.QueryEscape(use)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "search query failed: %s\n", resp.Status)
		os.Exit(1)
	}
	var result IssuseSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	resp.Body.Close()
	fmt.Println(result.TotalCount)
	fmt.Println()
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
