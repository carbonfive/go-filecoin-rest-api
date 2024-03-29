package v1

import (
	"net/http"

	"github.com/carbonfive/go-filecoin-rest-api/types"
)

// NodeHandler is the handler for the control/node endpoint
type NodeHandler struct {
	Callback func() (*types.Node, error)
}

func (nid *NodeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	node, err := nid.Callback()
	Respond(w, node, err)
}
