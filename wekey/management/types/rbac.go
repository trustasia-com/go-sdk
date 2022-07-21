package types

import "time"

// QueryRoleRequest query
type QueryRoleRequest struct {
	Name string `form:"name"` // 选填 角色名 最长64个字符
}

// QueryRoleResp resp
type QueryRoleResp struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`        // 名字
	Description string    `json:"description"` // 描述
	CreatedAt   time.Time `json:"created_at"`  // 创建时间
}

// RoleInfoResp resp
type RoleInfoResp struct {
	ID          string `json:"id"`          // ID
	Name        string `json:"name"`        // 名字
	Description string `json:"description"` // 描述
}

// CreateRoleRequest create role
type CreateRoleRequest struct {
	Name        string `json:"name"`        // 必填 名字 最长64个字符
	Description string `json:"description"` // 必填 描述 最长512个字符
}

// UpdateRoleRequest edit role
type UpdateRoleRequest struct {
	Name        string `json:"name"`        // 必填 名字 最长64个字符
	Description string `json:"description"` // 必填 描述 最长512个字符
}

// RolePermissionResp  resp
type RolePermissionResp struct {
	PermissionID          string `json:"permission_id"`          // 权限id
	PermissionScope       string `json:"permission_scope"`       // scope
	PermissionDescription string `json:"permission_description"` // 描述
	ResourceID            string `json:"resource_id"`            // 资源组id
	ResourceName          string `json:"resource_name"`          // 名字
	ResourceAudience      string `json:"resource_audience"`      // '资源组地址,当前空间唯一'
}

// RoleUserResponseItem role user item
type RoleUserResponseItem struct {
	ID       string `json:"id"`       // 记录ID，用于分页
	UserID   string `json:"user_id"`  // 用户id
	Picture  string `json:"picture"`  // 头像
	Username string `json:"username"` // username
	Email    string `json:"email"`    // 邮箱
	Nickname string `json:"nickname"` // 昵称
}

// RoleUsersListResp 角色用户列表响应
type RoleUsersListResp struct {
	Total int                    `json:"total"` // 总数据
	List  []RoleUserResponseItem `json:"list"`  // 数据列表
}

// AddRolePermissionRequest add role permission
type AddRolePermissionRequest struct {
	PermissionIDs []string `json:"permission_ids"` // 必填 权限ids
}

// RoleAssignUsersReq 分配角色
type RoleAssignUsersReq struct {
	UserIDs []string `json:"user_ids"` // 必填 用户ids
}

// QueryPermissionsWithoutRoleRequest request
type QueryPermissionsWithoutRoleRequest struct {
	ResourceID string `form:"resource_id"` // 必填 资源组id
	Scope      string `form:"scope"`       // 选填 权限
}

// QueryPermissionsWithoutRoleResp item
type QueryPermissionsWithoutRoleResp struct {
	ID          string `json:"id"`          // 权限id
	Scope       string `json:"scope"`       // 权限
	Description string `json:"description"` // 描述
}

// QueryUserWithoutRoleRequest query
type QueryUserWithoutRoleRequest struct {
	Search string `form:"search"` // 查询条件
}

// QueryUserWithoutRoleResp  用户列表
type QueryUserWithoutRoleResp struct {
	ID          string `json:"id"`           // 用户id
	Username    string `json:"username"`     // 用户名
	Nickname    string `json:"nickname"`     // 昵称
	PhoneNumber string `json:"phone_number"` // 电话号码
	Email       string `json:"email"`        // 邮箱
	Picture     string `json:"picture"`      // 头像
}
