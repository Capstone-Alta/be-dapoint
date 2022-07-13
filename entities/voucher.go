package entities

import (
	"gorm.io/gorm"
)

type Voucher struct {
	gorm.Model
	ID              uint64 `gorm:"primaryKey"`
	Name            string `json:"name" form:"name"`
	Stock           uint   `json:"stock" form:"stock"`
	HargaPoint      uint   `json:"harga_point" form:"harga_point"`
	VoucherDetailID uint   `json:"voucher_detail_id" form:"voucher_detail_id"`
	UserVouchers    []UserVoucher
}

type VoucherDTO struct {
	Name        string `json:"name" form:"name"`
	Stock       uint   `json:"stock" form:"stock"`
	HargaPoint  uint   `json:"harga_point" form:"harga_point"`
	TipeVoucher string `json:"tipe_voucher" form:"tipe_voucher"`
}

func ObjVoucher(dataName string, dataStock uint, dataHargaPoint uint) (voucher *Voucher) {
	return &Voucher{
		Name:       dataName,
		Stock:      dataStock,
		HargaPoint: dataHargaPoint,
	}
}

type VoucherRepository interface {
	FindById(id uint64) (voucher Voucher, err error)
	FindAll() (vouchers []Voucher, err error)
	FindByQuery(key string, value interface{}) (voucher Voucher, err error)
	Insert(data VoucherDTO) (id uint64, err error)
	Update(id int, data Voucher) (voucher Voucher, err error)
	FindByParam(value interface{}) (vouchers []Voucher, err error)
}

type VoucherService interface {
	GetById(id uint64) (voucher Voucher, err error)
	GetAll() (vouchers []Voucher, err error)
	Create(data VoucherDTO) (id uint64, err error)
	Modify(id int, data Voucher) (voucher Voucher, err error)
	GetByParam(value interface{}) (vouchers []Voucher, err error)
}
