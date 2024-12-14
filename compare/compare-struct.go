package compare

import "reflect"

func CompareStruct(x, y any) (identical bool) {
	valX := reflect.ValueOf(x)
	valY := reflect.ValueOf(y)
	typeX := reflect.TypeOf(x)

	// Ensure both values are structs
	if valX.Kind() != reflect.Struct || valY.Kind() != reflect.Struct {
		return false
	}

	// Iterate over all fields in x (assuming x and y are of the same type)
	for i := 0; i < valX.NumField(); i++ {
		fieldX := valX.Field(i)
		fieldY := valY.Field(i)
		fieldType := typeX.Field(i)

		// Check if the field is a struct itself (but not a time.Time, for example)
		if fieldX.Kind() == reflect.Struct && fieldType.Type != reflect.TypeOf(struct{}{}) {
			// Recursively compare the struct fields
			if !CompareStruct(fieldX.Interface(), fieldY.Interface()) {
				return false
			}
		} else {
			if !reflect.DeepEqual(fieldX.Interface(), fieldY.Interface()) {
				return false
			}
		}
	}

	return true
}
