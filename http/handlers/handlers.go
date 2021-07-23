package handlers

import (
	"github.com/advita-comics/advita-comics-backend/db"
)

// ErrorResp - структура стандартной ошибки
type ErrorResp struct {
	Error string `json:"error"`
}

// OkResp - структура стандартного удачного ответа
type OkResp struct {
	Message string `json:"message"`
}

// Handlers - контейнер хендлеров
type Handlers struct {
	Donations DonationHandler
	Company   CompanyHandler
}

// NewHandlers - конструктор
func NewHandlers(db db.DB) *Handlers {
	return &Handlers{
		Donations: NewDonationHandler(db),
		Company:   NewCompanyHandler(db),
	}
}
