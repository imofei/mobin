package validator

import (
	"encoding/json"
	"errors"
)

type RewardPlanCreateUpdate struct {
	Name               string `form:"name" json:"name" binding:"required"`
	Currency           string `form:"currency" json:"currency" binding:"required"`
	//Number             uint64 `form:"number" json:"number" binding:"required"`
	GrantNumber        uint64 `form:"grant_number" json:"grant_number" binding:"required"`
	GrantIntervalType  uint16 `form:"grant_interval_type" json:"grant_interval_type" binding:"required"`
	GrantIntervalValue uint64 `form:"grant_interval_value" json:"grant_interval_value" binding:"required"`
	GrantRules         []uint `form:"grant_rules" json:"grant_rules" binding:"required"`
}

func (f RewardPlanCreateUpdate) GrantRuleString() string {
	b, err := json.Marshal(f.GrantRules)
	if err != nil {
		return ""
	}
	return string(b)
}

func (f RewardPlanCreateUpdate) ValidateRules() error {
	if len(f.GrantRules) != int(f.GrantNumber) {
		return errors.New("发放次数不匹配")
	}

	var total uint = 0
	for _, per := range f.GrantRules {
		total += per
	}
	if total != 100 {
		return errors.New("发放比例不正确")
	}
	return nil
}

type RewardPlanList struct {
	Ids     []uint64
	Keyword string `form:"keyword"`
	Status  uint16 `form:"status"`
}

type RewardPlanDetail struct {
	Id string `form:"id"`
}
