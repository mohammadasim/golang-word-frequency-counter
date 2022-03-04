package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

var filePath = flag.String("fpath", "", "The path of the file to be read.")

func main() {
	var counts = make(map[string]int)
	flag.Parse()
	if *filePath == "" {
		fmt.Println("No file path provided.")
		os.Exit(1)
	}
	f, err := os.Open(*filePath)
	if err != nil {
		fmt.Printf("File read error: %v", err)
		os.Exit(1)
	}
	// Anonymous func to close the file at the end.
	defer func() {
		f.Close()
	}()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	// we write data into a table
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Words\tFrequency\t")
	for c, n := range counts {
		fmt.Fprintln(w, c, "\t", n, "\t")
	}
	w.Flush()

}
