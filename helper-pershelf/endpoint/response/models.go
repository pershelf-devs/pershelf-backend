package response

import "github.com/pershelf/pershelf/crud"

type ResponseMessage struct {
	Code   string   `json:"code"`
	Values []string `json:"values"`
}

type UsersResp struct {
	Status ResponseMessage `json:"status"`
	Users  []crud.User     `json:"users"`
}
