package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	// Criar arquivo de saída
	file, err := os.Create("links_coletados.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	c := colly.NewCollector(
		colly.AllowedDomains("eteccamargoaranha.cps.sp.gov.br"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		linha := fmt.Sprintf("Link found: %q -> %s\n", e.Text, link)
		file.WriteString(linha)
		fmt.Print(linha)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		linha := fmt.Sprintf("Acessei o link: %s\n", r.URL.String())
		file.WriteString(linha)
		fmt.Print(linha)
	})

	c.Visit("https://eteccamargoaranha.cps.sp.gov.br/")

	fmt.Println("\nDados salvos em links_coletados.txt")
}
