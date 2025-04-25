package cherryReflect

import (
	"fmt"
	"reflect"
	"runtime"

	cstring "github.com/cherry-game/cherry/extend/string"
)

func ReflectTry(f reflect.Value, args []reflect.Value, handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("-------------panic recover---------------")
			if handler != nil {
				handler(err)
			}
		}
	}()
	f.Call(args)
}

func GetStructName(v interface{}) string {
	return reflect.Indirect(reflect.ValueOf(v)).Type().Name()
}

func GetFuncName(fn interface{}) string {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		panic(fmt.Sprintf("[fn = %v] is not func type.", fn))
	}

	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	return cstring.CutLastString(fullName, ".", "-")
}

////GetInvokeFunc reflect function convert to MethodInfo
//func GetInvokeFunc(name string, fn interface{}) (*cfacade.MethodInfo, error) {
//	if name == "" {
//		return nil, cerr.Error("func name is nil")
//	}
//
//	if fn == nil {
//		return nil, cerr.Errorf("func is nil. name = %s", name)
//	}
//
//	typ := reflect.TypeOf(fn)
//	val := reflect.ValueOf(fn)
//
//	if typ.Kind() != reflect.Func {
//		return nil, cerr.Errorf("name = %s is not func type.", name)
//	}
//
//	var inArgs []reflect.Type
//	for i := 0; i < typ.NumIn(); i++ {
//		t := typ.In(i)
//		inArgs = append(inArgs, t)
//	}
//
//	var outArgs []reflect.Type
//	for i := 0; i < typ.NumOut(); i++ {
//		t := typ.Out(i)
//		outArgs = append(outArgs, t)
//	}
//
//	invoke := &cfacade.MethodInfo{
//		Type:    typ,
//		Value:   val,
//		InArgs:  inArgs,
//		OutArgs: outArgs,
//	}
//
//	return invoke, nil
//}

func IsPtr(val interface{}) bool {
	if val == nil {
		return false
	}

	return reflect.TypeOf(val).Kind() == reflect.Ptr
}

func IsNotPtr(val interface{}) bool {
	if val == nil {
		return false
	}

	return reflect.TypeOf(val).Kind() != reflect.Ptr
}

func GetField(obj interface{}, fieldName string) (*reflect.Value, error) {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("target is not struct")
	}

	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, fmt.Errorf("field %s not exist", fieldName)
	}

	return &field, nil
}

func GetFieldValue(obj interface{}, fieldName string) (interface{}, error) {
	field, err := GetField(obj, fieldName)
	if err != nil {
		return nil, err
	}
	return field.Interface(), nil
}

func GetFieldPtr(obj interface{}, fieldName string) (interface{}, error) {
	field, err := GetField(obj, fieldName)
	if err != nil {
		return nil, err
	}
	return field.Addr().Interface(), nil
}
