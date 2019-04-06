# Mocks

Mocks are structs that implement an interface required for testing, record the values they are called with and return known values.

Examine the following code:

```Go
type Handler struct {
    database *postgres.Conn
}

func (h *Handler) Admin(ctx context.Context, req *pb.Request) (*pb.Response, error) {
    user, err := h.database.GetUser(req.GetUserID())
    if err != nil {
        return nil, err
    }
    ...
}
```

Imagine we want to write a test for this where it returns an error. How would that even be done? Or maybe we would like to handle a specific error differently. It would be extremely difficult to test this.

If we modify the struct to use the Database interface from the previous section.

```Go
type Database interface {
    GetUser(id int) (User, error)
}

type Handler struct {
    database Database
}
```

We can then create a Handler in a test and give it a mock database that returns an error. A mock could be as simple as:

```Go
type MockDatabase struct {
    GetUserCallReceivesID    int
    GetUserCallReturnsUser   *User
    GetUserCallReturnsErr    error
}

func (m *MockDatabase) GetUser(id int) (*User, error) {
    m.GetUserCallReceivesID = id
    return m.GetUserCallReturnsUser, m.GetUserCallReturnsErr
}
```

This new type MockDatabase has a method GetUser with a definition that satisfies the Database interface. When we create it in a test, we can assign exactly what values it will return. Also after it is used, we will be able to check the GetUserCallReceivesID field and see what arguments the GetUser method was called with.

**Note: Do not write mocks like this. We will cover the correct way to write/generate them later.**

Using this mock, we can create a test case for the Handler where the database returns an error.

```Go
func TestHandler_Admin_DatabaseReturnsError(t *testing.T) {
    // Create the mock.
    mock := &MockDatabase{
        // Assign an error to be returned by the mock when GetUser is called.
        GetUserCallReturnsErr = errors.New("error!")
    }
    // We now have an implementation of the Database interface that will return an error when 
    // GetUser is called.

    // Create the handler with the mock.
    handler := Handler{
        database: mock
    }

    expectedUserID := 1
    req := &pb.Request{UserID: expectedUserID}

    // Call the Admin method being tested on the handler.
    response, err := handler.Admin(context.Background(), req)

    // Check that the database was called with the expected user id.
    if mock.GetUserCallReceives != expectedUserID {
        t.Errorf("Database was called with incorrect user ID. Expected: %d, Got: %d.", 
                 expectedUserID, 
                 mock.GetUserCallReceives,
                 )
    }

    // We also expect that because the database returned an error, the method should return an error.
    if err == nil {
        t.Error("Error expected but not returned.")
    }
}
```

Using the interface in this handler allows us to write a test case where we are only testing the Handler with no other dependencies. This is crucial in Unit Testing. Unit tests should only test code in one package. Each test or set of tests should generally only test one object. Because of Go's type system, if interfaces are used between objects (essentially as contracts), and the code compiles, and each individual piece is tested well, it should all work when put together. An often used argument against this is, "The test cases seem too simple. I feel like I'm not actually testing anything." That's a good thing if the test cases seem to simple. In highly decoupled code with small objects, test cases should be simple.