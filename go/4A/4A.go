package main
import(
//	"bufio"
	"fmt"

)

func main()  {
	var w int
	fmt.Scanf("%d\n",&w)
	if w<4 {
		fmt.Println("NO")
	} else if w%2 == 0{
		fmt.Println("YES")
	} else{
		fmt.Println("NO")
	}

}