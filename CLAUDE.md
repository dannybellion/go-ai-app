# Go AI App - Development Guidelines

## Build & Run Commands
- Build: `go build`
- Run: `go run main.go`
- Test: `go test ./...` or `go test ./package/...` for specific package
- Format: `go fmt ./...`
- Lint: `go vet ./...`

## Code Style
- **Imports**: Standard library first, then third-party, then local packages
- **Naming**: PascalCase for exported, camelCase for unexported
- **Error Handling**: Check errors immediately, log appropriately
- **Formatting**: Follow Go standards (gofmt)
- **Documentation**: Add comments for exported functions
- **Function Size**: Keep functions small and focused

## Project Structure
- `/config`: Application configuration
- `/database`: Database connections and schema
- `/handlers`: HTTP request handlers
- `/models`: Data structures and operations
- `/routes`: API routes definition
- `/utils`: Helper functions