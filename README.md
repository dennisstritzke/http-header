[![Documentation](https://godoc.org/github.com/dennisstritzke/httpheader?status.svg)](https://godoc.org/github.com/dennisstritzke/httpheader)

# HTTP Header
Golang library to bind HTTP headers into struct values.

## Example
```
package main

import (
	"fmt"
	"github.com/dennisstritzke/httpheader"
	"net/http"
)

type GitHubResponse struct {
	Date string `header:"Date"`
}

func main() {
	resp, _ := http.Get("https://github.com/")

	var githubResponse GitHubResponse
	err := httpheader.Bind(resp.Header, &githubResponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Date: %s\n", githubResponse.Date)
}
```