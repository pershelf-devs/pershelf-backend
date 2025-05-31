package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pershelf/pershelf/crud"
	"github.com/pershelf/pershelf/endpoint/response"
	"github.com/valyala/fasthttp"
)

// GetAllFollowsHandler retrieves all follow relationships from the database.
func GetAllFollowsHandler(ctx *fasthttp.RequestCtx) {
	var follows []crud.Follow
	if follows = crud.GetAllFollows(); follows == nil {
		log.Printf("(Error): error retrieving follow list at endpoint (%s).", string(ctx.Path()))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): follow list retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(response.FollowsResp{
		Status:  response.ResponseMessage{Code: "0", Values: nil},
		Follows: follows,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
	}
}

// GetFollowByIDHandler retrieves a follow relationship by ID from the database.
func GetFollowByIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth           = ctx.Path()
		followID, err = strconv.Atoi(ctx.UserValue("id").(string))
		follow        crud.Follow
	)

	if err != nil {
		log.Printf("(Error): error converting follow ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	follow = crud.GetFollowByID(followID)
	if follow.ID == 0 {
		log.Printf("(Error): follow not found by ID at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	log.Printf("(Information): follow retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(follow); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// CreateFollowHandler creates a new follow relationship in the database.
func CreateFollowHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth    = ctx.Path()
		follow crud.Follow
	)

	if err := json.Unmarshal(ctx.Request.Body(), &follow); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	follow = crud.CreateFollow(&follow)
	if follow.ID == 0 {
		log.Printf("(Error): error creating follow at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): follow created successfully.")
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

// UpdateFollowHandler updates an existing follow relationship in the database.
func UpdateFollowHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth    = ctx.Path()
		follow crud.Follow
	)

	if err := json.Unmarshal(ctx.Request.Body(), &follow); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	follow = crud.UpdateFollow(follow)
	if follow.ID == 0 {
		log.Printf("(Error): error updating follow at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): follow updated successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// DeleteFollowHandler deletes a follow relationship from the database.
func DeleteFollowHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth           = ctx.Path()
		followID, err = strconv.Atoi(ctx.UserValue("id").(string))
	)

	if err != nil {
		log.Printf("(Error): error converting follow ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if err := crud.DeleteFollow(followID); err != nil {
		log.Printf("(Error): error deleting follow at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): follow deleted successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
