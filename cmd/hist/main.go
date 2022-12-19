package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"flag"
	"gostats"
)


func main() {

	flag.Usage = func() {
		w := flag.CommandLine.Output()

		fmt.Fprintf(w, "Usage of %s:\n", os.Args[0])

		flag.PrintDefaults()

		fmt.Fprintf(w, "hist plots an histogram of data passed from standard input. \n")
	}

	nbins := flag.Int("nbins", 10, "number of intervals")
	flag.Parse()

	
	scanner := bufio.NewScanner(os.Stdin)
	data, err := gostats.ReadFloats(scanner)
	if err != nil {
		log.Fatal(err)
	}
	
	min, max := gostats.GetMinMax(data)

	breaks := gostats.Linspace(math.Floor(min),
		math.Floor(max)+1,
		*nbins)
	h := gostats.NewHistogram(breaks)
	h.AddValues(data)

	midp := gostats.GetMidPoints(h.Breaks)
	_, maxc := h.GetMinMaxCounts()
	var scale float64
	if maxc < 80 {
		scale = 1.0
	} else {
		scale = 80.0 / float64(maxc)
	}

	for i := 0; i < *nbins; i++ {
		fmt.Printf("%8.2f %4.0v ", midp[i], h.Counts[i])
		for j := 0; j < int(float64(h.Counts[i]) * scale); j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

}
