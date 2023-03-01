package astro

import (
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// New creates a new Astro server object.
func New() *Server {
	s := &Server{Node: maelstrom.NewNode()}
	s.Handlers()
	return s
}


// Server wraps an Astro server and it's dependencies.
type Server struct {
	Node *maelstrom.Node
}

// Run launches the Astro servers main event-handling loop. It reads messages 
// from STDIN and delegates them to their corresponding registered handler. 
func (s *Server) Run() error {
	return s.Node.Run()
}
