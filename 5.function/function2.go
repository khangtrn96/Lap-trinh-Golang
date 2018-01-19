package main

import (
	"fmt"
)

type yournation struct {
	capital    string
	nation     string
	population string
}

func showyourinfo(identify yournation) {
	fmt.Println(expressinfoagain(identify.capital, identify.nation, identify.population))
}

func expressinfoagain(capitalname, nationname, populationreal string) string {
	return nationname + " có thủ đô là " + capitalname + " với dân số " + populationreal
}

func main() {
	var whereyoulive = yournation{"Ha Noi", "Viet Nam", "90000000"}
	showyourinfo(whereyoulive)
}
