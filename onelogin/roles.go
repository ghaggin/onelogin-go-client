package onelogin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type Role struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Admins []int  `json:"admins,omitempty"`
	Apps   []int  `json:"apps,omitempty"`
	Users  []int  `json:"users,omitempty"`
}

func (c *Client) ListRoles() ([]*Role, error) {
	return nil, errors.New("not implemented")
}

func (c *Client) CreateRole(role *Role) (*Role, error) {
	body, err := json.Marshal(role)
	if err != nil {
		return nil, err
	}

	var newRole Role
	err = c.exec(POST, "/api/2/roles", bytes.NewReader(body), &newRole)
	if err != nil {
		return nil, err
	}

	role.ID = newRole.ID
	return role, nil
}

func (c *Client) GetRole(id int) (*Role, error) {
	var role Role
	err := c.exec(GET, fmt.Sprintf("/api/2/roles/%v", id), nil, &role)
	return &role, err
}

func (c *Client) UpdateRole(role *Role) (*Role, error) {
	// get the current state
	currentRole, err := c.GetRole(role.ID)
	if err != nil {
		return nil, err
	}

	// update name
	if role.Name != currentRole.Name {
		body, err := json.Marshal(map[string]string{
			"name": role.Name,
		})
		if err != nil {
			return nil, err
		}
		err = c.exec(PUT, fmt.Sprintf("/api/2/roles/%v", role.ID), bytes.NewReader(body), nil)
		if err != nil {
			return nil, err
		}
	}

	// update apps
	if !sliceEqual(role.Apps, currentRole.Apps) {
		err = c.setRoleApps(role.ID, role.Apps)
		if err != nil {
			return nil, err
		}
	}

	// update users
	add, remove := sliceDiff(role.Users, currentRole.Users)
	if len(add) > 0 {
		err = c.addRoleUsers(role.ID, add)
		if err != nil {
			return nil, err
		}
	}
	if len(remove) > 0 {
		err = c.removeRoleUsers(role.ID, remove)
		if err != nil {
			return nil, err
		}
	}

	// update admins
	add, remove = sliceDiff(role.Admins, currentRole.Admins)
	if len(add) > 0 {
		err = c.addRoleAdmins(role.ID, add)
		if err != nil {
			return nil, err
		}
	}
	if len(remove) > 0 {
		err = c.removeRoleAdmins(role.ID, remove)
		if err != nil {
			return nil, err
		}
	}

	return role, nil
}

func (c *Client) DeleteRole(id int) error {
	return c.exec(DELETE, fmt.Sprintf("/api/2/roles/%v", id), nil, nil)
}

func (c *Client) setRoleApps(id int, apps []int) error {
	body, err := json.Marshal(apps)
	if err != nil {
		return err
	}
	return c.exec(PUT, fmt.Sprintf("/api/2/roles/%v/apps", id), bytes.NewReader(body), nil)
}

func (c *Client) addRoleUsers(id int, users []int) error {
	return c.modifyRoleUsers(POST, id, users)
}

func (c *Client) removeRoleUsers(id int, users []int) error {
	return c.modifyRoleUsers(DELETE, id, users)
}

func (c *Client) modifyRoleUsers(op method, id int, users []int) error {
	body, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return c.exec(op, fmt.Sprintf("/api/2/roles/%v/users", id), bytes.NewReader(body), nil)
}

func (c *Client) addRoleAdmins(id int, users []int) error {
	return c.modifyRoleAdmins(POST, id, users)
}

func (c *Client) removeRoleAdmins(id int, users []int) error {
	return c.modifyRoleAdmins(DELETE, id, users)
}

func (c *Client) modifyRoleAdmins(op method, id int, users []int) error {
	body, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return c.exec(op, fmt.Sprintf("/api/2/roles/%v/admins", id), bytes.NewReader(body), nil)
}
