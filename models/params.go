package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	Username   string `json:"username" binding:"required,min=3,max=64"`
	Password   string `json:"password" binding:"required,min=8,max=72"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required,min=3,max=64"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}

type ParamVote struct {
	PostID    string `json:"post_id" binding:"required,numeric"`
	Direction int8   `json:"direction" binding:"oneof=-1 0 1"`
}

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id" binding:"omitempty,gt=0"`               // 可以为空
	Page        int64  `json:"page" form:"page" binding:"omitempty,min=1,max=1000000" example:"1"`      // 页码
	Size        int64  `json:"size" form:"size" binding:"omitempty,min=1,max=100" example:"10"`         // 每页数据量
	Order       string `json:"order" form:"order" binding:"omitempty,oneof=time score" example:"score"` // 排序依据
}
