package dto

type GetListUserResponse struct {
	NPage    int    `json:"n_page"`
	PageList []Page `json:"page_list"`
}
type Page struct {
	NUser         int                 `json:"n_user"`
	Recomnendlist []ResponseRecommend `json:"user_list"`
}

type ResponseRecommend struct {
	ID       uint
	Username string
	Name     string
	Gender   string
	Location string
}
