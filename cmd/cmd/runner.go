package cmd

type Runner interface {
	Run() error
	Name() string
}
