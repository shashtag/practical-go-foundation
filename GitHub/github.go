package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	fmt.Println(githubInfo(ctx, "tebeka"))
}

// githubInfo returns name and number of public repos for login
func githubInfo(ctx context.Context, login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login)
	// resp,err := http.Get(url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}

	defer resp.Body.Close()

	// fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	// var r Reply

	var r struct { // anonymous struct
		Name string
		// Public_Repos int
		NumRepos int `json:"public_repos"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}

	return r.Name, r.NumRepos, nil
}
