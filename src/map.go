package main

func main() {
	var idMap map[int]string     // Nil Map
	idMap = make(map[int]string) // Init Map

	idMap[901] = "Apple"
	idMap[134] = "Grape"
	idMap[532] = "Tomato"
	// str = idMap[532]
	// println(str)

	// noData := idMap[999]
	// println(noData)
	// println(noData == "")

	delete(idMap, 134)
	// fmt.Println(idMap)

	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"AP":   "Apple",
		"MT":   "Meta",
	}
	//str := tickers["AP"]
	//println(str)

	v, est := tickers["AP"]
	if !est {
		println("No AP ticker", v)
	}
	// println(v, est)

	for key, val := range tickers {
		// fmt.Println(key, val)
		println(key, val)
	}

}
