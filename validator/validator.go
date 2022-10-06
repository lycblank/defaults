package validator

import "github.com/lycblank/defaults"

type StructValidator interface {
    // ValidateStruct can receive any kind of type and it should never panic, even if the configuration is not right.
    // If the received type is a slice|array, the validation should be performed travel on every element.
    // If the received type is not a struct or slice|array, any validation should be skipped and nil must be returned.
    // If the received type is a struct or pointer to a struct, the validation should be performed.
    // If the struct is not valid or the validation itself fails, a descriptive error should be returned.
    // Otherwise nil must be returned.
    ValidateStruct(interface{}) error

    // Engine returns the underlying validator engine which powers the
    // StructValidator implementation.
    Engine() interface{}
}

type structValidatorWithDefault struct {
    StructValidator
}

func NewStructValidatorWithDefault(s StructValidator) StructValidator {
    return &structValidatorWithDefault{
        StructValidator: s,
    }
}

func (s *structValidatorWithDefault) ValidateStruct(v interface{}) error {
    if err := defaults.Apply(v); err != nil {
        return err
    }
    return s.StructValidator.ValidateStruct(v)
}
