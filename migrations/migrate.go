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

}
