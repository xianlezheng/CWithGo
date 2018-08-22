package main

/*
struct Demo {
	int a;
    int b;
};
*/
import "C"
import "fmt"


//export GoCall
func GoCall(buffer *C.char,out* *C.char) {
	*out = buffer
}

//export PrintGoString
func PrintGoString(bf *C.char){
	fmt.Println(C.GoString(bf))
}

//export GetStructDemo
func GetStructDemo(a,b C.int) C.struct_Demo {
	return C.struct_Demo{a,b}
}

//export PrintStrutDemo
func PrintStrutDemo(s C.struct_Demo){
	fmt.Printf("Struct a:%d,b:%d\n",s.a,s.b)
}

//export PrintGoInt
func PrintGoInt(a,b C.int){
	fmt.Printf("a:%d,b:%d\n",a,b)
}

func main() {

}
