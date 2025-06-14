package customizedmodels

import (
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

type DetailedReview struct {
	Review tablesModels.Review `json:"review,omitempty"`
	User   tablesModels.User   `json:"user,omitempty"`
	Book   tablesModels.Book   `json:"book,omitempty"`
}
