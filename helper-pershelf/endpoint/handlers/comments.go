package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pershelf/pershelf/crud"
	"github.com/pershelf/pershelf/endpoint/response"
	"github.com/valyala/fasthttp"
)

// GetAllCommentsHandler retrieves all comments from the database.
func GetAllCommentsHandler(ctx *fasthttp.RequestCtx) {
	var comments []crud.Comment
	if comments = crud.GetAllComments(); comments == nil {
		log.Printf("(Error): error retrieving comments list at endpoint (%s).", string(ctx.Path()))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): comments list retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(response.CommentsResp{
		Status:   response.ResponseMessage{Code: "0", Values: nil},
		Comments: comments,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
	}
}

// GetCommentByIDHandler retrieves a comment by ID from the database.
func GetCommentByIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth      = ctx.Path()
		commentID, err = strconv.Atoi(ctx.UserValue("id").(string))
		comment  crud.Comment
	)

	if err != nil {
		log.Printf("(Error): error converting comment ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	comment = crud.GetCommentByID(commentID)
	if comment.ID == 0 {
		log.Printf("(Error): comment not found by ID at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	log.Printf("(Information): comment retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(comment); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// CreateCommentHandler creates a new comment in the database.
func CreateCommentHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth     = ctx.Path()
		comment crud.Comment
	)

	if err := json.Unmarshal(ctx.Request.Body(), &comment); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	comment = crud.CreateComment(&comment)
	if comment.ID == 0 {
		log.Printf("(Error): error creating comment at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): comment created successfully.")
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

// UpdateCommentHandler updates an existing comment in the database.
func UpdateCommentHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth     = ctx.Path()
		comment crud.Comment
	)

	if err := json.Unmarshal(ctx.Request.Body(), &comment); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	comment = crud.UpdateComment(comment)
	if comment.ID == 0 {
		log.Printf("(Error): error updating comment at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): comment updated successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// DeleteCommentHandler deletes a comment from the database.
func DeleteCommentHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth      = ctx.Path()
		commentID, err = strconv.Atoi(ctx.UserValue("id").(string))
	)

	if err != nil {
		log.Printf("(Error): error converting comment ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if err := crud.DeleteComment(commentID); err != nil {
		log.Printf("(Error): error deleting comment at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): comment deleted successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
