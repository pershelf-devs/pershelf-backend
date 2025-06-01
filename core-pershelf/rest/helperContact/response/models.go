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

type RefreshTokensResp struct {
	Status        ResponseMessage             `json:"status"`
	RefreshTokens []tablesModels.RefreshToken `json:"refreshTokens"`
}

type UserBooksResp struct {
	Status    ResponseMessage         `json:"status"`
	UserBooks []tablesModels.UserBook `json:"userBooks"`
}

type ReviewsResp struct {
	Status  ResponseMessage       `json:"status"`
	Reviews []tablesModels.Review `json:"reviews"`
}

type BooksResp struct {
	Status ResponseMessage     `json:"status"`
	Books  []tablesModels.Book `json:"books"`
}

type ShelfBooksResp struct {
	Status     ResponseMessage          `json:"status"`
	ShelfBooks []tablesModels.ShelfBook `json:"shelfBooks"`
}

type UserShelfsResp struct {
	Status     ResponseMessage          `json:"status"`
	UserShelfs []tablesModels.UserShelf `json:"userShelfs"`
}

type FollowsResp struct {
	Status  ResponseMessage       `json:"status"`
	Follows []tablesModels.Follow `json:"follows"`
}

type CommentsResp struct {
	Status   ResponseMessage        `json:"status"`
	Comments []tablesModels.Comment `json:"comments"`
}

type BookLikesResp struct {
	Status    ResponseMessage         `json:"status"`
	BookLikes []tablesModels.BookLike `json:"bookLikes"`
}
