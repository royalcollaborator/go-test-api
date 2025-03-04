# Project Structure

```
github.com/royalcollaborator/go-test-api
│
├── main.go                         // Main handler for retrieval of the Last Traded Price of Bitcoin function
├── api
│   └── controller                  // API Controller for API response codes and messages
├── bootstrap
│   └── app.go                      // App Instance handler
│   └── env.go                      // Env Instance handler
├── domain
│   └── mock                        // Generated Uber Mock Package
│   └── error_response.go           // Error Message Struct
│   └── trade_service.go            // Redefination of Interface function
│   └── trade.go                    // TradeUsecase interface
└── route                           // API Router Set
```

# How to manage dependency
Adding dependencies in a Go project is typically done using the Go Modules system. Here’s the best way to manage and add dependencies:

### 1. Initialize a Go Module

If you haven't already, you need to initialize a Go module in your project directory. This creates a `go.mod` file which will track your dependencies.

```bash
go mod init <module-name>
```

Replace `<module-name>` with your module's name, often structured like a URL (e.g., `github.com/username/project`).

### 2. Add a Dependency

To add a dependency, you can use the `go get` command followed by the package path. For example:

```bash
go get github.com/some/package
e.g.
go get github.com/gin-gonic/gin
```

This command will:

- Download the specified package and its dependencies.
- Update your `go.mod` file to include the new dependency.
- Update your `go.sum` file with checksums for the new dependencies.

### 3. Specify Version Constraints (Optional)

If you want to specify a particular version or version range, you can do so with `go get`. For example:

```bash
go get github.com/some/package@v1.2.3
```

You can also use other versioning formats, such as:

- `latest`: Get the latest version.
- `v1.2.x`: Get the latest patch version of 1.2.
- `v1.2.3`: Get a specific version.
- `v1.2.3-beta`: Get a pre-release version.

### 4. Verify and Tidy Up

After adding dependencies, it’s good practice to run:

```bash
go mod tidy
```

This command will ensure that your `go.mod` and `go.sum` files are clean, removing any unused dependencies and adding any missing ones.

### 5. Import the Dependency in Your Code

After adding the dependency, import it in your Go files like so:

```go
import "github.com/some/package"
```

### Summary

- **Use `go mod init`** to create a module.
- **Use `go get`** to add dependencies.
- **Specify versions** as needed.
- **Run `go mod tidy`** to clean up your module files.
- **Import the package** in your code.

# How to generate mock
```
go generate ./...
```

# How to run unit test
```
go test ./...
```

# How to get coverage check
```
go test $(go list ./... | grep -v -E 'mocks|interfaces') -coverprofile=coverage.out
go tool cover -func=coverage.out
```

# How to build project
```
go build -o main main.go
```

# How to run project
```
go run .
```
