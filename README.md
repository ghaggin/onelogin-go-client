# OneLogin Go Client

This client implements a go client for the OneLogin API as documented [here](https://developers.onelogin.com/api-docs/2/getting-started/dev-overview).

## Developing

### Run Tests
OneLogin instance variables are required to run the test.  These tests are integration tests and will modify the state of the OneLogin instance provided, so care should be taken in deciding which environment to run these tests against.  Ideally the tests will clean up any modificiations that they make, but this is not guaranteed and errors in logic could result in changes being made to the remote state that persist after the tests completes

Prior to running the tests, export these variables to the shell that the test will be run in
```
export CLIENT_ID=<client-id>
export CLIENT_SECRET=<client-secret>
export SUBDOMAIN=<instance-subdomain>
```

Run the test suite and observe the output
```
go test -v
```

If you are using VSCode with the golang plugin, add the following settings to `.vscode/settings.json`.  This will allow you to run individual tests from the UI.
```
{
    ...
    "go.testEnvVars": {
        "CLIENT_ID": "<client-id>",
        "CLIENT_SECRET": "<client-secret>",
        "SUBDOMAIN": "<instance-subdomain>",
    }
}
```

### TODO
- ListUsers should take pointer
