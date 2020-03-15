package main
import(
"fmt"
"time"
)
func main(){
go count("sleep")
 count("go")

fmt.Scanln()


}

func count(thing string){
for i:=1;true;i++{
fmt.Println(i,thing)
time.Sleep(time.Millisecond*500)

}

}
