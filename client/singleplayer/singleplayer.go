package singleplayer

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/z46-dev/arctic"
	"github.com/z46-dev/genn/client/render"
	"github.com/z46-dev/genn/game"
	"github.com/z46-dev/genn/game/simulation"
	"github.com/z46-dev/genn/shared"
	"github.com/z46-dev/golog"
)

const (
	singleplayerLoopbackAddress string        = "127.0.0.1:0"
	listenerReadyTimeout        time.Duration = 5 * time.Second
)

type LocalServer struct {
	server     *arctic.Server
	games      sync.Map
	listenDone chan error
	closeOnce  sync.Once
	closeErr   error
	log        *golog.Logger
}

// NewSingleplayer creates a normal network client connected to the local server.
func NewSingleplayer(renderer *render.Renderer, serverAddress string) (sim *simulation.Simulation, err error) {
	var (
		client *arctic.Client
		log    *golog.Logger = golog.New().Prefix("[SPCLT]", golog.BoldGreen).Timestamp()
	)

	if client, err = arctic.NewClient(arctic.ClientConfig{
		ServerAddress: serverAddress,
		BufferSize:    shared.NetworkBufferSize,
		Metadata: map[string]any{
			"key": MetaAllowedKey,
		},
	}); err != nil {
		err = fmt.Errorf("create client: %w", err)
		return
	}

	if sim, err = simulation.NewSimulation(renderer, client); err != nil {
		err = fmt.Errorf("create game view: %w", err)
		return
	}

	client.OnMessage(func(data []byte) {
		if len(data) == 0 {
			log.Warningf("discard empty message")
			return
		}

		var messageType uint8 = data[0]
		if messageType <= shared.MsgTypeGameCutoffConst {
			return
		}

		sim.OnMessage(shared.NewReader(data))
	})

	client.OnClose(func() {
		log.Info("disconnected")
	})

	client.OnError(func(clientErr error) {
		log.Errorf("connection error: %v", clientErr)
	})

	if err = client.Connect(); err != nil {
		sim.Release()
		sim = nil
		err = fmt.Errorf("connect client: %w", err)
		return
	}

	log.Info("connected")
	return
}

// ListenForSingleplayerLobbies starts a ready local server on an ephemeral loopback port.
func ListenForSingleplayerLobbies() (local *LocalServer, err error) {
	local = &LocalServer{
		listenDone: make(chan error, 1),
		log:        golog.New().Prefix("[SPSRV]", golog.BoldYellow).Timestamp(),
	}

	if local.server, err = arctic.NewServer(arctic.ServerConfig{
		BindAddress: singleplayerLoopbackAddress,
		BufferSize:  shared.NetworkBufferSize,
	}); err != nil {
		local = nil
		err = fmt.Errorf("create server: %w", err)
		return
	}

	local.bindHandlers()
	go func() {
		local.listenDone <- local.server.Listen()
	}()

	var timeout *time.Timer = time.NewTimer(listenerReadyTimeout)
	var poll *time.Ticker = time.NewTicker(time.Millisecond)
	defer timeout.Stop()
	defer poll.Stop()

	for local.server.Addr() == nil {
		select {
		case err = <-local.listenDone:
			local = nil
			if err == nil {
				err = fmt.Errorf("server stopped before accepting connections")
			} else {
				err = fmt.Errorf("listen: %w", err)
			}

			return
		case <-timeout.C:
			var _ error = local.server.Close()
			local = nil
			err = fmt.Errorf("wait for local server: %w", context.DeadlineExceeded)
			return
		case <-poll.C:
		}
	}

	return
}

// Address returns the local address clients should connect to.
func (s *LocalServer) Address() (address string) {
	if s != nil && s.server != nil && s.server.Addr() != nil {
		address = s.server.Addr().String()
	}

	return
}

// Close stops every local game and closes the listening transport.
func (s *LocalServer) Close() (err error) {
	if s == nil {
		return
	}

	s.closeOnce.Do(func() {
		s.games.Range(func(_, value any) bool {
			value.(*game.Game).Stop()
			return true
		})

		s.closeErr = s.server.Close()
		var listenErr error = <-s.listenDone
		if s.closeErr == nil && listenErr != nil && !errors.Is(listenErr, net.ErrClosed) {
			s.closeErr = listenErr
		}
	})

	err = s.closeErr
	return
}

// bindHandlers attaches transport events to isolated local game sessions.
func (s *LocalServer) bindHandlers() {
	s.server.OnClient(func(serverClient *arctic.ServerClient) {
		var metadata map[string]any = serverClient.Metadata()
		if fmt.Sprintf("%v", metadata["key"]) != MetaAllowedKey {
			s.log.Warningf("client %d provided an invalid local key", serverClient.ID())
			var _ error = serverClient.Close()
			return
		}

		var (
			authoritative *game.Game   = game.NewGame(game.Configuration{})
			gameClient    *game.Client = game.NewClient(authoritative, serverClient)
		)

		s.games.Store(serverClient.ID(), authoritative)
		authoritative.Start()
		s.log.Infof("client %d connected\n", serverClient.ID())

		serverClient.OnMessage(func(data []byte) {
			if len(data) == 0 || data[0] <= shared.MsgTypeGameCutoffConst {
				return
			}

			gameClient.OnMessage(shared.NewReader(data))
		})

		serverClient.OnClose(func() {
			gameClient.Remove()
			authoritative.Stop()
			s.games.Delete(serverClient.ID())
			s.log.Infof("client %d disconnected\n", serverClient.ID())
		})

		serverClient.OnError(func(clientErr error) {
			s.log.Errorf("client %d error: %v", serverClient.ID(), clientErr)
		})
	})

	s.server.OnError(func(serverErr error) {
		s.log.Errorf("server error: %v", serverErr)
	})
}
