//package main
//
//import (
//	mcc "./test"
//	"fmt"
//	"reflect"
//	"time"
//	"unsafe"
//)
//const imooc string ="慕课网"
//const name="我是zhang不大的先生"
//const (
//	cat string="猫"
//	dog="狗"
//)
//var r,t=1111,2222
//type hcc uint
//var(
//	c string ="mcc"
//	d int =12
//)
//const sr=iota
//const ji=iota
//const (
//	gi=iota
//	_=iota
//	gp=iota
//)
//const (
//	si=iota
//	sp=3.14
//	sl=iota*2
//)
//func main() {
//	mcc.Show()
//	var i uint
//	fmt.Print("\n")
//	fmt.Print("uint:")
//	fmt.Print(i)
//	fmt.Print("\n")
//
//	fmt.Print(unsafe.Sizeof(i))
//	var f float32
//	fmt.Print("\nfloat32:")
//	fmt.Print(f)
//	fmt.Print("\n")
//	fmt.Print(unsafe.Sizeof(f))
//
//	var b bool
//	fmt.Print("\n")
//	fmt.Print("bool:")
//	fmt.Print(b)
//	fmt.Print(unsafe.Sizeof(b))
//	fmt.Print("\n")
//	var src hcc = 1
//
//	var src2 hcc
//	fmt.Print(src2 + src)
//
//	var rt, c = 11213213.1, 2123212
//	fmt.Print(rt, "\n", c, "\n")
//	fmt.Print(r, "\n", t)
//	fmt.Print("\n", reflect.TypeOf(rt))
//	var ru int = int(i)
//	fmt.Print("\n", ru, "\n")
//
//	su, ss := 1, 2
//	fmt.Print(su, "\n", ss)
//	var ds, dd, _ = 12, 22, 32
//
//	var df=int64(ds)
//	fmt.Print("\n",ds,"\n",dd,"\n",df,"\n")
//	fmt.Print(reflect.TypeOf(df))
//	fmt.Print("\n",mcc.B)
//
//
//	const apple,banana string="苹果","香蕉"
//	fmt.Print("\n",apple,"\n",banana,"\n")
//	fmt.Print("\n",dog,"\n",cat,"\n")
//
//	const alen=len(apple)
//	fmt.Print(alen)
//	var um=ic(apple)
//	fmt.Print("\n",um)
//	fmt.Print("gp的值为：",gp)
//	fmt.Print("\n",si,"\n",sp,"\n",sl,"\n")
//	var a=1
//	if a>0{
//		fmt.Print("a>0\n")
//		if a<4{
//			fmt.Print("a<4\n")
//
//		}else{
//			fmt.Print("a>=4\n")
//		}
//	} else {
//		fmt.Print("a<=0\n")
//	}
//	var gui interface{}
//	gui="afd"
//	switch gui.(type){
//	case int:
//		fmt.Print("\n值为int")
//	case string:
//		fmt.Print("\n值为string")
//	default:
//		fmt.Print("\n值为default\n")
//	}
//
//	for i:=1;i<2;i++{
//		fmt.Print("mcc",i,"\n")
//		time.Sleep(1*time.Second)
//
//	}
//	var sss=[]string{"香蕉","苹果","梨子"}
//
//	for key,value:=range sss{
//		fmt.Print("key的值为：")
//		fmt.Print(key)
//		fmt.Print("\n")
//
//		fmt.Print("value的值为：")
//		fmt.Print(value)
//		fmt.Print("\n")
//
//	}
//	for _,value:=range sss{
//		fmt.Print("value的值为：")
//		fmt.Print(value)
//		fmt.Print("\n")
//
//	}
//
//
//	fmt.Print("中间代码块\n")
//	One:
//		fmt.Print("这里是代码块一\n")
//		time.Sleep(2*time.Second)
//		a++
//		if a>3{
//			goto two
//	}
//
//
//	goto One
//	two:
//		fmt.Print("this is two")
//}
//func ic(b string) string{
//	return "asd"
//}
package main

import "fmt"

func main(){
	fmt.Print("mcc")
}