# GO comments
A comment is a text that is ignored when it is execution, used to explain the code, and to make it more readable. it can also be used to prevent code execution when testing an alternative code. 
Go also supports single-line or multi-line comments.


# types of comment
1. single line comment : the forward slash "//". any code between the comment and end of line will not be executed.

```go
// This is a comment
package main
import ("fmt")

func main() {
  // This is a comment
  fmt.Println("Hello World!")
}
```

2. multiple line comment: starts with the /* and ends with */, any code in between will be ignored and not executed.

```go 
package main
import ("fmt")

func main() {
  /* The code below will print Hello World
  to the screen, and it is amazing */
  fmt.Println("Hello World!")
} 
```