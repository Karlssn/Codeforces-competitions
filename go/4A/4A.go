package main

import (
  "bufio"
  "fmt"
  "os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
  
  defer writer.Flush()
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