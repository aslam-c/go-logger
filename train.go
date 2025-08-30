package main

import (
	"errors"
	"fmt"
)

type car struct{
	make int
	company string
}

type truck struct{
	wheelbase int
	car
}

type cart interface{
	getMakeYear() string
}

func (t truck) getMakeYear() string{
	return "Year "+ t.company
}

type nameError struct{
	name string
}

func (ne nameError) Error(name string) string{
	return fmt.Sprintf("%v is too long")
}


func add(a,b int) (int,error){
	if(a==0||b==0){
		return 0,errors.New("Not Zero")
	}
	return a+b,nil
}


func concatter() func(string)string{
	defer fmt.Println("++++++++++++++")
	sumstring:=""
	return func (adder string) string{
	defer fmt.Println("-----------")
		sumstring+=adder+" "
		return sumstring
	}
}



func main(){
	concatterFun:=concatter()
	arrt:=[]string{"1","2","3"}
	arrt=append(arrt,"3")
	for _,item:=range arrt{
		concatterFun(item)
	}
	fmt.Println("String is ",concatterFun("111"))			

}




