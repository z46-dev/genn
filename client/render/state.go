package render

func (s *RendererState) Release() {
	if s.Cleanup != nil {
		s.Cleanup()
		s.Cleanup = nil
	}

	if s.Ctx != nil {
		s.Ctx.Release()
		s.Ctx = nil
	}
}
