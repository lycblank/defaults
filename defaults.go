package defaults

import "reflect"

// Apply set default values for variables
func Apply(v interface{}) error {
    return apply(reflect.ValueOf(v))
}

func apply(rv reflect.Value) error {
    switch rv.Kind() {
    case reflect.Ptr:
        return apply(rv.Elem())
    case reflect.Struct:
        return applyStruct(rv)
    case reflect.Slice, reflect.Array:
        count := rv.Len()
        for i := 0; i < count; i++ {
            if err := apply(rv.Index(i)); err != nil {
                return err
            }
        }
        return nil
    default:
        return nil
    }
}

// set default values for struct
func applyStruct(v reflect.Value) error {
    field := &field{}

    numField := v.NumField()
    t := v.Type()
    for i := 0; i < numField; i++ {
        if !t.Field(i).IsExported() {
            continue
        }
        field.v = v.Field(i)
        if err := field.apply(t.Field(i).Tag); err != nil {
            return err
        }
    }
    return nil
}



