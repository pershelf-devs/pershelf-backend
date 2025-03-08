package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pershelf/pershelf/crud"
	"github.com/pershelf/pershelf/endpoint/response"
	"github.com/valyala/fasthttp"
)

// GetAllRefreshTokensHandler retrieves all refresh tokens from the database and sends them in a response to client's request.
func GetAllRefreshTokensHandler(ctx *fasthttp.RequestCtx) {
	var refreshTokens []crud.RefreshToken
	if refreshTokens = crud.GetAllRefreshTokens(); refreshTokens == nil {
		log.Printf("(Error): error retrieving refresh tokens list at endpoint (%s).", string(ctx.Path()))
		if err := json.NewEncoder(ctx).Encode(response.RefreshTokensResp{
			Status:        response.ResponseMessage{Code: "4", Values: nil},
			RefreshTokens: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
		}
		return
	}

	log.Printf("(Information): refresh tokens list retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.RefreshTokensResp{
		Status:        response.ResponseMessage{Code: "0", Values: nil},
		RefreshTokens: refreshTokens,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
	}
}

// GetRefreshTokenByIDHandler retrieves a refresh token by id from the database and sends them in a response to client's request.
func GetRefreshTokenByIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth                 = ctx.Path()
		refreshTokenID, err = strconv.Atoi(ctx.UserValue("id").(string))
		RefreshToken        crud.RefreshToken
	)

	if err != nil {
		log.Printf("(Error): error converting refresh token id to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.RefreshTokensResp{
			Status:        response.ResponseMessage{Code: "4", Values: nil},
			RefreshTokens: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	RefreshToken = crud.GetRefreshTokenByID(refreshTokenID)
	if RefreshToken.ID == 0 {
		log.Printf("(Error): error retrieving refresh token by id at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.RefreshTokensResp{
			Status:        response.ResponseMessage{Code: "3", Values: nil},
			RefreshTokens: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): refresh token retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.RefreshTokensResp{
		Status:        response.ResponseMessage{Code: "0", Values: nil},
		RefreshTokens: []crud.RefreshToken{},
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// GetRefreshTokenByUserIDHandler retrieves a refresh token by user id from the database and send them in a response to client's request.
func GetRefreshTokenByUserIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth                 = ctx.Path()
		refreshTokenID, err = strconv.Atoi(ctx.UserValue("user-id").(string))
		RefreshToken        crud.RefreshToken
	)

	if err != nil {
		log.Printf("(Error): error converting refresh token refreshToken id to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.RefreshTokensResp{
			Status:        response.ResponseMessage{Code: "4", Values: nil},
			RefreshTokens: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	RefreshToken = crud.GetRefreshTokenByUserID(refreshTokenID)
	if RefreshToken.ID == 0 {
		log.Printf("(Error): error retrieving refresh token by refreshToken id at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.RefreshTokensResp{
			Status:        response.ResponseMessage{Code: "3", Values: nil},
			RefreshTokens: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): refresh token retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.RefreshTokensResp{
		Status:        response.ResponseMessage{Code: "0", Values: nil},
		RefreshTokens: []crud.RefreshToken{},
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// CreateRefreshTokenHandler creates a new refresh token in the database and sends them in a response to client's request.
func CreateRefreshTokenHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth          = ctx.Path()
		refreshToken crud.RefreshToken
	)

	if err := json.Unmarshal(ctx.Request.Body(), &refreshToken); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		if err := json.NewEncoder(ctx).Encode(response.RefreshTokensResp{Status: response.ResponseMessage{Code: "3", Values: nil}}); err != nil {
			log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
		}
		return
	}

	refreshToken = crud.CreateRefreshToken(&refreshToken)
	if refreshToken.ID == 0 {
		log.Printf("(Error): error creating refresh token at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.RefreshTokensResp{
			Status:        response.ResponseMessage{Code: "4", Values: nil},
			RefreshTokens: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): refresh token created successfully.")
	if err := json.NewEncoder(ctx).Encode(response.RefreshTokensResp{
		Status:        response.ResponseMessage{Code: "0", Values: nil},
		RefreshTokens: []crud.RefreshToken{refreshToken},
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// UpdateRefreshTokenHandler updates a refresh token in the database and sends them in a response to client's request.
func UpdateRefreshTokenHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth          = ctx.Path()
		refreshToken crud.RefreshToken
	)

	if err := json.Unmarshal(ctx.Request.Body(), &refreshToken); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: nil}); err != nil {
			log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
		}
		return
	}

	refreshToken = crud.UpdateRefreshToken(refreshToken)
	if refreshToken.ID == 0 {
		log.Printf("(Error): error updating refreshToken at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "4", Values: nil}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): refreshToken updated successfully.")
	if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: nil}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// DeleteRefreshTokenHandler deletes a refresh token from the database and sends them in a response to client's request.
func DeleteRefreshTokenHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth                 = ctx.Path()
		refreshTokenID, err = strconv.Atoi(ctx.UserValue("id").(string))
	)

	if err != nil {
		log.Printf("(Error): error converting refresh token id to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: nil}); err != nil {
			log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
		}
		return
	}

	if refreshTokenID == 0 {
		log.Printf("(Error): error retrieving refresh token by id at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: nil}); err != nil {
			log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
		}
		return
	}

	if err := crud.DeleteRefreshToken(refreshTokenID); err != nil {
		log.Printf("(Error): error deleting refresh token at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "4", Values: nil}); err != nil {
			log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
		}
		return
	}

	log.Printf("(Information): refresh token deleted successfully.")
	if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: nil}); err != nil {
		log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
	}
}
