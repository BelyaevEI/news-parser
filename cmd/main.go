package main

import "github.com/BelyaevEI/news-parser/internal/initialization"

func main() {

	services := initialization.New()
	services.RunServices()

}
