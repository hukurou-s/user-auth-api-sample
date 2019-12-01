package controllers

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{}) error
	Get(string) interface{}
}
