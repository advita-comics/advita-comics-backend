package handlers

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/advita-comics/advita-comics-backend/db"
	"github.com/advita-comics/advita-comics-backend/types"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// CompanyHandler - ручки компаний
type CompanyHandler interface {
	CompanyInfo(c echo.Context) error
}

type companyHandler struct {
	db db.DB
}

// NewCompanyHandler - конструктор
func NewCompanyHandler(db db.DB) CompanyHandler {
	return companyHandler{db: db}
}

// CompanyInfo - ручка отдает информацию по активной компании
// активная компания может быть только одна, иначе вернется ошибка
func (d companyHandler) CompanyInfo(c echo.Context) error {
	ctx := context.Background()

	companies, err := d.db.Dao().CompanyDao().List(ctx, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResp{
			Error: errors.Wrap(err, "d.getDonationType").Error(),
		})
	}
	if len(companies) != 1 {
		return c.JSON(http.StatusInternalServerError, ErrorResp{
			Error: fmt.Errorf("колличество активных компаний <%d> вместо 1", len(companies)).Error(),
		})
	}

	var sum float64

	donations, err := d.db.Dao().DonationDao().Donations(ctx, companies[0].ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResp{
			Error: errors.Wrap(err, "d.db.Dao().DonationDao().Donations()").Error(),
		})
	}
	for _, donation := range donations {
		sum += donation.Amount
	}

	remainDays := math.Round(time.Until(companies[0].ExpirationDate).Hours() / 24)

	return c.JSON(http.StatusOK, types.CompanyInfo{
		TerminationAmount: companies[0].TerminationAmount,
		CollectedAmount:   sum,
		DayRemains:        remainDays,
		DonationCount:     len(donations),
	})
}
