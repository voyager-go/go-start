package repository

import (
	"errors"
	"github.com/jinzhu/copier"
	"go-start/internal/model/entity"
	"go-start/internal/pkg/helper"
	"go-start/internal/pkg/mysql"
	"go-start/internal/request"
	"go-start/internal/response"
	"go-start/internal/service"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

var (
	_ service.UserInfoService = (*userRepository)(nil)
)

type userRepository struct {
	db *gorm.DB
}

func newUserInfoRepository() service.UserInfoService {
	return &userRepository{db: mysql.Conn}
}

func (r *userRepository) Show(req request.UserInfoShowReq) (res *response.UserInfoShowRes, err error) {
	if req.Id == "" && req.Passport == "" {
		err = errors.New("查询条件缺失")
		return
	}
	query := r.db.Model(&entity.UserInfo{})
	if req.Id != "" {
		query.First(&res, req.Id)
		return
	}
	if req.Passport != "" {
		query.Where("passport = ?", req.Passport).First(&res)
	}
	return
}

func (r *userRepository) List(req request.UserInfoListReq) *[]entity.UserInfo {
	return nil
}

func (r *userRepository) Create(req request.UserInfoCreateReq) error {
	var (
		user entity.UserInfo
		res  *response.UserInfoShowRes
		err  error
	)
	res, err = r.Show(request.UserInfoShowReq{Passport: req.Passport})
	if err != nil {
		return err
	}
	if res.Id != "" {
		return errors.New("passport 已存在")
	}
	if err := copier.Copy(&user, req); err != nil {
		return err
	}
	by, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(by)
	user.CreatedAt = helper.FTime{Time: time.Now()}
	return r.db.Create(&user).Error
}
