package main

import (
	"fmt"
	"reflect"
)

type People interface {
	GetName()
}

type Student struct {
	name string
	id   int
}

func (stu Student) GetName() {
	fmt.Println(stu.name)
}

func printMethodAndField() {
	stu := Student{
		name: "gongqi",
		id:   13101104,
	}
	var p People = stu
	ty := reflect.TypeOf(p)
	for i := 0; i < ty.NumMethod(); i++ {
		methodName := ty.Method(i).Name
		methodType := ty.Method(i).Type
		fmt.Println("method:", i, methodName, methodType)
	}
	for i := 0; i < ty.NumField(); i++ {
		fieldName := ty.Field(i).Name
		fieldType := ty.Field(i).Type
		fmt.Println("field:", i, fieldName, fieldType)
	}
	val := reflect.ValueOf(p)
	fmt.Println(val)
}

//通过反射修改切片的容量和长度
func modifySliceByReflect() {
	sli := make([]int, 10)
	fmt.Println("修改前的容量和长度:", cap(sli), len(sli))
	sliceValue := reflect.ValueOf(&sli)
	fmt.Println("type of sliceValue:", sliceValue.Type())
	fmt.Println("settable of sliceValue:", sliceValue.CanSet())
	sliceValueElem := sliceValue.Elem()
	fmt.Println("type of sliceValueElem:", sliceValueElem.Type())
	fmt.Println("settable of sliceValueElem:", sliceValueElem.CanSet())
	sliceValueElem.SetLen(8)
	sliceValueElem.SetCap(8)
	fmt.Println("修改后的容量和长度:", cap(sli), len(sli))
}

func deepEqual() {
	m1 := make(map[int]string)
	m1[1] = "xxx"
	m2 := make(map[int]string)
	m2[1] = "xxx"
	fmt.Println(reflect.DeepEqual(m1, m2))
}

type CustomError struct{}

func (*CustomError) Error() string {
	return ""
}

func implement() {
	typeOfError := reflect.TypeOf((*error)(nil)).Elem()
	customErrorPtr := reflect.TypeOf(&CustomError{})
	customError := reflect.TypeOf(CustomError{})

	fmt.Println(customErrorPtr.Implements(typeOfError)) // #=> true
	fmt.Println(customError.Implements(typeOfError))    // #=> false
}

func Add(a, b int) int { return a + b }

func addByReflect() {
	v := reflect.ValueOf(Add)
	if v.Kind() != reflect.Func {
		return
	}
	t := v.Type()
	argv := make([]reflect.Value, t.NumIn())
	for i := range argv {
		if t.In(i).Kind() != reflect.Int {
			return
		}
		argv[i] = reflect.ValueOf(3)
	}
	result := v.Call(argv)
	if len(result) != 1 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println(result[0].Int()) // #=> 6
}
