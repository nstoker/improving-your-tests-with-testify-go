# README

Following a tutorial by Elliot Forbes [improving-your-tests-with-testify-go](https://tutorialedge.net/golang/improving-your-tests-with-testify-go/).

## Assertions

Created a table driven test using `assert.Equal()`.

## Mocking

There is an error in the code examples for main/main_test. In the test it is setting up the service

```go
func TestChargeCustomer(t *testing.T) {
    smsService := new(smsServiceMock)

    // we then define what should be returned from SendChargeNotification
    // when we pass in the value 100 to it. In this case, we want to return
    // true as it was successful in sending a notification
    smsService.On("SendChargeNotification", 100).Return(true)

    // next we want to define the service we wish to test
    myService := MyService{smsService}
    // ...
```

fails to compile with

```go
# github.com/nstoker/improving-your-tests-with-testify-go [github.com/nstoker/improving-your-tests-with-testify-go.test]
./main_test.go:66:25: cannot use smsService (type *smsServiceMock) as type MessageService in field value:
  *smsServiceMock does not implement MessageService (wrong type for SendChargeNotification method)
    have SendChargeNotification(int) bool
    want SendChargeNotification(int) error
FAIL     github.com/nstoker/improving-your-tests-with-testify-go [build failed]
FAIL
```
