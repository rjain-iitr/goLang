package main
import (
"fmt"
) 


func main(){
var x int

fmt.Println(x)

var y int=5
fmt.Println(y)
y=6
fmt.Println(y)
z:=7
fmt.Println(z)
sum :=z+y
fmt.Println(sum)
sum=sum+2
fmt.Println(sum)

if  x>6{

fmt.Println("More than 6")
} else if y>10{
fmt.Println("y is greater")
}else {
fmt.Println("z is greater")
}

}
