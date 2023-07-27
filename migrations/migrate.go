package migrations

import (
	"base-go/common/logger"
	"base-go/domain/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	logger.Info("Running migrations for Cat model...")
	db.AutoMigrate(&model.Blog{}, &model.User{}, &model.Event{}, &model.Notification{}, &model.NotificationType{}, &model.Invoice{}, &model.Map{}, &model.Voucher{}, &model.Report{})

	db.AutoMigrate(&model.Blog{}, &model.User{}, &model.Event{}, &model.Notification{}, &model.NotificationType{}, &model.Invoice{}, &model.Map{}, &model.Voucher{}, &model.Report{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Blog{}, &model.User{}, &model.Event{}, &model.Notification{}, &model.NotificationType{}, &model.Invoice{}, &model.Map{}, &model.Voucher{}, &model.Report{})
	defaultUser := model.User{
		UserName: "admin",
		FullName: "admin",
		Password: "$2a$14$L3q6xZ3aHDklI2duNTjA1OmWKYULW5eArNStdZJMHhXw6CX.4r5SG",
	}
	existingUser := model.User{}
	err := db.Where("user_name = ?", defaultUser.UserName).First(&existingUser).Error
	if err == nil {
		return
	}
	if existingUser.UserName == "" {
		db.Create(&defaultUser)
	}
}
