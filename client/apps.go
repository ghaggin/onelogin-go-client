package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ghaggin/onelogin-go-client/models"
)

func (c *Client) ListApps() ([]*models.App, error) {
	return nil, errors.New("not implemented")
}

func (c *Client) GetApp(id int) (*models.App, error) {
	var app models.App
	err := c.exec(GET, fmt.Sprintf("/api/2/apps/%v", id), nil, &app)
	return &app, err
}

func (c *Client) CreateApp(app *models.App) (*models.App, error) {
	body, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}

	var newApp models.App
	err = c.exec(POST, "/api/2/apps", bytes.NewReader(body), &newApp)
	if err != nil {
		return nil, err
	}

	app.ID = newApp.ID
	return app, nil
}

func (c *Client) UpdateApp(app *models.App) error {
	if app.ID == 0 {
		return errors.New("app ID is required")
	}

	// TODO: fix delete parameters when I get a response
	// from OneLogin reps
	//
	// oldApp, err := c.GetApp(app.ID)
	// if err != nil {
	// 	return err
	// }

	// for parameterKey, parameter := range oldApp.Parameters {
	// 	if _, ok := app.Parameters[parameterKey]; !ok {
	// 		fmt.Println("deleting parameter:", parameterKey)
	// 		err = c.deleteAppParameter(app.ID, parameter.ID)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	body, err := json.Marshal(app)
	if err != nil {
		return err
	}

	var newApp models.App
	err = c.exec(PUT, fmt.Sprintf("/api/2/apps/%v", app.ID), bytes.NewReader(body), &newApp)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteApp(id int) error {
	return c.exec(DELETE, fmt.Sprintf("/api/2/apps/%v", id), nil, nil)
}

func (c *Client) deleteAppParameter(appID, parameterID int) error {
	return c.exec(DELETE, fmt.Sprintf("/api/2/apps/%v/parameters/%v", appID, parameterID), nil, nil)
}

func (c *Client) listAppUsers(appID int) ([]int, error) {
	return nil, errors.New("not implemented")
}
