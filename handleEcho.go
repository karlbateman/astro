package astro

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// HandleEcho responds to echo requests from Maelstrom.
func (s *Server) HandleEcho() {
	s.Node.Handle("echo", func(m maelstrom.Message) error {
		var b map[string]any
		if err := json.Unmarshal(m.Body, &b); err != nil {
			return err
		}
		b["type"] = "echo_ok"
		return s.Node.Reply(m, b)
	})
}
