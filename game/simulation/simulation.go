package simulation

import (
	"github.com/z46-dev/arctic"
	"github.com/z46-dev/genn/client/render"
	"github.com/z46-dev/genn/shared"
)

func NewSimulation(renderer *render.Renderer, client *arctic.Client) (sim *Simulation, err error) {
	return
}

func (s *Simulation) OnMessage(msg *shared.Reader) {}

func (s *Simulation) Release() {}