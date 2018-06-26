package ch

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"time"
)

type Job struct {
	filename string
	results  chan<- Result
}

type Result struct {
	filename string
	line     string
	lino     int
}

var worker = runtime.NumCPU()

func CHHHHHHHHH() {
	ch := make(chan int, 1)
	tt := time.NewTicker(time.Second * 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		case <-tt.C:
			ch <- 3
		}
		i := <-ch
		fmt.Println("Value received:", i)

	}
}

func GOGOGOGOG() {
	// config cpu number
	runtime.GOMAXPROCS(worker)
	files := os.Args[2:]
	regex, err := regexp.Compile(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	// 任务列表, 并发数目为CPU个数
	jobs := make(chan Job, worker)
	// 结果
	results := make(chan Result, minimum(1000, len(files)))
	// 标记完成
	dones := make(chan struct{}, worker)

	go addJob(files, jobs, results)
	for i := 0; i < worker; i++ {
		go doJob(jobs, regex, dones)
	}
	awaitForCloseResult(dones, results)
}

func addJob(files []string, jobs chan<- Job, results chan<- Result) {
	for _, filename := range files {
		jobs <- Job{filename, results}
	}
	close(jobs)
}

func doJob(jobs <-chan Job, regex *regexp.Regexp, dones chan<- struct{}) {
	for job := range jobs {
		job.Do(regex)
	}
	dones <- struct{}{}
}

func awaitForCloseResult(dones <-chan struct{}, results chan Result) {
	working := worker
	done := false
	for {
		select {
		case result := <-results:
			println(result)
		case <-dones:
			working -= 1
			if working <= 0 {
				done = true
			}
		default:
			if done {
				return
			}
		}
	}
}

func (j *Job) Do(re *regexp.Regexp) {
	f, err := os.Open(j.filename)
	if err != nil {
		println(err)
		return
	}
	defer f.Close()

	b := bufio.NewReader(f)
	lino := 0
	for {
		line, _, err := b.ReadLine()
		if re.Match(line) {
			j.results <- Result{j.filename, string(line), lino}
		}

		if err != nil {
			break
		}
		lino += 1
	}
}

func minimum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func println(o ...interface{}) {
	fmt.Println(o...)
}
