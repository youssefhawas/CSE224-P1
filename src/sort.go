package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v inputfile outputfile\n", os.Args[0])
	}

	var infile string = os.Args[1]
	var outfile string = os.Args[2]

	f, err := os.Open(infile)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	records := [][]byte{}
	for {
		buf := make([]byte, 100)
		n, err := f.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		records = append(records, buf[0:n])
	}

	sort.Slice(records, func(i, j int) bool { return bytes.Compare(records[i][:10], records[j][:10]) < 0 })

	f, create_err := os.Create(outfile)
	if create_err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, record := range records {
		_, write_err := f.Write(record)
		if write_err != nil {
			log.Fatal(err)
		}
	}

	f.Sync()
}
