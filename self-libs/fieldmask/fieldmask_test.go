package fieldmask_test

import (
	"reflect"
	"testing"

	"google.golang.org/genproto/protobuf/field_mask"

	"roshar/utils/fieldmask"
	"roshar/utils/ptr"
)

func TestGenerateFieldMask(t *testing.T) {
	type TestFoo struct {
		AInt    *int
		AString *string
	}

	type Test struct {
		Foo     *TestFoo
		AInt    *int
		AString *string
	}

	cases := []struct {
		object interface{}
		want   []string
	}{
		{nil, []string{}},
		{&Test{nil, ptr.Int(1), ptr.String("a")}, []string{"test.a_int", "test.a_string"}},
		{&Test{&TestFoo{ptr.Int(1), ptr.String("a")}, ptr.Int(1), ptr.String("a")}, []string{"test.foo.a_int", "test.foo.a_string", "test.a_int", "test.a_string"}},
	}

	for _, c := range cases {
		mask := fieldmask.GenerateFieldMask(c.object)
		if !reflect.DeepEqual(mask.Paths, c.want) {
			t.Errorf("want %v, got %v", c.want, mask.Paths)
		}
	}
}

func TestGenerateMaskedMap(t *testing.T) {
	type TestFoo struct {
		AInt    int
		AString string
	}

	type Test struct {
		Foo     TestFoo
		AInt    int
		AString string
	}

	cases := []struct {
		object interface{}
		mask   []string
		want   map[string]interface{}
		desc   string
	}{
		{nil, []string{}, map[string]interface{}{}, "test nil"},

		{&Test{TestFoo{2, "b"}, 1, "a"}, []string{"test.a_int"},
			map[string]interface{}{"a_int": 1},
			"test single field mask on attributes"},

		{&Test{TestFoo{2, "b"}, 1, "a"}, []string{"test.a_int", "test.a_string"},
			map[string]interface{}{"a_int": 1, "a_string": "a"},
			"test multiple field masks on attirbutes"},

		{&Test{TestFoo{2, "b"}, 1, "a"}, []string{"test.foo.a_int"},
			map[string]interface{}{"foo": map[string]interface{}{"a_int": 2}},
			"test single field mask in nested struct attribute"},

		{&Test{TestFoo{2, "b"}, 1, "a"}, []string{"test.foo.a_int", "test.foo.a_string"},
			map[string]interface{}{"foo": map[string]interface{}{"a_int": 2, "a_string": "b"}},
			"test multiple field masks in nested struct attribute"},

		{&Test{TestFoo{2, "b"}, 1, "a"}, []string{"test.foo.a_int", "test.foo.a_string", "test.a_int"},
			map[string]interface{}{"foo": map[string]interface{}{"a_int": 2, "a_string": "b"}, "a_int": 1},
			"test multiple field masks in both direct and nested struct attribute"},

		{&Test{TestFoo{2, "b"}, 1, "a"}, []string{"test.foo.a_int", "test.foo.a_string", "test.a_int", "test.a_string"},
			map[string]interface{}{"foo": map[string]interface{}{"a_int": 2, "a_string": "b"}, "a_int": 1, "a_string": "a"},
			"test multiple field masks in all fields"},
	}

	for _, c := range cases {
		got := fieldmask.GenerateMaskedMap(c.object, field_mask.FieldMask{Paths: c.mask})
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("want %v, got %v for %s", c.want, got, c.desc)
		}
	}
}
