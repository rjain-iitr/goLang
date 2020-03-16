package main
import(
"fmt"
)

type student struct{
	name string
	age int
	subject1marks float64
	subject2marks float64
	subject3marks float64
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

func computePercentage (a student)  float64 {
    percentageChan := percentageFuture(a)
    percentage := <-percentageChan
    return percentage
}

func percentageFuture (a student) chan float64 {
    future := make (chan float64)
    go func () { future <- percentage(a)  }()
    return future;
}

func percentage (a student)float64{
	var percentage float64
	percentage=(a.subject1marks+a.subject2marks+a.subject3marks)*100/300
	return percentage
}


func computeAverage (a student,b student)  float64 {
    averageChan := averageFuture(a,b)
    average := <-averageChan
    return average
}

func averageFuture (a student,b student) chan float64 {
    future := make (chan float64)
    go func () { future <- average(a,b)  }()
    return future;
}

func average (a student,b student)float64{
	averageA:=(a.subject1marks+a.subject2marks+a.subject3marks)/3
	averageB:=(b.subject1marks+b.subject2marks+b.subject3marks)/3
	return (averageA+averageB)/2
}



