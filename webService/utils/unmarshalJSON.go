package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
)

type SpecialJSON struct {
	Data interface{}
}

func (r *SpecialJSON) UnmarshalJSON(data []byte) error {
	v, err := JSONSliceToStr(data, r.Data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(v, r.Data); err != nil {
		return err
	}
	return nil
}
func JSONSliceToStr(data []byte, v interface{}) ([]byte, error) {
	rList := &[]string{}
	genJSONSliceToStrRegexp(v, []string{}, rList)
	for i := 0; i < len(*rList); i++ {
		compile, err := regexp.Compile((*rList)[i])
		if err != nil {
			return nil, err
		}
		for compile.Match(data) {
			idx := compile.FindSubmatchIndex(data)
			data[idx[2]] = '"'
			data[idx[2]+1] = '"'
		}
	}
	return data, nil
}

// 生成提取[]的正则表达式
func genJSONSliceToStrRegexp(target interface{}, root []string, rList *[]string) {
	var tType reflect.Type
	if _, ok := target.(reflect.Type); ok {
		tType = target.(reflect.Type)
	} else {
		tType = reflect.TypeOf(target)
	}
	if tType.Kind() == reflect.Ptr {
		tType = tType.Elem()
	}
	num := tType.NumField()
	for i := 0; i < num; i++ {
		newRoot := make([]string, len(root)+1)
		copy(newRoot, append(root, tType.Field(i).Tag.Get("json")))
		if tType.Field(i).Type.Kind() == reflect.Struct {
			genJSONSliceToStrRegexp(tType.Field(i).Type, newRoot, rList)
		} else if tType.Field(i).Type.Kind() == reflect.Ptr &&
			tType.Field(i).Type.Elem().Kind() == reflect.Struct {
			genJSONSliceToStrRegexp(tType.Field(i).Type.Elem(), newRoot, rList)
		} else if tType.Field(i).Type.Kind() == reflect.Slice {
			if tType.Field(i).Type.Elem().Kind() == reflect.Ptr &&
				tType.Field(i).Type.Elem().Elem().Kind() == reflect.Struct {
				genJSONSliceToStrRegexp(tType.Field(i).Type.Elem().Elem(), newRoot, rList)
			} else if tType.Field(i).Type.Elem().Kind() == reflect.Struct {
				genJSONSliceToStrRegexp(tType.Field(i).Type.Elem(), newRoot, rList)
			}
		} else if tType.Field(i).Type.Kind() == reflect.String {
			if len(root) >= 1 {
				*rList = append(*rList, fmt.Sprintf(`"%s"[^{]*{.*"%s":[ ]?(\[\])[^",]*`, root[len(root)-1], tType.Field(i).Tag.Get("json")))
			} else if len(root) == 0 {
				*rList = append(*rList, fmt.Sprintf(`^[^{]*{.*"%s":[ ]?(\[\])[^",]*`, tType.Field(i).Tag.Get("json")))
			}
		}
	}
}
