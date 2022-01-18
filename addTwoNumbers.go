package main

import "fmt"

func main(){
	fmt.Print("Enter First number: ")
	var firstNo int
	fmt.Scanln(&firstNo)
  fmt.print("Enter Second number")
  var secondNo int
  fmt.Scanln(&secondNo)
	
	fmt.Printf("%d + %d = %d \n",firstNo,secondNo,firstNo+secondNo)
}
