package fieldmask

import (
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/genproto/protobuf/field_mask"

	"roshar/utils/collection"
	"roshar/utils/ptr"
	"roshar/utils/reflection"
	"roshar/utils/stringutil"
	"roshar/utils/structutil"
)

// GenerateFieldMask generates field mask for a struct, ignore fields whose value is nil
// All fields of the object should be pointers, otherwise unable to check whether it is nil
func GenerateFieldMask(object interface{}) field_mask.FieldMask {
	mask := field_mask.FieldMask{Paths: []string{}}
	var (
		val        reflect.Value
		structName string
	)

	if object == nil {
		return mask
	}

	t := reflection.IgnorePointerType(reflect.TypeOf(object))
	// If type is reflect.Value, the object is a nested struct of original object,
	// which needs to be handled differently
	if t.String() == "reflect.Value" {
		val = object.(reflect.Value)
		t = val.Type()
		structName = ""
	} else {
		val = reflection.IgnorePointerValue(reflect.ValueOf(object))
		structName = structutil.UnderscoreStructName(object)
		structName += "."
	}

	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Kind() == reflect.Ptr && val.Field(i).IsNil() {
			continue
		}

		fieldName := stringutil.UnderscoreCamelCase(t.Field(i).Name)

		innerVal := reflection.IgnorePointerValue(val.Field(i))
		// If field is a struct, also go through inner of it
		if innerVal.Kind() == reflect.Struct {
			innerMask := GenerateFieldMask(innerVal)
			fieldName += "."
			for _, path := range innerMask.Paths {
				mask.Paths = append(mask.Paths, fmt.Sprintf("%s%s%s", structName, fieldName, path))
			}
		} else {
			mask.Paths = append(mask.Paths, fmt.Sprintf("%s%s", structName, fieldName))
		}
	}

	return mask
}

// AppendFieldMask appends new field masks to existing onw
func AppendFieldMask(mask field_mask.FieldMask, structName string, paths []string) field_mask.FieldMask {
	for _, path := range paths {
		path = fmt.Sprintf("%s.%s", structName, path)
		if collection.IncludesStr(mask.Paths, path) {
			continue
		}
		mask.Paths = append(mask.Paths, path)
	}
	return mask
}

// GenerateMaskedMap generates a map from object, which is a struct, by filtering out the keys in mask
func GenerateMaskedMap(object interface{}, mask field_mask.FieldMask) map[string]interface{} {
	res := make(map[string]interface{})
	subMasks := make(map[string]field_mask.FieldMask)
	var val reflect.Value

	if object == nil {
		return res
	}

	t := reflection.IgnorePointerType(reflect.TypeOf(object))
	// If type is reflect.Value, the object is a nested struct of original object,
	// which needs to be handled differently
	if t.String() == "reflect.Value" {
		val = object.(reflect.Value)
		t = val.Type()
	} else {
		val = reflection.IgnorePointerValue(reflect.ValueOf(object))
	}

	for _, path := range mask.Paths {
		names := strings.Split(path, ".")
		if len(names) > 2 {
			subPath := strings.Join([]string(names[1:]), ".")
			subMask, ok := subMasks[names[1]]
			if !ok {
				subMasks[names[1]] = field_mask.FieldMask{Paths: []string{subPath}}
			} else {
				subMask.Paths = append(subMask.Paths, subPath)
				subMasks[names[1]] = subMask
			}
		}

		for i := 0; i < val.NumField(); i++ {
			fieldName := stringutil.UnderscoreCamelCase(t.Field(i).Name)
			if fieldName == names[1] {
				innerVal := reflection.IgnorePointerValue(ptr.UnWrap(val.Field(i)))
				res[fieldName] = reflection.GetReflectValue(innerVal)
				break
			}
		}
	}

	for key, subMask := range subMasks {
		names := strings.Split(key, ".")
		for i := 0; i < val.NumField(); i++ {
			fieldName := stringutil.UnderscoreCamelCase(t.Field(i).Name)
			if fieldName == names[0] {
				innerVal := reflection.IgnorePointerValue(ptr.UnWrap(val.Field(i)))
				innerMap := GenerateMaskedMap(innerVal, subMask)
				res[fieldName] = innerMap
				break
			}
		}
	}

	return res
}
