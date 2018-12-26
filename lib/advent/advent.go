package advent

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"
)

// Input is returned.
type Input struct {
	Name, Val string
	Cached    bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Inputs gets input files from cache or from Advent of Code website.
func Inputs(year, day int) <-chan Input {
	ch := make(chan Input)
	go func() {
		url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
		// url := "http://2b.ee/test.html"
		c := &http.Client{Timeout: 10 * time.Second}
		ms, err := filepath.Glob("inputs/*.cookie")
		check(err)
		for _, m := range ms {
			cookie, err := ioutil.ReadFile(m)
			check(err)

			m = filepath.Base(m)
			iext := len(m) - len(filepath.Ext(m))
			m = m[:iext]

			ckey := fmt.Sprintf("inputs/%d_%d_%s.txt", year, day, m)

			cache, err := ioutil.ReadFile(ckey)
			if err == nil {
				fmt.Println("cached", ckey)
				ch <- Input{Name: m, Val: string(cache), Cached: true}
				continue
			}

			fmt.Println("downloading", m, "from", url)
			req, err := http.NewRequest("GET", url, nil)
			req.AddCookie(&http.Cookie{
				Name:   "session",
				Value:  string(cookie),
				Path:   "/",
				Domain: ".adventofcode.com",
			})
			resp, err := c.Do(req)
			check(err)
			defer resp.Body.Close() // needed?

			if resp.StatusCode < 200 || resp.StatusCode > 299 {
				panic(resp.Status)
			}

			buf, err := ioutil.ReadAll(resp.Body)
			check(err)

			if len(buf) > 0 && buf[len(buf)-1] == 10 { // remove trailing newline. AoC bug or what?
				buf = buf[:len(buf)-1]
			}

			err = ioutil.WriteFile(ckey, buf, 600)
			check(err)

			ch <- Input{Name: m, Val: string(buf), Cached: false}
		}
		close(ch)
	}()
	return ch
}
