package voucher

import (
	"dapoint-api/entities"
	dapoint_api "dapoint-api/error"
	"gorm.io/gorm"
)

type MysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) entities.VoucherRepository {
	return &MysqlRepository{
		db: db,
	}
}

func (repo MysqlRepository) FindById(id uint64) (user entities.User, err error) {
	//TODO implement me
	if err = repo.db.Find(&user, id).Error; err != nil {
		return
	}

	return user, nil
}

func (repo MysqlRepository) FindAll() (users []entities.User, err error) {
	//TODO implement me
	if err = repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo MysqlRepository) FindByQuery(key string, value interface{}) (user entities.User, err error) {
	//TODO implement me

	err = repo.db.Where(key+" = ?", value).Find(&user).Error
	if err != nil {
		err = dapoint_api.ErrNotFound
		return
	}

	return user, nil
}

func (repo MysqlRepository) Insert(data entities.User) (id uint64, err error) {
	//TODO implement me

	err = repo.db.Create(&data).Error
	if err != nil {
		err = dapoint_api.ErrInternalServer
		return
	}
	return
}

func (repo MysqlRepository) Update(id int, data entities.User) (res entities.User, err error) {
	//TODO implement me
	var user entities.User
	repo.db.First(&user, "id = ?", id)

	if err = repo.db.Model(&user).Updates(map[string]interface{}{"name": data.Name, "email": data.Email, "password": data.Password, "photo": data.Photo}).Error; err != nil {
		return user, err
	}
	return user, err
}
