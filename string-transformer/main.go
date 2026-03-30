package main
import (
	"fmt"
	"strings"
	//"strconv"
)

func upperAll(p string) string {
	return strings.ToUpper(p)
}

func lowerAll(d string) string{
	return strings.ToLower(d)
}

func capFirst(g string) string{
	return strings.Title(g)
}

func Title(o string) string{
	return strings.Title(o)
}

func reverse(m string) string{
	h := ""
	for i := len(m) -1; i >= 0; i--{
		h += string(m[i])
	}
	return h
}

func snakeCase(r string) string{
	b := strings.Fields(r)
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
		fmt.Println("operation: \n1. upperAll \n2. lowerAll \n3. capFirst \n4. Title \n5. reverse \n6. snakeCase")
		fmt.Scanln(&operation)

		switch operation {

	case "upperAll":

	         fmt.Println(upperAll(""))
	         fmt.Scanln(upperAll("text"))

	case "lowerAll":
	         fmt.Println(lowerAll(""))
	         fmt.Scanln(lowerAll("text"))

	case "capFirst":
			fmt.Println(capFirst( ""))
	        fmt.Println(capFirst( "text"))


	case "Title":
	fmt.Println(Title(""))
	fmt.Println(Title("text"))

	case "snakeCase":
	 fmt.Println(snakeCase(""))
	 fmt.Println(reverse("text"))
}
}
}