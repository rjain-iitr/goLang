package main
import(
"fmt"
"errors"
"math"

)


func main(){
result:=sum(2,3)
fmt.Println(result)

result2,err :=sqrt(27)
print(result2,err)

result2,err =sqrt(-2)
print(result2,err)

}

func sum(x int,y int)int{
	return x+y
}


func sqrt(x float64)(float64,error){
	if x<0{
		return 0,errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x),nil

}

func print(result2 float64,err error) {
	if err!=nil{
	fmt.Println(err)
	}else{
		fmt.Println(result2)
	}
}
