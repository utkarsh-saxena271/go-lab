package main
import "fmt";

func sum (a int, b int) int {
	return a+b;
}
func sub (a int, b int) (result int) {
	result = a-b;
	return;
}

func main(){
	fmt.Println(sum(2,3));
	fmt.Println(sub(12,3));
}