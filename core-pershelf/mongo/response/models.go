package response

import "github.com/core-pershelf/mongo/tablesModels"

type ResponseMessage struct {
	Code   string   `json:"code"`
	Values []string `json:"values"`
}

type BooksResp struct {
	Status ResponseMessage     `json:"status"`
	Books  []tablesModels.Book `json:"books"`
}

type UserBooksResp struct {
	Status ResponseMessage         `json:"status"`
	Books  []tablesModels.UserBook `json:"books"`
}

type UsersResp struct {
	Status ResponseMessage     `json:"status"`
	Users  []tablesModels.User `json:"users"`
}
