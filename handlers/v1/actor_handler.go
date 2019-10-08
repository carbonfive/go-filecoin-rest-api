package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/carbonfive/go-filecoin-rest-api/types"
)

// ActorHandler is a handler for the actors/{actorId} endpoint
type ActorHandler struct {
	Callback func(actorId string) (*types.Actor, error)
}

// ServeHTTP handles an HTTP request.
func (a *ActorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	actorId := chi.URLParam(r, "actorId")
	var marshaled []byte

	actor, err := a.Callback(actorId)
	if err != nil {
		marshaled = types.MarshalErrors([]string{err.Error()})
	} else {
		actor.Kind = "actor"
		if marshaled, err = json.Marshal(actor); err != nil {
			log.Error(err)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if _,err = fmt.Fprint(w, string(marshaled[:])); err != nil {
		log.Error(err)
	}
}
