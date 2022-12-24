package build

import "fmt"

type Build struct {
	Time    string `json:"time"`
	User    string `json:"user"`
	Version string `json:"version"`
}

func (b *Build) Print() {
	fmt.Println("build.Version: \t", b.Version)
	fmt.Println("build.Time: \t", b.Time)
	fmt.Println("build.User: \t", b.User)
}
