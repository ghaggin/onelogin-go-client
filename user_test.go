package onelogin

import "time"

func (s *OneLoginTestSuite) Test_ListUsers_success() {
	users, err := s.client.ListUsers(UserQuery{
		CreatedSince: time.Now().Add(24 * time.Hour),
	})
	s.Nil(err)
	s.Equal(0, len(users))
}

// This function will test the successful operation of the user crud operations
// and will cleanup any state changes
func (s *OneLoginTestSuite) Test_UserOperations() {
	testUsername := "test_username"
	testEmail := "test_username@example.com"

	// Cleanup user if exists
	users, err := s.client.ListUsers(UserQuery{
		Username: testUsername,
	})
	s.Require().Nil(err)

	if len(users) == 1 {
		s.T().Logf("Cleaning up test user: %s", testUsername)
		err = s.client.DeleteUser(users[0].ID)
		s.Require().Nil(err)
	}

	// Create user
	user := &User{
		UserName: testUsername,
		Email:    testEmail,
	}
	newUser, err := s.client.CreateUser(user)
	s.Require().Nil(err)
	s.Equal(user.UserName, newUser.UserName)
	s.Equal(user.Email, newUser.Email)
	s.NotEqual(0, newUser.ID)

	// Update user
	updateUser := &User{
		ID:    newUser.ID,
		Email: "test_username@test.com",
	}
	updatedUser, err := s.client.UpdateUser(updateUser)
	s.Require().Nil(err)
	s.Equal(updateUser.Email, updatedUser.Email)
	s.Equal(user.UserName, updatedUser.UserName)
	s.Equal(newUser.ID, updatedUser.ID)

	// Get user
	gotUser, err := s.client.GetUser(newUser.ID)
	s.Require().Nil(err)
	s.Equal(updateUser.Email, gotUser.Email)
	s.Equal(user.UserName, gotUser.UserName)
	s.Equal(newUser.ID, gotUser.ID)

	// List users
	listUsers, err := s.client.ListUsers(UserQuery{
		Username: testUsername,
	})
	s.Require().Nil(err)
	s.Equal(1, len(listUsers))
	s.Equal(updateUser.Email, listUsers[0].Email)
	s.Equal(user.UserName, listUsers[0].UserName)
	s.Equal(newUser.ID, listUsers[0].ID)

	// Delete user
	err = s.client.DeleteUser(newUser.ID)
	s.Nil(err)
}

func (s *OneLoginTestSuite) Test_CreateUser_missing_fields() {
	_, err := s.client.CreateUser(&User{})
	s.Require().NotNil(err)
	s.Equal(err, ErrMissingField{"username"})

	_, err = s.client.CreateUser(&User{UserName: "test"})
	s.Require().NotNil(err)
	s.Equal(err, ErrMissingField{"email"})
}

func (s *OneLoginTestSuite) Test_UpdateUser_missing_fields() {
	_, err := s.client.UpdateUser(&User{})
	s.Require().NotNil(err)
	s.Equal(err, ErrMissingField{"id"})
}
