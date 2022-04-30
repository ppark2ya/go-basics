package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id				string
	title 		string
	location 	string
	summary 	string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python";

func main() {
	var jobs []extractedJob
	mainChannel := make(chan []extractedJob)
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		go getPage(i, mainChannel)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-mainChannel
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Location", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobsSlice := []string{
			"https://kr.indeed.com/viewjob?jk=" + job.id,
			job.title,
			job.location,
			job.summary,
		}
		jwErr := w.Write(jobsSlice)
		checkErr(jwErr)
	}
}

func getPage(page int, mainChannel chan<- []extractedJob) {
	var jobs []extractedJob
	channel := make(chan extractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(page * 10)
	fmt.Println("Requesting: ", pageURL)
	res, err := http.Get(pageURL)

	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find("#mosaic-provider-jobcards > a.resultWithShelf")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, channel)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-channel
		jobs = append(jobs, job)
	}

	mainChannel <- jobs
}

func extractJob(card *goquery.Selection, channel chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".companyName").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	summary := cleanString(card.Find(".job-snippet").Text())

	channel <- extractedJob {
		id: id,
		title: title,
		location: location,
		summary: summary,
	}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
