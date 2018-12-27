package source

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	. "github.com/logrusorgru/aurora"
)

// Input values.
type Input struct {
	Name, Val string
	parts     int      // part 1, 2 or 1+2.
	ex        []string // expected results.
	year, day int
}

// Inputs returned by all getters.
type Inputs []Input

// Test values before real inputs.
func Test(part int, in string, ex ...string) Inputs {
	return Inputs{Input{Name: "test0", Val: in, parts: part, ex: ex}}
}

// Test values before real inputs.
func (ins Inputs) Test(part int, in string, ex ...string) Inputs {
	if part&1 == 0 { // test for part 2 only.
		ex = append([]string{""}, ex...)
	}
	return append(ins, Input{Name: "test" + strconv.Itoa(len(ins)), Val: in, parts: part, ex: ex})
}

// Advent sources.
func Advent(year, day int) <-chan Parser {
	return getInputs(year, day)
}

// Advent sources.
func (ins Inputs) Advent(year, day int) <-chan Parser {
	ch := make(chan Parser)
	go func() {
		for _, in := range ins {
			ch <- Parser{in}
		}
		for p := range getInputs(year, day) {
			ch <- p
		}
		close(ch)
	}()
	return ch
}

// Part filter.
func (p Parser) Part(part int) bool {
	return p.parts&part != 0
}

// Result without submitting.
func (p Parser) Result(part int, v string) {
	prefix := fmt.Sprintf("part%d:", part)
	if p.Part(part) && part <= len(p.ex) {
		ok := Green("✓").String()
		if v != p.ex[part-1] {
			ok = Red("✗ expected " + p.ex[part-1]).String()
		}
		fmt.Println(prefix, Cyan(v), ok)
		return
	}
	fmt.Println(prefix, Green(v))
}

// ResultInt without submitting.
func (p Parser) ResultInt(part, n int) {
	p.Result(part, strconv.Itoa(n))
}

// Submit result.
func (p Parser) Submit(part int, v string) {
	p.Result(part, v)
	trySubmit(p.Name, p.year, p.day, part, v)
}

// SubmitInt result.
func (p Parser) SubmitInt(part, n int) {
	p.Submit(part, strconv.Itoa(n))
}

// Submit1 part 1.
func (p Parser) Submit1(v string) {
	p.Submit(1, v)
}

// Submit2 part 2.
func (p Parser) Submit2(v string) {
	p.Submit(2, v)
}

// SubmitInt1 part 1.
func (p Parser) SubmitInt1(n int) {
	p.Submit(1, strconv.Itoa(n))
}

// SubmitInt2 part 2.
func (p Parser) SubmitInt2(n int) {
	p.Submit(2, strconv.Itoa(n))
}

////////////////////////////////////////////////////////////////////////
// Implementation
//

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// getInputs gets input files from cache or from Advent of Code website.
func getInputs(year, day int) <-chan Parser {
	ch := make(chan Parser)
	go func() {
		urlGet := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

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
				//fmt.Println("cached", ckey)
				ch <- Parser{Input{Name: m, Val: string(cache), parts: 1 + 2, year: year, day: day}}
				continue
			}

			fmt.Println("downloading", m, "from", urlGet)
			req, err := http.NewRequest("GET", urlGet, nil)
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

			ch <- Parser{Input{Name: m, Val: string(buf), parts: 1 + 2, year: year, day: day}}
		}
		close(ch)
	}()
	return ch
}

var lastSubmit time.Time

const submitThrottle time.Duration = 5 * time.Second

func trySubmit(name string, year, day, part int, v string) {
	ckey := fmt.Sprintf("results/%d_%d_%d_%s.txt", year, day, part, name)
	result, err := ioutil.ReadFile(ckey)
	if err == nil {
		fmt.Println("already submitted:", Green(string(result)))
		return
	}

	cookie, err := ioutil.ReadFile("results/" + name + ".cookie")
	if err != nil { // no cookie -> no submit.
		//fmt.Println("skipping")
		return
	}

	urlPost := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)
	fmt.Println("submitting to:", urlPost, "for", name)

	wait := submitThrottle - time.Since(lastSubmit)
	if wait > 0 {
		fmt.Println("waiting", Cyan(wait))
		time.Sleep(wait)
	}
	data := url.Values{}
	data.Set("level", strconv.Itoa(part))
	data.Set("answer", v)
	encoded := data.Encode()
	req, err := http.NewRequest("POST", urlPost, strings.NewReader(encoded))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.AddCookie(&http.Cookie{
		Name:   "session",
		Value:  string(cookie),
		Path:   "/",
		Domain: ".adventofcode.com",
	})

	c := &http.Client{Timeout: 10 * time.Second}
	resp, err := c.Do(req)
	check(err)
	defer resp.Body.Close() // needed?

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		panic(resp.Status)
	}

	lastSubmit = time.Now()

	fmt.Println("status:", resp.Status)
	buf, err := ioutil.ReadAll(resp.Body)
	check(err)
	html := string(buf)
	reg := regexp.MustCompile("(?s)<main>\\s*<article><p>(.*)</p></article>\\s*</main>")
	m := reg.FindStringSubmatch(html)
	main := html
	if len(m) > 1 {
		main = m[1]
		fmt.Println("submatch:", main)
	}
	if strings.Contains(main, "You don't seem to be solving the right level.") {
		fmt.Println("Already submitted.")
		ioutil.WriteFile(ckey, []byte("Unknown value"), 600)
		return
	}
	if strings.Contains(main, "That's the right answer!") {
		fmt.Println(Green("Correct answer."))
		ioutil.WriteFile(ckey, []byte(v), 600)
		return
	}
	if strings.Contains(main, "That's not the right answer") {
		fmt.Println(Red("Incorrect answer."))
		if strings.Contains(main, "your answer is too low") {
			fmt.Println(Red("- too low."))
		} else if strings.Contains(main, "your answer is too high") {
			fmt.Println(Red("- too high."))
		}
		ioutil.WriteFile(ckey+".err.txt", []byte(main), 600)
		return
	}
	if strings.Contains(main, "You gave an answer too recently;") {
		fmt.Println(Brown("Submitting too soon. Wait some more."))
		ioutil.WriteFile(ckey+".err.txt", []byte(main), 600)
		return
	}
	ioutil.WriteFile(ckey+".err.txt", []byte(main), 600)
	fmt.Println("main:", main)
}
