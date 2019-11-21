package main

import (
  "bufio"
  "fmt"
  "os"
  "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
  
	defer writer.Flush()
	var n,m,a uint64
	scanf("%d %d %d\n",&n,&m,&a)
	var h,b = 1,1
	if(m<=a && n <= a){
		printf("%d\n",1);	
		return
	}	 
	if(n>a){
		h = int(math.Ceil(float64(n)/float64(a)))
	}
	if(m>a){
		b = int(math.Ceil(float64(m)/float64(a)))
	}
	printf("%d\n",uint64(h)*uint64(b));
}