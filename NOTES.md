# Course notes

- Types of tests: unit, integration, functional
- Unit tests should test one single unit of code (usually single functions), mocking all the external dependencies
- Integration tests should test the integration between multiple layers of the application
- Functional tests should test input/output of an entire system or high level functional requiremenet
- tests files should be in the same directory as the files they'r testing
- unit tests should be part of the same package as the package they'r testing. This also allows access to private, consts etc from the implementation
- unit tests - any test case should start with `Test`FunctionName
- unit tests - any test must have a `*testing.T` parameter
- unit test steps:
    ```golang
    // init
    // execution
    // validation
    ```
    much like `arrange, act, assert`
- `go test -cover` will run the tests and print code coverage
- code coverage can be missleading. you can have 100% code coverage without testing possible scenarios. code coverage only tells you which lines are executed when you run your tests 
- prefer white box testing over black box testing (e.g write your tests inside the same package you are testing)
- both integration and unit tests use the same interface `t *testing.T`. You can only tell the difference my looking at the test implementation
- go has native benchmark support. we can write a benchmark by creating a function starting with `Benchmark` receiving the `*testing.B` has parameter. benchmarks are written inside `_test` files
- `go test -bench=.` to run benchmarks cases



funny quote from the author:
`deploying to production without tests is like drinking and driving`