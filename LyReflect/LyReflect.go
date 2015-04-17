package LyReflect

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func GetTypeOf() {
	u := &User{"小米", 1}
	fmt.Println(reflect.TypeOf(u.Name))
	fmt.Println(reflect.TypeOf(u.Age))
	fmt.Println(reflect.TypeOf(u))
}

func GetValueOf() {
	u := &User{"小米", 13}
	fmt.Println(reflect.ValueOf(u.Age))
}

//Kind返回v持有的值的分类，如果v是Value零值，返回值为Invalid
func GetKind() {
	u := &User{"小米", 24}
	v := reflect.ValueOf(u.Age)
	k := v.Kind()
	fmt.Println(k)
}

func GetType() {
	u := &User{"小米", 23}
	v := reflect.ValueOf(u.Name)
	t := v.Type()
	fmt.Println(t)
}
