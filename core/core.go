/*
@Desc:
@Author: qy
@Time : 2020/12/16 9:48
@File : core
@Software: GoLand
*/
package core

import "reflect"

// 将验证参数后的结构体参数值赋给对应的model
func ToModel(_params interface{}, _model interface{}){
	tpeParams := reflect.TypeOf(_params)
	valParams := reflect.ValueOf(_params)
	valModel := reflect.ValueOf(_model)
	if valModel.Kind() != reflect.Ptr {
		panic("非指针参数")
	}
	valModel = valModel.Elem()
	for i :=0; i < tpeParams.NumField(); i++{
		key := tpeParams.Field(i)
		value := valParams.Field(i)
		keyModel := valModel.FieldByName(key.Name)
		keyModel.Set(value)
	}
}

