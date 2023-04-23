package entities

import (
	"github.com/shopspring/decimal"
	"time"
)

type User struct {
	Id                string          `json:"id"`
	UserCode          string          `json:"userCode"`
	Username          string          `json:"username"`
	RealName          string          `json:"realName"`
	Phone             string          `json:"phone"`
	Email             string          `json:"email"`
	Gender            int             `json:"gender"`
	Birthday          *time.Time      `json:"birthday"`
	HeadImage         string          `json:"headImage"`
	Balance           decimal.Decimal `json:"balance"`
	CommissionBalance decimal.Decimal `json:"commissionBalance"`
	CommissionFreeze  decimal.Decimal `json:"commissionFreeze"`
	Vip               int             `json:"vip"`
	Activate          int             `json:"activate"`
	Locked            int             `json:"locked"`
	OnlineStatus      int             `json:"onlineStatus"`
	BindingLoginId    string          `json:"bindingLoginId"`
	WhatsappLink      string          `json:"whatsappLink"`
	Password          string          `json:"password" gorm:"<-:false"`
	PayPassword       string          `json:"payPassword" gorm:"<-:false"`
	CreateTime        *time.Time      `json:"createTime"`
	UpdateTime        *time.Time      `json:"updateTime"`
	LastConsumeTime   *time.Time      `json:"lastConsumeTime"`
	LastRechargeTime  *time.Time      `json:"lastRechargeTime"`
}

func (*User) TableName() string {
	return "lianfei_auth.lf_user"
}
