package main

import "fmt";

func main(){
	for i := 0; i < 5; i++ {
		fmt.Println(i);
	}

	for i := range 5 {
		fmt.Println(i);
	}

	b := true;
	c := 0;
	for b {
		if c == 10 {
			b = false;
		}
		c++;
		fmt.Println(c);
	}
}