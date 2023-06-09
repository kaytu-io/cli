package main

type ServiceTemplate struct {
	ServiceName      string
	ServiceNameSnake string
}

type ChildCmdTemplate struct {
	NameCamel        string
	NameSnake        string
	NameCommand      string
	APIName          string
	ServiceName      string
	ServiceNameSnake string
	//Params           []Param
	ParamString string
}

type Param struct {
	Name       string
	Type       string
	FlagName   string
	StructName string
	Required   bool
}
