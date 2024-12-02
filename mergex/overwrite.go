package mergex

import "reflect"

// OverWriteTo 将给定的结构体覆盖到给定的指针指向的结构体。
func OverWriteTo[S any, D any](src S, dest *D) {
	s := reflect.ValueOf(src)
	d := reflect.ValueOf(dest).Elem()
	if s.Kind() != reflect.Struct && d.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < s.NumField(); i++ {
		if dField, sField := d.FieldByName(s.Type().Field(i).Name), s.Field(i); dField.IsValid() && dField.Kind() == sField.Kind() {
			dField.Set(sField)
		}
	}
	newDest := d.Interface().(D)
	dest = &newDest
	return
}
