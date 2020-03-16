package main
import(
"fmt"
)

type student struct{
	name string
	age int
	subject1marks int
	subject2marks int
	subject3marks int
}

func main(){
	s1:=student{ name:"Rahul",age:29,subject1marks:95,subject2marks:92,subject3marks:90 }
	s2:=student{ name:"Mohit",age:30,subject1marks:93,subject2marks:94,subject3marks:87 }
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1.name,"percentage",computePercentage(s1))
	fmt.Println(s2.name,"percentage",computePercentage(s2))
	fmt.Println("class Average is",computeAverage(s1,s2))

	
}

func computePercentage (a student)  int {
    percentageChan := percentageFuture(a)
    percentage := <-percentageChan
    return percentage
}

func percentageFuture (a student) chan int {
    future := make (chan int)
    go func () { future <- percentage(a)  }()
    return future;
}

func percentage (a student)int{
	var percentage int
	percentage=(a.subject1marks+a.subject2marks+a.subject3marks)/3
	return percentage
}


func computeAverage (a student,b student)  int {
    averageChan := averageFuture(a,b)
    average := <-averageChan
    return average
}

func averageFuture (a student,b student) chan int {
    future := make (chan int)
    go func () { future <- average(a,b)  }()
    return future;
}

func average (a student,b student)int{
	averageA:=(a.subject1marks+a.subject2marks+a.subject3marks)/3
	averageB:=(b.subject1marks+b.subject2marks+b.subject3marks)/3
	return (averageA+averageB)/2
}

