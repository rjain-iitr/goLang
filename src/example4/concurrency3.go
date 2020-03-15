package main
import(
"fmt"
"time"
)
func main(){
	c :=make(chan string)
	go count("rahul",c)

	for {
	msg,open:= <-c

	if !open{
	break
	}
	fmt.Println(msg)
	}
	
	ch:=make(chan string)
	go count2("jain",ch)
	for msg2:= range ch{

		fmt.Println(msg2)
	}
}

func count(thing string,c chan string){
	for i:=1;i<=5;i++{
	c<-thing
	time.Sleep(time.Millisecond*500)

	}	

	close(c)
}



func count2(thing string,c chan string){
	for i:=1;i<=5;i++{
	c<-thing
	time.Sleep(time.Millisecond*500)

	}	

	close(c)
}
