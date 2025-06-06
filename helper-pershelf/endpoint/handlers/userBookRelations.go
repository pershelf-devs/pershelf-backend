package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/valyala/fasthttp"

	"github.com/pershelf/pershelf/crud"
	"github.com/pershelf/pershelf/endpoint/response"
)

// GetAllUserBookRelationsHandler gets all user book relations
func GetAllUserBookRelationsHandler(ctx *fasthttp.RequestCtx) {
	var userBookRelations []crud.UserBookRelation
	if userBookRelations = crud.GetAllUserBookRelations(); userBookRelations == nil {
		log.Printf("(Error): error retrieving user book relations list at endpoint (%s).", string(ctx.Path()))
		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "3", Values: []string{"Error retrieving user book relations list"}},
			UserBookRelations: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
		}
		return
	}

	log.Printf("(Information): User book relations list retrieved successfully (length: %d).", len(userBookRelations))
	if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
		Status:            response.ResponseMessage{Code: "0", Values: nil},
		UserBookRelations: userBookRelations,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
	}
}

// GetUserBookRelationByIDHandler gets a user book relation by id
func GetUserBookRelationByIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth                     = ctx.Path()
		userBookRelationID, err = strconv.Atoi(ctx.UserValue("id").(string))
		userBookRelation        crud.UserBookRelation
	)

	if err != nil {
		log.Printf("(Error): error converting user book relation ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "3", Values: []string{"Error converting user book relation ID to int"}},
			UserBookRelations: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if userBookRelationID <= 0 {
		log.Printf("(Error): invalid user book relation ID (retrieved: %d) at endpoint (%s).", userBookRelationID, string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "3", Values: []string{"Invalid user book relation ID"}},
			UserBookRelations: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if userBookRelation = crud.GetUserBookRelationByID(userBookRelationID); userBookRelation.ID == 0 {
		log.Printf("(Error): error retrieving user book relation at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "3", Values: []string{"Error retrieving user book relation"}},
			UserBookRelations: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): User book relation retrieved successfully (id: %d).", userBookRelation.ID)
	if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
		Status:            response.ResponseMessage{Code: "0", Values: nil},
		UserBookRelations: []crud.UserBookRelation{userBookRelation},
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// GetUserBookRelationsByUserIDHandler gets all user book relations by user id
func GetUserBookRelationsByUserIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth               = ctx.Path()
		userID, err       = strconv.Atoi(ctx.UserValue("user-id").(string))
		userBookRelations []crud.UserBookRelation
	)

	if err != nil {
		log.Printf("(Error): error converting user ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "3", Values: []string{"Error converting user ID to int"}},
			UserBookRelations: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if userID <= 0 {
		log.Printf("(Error): invalid user ID (retrieved: %d) at endpoint (%s).", userID, string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "3", Values: []string{"Invalid user ID"}},
			UserBookRelations: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if userBookRelations = crud.GetUserBookRelationsByUserID(userID); userBookRelations == nil {
		log.Printf("(Error): error retrieving user book relations list at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "3", Values: []string{"Error retrieving user book relations list"}},
			UserBookRelations: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): User book relations list retrieved successfully (length: %d).", len(userBookRelations))
	if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
		Status:            response.ResponseMessage{Code: "0", Values: nil},
		UserBookRelations: userBookRelations,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// GetUserBookRelationsByBookIDHandler gets all user book relations by book id
func GetUserBookRelationsByBookIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth               = ctx.Path()
		bookID, err       = strconv.Atoi(ctx.UserValue("book-id").(string))
		userBookRelations []crud.UserBookRelation
	)

	if err != nil {
		log.Printf("(Error): error converting book ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "3", Values: []string{"Error converting book ID to int"}},
			UserBookRelations: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if bookID <= 0 {
		log.Printf("(Error): invalid book ID (retrieved: %d) at endpoint (%s).", bookID, string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "3", Values: []string{"Invalid book ID"}},
			UserBookRelations: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if userBookRelations = crud.GetUserBookRelationsByBookID(bookID); userBookRelations == nil {
		log.Printf("(Error): error retrieving user book relations list at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "3", Values: []string{"Error retrieving user book relations list"}},
			UserBookRelations: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): User book relations list retrieved successfully (length: %d).", len(userBookRelations))
	if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
		Status:            response.ResponseMessage{Code: "0", Values: nil},
		UserBookRelations: userBookRelations,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// CreateUserBookRelationHandler creates a user book relation
func CreateUserBookRelationHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth              = ctx.Path()
		userBookRelation crud.UserBookRelation
	)

	if err := json.Unmarshal(ctx.Request.Body(), &userBookRelation); err != nil {
		log.Printf("(Error): error decoding request body at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "3", Values: []string{"Error decoding request body"}},
			UserBookRelations: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if err := crud.CreateUserBookRelation(&userBookRelation); err != nil {
		log.Printf("(Error): error creating user book relation at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "3", Values: []string{"Error creating user book relation"}},
			UserBookRelations: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): User book relation created successfully (id: %d).", userBookRelation.ID)
	if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
		Status:            response.ResponseMessage{Code: "0", Values: nil},
		UserBookRelations: []crud.UserBookRelation{userBookRelation},
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// UpdateUserBookRelationHandler updates a user book relation
func UpdateUserBookRelationHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth              = ctx.Path()
		userBookRelation crud.UserBookRelation
	)

	if err := json.Unmarshal(ctx.Request.Body(), &userBookRelation); err != nil {
		log.Printf("(Error): error decoding request body at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error decoding request body"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if err := crud.UpdateUserBookRelation(userBookRelation); err != nil {
		log.Printf("(Error): error updating user book relation at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error updating user book relation"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): User book relation updated successfully (id: %d).", userBookRelation.ID)
	if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
		Status:            response.ResponseMessage{Code: "0", Values: nil},
		UserBookRelations: []crud.UserBookRelation{userBookRelation},
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// DeleteUserBookRelationByIDHandler deletes a user book relation by id
func DeleteUserBookRelationByIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth                     = ctx.Path()
		userBookRelationID, err = strconv.Atoi(ctx.UserValue("id").(string))
	)

	if err != nil {
		log.Printf("(Error): error converting user book relation ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error converting user book relation ID to int"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if userBookRelationID <= 0 {
		log.Printf("(Error): invalid user book relation ID (retrieved: %d) at endpoint (%s).", userBookRelationID, string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Invalid user book relation ID"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if err := crud.DeleteUserBookRelationByID(userBookRelationID); err != nil {
		log.Printf("(Error): error deleting user book relation at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error deleting user book relation"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}
}

// DeleteUserBookRelationByUserIDHandler deletes a user book relation by user id
func DeleteUserBookRelationByUserIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		userID, err = strconv.Atoi(ctx.UserValue("user-id").(string))
	)

	if err != nil {
		log.Printf("(Error): error converting user ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error converting user ID to int"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if userID <= 0 {
		log.Printf("(Error): invalid user ID (retrieved: %d) at endpoint (%s).", userID, string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Invalid user ID"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if err := crud.DeleteUserBookRelationByUserID(userID); err != nil {
		log.Printf("(Error): error deleting user book relation at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error deleting user book relation"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): User book relation deleted successfully (user ID: %d).", userID)
	if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: nil}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// DeleteUserBookRelationByBookIDHandler deletes a user book relation by book id
func DeleteUserBookRelationByBookIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		bookID, err = strconv.Atoi(ctx.UserValue("book-id").(string))
	)

	if err != nil {
		log.Printf("(Error): error converting book ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error converting book ID to int"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if bookID <= 0 {
		log.Printf("(Error): invalid book ID (retrieved: %d) at endpoint (%s).", bookID, string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Invalid book ID"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if err := crud.DeleteUserBookRelationByBookID(bookID); err != nil {
		log.Printf("(Error): error deleting user book relation at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error deleting user book relation"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): User book relation deleted successfully (book ID: %d).", bookID)
	if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: nil}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// DeleteUserBookRelationByUserIDAndBookIDHandler deletes a user book relation by user id and book id
func DeleteUserBookRelationByUserIDAndBookIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth          = ctx.Path()
		userID, err1 = strconv.Atoi(ctx.UserValue("user-id").(string))
		bookID, err2 = strconv.Atoi(ctx.UserValue("book-id").(string))
	)

	if err1 != nil {
		log.Printf("(Error): error converting user ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error converting user ID to int"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if err2 != nil {
		log.Printf("(Error): error converting book ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error converting book ID to int"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if userID <= 0 {
		log.Printf("(Error): invalid user ID (retrieved: %d) at endpoint (%s).", userID, string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Invalid user ID"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if bookID <= 0 {
		log.Printf("(Error): invalid book ID (retrieved: %d) at endpoint (%s).", bookID, string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Invalid book ID"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if err := crud.DeleteUserBookRelationByUserIDAndBookID(userID, bookID); err != nil {
		log.Printf("(Error): error deleting user book relation at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error deleting user book relation"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): User book relation deleted successfully (user ID: %d, book ID: %d).", userID, bookID)
	if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: nil}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}
