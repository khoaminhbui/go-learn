package main

import (
	"github.com/khoaminhbui/go-learn/domain"
	"github.com/khoaminhbui/go-learn/infra"
)

func main() {
	domain.ShowCores()
	infra.StartSimpleServer()
}
