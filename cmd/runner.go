package main

type Runner interface {
	Run() error
	Name() string
}
