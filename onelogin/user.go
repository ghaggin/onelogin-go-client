package onelogin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID                   int                    `json:"id,omitempty"`
	UserName             string                 `json:"username,omitempty"`
	Email                string                 `json:"email,omitempty"`
	FirstName            string                 `json:"firstname,omitempty"`
	LastName             string                 `json:"lastname,omitempty"`
	Password             string                 `json:"password,omitempty"`
	PasswordConfirmation string                 `json:"password_confirmation,omitempty"`
	PasswordAlgorithm    string                 `json:"password_algorithm,omitempty"`
	Salt                 string                 `json:"salt,omitempty"`
	Title                string                 `json:"title,omitempty"`
	Department           string                 `json:"department,omitempty"`
	Company              string                 `json:"company,omitempty"`
	Comment              string                 `json:"comment,omitempty"`
	GroupID              int                    `json:"group_id,omitempty"`
	RoleIDs              []int                  `json:"role_ids,omitempty"`
	Phone                string                 `json:"phone,omitempty"`
	State                int                    `json:"state,omitempty"`
	Status               int                    `json:"status,omitempty"`
	DirectoryID          int                    `json:"directory_id,omitempty"`
	TrustedIDPID         int                    `json:"trusted_idp_id,omitempty"`
	ManagerADID          int                    `json:"manager_ad_id,omitempty"`
	Samaccountname       string                 `json:"samaccountname,omitempty"`
	MemberOf             string                 `json:"member_of,omitempty"`
	UserPrincipalName    string                 `json:"userprincipalname,omitempty"`
	DistinguishedName    string                 `json:"distinguished_name,omitempty"`
	ExternalID           string                 `json:"external_id,omitempty"`
	OpenidName           string                 `json:"openid_name,omitempty"`
	InvalidLoginAttempts int                    `json:"invalid_login_attempts,omitempty"`
	CustomAttributes     map[string]interface{} `json:"custom_attributes,omitempty"`
}

type UserQuery struct {
	CreatedSince     time.Time              `json:"created_since,omitempty"`
	CreatedUntil     time.Time              `json:"created_until,omitempty"`
	UpdatedSince     time.Time              `json:"updated_since,omitempty"`
	UpdatedUntil     time.Time              `json:"updated_until,omitempty"`
	LastLoginSince   time.Time              `json:"last_login_since,omitempty"`
	LastLoginUntil   time.Time              `json:"last_login_until,omitempty"`
	FirstName        string                 `json:"firstname,omitempty"`
	LastName         string                 `json:"lastname,omitempty"`
	Email            string                 `json:"email,omitempty"`
	Username         string                 `json:"username,omitempty"`
	Samaccountname   string                 `json:"samaccountname,omitempty"`
	DirectoryID      string                 `json:"directory_id,omitempty"`
	ExternalID       string                 `json:"external_id,omitempty"`
	AppID            string                 `json:"app_id,omitempty"`
	UserIDs          []int                  `json:"user_ids,omitempty"`
	CustomAttributes map[string]interface{} `json:"custom_attributes,omitempty"`
	Fields           []string               `json:"fields,omitempty"`
}

// https://developers.onelogin.com/api-docs/2/users/list-users
func (c *Client) ListUsers(query *UserQuery) ([]*User, error) {
	var users []*User
	err := c.execRequest(&oneloginRequest{
		method:      GET,
		path:        "/api/2/users",
		respModel:   &users,
		queryParams: userQueryToParams(query),
	})
	return users, err
}

// https://developers.onelogin.com/api-docs/2/users/get-user
func (c *Client) GetUser(id int) (*User, error) {
	var user User
	err := c.execRequest(&oneloginRequest{
		method:    GET,
		path:      fmt.Sprintf("/api/2/users/%v", id),
		respModel: &user,
	})
	return &user, err
}

// https://developers.onelogin.com/api-docs/2/users/create-user
func (c *Client) CreateUser(user *User) (*User, error) {
	if user.UserName == "" {
		return nil, ErrMissingField{"username"}
	}
	if user.Email == "" {
		return nil, ErrMissingField{"email"}
	}

	body, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	var newUser User
	err = c.execRequest(&oneloginRequest{
		method: POST,
		path:   "/api/2/users",
		body:   bytes.NewReader(body),
		queryParams: map[string]string{
			"mappings":        "async", // default
			"validate_policy": "true",  // default
		},
		respModel: &newUser,
	})

	return &newUser, err
}

// https://developers.onelogin.com/api-docs/2/users/update-user
func (c *Client) UpdateUser(user *User) (*User, error) {
	if user.ID == 0 {
		return nil, ErrMissingField{"id"}
	}

	body, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	var updatedUser User
	err = c.execRequest(&oneloginRequest{
		method:    PUT,
		path:      fmt.Sprintf("/api/2/users/%v", user.ID),
		body:      bytes.NewReader(body),
		respModel: &updatedUser,
		queryParams: map[string]string{
			"mappings":        "async", // default
			"validate_policy": "true",  // default
		},
	})

	return &updatedUser, err
}

// https://developers.onelogin.com/api-docs/2/users/delete-user
func (c *Client) DeleteUser(id int) error {
	return c.execRequest(&oneloginRequest{
		method: DELETE,
		path:   fmt.Sprintf("/api/2/users/%v", id),
	})
}

// https://developers.onelogin.com/api-docs/2/users/get-user-apps
func (c *Client) GetUserApps(id int) ([]int, error) {
	return nil, ErrNotImplemented{}
}

func userQueryToParams(query *UserQuery) map[string]string {
	params := map[string]string{}

	if !query.CreatedSince.IsZero() {
		params["created_since"] = query.CreatedSince.UTC().Format(time.RFC3339)
	}
	if !query.CreatedUntil.IsZero() {
		params["created_until"] = query.CreatedUntil.UTC().Format(time.RFC3339)
	}
	if !query.UpdatedSince.IsZero() {
		params["updated_since"] = query.UpdatedSince.UTC().Format(time.RFC3339)
	}
	if !query.UpdatedUntil.IsZero() {
		params["updated_until"] = query.UpdatedUntil.UTC().Format(time.RFC3339)
	}
	if !query.LastLoginSince.IsZero() {
		params["last_login_since"] = query.LastLoginSince.UTC().Format(time.RFC3339)
	}
	if !query.LastLoginUntil.IsZero() {
		params["last_login_until"] = query.LastLoginUntil.UTC().Format(time.RFC3339)
	}
	if query.FirstName != "" {
		params["firstname"] = query.FirstName
	}
	if query.LastName != "" {
		params["lastname"] = query.LastName
	}
	if query.Email != "" {
		params["email"] = query.Email
	}
	if query.Username != "" {
		params["username"] = query.Username
	}
	if query.Samaccountname != "" {
		params["samaccountname"] = query.Samaccountname
	}
	if query.DirectoryID != "" {
		params["directory_id"] = query.DirectoryID
	}
	if query.ExternalID != "" {
		params["external_id"] = query.ExternalID
	}
	if query.AppID != "" {
		params["app_id"] = query.AppID
	}
	if len(query.UserIDs) > 0 {
		params["user_ids"] = intSliceToString(query.UserIDs, ",")
	}
	if len(query.CustomAttributes) > 0 {
		for key, value := range query.CustomAttributes {
			params[key] = value.(string)
		}
	}
	if len(query.Fields) > 0 {
		params["fields"] = strings.Join(query.Fields, ",")
	}

	return params
}
