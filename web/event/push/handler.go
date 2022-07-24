package push

import (
	"encoding/json"
	"gitrabbit/event"
)

func Handle(j *[]byte) (*event.MessageNotify, error) {
	var pr PushRequest
	err := json.Unmarshal(*j, &pr)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
