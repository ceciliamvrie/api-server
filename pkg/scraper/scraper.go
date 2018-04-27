package scraper

import (
	"log"
	"regexp"
	"strconv"

	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/techmexdev/lineuplist"
)

func Festivals() ([]lineuplist.Festival, error) {
	ff := []lineuplist.Festival{}

	c := colly.NewCollector(colly.Async(true))
	wg := sync.WaitGroup{}

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("colly response error: %s", err)
	})

	c.OnHTML(".festivaltitle", func(e *colly.HTMLElement) {
		c.Visit(e.ChildAttr("a", "href"))
	})

	scrapeONFestivalPage(c, &ff, &wg)

	pc := pageCount()
	for i := 1; i <= pc; i++ {
		pageLink := "https://www.musicfestivalwizard.com/all-festivals/page/" + strconv.Itoa(i) + "/?festival_guide=us-festivals"
		log.Println("pageLink: ", pageLink)
		go c.Visit(pageLink)
	}

	c.Wait()

	return ff, nil
}

func pageCount() int {
	var lastPageNum int

	c := colly.NewCollector()

	c.OnHTML("ul.page-numbers", func(e *colly.HTMLElement) {
		pageNumEls := e.DOM.Find("a.page-numbers")
		var err error
		lastPageNum, err = strconv.Atoi(pageNumEls.Eq(pageNumEls.Length() - 2).Text())
		if err != nil {
			log.Fatalf("could not parse number of pages: %s", err.Error())
		}
	})

	c.Visit("https://www.musicfestivalwizard.com/all-festivals/?festival_guide=us-festivals")
	c.Wait()

	return lastPageNum
}

func scrapeONFestivalPage(c *colly.Collector, ff *[]lineuplist.Festival, wg *sync.WaitGroup) {
	c.OnHTML("#inner-wrap", func(e *colly.HTMLElement) {
		wg.Add(1)
		var f lineuplist.Festival

		text := e.DOM.Find("#festival-basics").Text()
		parsedText := regexp.MustCompile(`WHERE:(\w+,\s\w.)\sWHEN:(.*)`).FindStringSubmatch(text)
		if len(parsedText) != 3 {
			return
		}

		f.Name = e.ChildText("h1.entry-title span")

		var err error
		f.StartDate, f.EndDate, err = parseDate(parsedText[2])
		if err != nil {
			log.Println("error parsing date for: ", f.Name, err)
			return
		}

		f.Country, f.State, f.City, err = parseLocation(parsedText[1])
		if err != nil {
			log.Println("error parsing location for ", f.Name, err)
			return
		}

		aNames := e.DOM.Find(".lineupguide li").Map(func(_ int, s *goquery.Selection) string {
			return s.Text()
		})
		for _, name := range aNames {
			f.Lineup = append(f.Lineup, lineuplist.Artist{Name: name})
		}

		if len(f.Lineup) == 0 {
			return
		}

		log.Printf("scraped: %s\n", f.Name)
		*ff = append(*ff, f)
		wg.Done()
	})
}
