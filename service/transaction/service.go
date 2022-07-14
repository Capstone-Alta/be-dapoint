package transaction

import (
	"dapoint-api/entities"
	dapoint_api "dapoint-api/error"
	"github.com/go-playground/validator/v10"
)

type service struct {
	repository entities.TransactionRepository
	userRepo   entities.UserRepository
	validate   *validator.Validate
}

func NewService(repository entities.TransactionRepository, repoUser entities.UserRepository) entities.TransactionService {
	return &service{
		repository: repository,
		userRepo:   repoUser,
		validate:   validator.New(),
	}
}

func (s service) GetById(id uint64) (transaction entities.Transaction, err error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetAll() (transactions []entities.Transaction, err error) {
	//TODO implement me
	transactions, err = s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s service) Create(data entities.Transaction) (transaction entities.Transaction, err error) {
	//TODO implement me

	err = s.validate.Struct(&data)
	if err != nil {
		err = dapoint_api.ErrBadRequest
		return
	}

	totalBelanja := data.TotalBelanja
	var pointEarn uint
	if totalBelanja >= 1000 {
		pointEarn = totalBelanja / 1000
	}
	data.PointEarn = uint(pointEarn)
	var userData entities.User

	userData, err = s.userRepo.FindById(uint64(data.UserID))

	userData.TotalPoint = userData.TotalPoint + uint64(pointEarn)
	_, err = s.userRepo.PointUpdate(int(data.UserID), userData)

	transaction, err = s.repository.Insert(data)
	if err != nil {
		return
	}
	return

}

func (s service) GetByQuery(key string, value interface{}) (transaction entities.Transaction, err error) {
	//TODO implement me
	transaction, err = s.repository.FindByQuery(key, value)
	if err != nil {
		return
	}

	return transaction, nil
}

func (s service) Modify(id int, data entities.Transaction) (transaction entities.Transaction, err error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetByAny(value interface{}) (res interface{}, err error) {
	//TODO implement me
	res, err = s.repository.FindByAny(value)
	if err != nil {
		return nil, err
	}
	return res, nil
}
