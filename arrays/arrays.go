package main

import "fmt";

// Notes

// Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

// var array_name = [length]datatype{values} 
// var array_name = [...]datatype{values}
// array_name := [length]datatype{values} 
// array_name := [...]datatype{values}

// length of array ->     len(Array);
// initialize only specific elements :    [length]datatype{index:val, index:val};

func main(){
	var arr1 = [4]int{1,2,3,4};
	var arr2 = [...]int{9,8,7};
	arr3 := [2]bool{true, false};
	arr4 := [...]string{"hello", "world"};

	fmt.Println(arr1);
	fmt.Println(arr2);
	fmt.Println(arr3);
	fmt.Println(arr4);

	// length of array
	var a int = len(arr1);
	fmt.Println(a);

	// access array elements
	fmt.Println(arr1[0]);
  	fmt.Println(arr1[2]);

	// initialize only specific elements
	arr5 := [5]int{1:10,2:40};
	fmt.Println(arr5);
}
