package main

import (
	"fmt"
	"reflect"
)

type customInt int64

type person struct {
	Name  string   `json:"name"`
	Age   uint     `json:"age"`
	Hobby []string `json:"hobby"`
}

func main() {
	fmt.Println("**** TypeOf start ****")
	// Type
	var floatVariable float64 = 3.14
	reflectType(floatVariable)
	var intVariable int = 3
	reflectType(intVariable)
	var boolVariable bool = false
	reflectType(boolVariable)
	// type kind
	// 反射中，数组，切片，map，指针 等类型的指针 ‘.Name()’ 返回的值为空
	var int64Variable int64 = 3
	reflectTypeAndKind(int64Variable)
	var custonIntVariable customInt = 3
	reflectTypeAndKind(custonIntVariable)
	fmt.Println("**** TypeOf end ****")

	fmt.Println("**** ValueOf start ****")
	var intValueVariable int64 = 100
	reflectValue(intValueVariable)
	var boolValueVariable bool = true
	reflectValue(boolValueVariable)
	var floatValueVariable float64 = 3.14
	reflectValue(floatValueVariable)
	fmt.Println("**** ValueOf end ****")

	fmt.Println("**** 通过反射设置变量的值 start ****")
	var intSetValue int64 = 99
	reflectSetValue(&intSetValue)
	fmt.Printf("intSetValue value: %d\n", intSetValue)
	fmt.Println("**** 通过反射设置变量的值 end ****")

	reflectStruct()
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", v)
}

func reflectTypeAndKind(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v, kind: %v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()

	switch k {
	case reflect.Int64:
		fmt.Printf("type: %v, value: %v\n", k, int64(v.Int()))
	case reflect.Float64:
		fmt.Printf("type: %v, value: %v\n", k, float64(v.Float()))
	case reflect.Bool:
		fmt.Printf("type: %v, value: %v\n", k, bool(v.Bool()))
	case reflect.Float32:
		fmt.Printf("type: %v, value: %v\n", k, float32(v.Float()))
	default:
		fmt.Printf("type: %v, value: %v\n", k, v)
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Elem().Kind()
	switch k {
	case reflect.Int64:
		v.Elem().SetInt(666)
	case reflect.Float64:
		v.Elem().SetFloat(3.1415926)
	case reflect.Bool:
		v.Elem().SetBool(false)
	default:
		fmt.Printf("type: %v, value: %v\n", k, v)
	}
}

func reflectStruct() {
	per := person{
		Name:  "曹操",
		Age:   100,
		Hobby: []string{"军事", "你懂的"},
	}
	v := reflect.TypeOf(per)
	fmt.Println(v.Name(), v.Kind())

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
}
