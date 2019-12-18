package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gocolly/colly"
	"github.com/tealeg/xlsx"
)

//FBUser è composto dal nome e dal link del profilo
type FBUser struct {
	Nome string
	Link string
	Data string
}

func main() {
			var file *xlsx.File
    var sheet *xlsx.Sheet
    var row *xlsx.Row
    var cell *xlsx.Cell
		var err error
		
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	c := colly.NewCollector(colly.MaxDepth(2),
		colly.Async(true),)
		c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 20})
	c.WithTransport(t)

	pages := []FBUser{}

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, elem *colly.HTMLElement) {
			nome := elem.Text
			link := elem.Attr("href")
			data := e.DOM.Find(".livetimestamp").Text()
			fmt.Println(data)
			//fmt.Println(link, e.Text)
			fb := FBUser{
					Nome:    nome,
					Link: link,
					Data: data,
				}
				pages = append(pages, fb)
		})
		
	})

	//c.OnHTML("a", func(e *colly.HTMLElement) {
	//	c.Visit("file://" + dir + "/html" + e.Attr("href"))
	//})

	fmt.Println("file://" + dir + "/to_extract.html")
	c.Visit("file://" + dir + "/to_extract.html")
	c.Wait()

	    file = xlsx.NewFile()
    sheet, err = file.AddSheet("Sheet1")
    if err != nil {
        fmt.Printf(err.Error())
		}
		
		row = sheet.AddRow()
    cell = row.AddCell()
		cell.Value = "Name"
		cell = row.AddCell()
		cell.Value = "Added on date"
		cell = row.AddCell()
		cell.Value = "Link on profile"
		


	for i, p := range pages {
		fmt.Printf("%d : %s il %s - %s\n", i, p.Nome, p.Data, p.Link)



    row = sheet.AddRow()
    cell = row.AddCell()
		cell.Value = p.Nome
		cell = row.AddCell()
		cell.Value = p.Data
		cell = row.AddCell()
		cell.Value = p.Link
    
	}

	err = file.Save("List.xlsx")
    if err != nil {
        fmt.Printf(err.Error())
    }
}
