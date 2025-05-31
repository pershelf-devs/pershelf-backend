package books

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookController interface {
	CreateBook(ctx *fasthttp.RequestCtx)
	GetBookByID(ctx *fasthttp.RequestCtx)
	GetBooksByOwnerID(ctx *fasthttp.RequestCtx)
	GetBooksByStatus(ctx *fasthttp.RequestCtx)
	UpdateBook(ctx *fasthttp.RequestCtx)
	DeleteBook(ctx *fasthttp.RequestCtx)
}

type BookControllerImpl struct {
	service BookService
}

func NewBookController(service BookService) *BookControllerImpl {
	return &BookControllerImpl{service: service}
}

func (c *BookControllerImpl) CreateBook(ctx *fasthttp.RequestCtx) {
	log.Printf("DEBUG: Controller layer - CreateBook endpoint called")

	var book Book
	if err := json.Unmarshal(ctx.PostBody(), &book); err != nil {
		log.Printf("ERROR: Failed to unmarshal request body: %v", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid request body"}`)
		return
	}

	if err := c.service.CreateBook(&book); err != nil {
		log.Printf("ERROR: Failed to create book: %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"error": "Failed to create book"}`)
		return
	}

	log.Printf("DEBUG: Book created successfully with ID: %v", book.ID)
	ctx.SetStatusCode(fasthttp.StatusCreated)
	json.NewEncoder(ctx).Encode(book)
}

func (c *BookControllerImpl) GetBookByID(ctx *fasthttp.RequestCtx) {
	log.Printf("DEBUG: Controller layer - GetBookByID endpoint called")

	idStr := ctx.UserValue("id").(string)
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Printf("ERROR: Invalid ID format: %v", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid ID format"}`)
		return
	}

	book, err := c.service.GetBookByID(id)
	if err != nil {
		log.Printf("ERROR: Failed to get book: %v", err)
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetBodyString(`{"error": "Book not found"}`)
		return
	}

	log.Printf("DEBUG: Book found with ID: %v", id)
	json.NewEncoder(ctx).Encode(book)
}

func (c *BookControllerImpl) GetBooksByOwnerID(ctx *fasthttp.RequestCtx) {
	log.Printf("DEBUG: Controller layer - GetBooksByOwnerID endpoint called")

	ownerIDStr := ctx.UserValue("owner_id").(string)
	ownerID, err := primitive.ObjectIDFromHex(ownerIDStr)
	if err != nil {
		log.Printf("ERROR: Invalid owner ID format: %v", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid owner ID format"}`)
		return
	}

	books, err := c.service.GetBooksByOwnerID(ownerID)
	if err != nil {
		log.Printf("ERROR: Failed to get books: %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"error": "Failed to get books"}`)
		return
	}

	log.Printf("DEBUG: Found %d books for owner ID: %v", len(books), ownerID)
	json.NewEncoder(ctx).Encode(books)
}

func (c *BookControllerImpl) GetBooksByStatus(ctx *fasthttp.RequestCtx) {
	log.Printf("DEBUG: Controller layer - GetBooksByStatus endpoint called")

	status := ctx.UserValue("status").(string)
	books, err := c.service.GetBooksByStatus(status)
	if err != nil {
		log.Printf("ERROR: Failed to get books: %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"error": "Failed to get books"}`)
		return
	}

	log.Printf("DEBUG: Found %d books with status: %s", len(books), status)
	json.NewEncoder(ctx).Encode(books)
}

func (c *BookControllerImpl) UpdateBook(ctx *fasthttp.RequestCtx) {
	log.Printf("DEBUG: Controller layer - UpdateBook endpoint called")

	var book Book
	if err := json.Unmarshal(ctx.PostBody(), &book); err != nil {
		log.Printf("ERROR: Failed to unmarshal request body: %v", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid request body"}`)
		return
	}

	if err := c.service.UpdateBook(&book); err != nil {
		log.Printf("ERROR: Failed to update book: %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"error": "Failed to update book"}`)
		return
	}

	log.Printf("DEBUG: Book updated successfully with ID: %v", book.ID)
	ctx.SetStatusCode(fasthttp.StatusOK)
	json.NewEncoder(ctx).Encode(book)
}

func (c *BookControllerImpl) DeleteBook(ctx *fasthttp.RequestCtx) {
	log.Printf("DEBUG: Controller layer - DeleteBook endpoint called")

	idStr := ctx.UserValue("id").(string)
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Printf("ERROR: Invalid ID format: %v", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid ID format"}`)
		return
	}

	if err := c.service.DeleteBook(id); err != nil {
		log.Printf("ERROR: Failed to delete book: %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"error": "Failed to delete book"}`)
		return
	}

	log.Printf("DEBUG: Book deleted successfully with ID: %v", id)
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBodyString(`{"message": "Book deleted successfully"}`)
}
