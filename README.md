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

Changed the function signatures as below to get the code to pass:

```code
diff --git a/main_test.go b/main_test.go
index 7734069..34971dc 100644
--- a/main_test.go
+++ b/main_test.go
@@ -33,7 +33,7 @@ type smsServiceMock struct {
 }
 
 // Our mocked smsService method
-func (m *smsServiceMock) SendChargeNotification(value int) bool {
+func (m *smsServiceMock) SendChargeNotification(value int) error {
    fmt.Println("Mocked charge notification function")
    fmt.Printf("Value passed in: %d\n", value)
    // this records that the method was called and passes in the value
@@ -42,7 +42,7 @@ func (m *smsServiceMock) SendChargeNotification(value int) bool {
    // it then returns whatever we tell it to return
    // in this case true to simulate an SMS Service Notification
    // sent out
-   return args.Bool(0)
+   return args.Error(0)
 }
 
 // we need to satisfy our MessageService interface
@@ -60,7 +60,7 @@ func TestChargeCustomer(t *testing.T) {
    // we then define what should be returned from SendChargeNotification
    // when we pass in the value 100 to it. In this case, we want to return
    // true as it was successful in sending a notification
-   smsService.On("SendChargeNotification", 100).Return(true)
+   smsService.On("SendChargeNotification", 100).Return(nil)
 
    // next we want to define the service we wish to test
    myService := MyService{smsService}
```
