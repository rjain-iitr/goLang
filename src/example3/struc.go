package main
import(
"fmt"
)

type person struct{
	name string
	age int
}

func main(){
	p:=person{ name:"rahul",age:29 }
	fmt.Println(p)
	fmt.Println(p.name)
	i:=34
	inc(i)	
        fmt.Println(i,&i)
	inc2(&i)	
        fmt.Println(i,&i)

}

func inc(i int ){
i++

}


func inc2(i *int ){
*i++
}


