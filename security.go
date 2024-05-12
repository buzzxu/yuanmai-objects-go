package main

import (
	"context"
	"fmt"
)

type _userkey int

const (
	userKey _userkey = iota
)

// SYS is a special operator that represents the system.
var OPER_SYS = Operator{
	ID:       0,
	UserName: "sys",
	Type:     99,
	RealName: "系统",
	Email:    "downloadxu@163.com",
	Mobile:   "15333819906",
	NickName: "sys",
}

var ANONYMOUS = UserInfo{}

type (
	OAuthUser struct {
		Id[uint]   `mapstructure:",squash"`
		UserId     int32     `json:"userId"`
		Type       OAuthType `json:"oAuthType"`
		OAuthId    string    `json:"OAuthId"`
		Unionid    string    `json:"unionid"`
		Credential string    `json:"credential"`
	}
	User struct {
		Id[uint]   `mapstructure:",squash"`
		Type       int32  `json:"type"`
		OrgId      int32  `json:"orgId"`
		RoleId     int32  `json:"roleId"`
		RealName   string `json:"realName"`
		Avatar     string `json:"avatar"`
		UserName   string `json:"userName"`
		NickName   string `json:"nickName"`
		Mobile     string `json:"mobile"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		Salt       string `json:"salt"`
		Status     int32  `json:"status"`
		FirstLogin bool   `json:"firstLogin" gorm:"default:false"`
		Merge      bool   `json:"merge" gorm:"default:false"`
		Use2FA     bool   `json:"use2FA" gorm:"default:false"`
		Secret2FA  bool   `json:"-"`
		Verified   bool   `json:"verified"`
		Gender     int32  `json:"gender"`
		Source     int32  `json:"source"`
		Lanaguage  string `json:"lanaguage"`
		Anonymous  bool   `json:"anonymous"`
	}

	UserInfo struct {
		*User     `mapstructure:",squash"` // 匿名字段
		RoleCodes []string                 `json:"roleCodes"`
		OAuths    *[]OAuthUser             `json:"oauths"`
	}

	Operator struct {
		ID       uint   `json:"id"`
		OrgId    int32  `json:"orgId"`
		RoleId   int32  `json:"roleId"`
		Type     int32  `json:"type"`
		UserName string `json:"userName"`
		RealName string `json:"realName"`
		NickName string `json:"nickName"`
		Mobile   string `json:"mobile"`
		Email    string `json:"email"`
		Avatar   string `json:"avatar"`
	}

	Menu struct {
		Id[uint] `mapstructure:",squash"`
		ParentId int32                  `json:"parentId"`
		Region   string                 `json:"region"`
		Name     string                 `json:"name"`
		Code     string                 `json:"code"`
		Target   string                 `json:"target"`
		Path     string                 `json:"path"`
		Icon     string                 `json:"icon"`
		Enable   bool                   `json:"enable"`
		Depth    int32                  `json:"depth"`
		Reamrk   string                 `json:"reamrk"`
		Ext      map[string]interface{} `json:"ext"`
		Sort     int32                  `json:"sort"`
		Childs   []Menu                 `json:"childs"`
	}

	Permisson struct {
		Id[uint]
		Type        int32  `json:"type"`
		Name        string `json:"name"`
		Code        string `json:"code"`
		Target      string `json:"target"`
		Description string `json:"description"`
		Path        string `json:"path"`
	}
	Role struct {
		Id[uint]
		ParentId int32                  `json:"parentId"`
		Region   int32                  `json:"region"`
		Name     string                 `json:"name"`
		Code     string                 `json:"code"`
		PermIds  []int32                `json:"permIds"`
		Perms    []Permisson            `json:"perms"`
		Ext      map[string]interface{} `json:"ext"`
	}

	PrivilegeInfo struct {
		Id[uint]
		UserName    string   `json:"userName"`
		Mobile      string   `json:"mobile"`
		Email       string   `json:"email"`
		NickName    string   `json:"nickName"`
		RealName    string   `json:"realName"`
		Avatar      string   `json:"avatar"`
		Type        int32    `json:"type"`
		Gender      int32    `json:"gender"`
		Source      int32    `json:"source"`
		RoleIds     []int32  `json:"roleIds"`
		PermIds     []int32  `json:"permIds"`
		Roles       []string `json:"roles"`
		Permissions []string `json:"permissions"`
		OrgId       int32    `json:"orgId"`
	}

	Department struct {
		Id[uint]
		ParentId int32        `json:"parentId"`
		Parent   *Department  `json:"parent"`
		Name     string       `json:"name"`
		Code     string       `json:"code"`
		Icon     string       `json:"icon"`
		PathIds  []string     `json:"pathIds"`
		Paths    []string     `json:"paths"`
		Level    int32        `json:"level"`
		Manager  bool         `json:"manager"`
		Childs   []Department `json:"children"`
	}
	Organization struct {
		Id[uint]
		ParentId int32          `json:"parentId"`
		Name     string         `json:"name"`
		Code     string         `json:"code"`
		Icon     string         `json:"icon"`
		PathIds  []string       `json:"pathIds"`
		Paths    []string       `json:"paths"`
		Path     string         `json:"path"`
		PathId   string         `json:"pathId"`
		Level    int32          `json:"level"`
		Areas    []Region       `json:"areas"`
		Childs   []Organization `json:"children"`
		ChildIds []int32        `json:"childIds"`
	}
)

func WithUser[USER User](ctx context.Context, user USER) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func GetUser[USER User](ctx context.Context) (USER, bool) {
	user, ok := ctx.Value(userKey).(USER)
	return user, ok
}

// IsAnonymous returns true if the user info is anonymous.
func (u *UserInfo) IsAnonymous() bool {
	return u == &ANONYMOUS
}

// Anonymous returns the ANONYMOUS user info.
func Anonymous() *UserInfo {
	return &ANONYMOUS
}

// Of returns the operator for the user info.
func (u *UserInfo) Of() *Operator {
	return &Operator{
		ID:       u.ID,
		Type:     u.Type,
		OrgId:    u.OrgId,
		RoleId:   u.RoleId,
		UserName: u.UserName,
		NickName: u.NickName,
		RealName: u.RealName,
		Email:    u.Email,
		Mobile:   u.Mobile,
		Avatar:   u.Avatar,
	}
}

// Name returns the name of the operator.
func (o *Operator) Name() string {
	if o.RealName != "" {
		return o.RealName
	} else if o.UserName != "" {
		return o.UserName
	} else if o.Mobile != "" {
		return o.Mobile
	} else if o.NickName != "" {
		return o.NickName
	}
	return "未知"
}

// String returns the string representation of the operator.
func (o *Operator) String() string {
	return fmt.Sprintf("Operator{id=%d, userName=%s, realName=%s, nickName=%s, mobile=%s, email=%s, avatar=%s, type=%d}", o.ID, o.UserName, o.RealName, o.NickName, o.Mobile, o.Email, o.Avatar, o.Type)
}

// Fill fills the privilege info with the given user information, roles, and permissions.
func (p *PrivilegeInfo) Fill(user *User, roles []string, permissions []string) *PrivilegeInfo {
	p.ID = user.ID
	p.UserName = user.UserName
	p.Mobile = user.Mobile
	p.Email = user.Email
	p.RealName = user.RealName
	p.Avatar = user.Avatar
	p.Gender = user.Gender
	p.Type = user.Type
	p.Source = user.Source
	p.Roles = roles
	p.Permissions = permissions
	return p
}
