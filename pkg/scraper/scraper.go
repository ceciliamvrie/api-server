package scraper

import (
	"regexp"

	"github.com/techmexdev/lineuplist/pkg/model"

	"github.com/gocolly/colly"
)

// GetFestivals scrapes and returns festivals
func GetFestivals() ([]model.Festival, error) {
	c := colly.NewCollector()
	fests, err := scrapeFests(c)
	if err != nil {
		return []model.Festival{{}}, err
	}
	c.Visit("https://www.musicfestivalwizard.com/festival-guide/us-festivals/")
	c.Wait()

	return fests, nil
}

func scrapeFests(c *colly.Collector) ([]model.Festival, error) {
	var fests = make([]model.Festival, 18, 18)
	var err error
	var nameReg *regexp.Regexp
	index := 0

	c.OnHTML(".singlefestlisting", func(e *colly.HTMLElement) {
		nameReg, err = regexp.Compile("(.*[^ \\d])")
		name := nameReg.FindString(e.ChildText(".festivaltitle"))
		fests[index] = model.Festival{Name: name}
		index++
		link := e.ChildAttr("a", "href")
		err = c.Visit(link)
	})

	if err != nil {
		return nil, err
	}

	return fests, nil
}
