package gonvs

import (
	"encoding/json"
	"errors"
)

func (c *Client) sendRequest(requestData []json.RawMessage, methodName string, resultPointer interface{}) error {
	response, err := c.RPC.RawRequest("", requestData)
	if err != nil {
		return errors.New("failed to get entrys: " + err.Error())
	}

	jsonBytes, err := response.MarshalJSON()
	if err != nil {
		return errors.New("failed to json encode client response: " + err.Error())
	}

	err = json.Unmarshal(jsonBytes, resultPointer)
	if err != nil {
		return errors.New("failed to decode json response: " + err.Error())
	}
	return nil
}
