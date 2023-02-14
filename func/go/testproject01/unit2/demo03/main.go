package main
// import "fmt"
// import "unsafe"
import(
	"fmt"
	"unsafe"
)

func main(){
	//溢出
	var num1 int8 = 100	//230
	fmt.Println(num1)

	var num2 uint = 230
	fmt.Println(num2)

	var num3 = 28
	fmt.Printf("num3的数据类型是: %T",num3)
	fmt.Println()
	fmt.Println("长度：",unsafe.Sizeof(num3))

	var age byte = 18
	fmt.Println(age)
}