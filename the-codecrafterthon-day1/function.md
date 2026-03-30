# function
A function is a block of statements that can be used repeatedly in a program. function will not execute automatically when a page loads.

A function will be executed by a call to the function.

# Naming Rules for Go Functions

    A function name must start with a letter
    A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
    Function names are case-sensitive
    A function name cannot contain spaces
    If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used

# Go Function Parameters and Arguments
Parameters and Arguments

Information can be passed to functions as a parameter. Parameters act as variables inside the function.

Parameters and their types are specified after the function name, inside the parentheses. You can add as many parameters as you want, just separate them with a comma

# Return Values
If you want the function to return a value, you need to define the data type of the return value (such as int, string, etc), and also use the return keyword inside the function

# Recursion Functions

Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.

In the following example, testcount() is a function that calls itself. We use the x variable as the data, which increments with 1 (x + 1) every time we recurse. The recursion ends when the x variable equals to 11 (x == 11). 