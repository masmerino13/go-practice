package main

import "fmt"

type Vertexx struct {
	Lat, Long float64
}

var m map[string]Vertexx

func ShowMap() {
	m = make(map[string]Vertexx)
	m["Bell Labs"] = Vertexx{
		Lat:  40.984848,
		Long: 7.932434,
	}
	m["Otro"] = Vertexx{9.213, 43.9348}

	fmt.Println(m, m["Otro"].Lat)
}
