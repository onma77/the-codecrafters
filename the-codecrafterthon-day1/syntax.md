# What is a GO syntax?
A go syntax is how you write GO program for the computer to understand, more like rules guilding you on how to write your program 

# Examples:
1. package declaraton: In GO every program must be in a package. let's say it belongs to main.

2. Import: this helps us imports files.

3. blankline: we leave a blank line to make our code clean and readable.

4. func main: is a function that execute any code inside it's curly braces

5. fmt.Println():  the fmt is package is for input/output, while the Println() → prints text + new line

# Go statement.
### fmt.Println("Hello World!") is a statement.
In Go, statements are separated by ending a line (hitting the Enter key) or by a semicolon ";".

Hitting the Enter key adds ";" to the end of the line implicitly (does not show up in the source code).

The left curly bracket { cannot come at the start of a line.

# Go Compact Code:
You can write more compact code, like shown below (this is not recommended because it makes the code more difficult to read):

# Example
package main; import ("fmt"); func main() { fmt.Println("Hello World!");}