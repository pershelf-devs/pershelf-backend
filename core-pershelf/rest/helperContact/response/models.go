package response

import "github.com/core-pershelf/rest/helperContact/tablesModels"

type ResponseMessage struct {
	Code   string   `json:"code"`
	Values []string `json:"values"`
}

type UsersResp struct {
	Status ResponseMessage     `json:"status"`
	Users  []tablesModels.User `json:"users"`
}
