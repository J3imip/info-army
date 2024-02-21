package helpers

import (
	"encoding/json"
	"io"

	"github.com/J3imip/info-army/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func ParseMessage(request io.ReadCloser) (*types.Update, error) {
	message := &types.Update{}

	err := json.NewDecoder(request).Decode(&message)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode message")
	}

	return message, nil
}
