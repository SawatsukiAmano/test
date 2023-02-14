package main
import "fmt"

//全局变量

var n7 = 100
var n8 = 1.1

var(
	n9 = 500
	n10 = "jack"
)

func main(){
	//第一种	指定类型，并且复制
	var num int = 18
	fmt.Println(num)

	//第二种	指定类型，不赋值，使用默认值
	var num2 int
	fmt.Println(num2)

	//第三种	没有写类型，根据后面的值进行判定变量的类型
	var num3 = 10.23
	fmt.Println(num3)

	//第四种	省略var，  var——:=	，不能写为 =
	sex := "男"	
	fmt.Println(sex)

	fmt.Println("--------------------------------------------------------")
	//声明多个变量
	var n1,n2,n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	var n4,name,n5 =10,"alice",7.8
	fmt.Println(n4)
	fmt.Println(name)
	fmt.Println(n5)

	n6,height := 6.9,100.6
	fmt.Println(n6)
	fmt.Println(height)

	fmt.Println(n7)
	fmt.Println(n8)

	fmt.Println(n9)
	fmt.Println(n10)
}