package scan_manager

type EnvVariableScanObject struct {
	variable  string
	value     string
}

func (obj *EnvVariableScanObject) VariableName() string {
	return obj.variable
}

func (obj *EnvVariableScanObject) Value() string {
	return obj.value
}
