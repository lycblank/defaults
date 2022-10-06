package defaults

import "reflect"

var defaultDefault = &DefaultContainer{}

// Apply set default values for variables
func Apply(v interface{}) error {
    return defaultDefault.Apply(v)
}

type DefaultContainer struct {
    field field
}

// Apply set default values for variables
func (d *DefaultContainer) Apply(v interface{}) error {
    return d.apply(reflect.ValueOf(v), nil)
}

func (d *DefaultContainer) apply(rv reflect.Value, tag *reflect.StructTag) error {
    switch rv.Kind() {
    case reflect.Ptr,reflect.Interface:
        return d.apply(rv.Elem(), nil)
    case reflect.Struct:
        return d.applyStruct(rv)
    case reflect.Slice, reflect.Array:
        count := rv.Len()
        for i := 0; i < count; i++ {
            if err := d.apply(rv.Index(i), nil); err != nil {
                return err
            }
        }
        return nil
    case reflect.Map:
        for iter := rv.MapRange(); iter.Next(); {
           if iter.Value().Kind() == reflect.Ptr {
                if err := d.apply(iter.Value(), nil); err != nil {
                    return err
                }
           }
        }
        return nil
    }

    if tag == nil {
        return nil
    }

    d.field.v = rv
    return d.field.apply(*tag)
}

// set default values for struct
func (d *DefaultContainer) applyStruct(v reflect.Value) (err error) {
    numField := v.NumField()
    t := v.Type()
    for i := 0; i < numField; i++ {
        if !t.Field(i).IsExported() {
            continue
        }

        tag := t.Field(i).Tag
        if err := d.apply(v.Field(i), &tag); err != nil {
            return err
        }
    }
    return nil
}
