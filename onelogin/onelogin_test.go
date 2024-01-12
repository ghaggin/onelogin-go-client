package onelogin

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type OneLoginTestSuite struct {
	suite.Suite

	// client setup prior to the suite running
	client *Client
}

func TestOneLogin(t *testing.T) {
	// Perform setup that is common to all tests
	// in the package here
	oneloginTestSuite := &OneLoginTestSuite{}

	// Retrieve instance variables from environment
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	subdomain := os.Getenv("SUBDOMAIN")

	client, err := NewClient(ClientConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Subdomain:    subdomain,
		Timeout:      60 * time.Second, // app delete is incredibly slow
	})

	// Client is required for most tests
	if err != nil {
		t.Fatalf("Error creating client: %s", err.Error())
	}
	oneloginTestSuite.client = client

	suite.Run(t, oneloginTestSuite)
}

// SetupTest runs before each test
// Perform any setup required before each test here
func (suite *OneLoginTestSuite) SetupTest() {
}

// TearDownTest runs after each test
// Perform any teardown steps required after each test here
func (suite *OneLoginTestSuite) TearDownTest() {
}
