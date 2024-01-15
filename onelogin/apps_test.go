package onelogin

import "strconv"

func (s *OneLoginTestSuite) Test_ListConnectorIDs() {
	connectors, err := s.client.ListConnectorIDs(&AppConnectorQuery{
		Paging: Paging{
			Limit: 3,
			Page:  1,
		},
	})
	s.Require().Nil(err)
	s.Equal(3, len(connectors))
}

func (s *OneLoginTestSuite) Test_ListApps() {
	apps31, err := s.client.ListApps(&AppQuery{
		Paging: Paging{
			Limit: 3,
			Page:  1,
		},
	})
	s.Require().Nil(err)
	s.Require().Equal(3, len(apps31))

	apps22, err := s.client.ListApps(&AppQuery{
		Paging: Paging{
			Limit: 2,
			Page:  2,
		},
	})
	s.Require().Nil(err)
	s.Require().Equal(2, len(apps22))

	s.Equal(apps31[2].ID, apps22[0].ID)
}

func (s *OneLoginTestSuite) Test_GetApp() {
	apps, err := s.client.ListApps(&AppQuery{
		Paging: Paging{
			Limit: 1,
			Page:  1,
		},
	})
	s.Require().Nil(err)
	s.Require().Equal(1, len(apps))

	app, err := s.client.GetApp(apps[0].ID)
	s.Require().Nil(err)
	s.Equal(apps[0].ID, app.ID)
}

func (s *OneLoginTestSuite) Test_AppOperations() {
	app, err := s.client.CreateApp(&App{
		ConnectorID: 110016,
		Name:        "test_app",
	})
	s.Require().Nil(err)
	s.Equal("test_app", app.Name)
	s.Equal(110016, app.ConnectorID)
	s.NotEqual(0, app.ID)

	app.Name = "test_app_updated"
	if app.Parameters == nil {
		app.Parameters = map[string]*Parameter{}
	}
	app.Parameters["new_test_param"] = &Parameter{
		Label: "new_test_param",
	}

	err = s.client.UpdateApp(app)
	s.Require().Nil(err)

	updatedApp, err := s.client.GetApp(app.ID)
	s.Require().Nil(err)
	s.Equal("test_app_updated", updatedApp.Name)
	s.Equal(110016, updatedApp.ConnectorID)
	s.Equal(app.ID, updatedApp.ID)
	s.Require().Contains(updatedApp.Parameters, "new_test_param")
	s.Equal("new_test_param", updatedApp.Parameters["new_test_param"].Label)

	err = s.client.DeleteApp(app.ID)
	s.Nil(err)
}

func (s *OneLoginTestSuite) Test_appQueryToParams() {
	limit := 2
	page := 1
	name := "test_app"
	connectorID := 123
	authMethod := AppAuthMethodGoogle

	query := &AppQuery{
		Paging: Paging{
			Limit: 2,
			Page:  1,
		},
		Name:        "test_app",
		ConnectorID: 123,
		AuthMethod:  AppAuthMethodGoogle,
	}

	params := appQueryToParams(query)

	s.Require().Contains(params, "name")
	s.Equal(name, params["name"])

	s.Require().Contains(params, "limit")
	limitFromParam, err := strconv.Atoi(params["limit"])
	s.Require().Nil(err)
	s.Equal(limit, limitFromParam)

	s.Require().Contains(params, "page")
	pageFromParam, err := strconv.Atoi(params["page"])
	s.Require().Nil(err)
	s.Equal(page, pageFromParam)

	s.Require().Contains(params, "connector_id")
	connectorIDFromParam, err := strconv.Atoi(params["connector_id"])
	s.Require().Nil(err)
	s.Equal(connectorID, connectorIDFromParam)

	s.Require().Contains(params, "auth_method")
	authMethodFromParam, err := strconv.Atoi(params["auth_method"])
	s.Require().Nil(err)
	s.Equal(authMethod, intToAuthMethod(authMethodFromParam))
}

func (s *OneLoginTestSuite) Test_appConnectorQueryToParams() {
	limit := 2
	page := 1
	cursor := "test_cursor"
	name := "test_app"
	authMethod := AppAuthMethodGoogle

	query := &AppConnectorQuery{
		Paging: Paging{
			Limit:  2,
			Page:   1,
			Cursor: cursor,
		},
		Name:       "test_app",
		AuthMethod: AppAuthMethodGoogle,
	}

	params := appConnectorQueryToParams(query)

	s.Require().Contains(params, "name")
	s.Equal(name, params["name"])

	s.Require().Contains(params, "limit")
	limitFromParam, err := strconv.Atoi(params["limit"])
	s.Require().Nil(err)
	s.Equal(limit, limitFromParam)

	s.Require().Contains(params, "limit")
	s.Equal(cursor, params["cursor"])

	s.Require().Contains(params, "page")
	pageFromParam, err := strconv.Atoi(params["page"])
	s.Require().Nil(err)
	s.Equal(page, pageFromParam)

	s.Require().Contains(params, "auth_method")
	authMethodFromParam, err := strconv.Atoi(params["auth_method"])
	s.Require().Nil(err)
	s.Equal(authMethod, intToAuthMethod(authMethodFromParam))
}

func (s *OneLoginTestSuite) Test_UpdateApp_missing_parameters() {
	err := s.client.UpdateApp(&App{})
	s.Require().NotNil(err)
	s.Equal(ErrMissingField{"id"}, err)
}

func (s *OneLoginTestSuite) Test_deleteAppParameters() {
	err := s.client.deleteAppParameter(1, 1)
	s.Require().NotNil(err)
	s.Equal(ErrOneloginAPIBroken{}, err)
}
