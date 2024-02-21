package requests

import (
	"encoding/json"
	"net/http"

	"github.com/J3imip/info-army/types"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewUpdate(r *http.Request) (*types.Update, error) {
	var message types.Update

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		return nil, validation.Errors{
			"/": errors.Wrap(err, "failed to unmarshal message"),
		}.Filter()
	}

	return &message, nil
}
