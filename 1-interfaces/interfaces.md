# Interfaces

To understand how to unit test in Golang, you must understand interfaces. An interface is an agreed upon means for communication. It can be thought of as a contract. Imagine a function needs a database connection to fetch a users information from their ID. We would expect the database connection to have a method GetUser, that takes a user ID and returns the associated User and any error that occurred. This could be defined as:

```Go
type Database interface {
    GetUser(id int) (User, error)
}
```

In Go interfaces are implemented implicitly. A struct does not have to state that it implements the interface as in other languages. This is powerful because it makes code flexible. Any struct that has a method GetUser that takes an integer and returns a User and an error could be used as an implementation of this interface. For example, given the function:

```Go
func Foo(database *postgres.Conn, userID int) {
    user, err := database.GetUser(userID)
}
```

The type postgres.Conn implements the Database interface. So the function could be redefined as:

```Go
func Foo(database Database, userID int) {
    user, err := database.GetUser(userID)
}
```

Why would we want to do this? It makes the function more flexible. In the future, the database could change from Postgres to MySQL, Cassandra, DynamoDB or any other database. The implementation of that database would be done to satisfy the Database interface. And no changes would be required to the logic of the code.

What if there will only ever be one implementation of the database? Is an interface still required? Yes. How would the original Foo function be tested? You would have to start by setting up a postgres database and adding some data to it. Then what happens if there is a bug in the postgres.Conn implementation? A test for Foo could fail. This means Foo is too tightly coupled with the postgress.Conn. A test for Foo should only fail if there is a problem with Foo itself. Imagine running tests on an entire repository and seeing that Foo failed, then having to trace the bug through postgress or some/several other packages to get to the root. It's much easier to identify the issue if the packages are completely decoupled and not dependent on each other. A Unit Test should only ever fail if there is an issue with the package it is testing.

If you're totally new to testing you might be wondering still how we will test this without using the postgress.Conn implementation of the database, the answer is by using a mock, which will be covered in the next section.

In addition to making code more flexible and testable, interfaces often improve the readability of code. Imagine a case where the postgres.Conn has 50+ public methods and it's assigned to a struct in a field. As a reader, you have no idea what it's being used for. If only a few methods are actually being used, they could be defined in an interface just above the struct. The reader can then see that the struct depends on an interface, quickly jump to the interface and know exactly what it is doing. It is even worth while to convert larger interfaces into smaller ones for this reason! The smallest (sized by the number of methods) possible interface should always be used.

In summary, there are three reasons to use interfaces. They decouple your code, make the code more readable and allow passing unknown types with the empty interface (not covered here, but worth looking into if you don't already know). This is all you need to know about interfaces for testing. To learn more there's an excellent talk by [Francesc Campoy](https://www.youtube.com/watch?v=F4wUrj6pmSI).