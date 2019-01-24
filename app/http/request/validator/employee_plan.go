package validator

type EmployeePlanList struct {
	Keyword string `form:"keyword"`
}

type StaffReward struct {
	UserId 			string  `form:"user_id" json:"user_id" binding:"required"`
	RewardPlanId	uint64	`form:"reward_plan_id" json:"reward_plan_id" binding:"required"`
	Number          float64 `form:"number" json:"number" binding:"required"`
	FirstTime		string 	`form:"first_time" json:"first_time" binding:"required"`
	IsResignationContinue uint16 `form:"is_resignation_continue" json:"is_resignation_continue"`
}

type StaffRewardList struct {
	Keyword 		string `form:"keyword" json:"keyword"`
	UserId 			string `form:"user_id" json:"user_id"`
	RewardPlanId	uint64 `form:"reward_plan_id" json:"reward_plan_id"`
	BeginTime		string `form:"begin_time" json:"begin_time"`
	EndTime			string `form:"end_time" json:"end_time"`
}
