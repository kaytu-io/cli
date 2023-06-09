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
}
