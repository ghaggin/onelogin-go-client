package onelogin

func (s *OneLoginTestSuite) Test_RoleOperations() {
	apps, err := s.client.ListApps(&AppQuery{
		Paging: Paging{
			Limit: 2,
			Page:  1,
		},
	})
	s.Require().NoError(err)
	s.Require().Equal(2, len(apps))

	users, err := s.client.ListUsers(&UserQuery{
		Paging: Paging{
			Limit: 2,
			Page:  1,
		},
	})
	s.Require().NoError(err)
	s.Require().Equal(2, len(users))

	// create a role
	role, err := s.client.CreateRole(&Role{
		Name:   "test-role",
		Apps:   []int{apps[0].ID},
		Users:  []int{users[0].ID},
		Admins: []int{users[0].ID},
	})
	s.Require().NoError(err)
	s.Require().NotNil(role)
	s.NotZero(role.ID)
	s.Equal("test-role", role.Name)
	s.Require().Equal(1, len(role.Apps))
	s.Equal([]int{apps[0].ID}, role.Apps)
	s.Require().Equal(1, len(role.Users))
	s.Equal([]int{users[0].ID}, role.Users)
	s.Require().Equal(1, len(role.Admins))
	s.Equal([]int{users[0].ID}, role.Admins)

	// get the role
	role, err = s.client.GetRole(role.ID)
	s.Require().NoError(err)
	s.Require().NotNil(role)
	s.NotZero(role.ID)
	s.Equal("test-role", role.Name)

	// update the role
	roleUpdatedName := "test-role-updated"
	role, err = s.client.UpdateRole(&Role{
		ID:     role.ID,
		Name:   roleUpdatedName,
		Apps:   []int{apps[1].ID},
		Users:  []int{users[1].ID},
		Admins: []int{users[1].ID},
	})
	s.Require().NoError(err)
	s.Require().NotNil(role)
	s.NotZero(role.ID)
	s.Equal(roleUpdatedName, role.Name)
	s.Require().Equal(1, len(role.Apps))
	s.Equal([]int{apps[1].ID}, role.Apps)
	s.Require().Equal(1, len(role.Users))
	s.Equal([]int{users[1].ID}, role.Users)
	s.Require().Equal(1, len(role.Admins))
	s.Equal([]int{users[1].ID}, role.Admins)

	// delete the role
	err = s.client.DeleteRole(role.ID)
	s.NoError(err)
}

func (s *OneLoginTestSuite) Test_ListRoles() {
	roles, err := s.client.ListRoles(&RoleQuery{
		Paging: Paging{
			Limit: 2,
			Page:  1,
		},
	})
	s.Require().NoError(err)
	s.Equal(2, len(roles))
}
