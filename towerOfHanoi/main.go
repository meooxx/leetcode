package main 
import (
	"fmt"
)
func move(disk int, source, des, aux []int){
	if disk == 1  {
		des = append(des, disk)
	}
	move(disk - 1, source, aux, des)
	fmt.Println("move %v from %v to %v")
	move(disk - 1, aux, des, source)

}

func main(){
	source := []int{}
	des := []int{}
	aux :=[]int{}
	move(3, source, des, aux)
	fmt.Printf("1:%v, 2%v, 3:%v",source, des, aux)
}