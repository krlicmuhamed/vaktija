package main

import (
	"log"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("http://vaktija.ba")
  if err != nil {
    log.Fatal(err)
  }

	datum := doc.Find(".datumi span").Text()
	zora := doc.Find("#danas .td_3").Text()
	izl_sunca := doc.Find("#danas .td_4").Text()
	podne := doc.Find("#danas .td_5").Text()
	ikindija := doc.Find("#danas .td_6").Text()
	aksam := doc.Find("#danas .td_7").Text()
	jacija := doc.Find("#danas .td_8").Text()

	fmt.Printf(">%s\r\n", datum)
	fmt.Printf("Zora 				- %s\r\n", zora)
	fmt.Printf("Izlazak sunca  			- %s\r\n", izl_sunca)
	fmt.Printf("Podne 				- %s\r\n", podne)
	fmt.Printf("Ikindija 			- %s\r\n", ikindija)
	fmt.Printf("Aksam 				- %s\r\n", aksam)
	fmt.Printf("Jacija 				- %s\r\n", jacija)
}
