package main

import "fmt"

func main(){
	fmt.Print("Enter table number: ")
	var tableNo int
	fmt.Scanln(&tableNo)
	fmt.Println(" ")
	for i:=0;i<=10;i++{
		fmt.Printf("%d * %d = %d \n",tableNo,i,tableNo*i)
	}    
	
}




