package validator

// 待发放激励名单
type RewardMeetList struct {
	Ids     []uint64
	ListId    uint64 `form:"list_id" json:"list_id"`
	Name      string `form:"name" json:"name"`
	PlanName  string `form:"plan_name" json:"plan_name"`
	StartDate string `form:"start_date" json:"start_date"`
	EndDate   string `form:"end_date" json:"end_date"`
}

// 生成发放列表
type MakeRewardList struct {
	Ids         []uint64 `form:"ids" json:"ids" binding:"required"`
	Name        string `form:"name" json:"name" binding:"required"`
	Remark      string `form:"remark" json:"remark"`
	PlanDate    string `form:"plan_date" json:"plan_date"`
}

// 添加到发放列表
type AddSingleRewardList struct {
	Id          uint64 `form:"id" json:"id" binding:"required"`
	ListId      uint64 `form:"list_id" json:"list_id" binding:"required"`
	Remark      string `form:"remark" json:"remark"`
}

// 激励发放列表
type RewardLists struct {
	Ids     []uint64
	Name      string `form:"name" json:"name"`
	Status    string `form:"status" json:"status"`
	StartDate string `form:"start_date" json:"start_date"`
	EndDate   string `form:"end_date" json:"end_date"`
}

// 标记发放
type RewardListsMarkGrant struct {
	Id      uint64 `form:"id" json:"id" binding:"required"`
}

// 激励发放名单
type RewardGrantList struct {
	Id          uint64 `form:"id" json:"id" binding:"required"`
}