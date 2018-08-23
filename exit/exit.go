package exit

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"onion-router/comm"
)

/*
 * Handle() takes a message, makes a request to the
 * external website, and then returns the resulting
 * payload.
 */
func Handle(message comm.ExitMessage) (*comm.ExitMessage, error) {
	res, err := http.Get(message.Address)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to reach exit address")
	}
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read exit response body")
	}
	return &comm.ExitMessage{
		Address: message.Address,
		Payload: string(contents),
	}, nil
}
