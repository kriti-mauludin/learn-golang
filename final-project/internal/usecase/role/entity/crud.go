package entity

import "time"

type RoleReq struct {
	ID          int64  `json:"id,omitempty" swaggerignore:"true"`
	UserID      int64  `json:"user_id,omitempty" swaggerignore:"true"`
	Role        string `json:"role,omitempty" validate:"required" name:"Judul"`
	Description string `json:"description" validate:"required" name:"Deskripsi"`
	DoingAt     string `json:"doing_at" validate:"required" name:"Tanggal Aktifitas"`
}

type RoleResponse struct {
	ID          int64  `json:"id,omitempty"`
	Role        string `json:"role"`
	Description string `json:"description"`
	DoingAt     string `json:"doing_at"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type SendEmailReq struct {
	UserID    int64     `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (r *RoleReq) SetID(ID int64) {
	r.ID = ID
}

func (r *RoleReq) SetUserID(UserID int64) {
	r.UserID = UserID
}
