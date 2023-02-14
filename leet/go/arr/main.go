package main	//声明文件所在的包，每个go文件必须有归属的包
import "fmt"	//引入程序中需要用的包，为了使用包下的函数	比如：Println
func main(){	//main主函数	程序的入库
	fmt.Println("Hello World")	//在控制台打印出一句话，双引号中的内容会原样输出
}