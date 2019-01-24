package validator

type RoleList struct {
	Ids     []uint64
	Keyword string `form:"keyword"`
}

type SystemDetail struct {
	Id string `form:"id"`
}

type RoleCreate struct {
	Name             string `form:"name" json:"name" binding:"required"`
	MenuIds         []uint64 `form:"menu_ids" json:"menu_ids" binding:"required"`

}

