package main

import (
	//"io"
	"encoding/json"
	"fmt"
	"strings"
)

// func writeReplaced(writer io.Writer, str string, subs ...string) {
// replacer := strings.NewReplacer(subs...)
// replacer.WriteString(writer, str)
// }
func main() {
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProduct{
		Product:  &Kayak,
		Discount: 10.50,
	}
	namedItems := []Named{&dp, &Person{PersonName: "Alice"}}
	encoder.Encode(namedItems)
	fmt.Print(writer.String())
}
