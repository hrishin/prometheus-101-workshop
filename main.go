/*
Read evaluate print loop

Program keeps accepting the text input from STDIN,
process it (transform the text to upper case) and print
it on STDOUT

credit: opencensus.io
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	jobCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "repl_processed_jobs_total",
		Help: "The total number of processed job",
	})

	lastDataSize = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "repl_processed_jobs_last_data",
		Help: "The last job processed data size",
	})

	dataSizePerJob = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "repl_processed_jobs_data",
			Help: "The data size of job",
		},
		[]string{"job_name"},
	)
)

func main() {
	// expose HTTP endpoint for metrics
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":5050", nil)
	}()

	//   1. Read input
	//   2. process input
	br := bufio.NewReader(os.Stdin)

	// repl is the read, evaluate, print, loop
	for {
		if err := readEvaluateProcess(br); err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}

		// count processed jobs
		jobCounter.Inc()
	}
}

// readEvaluateProcess reads a line from the input reader and
// then processes it. It returns an error if any was encountered.
func readEvaluateProcess(br *bufio.Reader) (terr error) {
	fmt.Printf("> ")
	line, _, err := br.ReadLine()
	if err != nil {
		return err
	}

	out, err := processLine(line)
	if err != nil {
		return err
	}
	fmt.Printf("< %s\n\n", out)
	return nil
}

// processLine takes in a line of text and
// transforms it. Currently it just capitalizes it.
func processLine(in []byte) (out []byte, err error) {
	size := float64(len(in))
	lastDataSize.Set(size)
	dataSizePerJob.With(prometheus.Labels{"job_name": string(in)}).Set(size)
	return bytes.ToUpper(in), nil
}
