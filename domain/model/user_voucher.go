package model

type UserVoucher struct {
	VoucherId int  `gorm:"primaryKey" column:"voucher_id"`
	UserId    int  `gorm:"primaryKey" column:"user_id"`
	Status    bool `gorm:"column:status" json:"status"`
}
