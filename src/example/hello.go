package main
import (
"fmt"
) 


func main(){
fmt.Println("Hello World")
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

var a[5] int
fmt.Println(a)
a[4]=3
fmt.Println(a)
b:=[6]int{}
fmt.Println(b)
c:=[]int{}
fmt.Println(c)
c=append(c,12)
fmt.Println(c)
c[0]=5
fmt.Println(c)
var dict map[string]int 
fmt.Println(dict)

 map2:= make(map[string]int)
 map2["ss"]=1
 map2["s"]=2
fmt.Println(map2)
fmt.Println(map2["s"])
delete(map2,"s")
fmt.Println(map2)
fmt.Println(map2["s"])

 map2["his"]=233

for i:=0;i<5&&i<4;i++{
   fmt.Println("i",i)
}
i:=0
for i<5{
   fmt.Println("a[",i,"]",a[i])
i++
}

for index,value:=range a{
   fmt.Println("index:",index,"value:",value)
}

for key,value:=range map2{
   fmt.Println("key:",key,"value:",value)
}
for index,value:=range map2{
   fmt.Println("index:",index,"value:",value)
}


}
