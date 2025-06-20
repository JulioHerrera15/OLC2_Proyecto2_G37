package value

type FloatValue struct {
	InternalValue float64
}
func (f FloatValue) Value() interface{} {
	return f.InternalValue
}
func (f FloatValue) Type() string {
	return IVOR_FLOAT
}
func (f FloatValue) Copy() IVOR {
	return &FloatValue{f.InternalValue}
}
