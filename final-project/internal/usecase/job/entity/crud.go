package entity

type JobReq struct {
	ID          int64  `json:"id,omitempty" swaggerignore:"true"`
	UserID      int64  `json:"user_id,omitempty" swaggerignore:"true"`
	Name        string `json:"name,omitempty" validate:"required" name:"Judul"`
	Description string `json:"description" validate:"required" name:"Deskripsi"`
}

type JobResponse struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (r *JobReq) SetID(ID int64) {
	r.ID = ID
}

func (r *JobReq) SetUserID(UserID int64) {
	r.UserID = UserID
}
