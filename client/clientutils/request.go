package clientutils

import "encoding/json"

type UsrPwRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUsrPwRequest(username, password string) *UsrPwRequest {
	return &UsrPwRequest{
		Username: username,
		Password: password,
	}
}

func (r *UsrPwRequest) ToBytes() []byte {
	b, _ := json.Marshal(r)
	return b
}

func RequestFromBytes(b []byte) (*UsrPwRequest, error) {
	var req UsrPwRequest
	err := json.Unmarshal(b, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
