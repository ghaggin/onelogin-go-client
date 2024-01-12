package onelogin

func (s *OneLoginTestSuite) Test_NewClient_default_timeout() {
	// Test default timeout value set
	client, err := NewClient(ClientConfig{})
	s.Require().NotNil(err) // will receive error due to bad auth request with default config
	s.Equal(client.httpClient.Timeout, DefaultTimeout)

}

func (s *OneLoginTestSuite) Test_NewClient_success() {
	// Test accomplished in onelogin_test setup for use in other test routines
}
