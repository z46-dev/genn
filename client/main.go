package main

import (
	"github.com/z46-dev/genn/client/render"
	"github.com/z46-dev/genn/client/singleplayer"
	"github.com/z46-dev/golog"
)

func main() {
	var (
		renderer *render.Renderer
		err      error
		log      *golog.Logger = golog.New()
	)

	if renderer, err = render.NewRenderer(log); err != nil {
		log.Panicf("create renderer: %v", err)
	}

	var localServer *singleplayer.LocalServer
	if localServer, err = singleplayer.ListenForSingleplayerLobbies(); err != nil {
		log.Panicf("listen for singleplayer lobbies: %v", err)
	}

	// homeScreen.OnSubmit = func(username string) {
	// 	var (
	// 		sim *simulation.Simulation
	// 		err error
	// 	)

	// 	if sim, err = singleplayer.NewSingleplayer(renderer, localServer.Address()); err != nil {
	// 		err = fmt.Errorf("connect to local game: %w", err)
	// 		return
	// 	}

	// 	if err = sim.Spawn(username); err != nil {
	// 		sim.Release()
	// 		err = fmt.Errorf("spawn flower: %w", err)
	// 		return
	// 	}

	// 	renderer.SetState(sim.RenderState)
	// }

	if err = renderer.WaitAndCleanup(); err != nil {
		log.Panicf("wait and cleanup renderer: %v", err)
	}

	if err = localServer.Close(); err != nil {
		log.Errorf("close singleplayer server: %v", err)
	}
}
