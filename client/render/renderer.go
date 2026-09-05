package render

import (
	"fmt"
	"math"
	"time"

	"github.com/gogpu/gogpu"
	"github.com/gogpu/gpucontext"
	"github.com/gogpu/wgpu"
	"github.com/z46-dev/gctx2d"
	"github.com/z46-dev/genn/client/assets/fonts"
	"github.com/z46-dev/golog"
)

func NewRenderer(log *golog.Logger) (renderer *Renderer, err error) {
	var cfg gogpu.Config = gogpu.DefaultConfig().
		WithTitle("genn").
		WithSize(1280, 720).
		WithContinuousRender(true).
		// WithPowerPreference(gogpu.PowerPreferenceHighPerformance).
		WithVSync(true).
		WithBackend(gogpu.BackendAuto)

	renderer = &Renderer{
		App:        gogpu.NewApp(cfg),
		Provider:   nil,
		higherCtx:  nil,
		ready:      make(chan error, 1),
		activeWait: make(chan error, 1),
	}

	if renderer.LegacyFont, err = fonts.LoadUbuntu(); err != nil {
		renderer = nil
		return
	}

	renderer.initEventHandlers()
	renderer.initDrawLoop()

	go func() {
		var runErr error
		if runErr = renderer.App.Run(); runErr != nil {
			runErr = fmt.Errorf("run app: %w", runErr)
		}

		renderer.signalStopped(runErr)
	}()

	if err = <-renderer.ready; err != nil {
		renderer.App.Close()
		renderer = nil
	}

	return
}

func (r *Renderer) initEventHandlers() {
	var events gpucontext.EventSource = r.App.EventSource()
	events.OnMousePress(func(button gpucontext.MouseButton, x, y float64) {
		var state *RendererState = r.activeState.Load()
		if state != nil && state.OnMousePress != nil {
			state.OnMousePress(button, x, y)
		}
	})

	events.OnMouseRelease(func(button gpucontext.MouseButton, x, y float64) {
		var state *RendererState = r.activeState.Load()
		if state != nil && state.OnMouseRelease != nil {
			state.OnMouseRelease(button, x, y)
		}
	})

	events.OnMouseMove(func(x, y float64) {
		var state *RendererState = r.activeState.Load()
		if state != nil && state.OnMouseMove != nil {
			state.OnMouseMove(x, y)
		}
	})

	events.OnScroll(func(dx, dy float64) {
		var state *RendererState = r.activeState.Load()
		if state != nil && state.OnScroll != nil {
			state.OnScroll(dx, dy)
		}
	})

	events.OnKeyPress(func(key gpucontext.Key, mods gpucontext.Modifiers) {
		var state *RendererState = r.activeState.Load()
		if state != nil && state.OnKeyPress != nil {
			state.OnKeyPress(key, mods)
		}
	})

	events.OnKeyRelease(func(key gpucontext.Key, mods gpucontext.Modifiers) {
		var state *RendererState = r.activeState.Load()
		if state != nil && state.OnKeyRelease != nil {
			state.OnKeyRelease(key, mods)
		}
	})

	events.OnTextInput(func(text string) {
		var state *RendererState = r.activeState.Load()
		if state != nil && state.OnTextInput != nil {
			state.OnTextInput(text)
		}
	})

	r.App.OnUpdate(func(dt float64) {
		var state *RendererState = r.activeState.Load()
		if state != nil && state.OnUpdate != nil {
			state.OnUpdate(dt)
		}
	})

	r.App.OnResize(func(width, height int) {
		var state *RendererState = r.activeState.Load()
		if state != nil && state.OnResize != nil {
			state.OnResize(width, height)
		}
	})
}

// Sets up draw loop, first frame will be used to initialize some states
func (r *Renderer) initDrawLoop() {
	var ready bool = false
	r.App.OnDraw(func(dc *gogpu.Context) {
		if !ready {
			if r.Provider == nil {
				r.Provider = r.App.DeviceProvider()

				if r.Provider == nil {
					r.signalReady(fmt.Errorf("GPU context provider is not available"))
					return
				}
			}

			if r.higherCtx == nil {
				var err error
				if r.higherCtx, err = gctx2d.NewContext(r.Provider.Device(), r.Provider.SurfaceFormat()); err != nil {
					r.signalReady(fmt.Errorf("create higher context: %w", err))
					return
				}
			}

			r.signalReady(nil)
			ready = true
			return
		}

		r.frames++

		var now time.Time = time.Now()
		if now.Sub(r.lastFPSUpdate) >= time.Second {
			r.fps = r.frames
			r.frames = 0
			r.lastFPSUpdate = now
		}

		var state *RendererState = r.activeState.Load()
		if state == nil || state.Draw == nil {
			var message string = fmt.Sprintf("No active renderer state! (%d states loaded)", r.stateCount())
			if state != nil {
				message = "State has no draw function!"
			}

			var surface *wgpu.TextureView
			if surface = dc.SurfaceView(); surface == nil {
				return
			}

			var width, height uint32 = dc.SurfaceSize()
			r.higherCtx.Begin(int(width), int(height))

			r.higherCtx.SetFillStyle(gctx2d.ColorBlack)
			r.higherCtx.BeginPath()
			r.higherCtx.Rect(0, 0, float32(width), float32(height))
			r.higherCtx.Fill()

			// Text
			r.higherCtx.SetFillStyle(gctx2d.ColorWhite)
			r.higherCtx.SetFont(24, nil, gctx2d.FontWeightMedium)
			r.higherCtx.SetTextAlign(gctx2d.TextAlignCenter)
			r.higherCtx.SetTextBaseline(gctx2d.TextBaselineMiddle)
			r.higherCtx.FillText(message, float32(width/2), float32(height/2))

			// Frowny faces
			var face func(x float32) = func(x float32) {
				r.higherCtx.BeginPath()

				// Eyes
				r.higherCtx.MoveTo(x-5, float32(height)/2-85)
				r.higherCtx.Arc(x-10, float32(height)/2-85, 5, 0, math.Pi*2, false)

				r.higherCtx.MoveTo(x+15, float32(height)/2-85)
				r.higherCtx.Arc(x+10, float32(height)/2-85, 5, 0, math.Pi*2, false)

				// Mouth
				r.higherCtx.MoveTo(x-10, float32(height)/2-65)
				r.higherCtx.Arc(x, float32(height)/2-65, 10, math.Pi, 0, false)

				r.higherCtx.SetStrokeStyle(gctx2d.ColorWhite)
				r.higherCtx.SetLineWidth(2)
				r.higherCtx.Stroke()
			}

			face(float32(width)/2 - 100)
			face(float32(width)/2 + 100)

			r.higherCtx.Flush(dc.SurfaceView(), nil)
		} else {
			state.Draw(dc)
		}
	})
}

func (r *Renderer) NewState() (state *RendererState, err error) {
	if r.Provider == nil {
		err = fmt.Errorf("GPU context provider is not available")
		return
	}

	state = &RendererState{
		Renderer: r,
		Ctx:      nil,
		Target:   nil,
	}

	if state.Ctx, err = gctx2d.NewContext(r.Provider.Device(), r.Provider.SurfaceFormat()); err != nil {
		err = fmt.Errorf("create context: %w", err)
		return
	}

	r.statesMutex.Lock()
	r.states = append(r.states, state)
	r.statesMutex.Unlock()
	return
}

func (r *Renderer) SetState(state *RendererState) {
	r.activeState.Store(state)
}

func (r *Renderer) WaitAndCleanup() (err error) {
	r.App.OnClose(func() { r.signalStopped(nil) })

	if err = <-r.activeWait; err != nil {
		return fmt.Errorf("active wait: %w", err)
	}

	if r.higherCtx != nil {
		r.higherCtx.Release()
		r.higherCtx = nil
	}

	for _, state := range r.states {
		state.Release()
	}

	r.App.Close()
	return
}

// stateCount returns the number of renderer states currently owned by the renderer.
func (r *Renderer) stateCount() (count int) {
	r.statesMutex.Lock()
	count = len(r.states)
	r.statesMutex.Unlock()
	return
}

// signalReady publishes renderer initialization exactly once.
func (r *Renderer) signalReady(err error) {
	r.readyOnce.Do(func() {
		r.ready <- err
	})
}

// signalStopped publishes renderer termination exactly once.
func (r *Renderer) signalStopped(err error) {
	r.activeWaitOnce.Do(func() {
		r.activeWait <- err
	})
}

func (r *Renderer) GetFPS() int {
	return max(1, r.fps)
}
