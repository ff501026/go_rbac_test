package rbac

type Policy struct {
	//策略信息
	PType string `json:"ptype,omitempty" binding:"required" validate:"required"`
	//角色
	Role string `json:"role,omitempty" binding:"required" validate:"required"`
	//路由
	Object string `json:"object,omitempty" binding:"required" validate:"required"`
	//動作(GET、POST、PATCH、DELETE)
	Action string `json:"action,omitempty" binding:"required" validate:"required"`

}
