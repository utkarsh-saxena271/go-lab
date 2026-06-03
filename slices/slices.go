package main
import "fmt";

// Notes
// Slices are similar to arrays, but are more powerful and flexible.
// Like arrays, slices are also used to store multiple values of the same type in a single variable.
// However, unlike arrays, the length of a slice can grow and shrink as you see fit.

// In Go, there are several ways to create a slice:
// - Using the []datatype{values} format
// - Create a slice from an array
// - Using the make() function


// slice_name := []datatype{values}

// Slice made from array:
// myslice := myarray[start:end] 

// slice with make function 
// slice_name := make([]type, length, capacity)


// len() function - returns the length of the slice (the number of elements in the slice)
// cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

// append element
// slice_name = append(slice_name, element1, element2, ...)
// append slice to another slice
// slice3 = append(slice1, slice2...)

func main() {
  myslice1 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  // with omitted capacity
  myslice2 := make([]int, 5)
  fmt.Printf("myslice2 = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice2))
}