package main
import "fmt"



//包类变量
//这二种定义方式都是一样的
// var aa = 3
// var ss = "kkk"
// var bb = true
var (
	aa = 3
	ss = "kkk"
	bb = false
)


//定义变量，如果不赋初值，int的初始值为0,string的为空
func variableZeroValue(){
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

//也可以赋值，但定义一个变量就必须要使用，不然会报错
func variableInitialValue(){
	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a,s,b)
}

//不写类型可以定义不同类型的变量
func variableTypeDeduction(){
	var a,b, c,s = 3,4,true,"def"
	// var b = 5  //不能多次定义变量
	fmt.Println(a,b,c,s)
}


//用冒号定义与上面用var是一样的
func varialbeShorter(){
	a,b,c,s := 3,4,true,"def"
	fmt.Println(a,b,c,s)
}


func  main()  {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	varialbeShorter()
	fmt.Println(aa,bb,ss)
}