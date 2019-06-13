package blue

import "github.com/owais/demo-service/pkg/console"

type Listener struct{}

func (l Listener) Start() {
	console.Print("starting Blue")
}
