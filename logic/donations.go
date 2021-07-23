package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/advita-comics/advita-comics-backend/db"
	"github.com/advita-comics/advita-comics-backend/db/dao"
	"github.com/advita-comics/advita-comics-backend/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Donations - бизнесс логика пожертвований
type Donations interface {
	NewDonation(ctx context.Context, donation *types.Donation) error
}
type donations struct {
	db db.DB
}

// NewDonationManager - конструктор
func NewDonationManager(db db.DB) Donations {
	return donations{db: db}
}

// NewDonation - создает новое пожертвование
func (d donations) NewDonation(ctx context.Context, donation *types.Donation) error {
	comics, err := d.db.Dao().ComicsDao().List(ctx, map[string]interface{}{dao.ID: donation.ComicsID})
	if err != nil {
		return errors.Wrapf(err, "d.db.Dao().CompanyDao().List(id: <%d> ) ", donation.ComicsID)
	}
	if len(comics) == 0 {
		return fmt.Errorf("комикс с id <%d> не найден", donation.ComicsID)
	}

	users, err := d.db.Dao().UserDao().List(ctx, map[string]interface{}{dao.Email: donation.UserEmail})
	if err != nil {
		return errors.Wrapf(err, "d.db.Dao().UserDao().List(email: <%s> ) ", donation.UserEmail)
	}

	if err = d.db.Repo().Transaction(ctx, func(ctx context.Context) error {
		var user *dao.User
		if len(users) == 0 {
			user, err = d.createUser(ctx, donation)
			if err != nil {
				return errors.Wrapf(err, "d.createUser(email: <%s> ) ", donation.UserEmail)
			}
		} else {
			user = &users[0]
		}

		if user.FollowProcess != donation.SubscribedToProgress || user.GetReport != donation.SubscribedToReport {
			user, err = d.updateUser(ctx, user, donation)
			if err != nil {
				return errors.Wrapf(err, "d.update(email: <%s> ) ", donation.UserEmail)
			}
		}
		donationType, err := d.getDonationType(ctx, donation)
		if err != nil {
			return errors.Wrap(err, "d.getDonationType")
		}

		companies, err := d.db.Dao().CompanyDao().List(ctx, nil)
		if err != nil {
			return errors.Wrap(err, "d.getDonationType")
		}
		if len(companies) != 1 {
			return fmt.Errorf("колличество активных компаний <%d> вместо 1", len(companies))
		}

		personalisation, err := d.getPersonalisation(donation.Personalisation)
		if err != nil {
			return errors.Wrap(err, "d.getPersonalisation")
		}

		_, err = d.db.Dao().DonationDao().Create(ctx, &dao.Donation{
			Amount:          donation.Amount,
			CompanyID:       companies[0].ID,
			DonationTypeID:  donationType.ID,
			UserID:          user.ID,
			ComicsID:        donation.ComicsID,
			Personalisation: personalisation,
			Status:          dao.ACTIVE,
		})
		if err != nil {
			return errors.Wrap(err, "d.db.Dao().DonationDao().Create")
		}
		return nil
	}); err != nil {
		return err
	}

	logrus.Infof("new donation created email: <%s>, amount: <%f>", donation.UserEmail, donation.Amount)
	return nil
}

func (d donations) createUser(ctx context.Context, donation *types.Donation) (*dao.User, error) {
	logrus.Infof("creating new user <%s>", donation.UserEmail)

	return d.db.Dao().UserDao().Create(ctx, &dao.User{
		Role:          dao.USER,
		Name:          donation.UserName,
		Email:         donation.UserEmail,
		GetReport:     donation.SubscribedToReport,
		FollowProcess: donation.SubscribedToProgress,
		Status:        dao.ACTIVE,
	})
}

func (d donations) updateUser(ctx context.Context, user *dao.User, donation *types.Donation) (*dao.User, error) {
	logrus.Infof("updating user <%s>", donation.UserEmail)

	user.FollowProcess = donation.SubscribedToProgress
	user.GetReport = donation.SubscribedToReport
	return d.db.Dao().UserDao().Update(ctx, user)
}

func (d donations) getDonationType(ctx context.Context, donation *types.Donation) (*dao.DonationType, error) {
	dTypes, err := d.db.Dao().DonationDao().ListDonationTypes(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "d.db.Dao().DonationDao().ListDonationTypes")
	}

	for _, dt := range dTypes {
		if donation.Amount >= dt.MinAmount {
			return &dt, err
		}
	}

	return nil, fmt.Errorf("вариант донейшена для суммы <%f> не найден", donation.Amount)
}

func (d donations) getPersonalisation(p *types.Personalisation) (string, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return "", errors.Wrap(err, "json.Marshal")
	}

	return string(b), nil
}
