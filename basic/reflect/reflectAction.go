package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func main() {
	reflectGetType()
	reflectGetKind()
	reflectGetOriginValue(uint16(500))
	reflectGetOriginValue(st{
		"spider",
		123,
	})
	var i = 1
	reflectModifyValue(&i)

}

//反射修改值传指针
func reflectModifyValue(param interface{}) {
	paramMarshal, err := json.Marshal(param)
	if err !=nil{
		log.Println("param err",param)
	}
	log.Print("param is:",paramMarshal)
	ptr := reflect.ValueOf(param)
	pm := ptr.Elem()
	fmt.Println("是否可以修改:", pm.CanSet())
	fmt.Println("类型:", pm.Type())
	pm.SetInt(255)
	result, err := json.Marshal(param)
	log.Println("param value is:", result)

}

func reflectGetOriginValue(param interface{}) {
	paramValue := reflect.ValueOf(param)
	paramKind := paramValue.Kind()
	switch paramKind {
	case reflect.Uint8:
		fmt.Printf("reflectType int8 -> originValue: %v\n", int8(paramValue.Int()))
	case reflect.Uint16:
		fmt.Printf("reflectType int16 -> originValue: %v\n", uint16(paramValue.Uint()))
		fallthrough
	case reflect.Uint:
		fmt.Printf("reflectType int64 -> originValue: %v\n", paramValue.Uint())
	case reflect.Float32:
		fmt.Printf("reflectType float64 -> originValud: %v\n", float32(paramValue.Float()))
	case reflect.String:
		fmt.Printf("reflectType int8 -> originValue: %v\n", int8(paramValue.Int()))
	default:
		marshal, err := json.Marshal(param)
		if err != nil {
			fmt.Printf("param marshal fail，%s\n", err)
		}
		fmt.Printf("reflect get origin value fail:-> %v\n", string(marshal))
	}

}
func reflectGetType() {
	var f = 5.50
	var in int64 = 55
	reflectType(f)
	reflectType(in)
}

type dou float64
type st struct {
	Name string
	Age  dou
}

func
reflectGetKind() {
	st := new(st)
	st.Name = "spider"
	st.Age = 123
	k := reflect.TypeOf(st)
	typeName := k.Name()
	typeKind := k.Kind()
	fmt.Printf("reflect get type&kind [%v:%v]\n", typeName, typeKind)
	m := make(map[string]string)
	m["name"] = "spider"
	mof := reflect.TypeOf(m)
	fmt.Printf("reflect from map,name:kind->[%v:%v]\n", mof.Name(), mof.Kind())
}
func
reflectType(i interface{}) {
	ty := reflect.TypeOf(i)
	fmt.Printf("type:%v\n", ty)
}

//在运行期间对本身访问和修改的能力，字段信息、类型信息、结构体信息
//任何接口值都是由一个具体的类型和具体的类型值组成 reflect.Type + reflect.Value

//typeOf
//Name Kind
//数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
//valueOf
//interface int uint float bool []byte string
