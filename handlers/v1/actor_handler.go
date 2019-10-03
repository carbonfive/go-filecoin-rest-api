package v1

import (
	"encoding/json"
	"fmt"
	"github.com/carbonfive/go-filecoin-rest-api/types"
	"github.com/carbonfive/go-filecoin-rest-api/types/api_errors"
	"github.com/gorilla/mux"
	"net/http"
)

// ActorHandler is a handler for the Actor endpoint
type ActorHandler struct {
	Callback func(actorId string) (*types.Actor, error)
}

// ServeHTTP handles an HTTP request.
func (a *ActorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var marshaled []byte

	result, err := a.Callback(vars["actorId"])
	if err != nil {
		marshaled = api_errors.MarshalErrors([]string{err.Error()})
	} else {
		marshaled, _ = json.Marshal(result)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(marshaled[:])) // nolint: errcheck
}
