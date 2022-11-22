package org

import (
	"errors"
	"strings"
	"time"

	"github.com/grafana/grafana/pkg/models/roletype"
	"github.com/grafana/grafana/pkg/services/user"
)

// Typed errors
var (
	ErrOrgNotFound  = errors.New("organization not found")
	ErrOrgNameTaken = errors.New("organization name is taken")
)

type Org struct {
	ID      int64 `xorm:"pk autoincr 'id'"`
	Version int
	Name    string

	Address1 string
	Address2 string
	City     string
	ZipCode  string
	State    string
	Country  string

	Created time.Time
	Updated time.Time
}
type OrgUser struct {
	ID      int64 `xorm:"pk autoincr 'id'"`
	OrgID   int64 `xorm:"org_id"`
	UserID  int64 `xorm:"user_id"`
	Role    RoleType
	Created time.Time
	Updated time.Time
}

type RoleType = roletype.RoleType

const (
	RoleViewer RoleType = "Viewer"
	RoleEditor RoleType = "Editor"
	RoleAdmin  RoleType = "Admin"
)

type CreateOrgCommand struct {
	Name string `json:"name" binding:"Required"`

	// initial admin user for account
	UserID int64 `json:"-" xorm:"user_id"`
}

type GetOrgIDForNewUserCommand struct {
	Email        string
	Login        string
	OrgID        int64
	OrgName      string
	SkipOrgSetup bool
}

type GetUserOrgListQuery struct {
	UserID int64 `xorm:"user_id"`
}

type UserOrgDTO struct {
	OrgID int64    `json:"orgId" xorm:"org_id"`
	Name  string   `json:"name"`
	Role  RoleType `json:"role"`
}

type UpdateOrgCommand struct {
	Name  string
	OrgId int64
}

type SearchOrgsQuery struct {
	Query string
	Name  string
	Limit int
	Page  int
	IDs   []int64 `xorm:"ids"`
}

type OrgDTO struct {
	ID   int64  `json:"id" xorm:"id"`
	Name string `json:"name"`
}

type GetOrgByIdQuery struct {
	ID int64
}

type GetOrgByNameQuery struct {
	Name string
}

type UpdateOrgAddressCommand struct {
	OrgID int64 `xorm:"org_id"`
	Address
}

type Address struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	ZipCode  string `json:"zipCode"`
	State    string `json:"state"`
	Country  string `json:"country"`
}

type DeleteOrgCommand struct {
	ID int64 `xorm:"id"`
}

type AddOrgUserCommand struct {
	LoginOrEmail string   `json:"loginOrEmail" binding:"Required"`
	Role         RoleType `json:"role" binding:"Required"`

	OrgID  int64 `json:"-"`
	UserID int64 `json:"-"`

	// internal use: avoid adding service accounts to orgs via user routes
	AllowAddingServiceAccount bool `json:"-"`
}

type UpdateOrgUserCommand struct {
	Role RoleType `json:"role" binding:"Required"`

	OrgID  int64 `json:"-"`
	UserID int64 `json:"-"`
}

type OrgUserDTO struct {
	OrgID         int64           `json:"orgId"`
	UserID        int64           `json:"userId"`
	Email         string          `json:"email"`
	Name          string          `json:"name"`
	AvatarURL     string          `json:"avatarUrl"`
	Login         string          `json:"login"`
	Role          string          `json:"role"`
	LastSeenAt    time.Time       `json:"lastSeenAt"`
	Updated       time.Time       `json:"-"`
	Created       time.Time       `json:"-"`
	LastSeenAtAge string          `json:"lastSeenAtAge"`
	AccessControl map[string]bool `json:"accessControl,omitempty"`
	IsDisabled    bool            `json:"isDisabled"`
}

type RemoveOrgUserCommand struct {
	UserID                   int64
	OrgID                    int64
	ShouldDeleteOrphanedUser bool
	UserWasDeleted           bool
}

type GetOrgUsersQuery struct {
	UserID int64
	OrgID  int64
	Query  string
	Limit  int
	// Flag used to allow oss edition to query users without access control
	DontEnforceAccessControl bool

	User *user.SignedInUser
}

type SearchOrgUsersQuery struct {
	OrgID int64
	Query string
	Page  int
	Limit int

	User *user.SignedInUser
}

type SearchOrgUsersQueryResult struct {
	TotalCount int64         `json:"totalCount"`
	OrgUsers   []*OrgUserDTO `json:"OrgUsers"`
	Page       int           `json:"page"`
	PerPage    int           `json:"perPage"`
}

type ByOrgName []*UserOrgDTO

// Len returns the length of an array of organisations.
func (o ByOrgName) Len() int {
	return len(o)
}

// Swap swaps two indices of an array of organizations.
func (o ByOrgName) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

// Less returns whether element i of an array of organizations is less than element j.
func (o ByOrgName) Less(i, j int) bool {
	if strings.ToLower(o[i].Name) < strings.ToLower(o[j].Name) {
		return true
	}

	return o[i].Name < o[j].Name
}
