package funcs

type templateVar struct {
	value interface{}
}

func NewVar(value interface{}) *templateVar {
	return &templateVar{value}
}

func SetVar(v *templateVar, value interface{}) string {
	v.value = value

	return ""
}

func GetVar(v *templateVar) interface{} {
	return v.value
}
