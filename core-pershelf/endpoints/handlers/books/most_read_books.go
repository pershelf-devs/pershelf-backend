package books

import (
	"context"
	"encoding/json"
	"log"

	"github.com/core-pershelf/globals"
	"github.com/core-pershelf/mongo/response"
	"github.com/core-pershelf/mongo/tablesModels"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
)

// GetMostReadBooksHandler returns the most read books (default limit is 10)
func GetMostReadBooksHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var request struct {
			Limit    int    `json:"limit"`
			Category string `json:"category"`
		}

		if err := json.Unmarshal(ctx.Request.Body(), &request); err != nil {
			request.Limit = 10
		}

		// MongoDB aggregation pipeline to get most read books
		pipeline := []bson.M{
			{
				"$lookup": bson.M{
					"from":         "books",
					"localField":   "book_id",
					"foreignField": "_id",
					"as":           "book",
				},
			},
			{
				"$unwind": "$book",
			},
			{
				"$group": bson.M{
					"_id": "$book",
					"read_count": bson.M{
						"$sum": 1,
					},
				},
			},
			{
				"$sort": bson.M{
					"read_count": -1,
				},
			},
		}

		// Add category filter if provided
		if request.Category != "" {
			pipeline = append([]bson.M{
				{
					"$match": bson.M{
						"book.category": request.Category,
					},
				},
			}, pipeline...)
		}

		// Add limit
		pipeline = append(pipeline, bson.M{
			"$limit": request.Limit,
		})

		// Execute aggregation
		cursor, err := globals.UserBooksCollection.Aggregate(context.Background(), pipeline)
		if err != nil {
			log.Printf("Error executing aggregation at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error executing aggregation"}}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}
		defer cursor.Close(context.Background())

		// Get book details from aggregation results
		var results []struct {
			Book      tablesModels.Book `bson:"_id"`
			ReadCount int               `bson:"read_count"`
		}
		if err := cursor.All(context.Background(), &results); err != nil {
			log.Printf("Error decoding aggregation results at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error decoding aggregation results"}}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Convert results to books
		var mostReadBooks []tablesModels.Book
		for _, result := range results {
			mostReadBooks = append(mostReadBooks, result.Book)
		}

		// Return the results
		if err := json.NewEncoder(ctx).Encode(response.BooksResp{
			Status: response.ResponseMessage{Code: "0", Values: []string{"Most read books fetched successfully"}},
			Books:  mostReadBooks,
		}); err != nil {
			log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
		}
	}
}
