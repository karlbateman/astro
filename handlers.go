package astro

// Handlers registers handlers which respond to messages received from STDIN.
func (s *Server) Handlers() {
	s.HandleEcho()
}
