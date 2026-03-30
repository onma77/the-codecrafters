# slice
are also used to store multiple values of the same type in a single variable. However, unlike arrays, the length of a slice can grow and shrink as you see fit.


* A common way of declaring a slice is like this:
myslice := []int{}; this is empty

* To initialize the slice during declaration, use this:
myslice := []int{1,2,3}


# modifing slice
You can access a specific slice element by referring to the index number.
In Go, indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.