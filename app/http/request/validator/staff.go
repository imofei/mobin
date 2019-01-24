package validator

type UpdateStaffRole struct {
	RoleId  uint64 `form:"role_id" json:"role_id" binding:"required"`
}

type StaffList struct {
	Ids     []uint64
	Keyword string `form:"keyword"`
	Status  string `form:"status"`
}

type StaffDetail struct {
	Id string `form:"id"`
}

type UpdateStaff struct {
	ResignationDate  string `form:"resignation_date" json:"resignation_date"`
	Status           uint16 `form:"status" json:"status"`
}