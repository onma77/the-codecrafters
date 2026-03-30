# what is variable
variable is like a container that holds or better still stores value

# Go Variable Types
1. int: deals with integer

2. boolean: stores values with two states: true or false

3. float: stores numbers with decimals like 19.99 or -19.99

4. string: stores text like "Hello Peace"

# Declaring (Creating) Variables
1. With the var keyword:(var p string = "student")

2. With the := sign:(p := "student")

 It is not possible to declare a variable using :=, without assigning a value to it.

# Variable Declaration With Initial Value:
If the value of a variable is known from the start, you can declare the variable and assign a value to it on one line:
```go
package main
import ("fmt")

func main() {
  var student1 string = "onma" //type is string
  var student2 = "peace" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
}
```
# Variable Declaration Without Initial Value
In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type: 
```go
package main
import ("fmt")

func main() {
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```
* These variables are declared but they have not been assigned initial values.

# Value Assignment After Declaration
It is possible to assign a value to a variable after it is declared.
```go
package main
import ("fmt")

func main() {
  var student1 string
  student1 = "John"
  fmt.Println(student1)
}
```
# Difference Between var and :=
var Can be used inside and outside of functions.
while
:= Can only be used inside functions.

# Go Variables:
 Use var name type = value. Inside functions, you can use the shorthand name := value.

# Declare Multiple Variables: 
You can group them: var a, b, c = 1, 2, 3.

# Naming Rules:
 Names must start with a letter. Case matters: MyVar is different from myVar.