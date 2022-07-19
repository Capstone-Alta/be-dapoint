package voucher

import (
	"dapoint-api/entities"
	dapoint_api "dapoint-api/error"
	"github.com/go-playground/validator/v10"
)

type service struct {
	repository entities.VoucherRepository
	validate   *validator.Validate
}

func NewService(repository entities.VoucherRepository) entities.VoucherService {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s service) GetById(id uint64) (user entities.Voucher, err error) {
	//TODO implement me
	user, err = s.repository.FindById(id)
	if err != nil {
		return user, err
	}
	return
}

func (s service) GetAll() (total int, users []entities.Voucher, err error) {
	//TODO implement me
	total, users, err = s.repository.FindAll()
	if err != nil {
		return total, nil, err
	}
	return
}

func (s service) Create(data entities.VoucherDTO) (id uint64, err error) {
	//TODO implement me
	err = s.validate.Struct(&data)
	if err != nil {
		err = dapoint_api.ErrBadRequest
		return
	}
	//var ok bool
	//if ok, err = s.validate(&data); !ok {
	//	return res, err
	//}
	//newUser := entities.ObjUser(data.Name, data.Email, data.Password)
	id, err = s.repository.Insert(data)
	if err != nil {
		return
	}
	return
}

func (s service) Modify(id int, data entities.Voucher) (user entities.Voucher, err error) {
	//TODO implement me
	res, err := s.repository.Update(id, data)
	if err != nil {
		return
	}

	return res, nil
}

func (s service) GetByType(value interface{}) (vouchers []entities.Voucher, err error) {
	//TODO implement me
	res, err := s.repository.FindByType(value)
	if err != nil {
		return
	}

	return res, nil
}

func (s service) GetTotal() (res interface{}, err error) {
	//TODO implement me
	res, err = s.repository.GetTotal()
	if err != nil {
		return
	}

	return res, nil
}

func (s service) DeleteVoucher(i int) (err error) {
	//TODO implement me
	err = s.repository.DeleteVoucher(i)
	if err != nil {
		return err
	}
	return
}
