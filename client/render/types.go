package render

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/gogpu/gogpu"
	"github.com/gogpu/gpucontext"
	"github.com/z46-dev/gctx2d"
)

type (
	Renderer struct {
		App               *gogpu.App
		Provider          gogpu.DeviceProvider
		higherCtx         *gctx2d.Context // Context for forcibly rendering things above the rest of the scenes...
		activeState       atomic.Pointer[RendererState]
		ready, activeWait chan error
		states            []*RendererState
		statesMutex       sync.Mutex
		readyOnce         sync.Once
		activeWaitOnce    sync.Once
		frames, fps       int
		lastFPSUpdate     time.Time
		LegacyFont        *gctx2d.FontHandle
	}

	RendererState struct {
		Renderer                     *Renderer
		Ctx                          *gctx2d.Context
		Target                       *gogpu.Texture
		Cleanup                      func()
		Draw                         func(dc *gogpu.Context)
		OnMousePress, OnMouseRelease func(button gpucontext.MouseButton, x, y float64)
		OnMouseMove                  func(x, y float64)
		OnScroll                     func(dx, dy float64)
		OnUpdate                     func(dt float64)
		OnResize                     func(width, height int)
		OnKeyPress                   func(key gpucontext.Key, mods gpucontext.Modifiers)
		OnKeyRelease                 func(key gpucontext.Key, mods gpucontext.Modifiers)
		OnTextInput                  func(text string)
	}
)
