package reviews

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/internal/helperUtils/reviewUtils"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/valyala/fasthttp"
)

// GetDetailedReviewsByBookIDHandler gets detailed reviews by book ID
func GetDetailedReviewsByBookIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var bookID int
		if err := json.Unmarshal(ctx.Request.Body(), &bookID); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.DetailedReviewsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling request body"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		detailedReviews, err := reviewUtils.GetDetailedReviewsByBookID(bookID)
		if err != nil {
			log.Printf("Error retrieving detailed reviews by book ID at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.DetailedReviewsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Error retrieving detailed reviews by book ID"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		for i, _ := range detailedReviews {
			detailedReviews[i].User.Password = ""
			detailedReviews[i].Book.ImageBase64 = ""
		}

		log.Printf("(Information): detailed reviews retrieved successfully.")
		if err := json.NewEncoder(ctx).Encode(response.DetailedReviewsResp{
			Status:          response.ResponseMessage{Code: "0", Values: nil},
			DetailedReviews: detailedReviews,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
		}
	}
}
