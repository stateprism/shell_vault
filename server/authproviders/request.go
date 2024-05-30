package authproviders

import "encoding/json"

type UsrPwRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RequestFromBytes(b []byte) (*UsrPwRequest, error) {
	var req UsrPwRequest
	err := json.Unmarshal(b, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
