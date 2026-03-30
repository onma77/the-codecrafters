package main
import (
	"fmt"
	"strings"
	//"strconv"
)

func upperAll(p string) string {
	return strings.ToUpper(p)
}

func lowerAll(p string) string{
	return strings.ToLower(p)
}

func capFirst(p string) string{
	return strings.Title(p)
}

func Title(p string) string{
	return strings.Title(p)
}

func reverse(p string) string{
	h := ""
	for i := len(p) -1; i >= 0; i--{
		h += string(p[i])
	}
	return h
}

func snakeCase(p string) string{
	b := strings.Fields(p)
return strings.Join(b, "_")
}
func main() {
	for {
		var k string
		fmt.Println("user input")
		fmt.Scanln(&k)

		// var k2 string
		// fmt.Println("user input2")
		// fmt.Scanln(&k2)

		var operation string
		fmt.Println("operation =\n upperAll,\n lowerAll,\n capFirst,\n Title,\n reverse,\n snakeCase")
		fmt.Scanln(&operation)

		switch operation {

	case "upperAll":

	         fmt.Println(" ")
	         fmt.Println(upperAll(text))

	case "lowerAll":
	         fmt.Println(" ")
	         fmt.Println(lowerAll(text))
	// fmt.Println(capFirst( "director adaeze okonkwo"))
	// fmt.Println(Title( "the fall of the western power grid" ))
	// fmt.Println(snakeCase( "Alert! Level 5 detected."))
	// fmt.Println(reverse("Lagos Nigeria"))
}
}
}