package main

import "fmt"

func main(){
	fmt.Print("Enter a number: ")
	var addNo int
	fmt.Scanln(&addNo)
	for i:=0;i<=10;i++{
		fmt.Printf("%d + %d = %d \n",addNo,i,addNo+i)	
	}


}




