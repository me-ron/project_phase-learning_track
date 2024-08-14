# Unit Test Suite Documentation

## Overview
This documentation provides an overview of the unit test suite for the `UserRepo` and `TaskRepo` packages in the `task_manager` project. The test suite is designed using the `testify` package in Go, particularly the `suite` package to manage and run the tests in a structured manner. It includes mocking dependencies such as MongoDB collections and operations, enabling isolated testing of repository methods.

## Repositry
## Structure
The test suite is organized into two main suites:
- **UserRepoTestSuite**: Tests for the `UserRepo` repository methods.
- **TaskRepoTestSuite**: Tests for the `TaskRepo` repository methods.

Each suite contains the following components:
- **SetupTest**: Initializes mocks and sets up any necessary preconditions before each test.
- **Test Methods**: Each method tests a specific functionality of the repository, ensuring the correct behavior of the repository layer.

## Test Cases
The test cases cover the following methods for both `UserRepo` and `TaskRepo`:

- **Create**: Tests the creation of a user or task.
- **FindById**: Tests fetching a user or task by ID.
- **FindByEmail** (UserRepo only): Tests fetching a user by email.
- **FindAllUsers** (UserRepo) and **GetAllTasks** (TaskRepo): Tests retrieving all users or tasks.
- **UpdateUserById** (UserRepo) and **UpdateTaskById** (TaskRepo): Tests updating a user or task by ID.
- **DeleteUserByID** (UserRepo) and **DeleteTaskByID** (TaskRepo): Tests deleting a user or task by ID.

## Instructions for Running Tests

1. **Ensure Dependencies**:
   - Install the necessary Go modules if they are not already installed.
     ```bash
     go get github.com/stretchr/testify
     ```

2. **Running the Tests**:
   - Navigate to the package directory containing the test files (`repository`).
   - Run the tests using the Go testing tool.
     ```bash
     go test ./...
     ```
   - The `go test` command will automatically discover and run all tests in the current directory and its subdirectories.

3. **Test Coverage**:
   - To obtain test coverage metrics, run the following command:
     ```bash
     go test -cover ./...
     ```
   - This will provide a coverage report indicating the percentage of code covered by the tests.

4. **Verbose Output (Optional)**:
   - For more detailed output, use the `-v` flag:
     ```bash
     go test -v ./...
     ```

## Test Coverage Metrics
- **Coverage Report**: The test suite is designed to maximize coverage of the repository layer. By using mocks, the suite simulates various scenarios, ensuring that edge cases and normal operations are well-tested.
- **Improving Coverage**: If coverage is below expectations, consider adding more test cases to cover untested branches or error conditions in the code.

## Conclusion
The unit test suite is comprehensive and structured, ensuring that the core functionality of the `UserRepo` and `TaskRepo` is thoroughly tested. Following the instructions provided, you can run the tests and assess the coverage to maintain high code quality in the `task_manager` project.

## Use case
## Structure
The test suite is organized into two main suites:
- **TaskUsecaseSuite**: Tests for the `TaskUC` use case methods.
- **UserUsecaseSuite**: Tests for the `UserUC` use case methods.

Each suite contains the following components:
- **SetupTest**: Initializes mocks and sets up any necessary preconditions before each test.
- **Test Methods**: Each method tests a specific functionality of the use case, ensuring the correct behavior of the application logic.

## Test Cases

### TaskUsecaseSuite
The test cases for the `TaskUsecaseSuite` cover the following methods:

- **PostTask**: Tests the creation of a task.
- **GetTasks**: Tests retrieving all tasks based on a filter.
- **GetTask**: Tests fetching a specific task by ID and user ID.
- **UpdateTask**: Tests updating a task by ID and user.
- **DeleteTask**: Tests deleting a task by ID and user ID.

### UserUsecaseSuite
The test cases for the `UserUsecaseSuite` cover the following methods:

- **Login**: Tests user login functionality, including password comparison and token creation.
- **Signup**: Tests user registration, including password hashing and user creation.
- **GetUsers**: Tests retrieving all users.
- **GetUser**: Tests fetching a specific user by ID.
- **MakeAdmin**: Tests promoting a user to an admin role.
- **UpdateUser**: Tests updating a user's information, including password hashing.
- **DeleteUser**: Tests deleting a user by ID.

## Instructions for Running Tests

1. **Ensure Dependencies**:
   - Install the necessary Go modules if they are not already installed.
     ```bash
     go get github.com/stretchr/testify
     ```

2. **Running the Tests**:
   - Navigate to the package directory containing the test files (`useCase`).
   - Run the tests using the Go testing tool.
     ```bash
     go test ./...
     ```
   - The `go test` command will automatically discover and run all tests in the current directory and its subdirectories.

3. **Test Coverage**:
   - To obtain test coverage metrics, run the following command:
     ```bash
     go test -cover ./...
     ```
   - This will provide a coverage report indicating the percentage of code covered by the tests.

4. **Verbose Output (Optional)**:
   - For more detailed output, use the `-v` flag:
     ```bash
     go test -v ./...
     ```

## Test Coverage Metrics
- **Coverage Report**: The test suite is designed to maximize coverage of the use case layer. By using mocks, the suite simulates various scenarios, ensuring that edge cases and normal operations are well-tested.
- **Improving Coverage**: If coverage is below expectations, consider adding more test cases to cover untested branches or error conditions in the code.

## Conclusion
The use case test suite is comprehensive and structured, ensuring that the core functionality of the `TaskUC` and `UserUC` is thoroughly tested. Following the instructions provided, you can run the tests and assess the coverage to maintain high code quality in the `task_manager` project.

## Controllers
# `TaskHandlerTestSuite` Documentation

The `TaskHandlerTestSuite` is a test suite designed to validate the functionality of the task-related HTTP handlers in a Gin web framework. The suite uses the `testify/suite` package to organize and execute tests in a structured and maintainable manner.

## Structure

### Fields

- **`router`**: 
  - Type: `*gin.Engine`
  - Description: The Gin engine instance used for routing during tests.
  
- **`Task_UC`**: 
  - Type: `*mocks.TaskUsecase`
  - Description: A mock implementation of the `TaskUsecase` interface, used to simulate the behavior of the actual use case layer.

### Methods

- **`SetupTest()`**: 
  - Description: Initializes the test environment by setting the Gin mode to `TestMode` and creating new instances of the router and the mocked use case.
  
- **`TestGetAllTasks()`**: 
  - Description: Tests the `GetAllTasks` handler, ensuring it retrieves all tasks and returns them in the correct format.
  
- **`TestGetTaskById()`**: 
  - Description: Tests the `GetTaskById` handler, verifying that it retrieves a specific task by its ID and returns the expected result.
  
- **`TestPostTask()`**: 
  - Description: Tests the `PostTask` handler, confirming that it successfully creates a new task and returns the appropriate response.
  
- **`TestDeleteTask()`**: 
  - Description: Tests the `DeleteTask` handler, ensuring it deletes a specific task and returns a success message.
  
- **`TestUpdateTask()`**: 
  - Description: Tests the `UpdateTask` handler, verifying that it updates an existing task and returns the updated task data.

## Usage

This test suite is used to validate the behavior of task-related endpoints in a Gin-based web application. It mocks the `TaskUsecase` to simulate different scenarios and uses the Gin test context to handle HTTP requests and responses.

---

# `UserHandlerTestSuite` Documentation

The `UserHandlerTestSuite` is a test suite created to validate the functionality of user-related HTTP handlers in a Gin web framework. Like the `TaskHandlerTestSuite`, it employs the `testify/suite` package to manage and execute tests systematically.

## Structure

### Fields

- **`router`**: 
  - Type: `*gin.Engine`
  - Description: The Gin engine instance used for routing during tests.
  
- **`User_UC`**: 
  - Type: `*mocks.UserUsecase`
  - Description: A mock implementation of the `UserUsecase` interface, used to simulate the behavior of the actual user-related use case layer.

### Methods

- **`SetupTest()`**: 
  - Description: Initializes the test environment by setting the Gin mode to `TestMode` and creating new instances of the router and the mocked use case.
  
- **`TestRegister()`**: 
  - Description: Tests the `Register` handler, ensuring it registers a new user and returns the appropriate success message.
  
- **`TestLogIn()`**: 
  - Description: Tests the `Login` handler, verifying that it logs in a user and returns a token and user data.
  
- **`TestGetAllTasks()`**: 
  - Description: Tests the `GetAllUsers` handler, ensuring it retrieves all users and returns them in the correct format.
  
- **`TestGetUserById()`**: 
  - Description: Tests the `GetUserById` handler, verifying that it retrieves a specific user by their ID and returns the expected result.
  
- **`TestMakeAdmin()`**: 
  - Description: Tests the `MakeAdmin` handler, ensuring it grants admin privileges to a user and returns the updated user data.
  
- **`TestDeleteUser()`**: 
  - Description: Tests the `DeleteUser` handler, ensuring it deletes a specific user and returns a success message.
  
- **`TestUpdateUser()`**: 
  - Description: Tests the `UpdateUser` handler, verifying that it updates an existing user's information and returns the updated data.

## Usage

This test suite is used to validate the behavior of user-related endpoints in a Gin-based web application. It mocks the `UserUsecase` to simulate various scenarios and uses the Gin test context to handle HTTP requests and responses.
