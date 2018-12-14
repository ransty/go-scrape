package main

import (
	"fmt"
	"os"
	"log"
	"flag"

	"github.com/gocolly/colly"
)

func checkDumpFile(inputFile string) {
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Println("File does not exist, generating...")
		f, err := os.Create(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}
}

func main() {
	filePtr := flag.String("file", "/tmp/data", "The file to dump the scraped text to")

	flag.Parse()

	fmt.Println(*filePtr)
	checkDumpFile(*filePtr)
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("keano.me"),
	)

	c.OnHTML("p", func(e *colly.HTMLElement) {
		fmt.Println("Dumping paragraph text to file..")
		f, err := os.OpenFile(*filePtr, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		n, err := f.WriteString(e.Text + "\n")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("wrote %d bytes\n", n)		
		f.Close()
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	
	c.Visit("https://keano.me")
}
