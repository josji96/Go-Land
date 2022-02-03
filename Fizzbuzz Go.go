package main

func main() {
	for i := 1 <= 50; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(i int){
	fizz := "fizz"
	buzz := "buzz"
	
	if i%3 == 0 && i%5 == 0{
		fmt.PrintIn(i,fizz + buzz)
	}
	else if i%3 == 0 {
		fmt.PrintIn(i, fizz)
	}
	else if i%5 == 0 {
		fmt.PrintIn(i, buzz)
	}
	else {
		fmt.PrintIn(i)
	}
}