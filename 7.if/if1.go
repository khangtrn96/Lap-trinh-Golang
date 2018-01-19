package main

import (
	"./customif1"
)
func main() {
	customif1.whereyoulive := customif1.Salutation{"Việt Nam", "Hồ Chí Minh"}
	customif1.showresult(whereyoulive, customif1.Printcustom(" @_@"))
}
