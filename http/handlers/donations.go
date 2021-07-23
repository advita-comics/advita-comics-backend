package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/advita-comics/advita-comics-backend/db"
	"github.com/advita-comics/advita-comics-backend/logic"
	"github.com/advita-comics/advita-comics-backend/types"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// DonationHandler - ручки пожертвований
type DonationHandler interface {
	Donation(c echo.Context) error
}

type donationHandler struct {
	manager logic.Donations
}

// NewDonationHandler - конструктор
func NewDonationHandler(db db.DB) DonationHandler {
	return donationHandler{manager: logic.NewDonationManager(db)}
}

// DonationInfo - ручка создает новое пожертвование
func (h donationHandler) Donation(c echo.Context) error {
	ctx := context.Background()

	r := &types.Donation{ComicsID: -1}
	if err := c.Bind(r); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, ErrorResp{
			Error: err.Error(),
		})
	}

	if err := h.validateDonation(r); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResp{
			Error: err.Error(),
		})
	}

	if err := h.manager.NewDonation(ctx, r); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResp{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, OkResp{
		Message: "Пожертвование создано",
	})
}

func (h *donationHandler) validateDonation(d *types.Donation) error {
	if d.ComicsID == -1 {
		return errors.New("'comicsId' не передан")
	}
	if d.Amount < 1 {
		return errors.New("пожертвование меньше 1")
	}
	if d.UserEmail == "" {
		return errors.New("'userEmail' не передан")
	}

	if d.UserName == "" {
		return errors.New("'userName' не передан")
	}

	return nil
}
