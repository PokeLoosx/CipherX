package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	InviteUserID      int    `gorm:"column:invite_user_id;default:null;comment:Inviter ID" json:"invite_user_id"`
	TelegramID        int64  `gorm:"column:telegram_id;default:null;comment:Telegram ID" json:"telegram_id"`
	Email             string `gorm:"column:email;type:varchar(64);unique;not null;comment:Email address" json:"email"`
	Password          string `gorm:"column:password;type:varchar(64);not null;comment:Password" json:"password"`
	PasswordSalt      string `gorm:"column:password_salt;type:char(10);default:null;comment:Password salt" json:"password_salt"`
	Balance           int    `gorm:"column:balance;default:0;comment:Balance" json:"balance"`
	Discount          int    `gorm:"column:discount;default:null;comment:Discount" json:"discount"`
	CommissionType    int    `gorm:"column:commission_type;type:tinyint;default:null;comment:Types of commission" json:"commission_type"`
	CommissionRate    int    `gorm:"column:commission_rate;comment:Proportion of commission" json:"commission_rate"`
	CommissionBalance int    `gorm:"column:commission_balance;default:0;comment:Balance of commission" json:"commission_balance"`
	T                 int64  `gorm:"column:t;default:0;comment:Last active timestamp" json:"t"`
	U                 int64  `gorm:"column:u;type:bigint;uniqueIndex:user_u_d_expired_at_group_id_banned_transfer_enable_index;default:0;comment:Upload" json:"u"`
	D                 int64  `gorm:"column:d;type:bigint;uniqueIndex:user_u_d_expired_at_group_id_banned_transfer_enable_index;default:0;comment:Download" json:"d"`
	TransferEnable    int64  `gorm:"column:transfer_enable;type:bigint;uniqueIndex:user_u_d_expired_at_group_id_banned_transfer_enable_index;default:0;comment:Traffic limits" json:"transfer_enable"`
	Banned            int    `gorm:"column:banned;type:tinyint(1);uniqueIndex:user_u_d_expired_at_group_id_banned_transfer_enable_index;default:0;comment:Ban or not" json:"banned"`
	IsAdmin           int    `gorm:"column:is_admin;type:tinyint(1);default:0;comment:Administrator or not" json:"is_admin"`
	IsStaff           int    `gorm:"column:is_staff;type:tinyint(1);default:0;comment:Staff or not" json:"is_staff"`
	LastLoginAt       int    `gorm:"column:last_login_at;default:null;comment:Last login time" json:"last_login_at"`
	LastLoginIP       int    `gorm:"column:last_login_ip;comment:Last login IP" json:"last_login_ip"`
	UUID              string `gorm:"column:uuid;type:char(36);not null;comment:UUID" json:"uuid"`
	GroupID           int    `gorm:"column:group_id;uniqueIndex:user_u_d_expired_at_group_id_banned_transfer_enable_index;default:null;comment:User group ID" json:"group_id"`
	PlanID            int    `gorm:"column:plan_id;default:null;comment:Plan ID" json:"plan_id"`
	SpeedLimit        int    `gorm:"column:speed_limit;default:null;comment:Speed limit" json:"speed_limit"`
	RemindExpire      int    `gorm:"column:remind_expire;type:tinyint;default:0;comment:Expired reminders" json:"remind_expire"`
	RemindTraffic     int    `gorm:"column:remind_traffic;type:tinyint;default:0;comment:Traffic alerts" json:"remind_traffic"`
	Token             string `gorm:"column:token;type:char(32);not null;comment:token" json:"token"`
	ExpiredAt         int    `gorm:"column:expired_at;type:bigint;uniqueIndex:user_u_d_expired_at_group_id_banned_transfer_enable_index;default:0;comment:Expiration time" json:"expired_at"`
	Remarks           string `gorm:"column:remarks;type:text;comment:Remarks" json:"remarks"`
}

// CommissionType User Commission Type
const (
	CommissionTypeSystem  = iota + 1 // System Settings
	CommissionTypePeriod             // Loop
	CommissionTypeOnetime            // First time
)
