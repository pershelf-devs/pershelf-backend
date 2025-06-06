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

type RefreshTokensResp struct {
	Status        ResponseMessage     `json:"status"`
	RefreshTokens []crud.RefreshToken `json:"refreshTokens"`
}

type UserBooksResp struct {
	Status    ResponseMessage `json:"status"`
	UserBooks []crud.UserBook `json:"userBooks"`
}

type ReviewsResp struct {
	Status  ResponseMessage `json:"status"`
	Reviews []crud.Review   `json:"reviews"`
}

type BooksResp struct {
	Status ResponseMessage `json:"status"`
	Books  []crud.Book     `json:"books"`
}

type ShelfBooksResp struct {
	Status     ResponseMessage  `json:"status"`
	ShelfBooks []crud.ShelfBook `json:"shelfBooks"`
}

type UserShelfsResp struct {
	Status     ResponseMessage  `json:"status"`
	UserShelfs []crud.UserShelf `json:"userShelfs"`
}

type FollowsResp struct {
	Status  ResponseMessage `json:"status"`
	Follows []crud.Follow   `json:"follows"`
}

type CommentsResp struct {
	Status   ResponseMessage `json:"status"`
	Comments []crud.Comment  `json:"comments"`
}

type BookLikesResp struct {
	Status    ResponseMessage `json:"status"`
	BookLikes []crud.BookLike `json:"bookLikes"`
}

type UserBookRelationsResp struct {
	Status            ResponseMessage         `json:"status"`
	UserBookRelations []crud.UserBookRelation `json:"userBookRelations"`
}
