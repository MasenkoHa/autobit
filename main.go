package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"io"
	"fmt"
	"github.com/logrusorgru/aurora"
	"log"
	"regexp"
)


type Movies struct {
	Movie string
	Imdb  string
	Quality string
	Link 	string
}


func main() {
	//instantiate struct
	movieList := make([]Movies, -0)

	//fetch url
	doc, err := getUrl(" ")
	if err != nil {
		log.Fatal(err)
	}
	//test
	fmt.Println(aurora.Cyan(doc))

	doc.Find("body").Find("table").Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		//grab the movie title
		movie := s.Find("td:nth-child(3)").Text()
		//split the title and quality from movie variable
		movieRegex := regexp.MustCompile(`(?P<movie>.+)\s+?\d\d\d\d\s+?(?P<quality>.+?)\s`)
		//grab the link of the torrent
		link, _ := s.Find("td:nth-child(3)").Find("a:nth-child(3)").Attr("href")

		//get imdb rating
		imdb := s.Find("td:nth-child(3)").Find("a:nth-child(7)")
	})


//send mail
	/*m := gomail.NewMessage()
	m.SetHeader("From", " ")
	m.SetHeader("To",  " ")
	m.SetHeader("Subject", fmt.Sprint("Movies Report %s", time.Now().Format("01-02-2006")))
	m.SetBody("text/html", "")
	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}*/

}




func getUrl(url string) (*goquery.Document, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return goquery.NewDocumentFromResponse(resp)
}
func downloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}