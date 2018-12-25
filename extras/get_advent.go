package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var profiles = map[string]string{
	// "<profile>": "<session cookie>",
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: get_advent <day number from 1 to 25> [<output name>]")
		return
	}

	in := os.Args[1]
	out := in
	if len(os.Args) > 2 {
		out = os.Args[2]
	}

	url := "https://adventofcode.com/2017/day/" + in + "/input"
	fmt.Println("Getting", url)

	c := &http.Client{Timeout: 10 * time.Second}

	for profile, cookie := range profiles {
		fmt.Printf("for %s/%s: ", profile, in)

		req, err := http.NewRequest("GET", url, nil)
		check(err)

		req.AddCookie(&http.Cookie{
			Name:   "session",
			Value:  cookie,
			Path:   "/",
			Domain: ".adventofcode.com",
		})

		resp, err := c.Do(req)
		check(err)
		fmt.Println(resp.Status)
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			f, err := os.Create(out + "_" + profile + ".txt")
			check(err)
			defer f.Close()

			io.Copy(f, resp.Body)
		}
	}

}
