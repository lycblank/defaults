package defaults

import "reflect"

var defaultDefault = &DefaultContainer{}

// Apply set default values for variables
func Apply(v interface{}) error {
    return defaultDefault.Apply(v)
}

type DefaultContainer struct {
}

// Apply set default values for variables
func  (d *DefaultContainer) Apply(v interface{}) error {
    return d.apply(reflect.ValueOf(v))
}

func (d *DefaultContainer) apply(rv reflect.Value) error {
    switch rv.Kind() {
    case reflect.Ptr:
        return d.apply(rv.Elem())
    case reflect.Struct:
        return d.applyStruct(rv)
    case reflect.Slice, reflect.Array:
        count := rv.Len()
        for i := 0; i < count; i++ {
            if err := d.apply(rv.Index(i)); err != nil {
                return err
            }
        }
        return nil
    default:
        return nil
    }
}

// set default values for struct
func (d *DefaultContainer) applyStruct(v reflect.Value) (err error) {
    field := &field{}

    numField := v.NumField()
    t := v.Type()
    for i := 0; i < numField; i++ {
        if !t.Field(i).IsExported() {
            continue
        }

        rv := v.Field(i)
        switch rv.Kind() {
        case reflect.Ptr:
            err = d.apply(rv.Elem())
        case reflect.Struct:
            err = d.applyStruct(rv)
        case reflect.Slice, reflect.Array:
            count := rv.Len()
            for i := 0; i < count; i++ {
                if err = d.apply(rv.Index(i)); err != nil {
                    return err
                }
            }
        default:
            field.v = rv
            err = field.apply(t.Field(i).Tag)
        }

        if err != nil {
            return err
        }
    }
    return nil
}
