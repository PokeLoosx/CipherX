package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	InviteUserID      int    `gorm:"column:invite_user_id;default:null;comment:邀请人ID" json:"invite_user_id"`
	TelegramID        int64  `gorm:"column:telegram_id;default:null;comment:Telegram ID" json:"telegram_id"`
	Email             string `gorm:"column:email;type:varchar(64);unique;not null;comment:邮箱" json:"email"`
	Password          string `gorm:"column:password;type:varchar(64);not null;comment:密码" json:"password"`
	PasswordAlgo      string `gorm:"column:password_algo;type:char(10);default:null;comment:密码加密算法" json:"password_algo"`
	PasswordSalt      string `gorm:"column:password_salt;type:char(10);default:null;comment:密码加盐" json:"password_salt"`
	Balance           int    `gorm:"column:balance;default:0;comment:余额" json:"balance"`
	Discount          int    `gorm:"column:discount;default:null;comment:折扣" json:"discount"`
	CommissionType    int    `gorm:"column:commission_type;type:tinyint;default:null;comment:佣金类型" json:"commission_type"`
	CommissionRate    int    `gorm:"column:commission_rate;comment:佣金比例" json:"commission_rate"`
	CommissionBalance int    `gorm:"column:commission_balance;default:0;comment:佣金余额" json:"commission_balance"`
	T                 int64  `gorm:"column:t;default:0;comment:T值" json:"t"`
	U                 int64  `gorm:"column:u;type:bigint;uniqueIndex:v2_user_u_d_expired_at_group_id_banned_transfer_enable_index;default:0;comment:上传" json:"u"`
	D                 int64  `gorm:"column:d;type:bigint;uniqueIndex:v2_user_u_d_expired_at_group_id_banned_transfer_enable_index;default:0;comment:下载" json:"d"`
	TransferEnable    int64  `gorm:"column:transfer_enable;type:bigint;uniqueIndex:v2_user_u_d_expired_at_group_id_banned_transfer_enable_index;default:0;comment:流量限制" json:"transfer_enable"`
	Banned            int    `gorm:"column:banned;type:tinyint(1);uniqueIndex:v2_user_u_d_expired_at_group_id_banned_transfer_enable_index;default:0;comment:是否封禁" json:"banned"`
	IsAdmin           int    `gorm:"column:is_admin;type:tinyint(1);default:0;comment:是否管理员" json:"is_admin"`
	IsStaff           int    `gorm:"column:is_staff;type:tinyint(1);default:0;comment:是否员工" json:"is_staff"`
	LastLoginAt       int    `gorm:"column:last_login_at;default:null;comment:最后登录时间" json:"last_login_at"`
	LastLoginIP       int    `gorm:"column:last_login_ip;comment:最后登录IP" json:"last_login_ip"`
	UUID              string `gorm:"column:uuid;type:char(36);not null;comment:UUID" json:"uuid"`
	GroupID           int    `gorm:"column:group_id;uniqueIndex:v2_user_u_d_expired_at_group_id_banned_transfer_enable_index;default:null;comment:用户组ID" json:"group_id"`
	PlanID            int    `gorm:"column:plan_id;default:null;comment:套餐ID" json:"plan_id"`
	SpeedLimit        int    `gorm:"column:speed_limit;default:null;comment:限速" json:"speed_limit"`
	RemindExpire      int    `gorm:"column:remind_expire;type:tinyint;default:0;comment:过期提醒" json:"remind_expire"`
	RemindTraffic     int    `gorm:"column:remind_traffic;type:tinyint;default:0;comment:流量提醒" json:"remind_traffic"`
	Token             string `gorm:"column:token;type:char(32);not null;comment:token" json:"token"`
	ExpiredAt         int    `gorm:"column:expired_at;type:bigint;uniqueIndex:v2_user_u_d_expired_at_group_id_banned_transfer_enable_index;default:0;comment:过期时间" json:"expired_at"`
	Remarks           string `gorm:"column:remarks;type:text;comment:备注" json:"remarks"`
}

// CommissionType 用户佣金类型
const (
	CommissionTypeSystem = iota + 1
	CommissionTypePeriod
	CommissionTypeOnetime
)
