package gonvs

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
)

func (c *Client) sendRequest(methodName string, resultPointer interface{}, requestData []interface{}) error {
	requestFields := []json.RawMessage{}
	for i := 0; i < len(requestData); i++ {
		fieldData := requestData[i]
		switch fieldData.(type) {
		default:
			return errors.New("unknown field value: " + reflect.ValueOf(fieldData).String())
		case string:
			requestFields = append(requestFields, wrapJSONParam(fieldData.(string)))
			break
		case int:
			requestFields = append(requestFields, json.RawMessage(strconv.Itoa(fieldData.(int))))
			break
		case int64:
			requestFields = append(requestFields, json.RawMessage(
				strconv.FormatInt(fieldData.(int64), 10)),
			)
			break
		case ValueType:
			requestFields = append(requestFields, wrapJSONParam(string(fieldData.(ValueType))))
			break
		}
	}

	response, err := c.RPC.RawRequest(methodName, requestFields)
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
