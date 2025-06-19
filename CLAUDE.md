# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go client library for the Twitch Helix API. The library provides comprehensive coverage of Twitch's API endpoints with a clean, tested interface.

## Development Commands

### Building
```bash
go build .
go build -v .  # verbose output
```

### Testing
```bash
# Run all tests with coverage
go test -v -parallel=10 -covermode=count -coverprofile=coverage.out

# Run tests for specific file
go test -v ./file_test.go

# Run with race detection
go test -race -v .
```

### Code Quality
```bash
# Static analysis
go vet .

# Format code
go fmt ./...
```

## Architecture

### File Structure
- Each Twitch API endpoint has its own file (e.g., `channels.go`, `users.go`, `streams.go`)
- Test files follow the pattern `*_test.go` for each implementation file
- Sub-features are in separate files (e.g., `channels_editors.go`, `channels_vips.go`)
- Documentation examples are in the `docs/` directory

### API Pattern
Each API implementation follows this structure:
1. Request parameter structs with `query` tags
2. Response structs embedding `ResponseCommon`
3. Main function calling `c.get()`, `c.post()`, `c.put()`, or `c.delete()`
4. Response hydration pattern for common fields

Example:
```go
type SomeParams struct {
    BroadcasterID string `query:"broadcaster_id"`
    After         string `query:"after"`
}

type SomeResponse struct {
    ResponseCommon
    Data SomeData
}

func (c *Client) GetSomething(params *SomeParams) (*SomeResponse, error) {
    resp, err := c.get("/endpoint", &SomeData{}, params)
    if err != nil {
        return nil, err
    }
    
    response := &SomeResponse{}
    resp.HydrateResponseCommon(&response.ResponseCommon)
    response.Data = *resp.Data.(*SomeData)
    
    return response, nil
}
```

### Testing Pattern
Tests use table-driven approach with mock HTTP client:
```go
func TestGetSomething(t *testing.T) {
    t.Parallel()
    
    testCases := []struct {
        statusCode int
        options    *Options
        params     *SomeParams
        respBody   string
    }{
        // test cases
    }
    
    for _, testCase := range testCases {
        c := newMockClient(testCase.options, newMockHandler(
            testCase.statusCode,
            testCase.respBody,
            nil,
        ))
        
        resp, err := c.GetSomething(testCase.params)
        // assertions
    }
}
```

## Key Dependencies
- Go 1.24 or higher
- `github.com/golang-jwt/jwt/v4` - JWT token handling for Extensions

## CI/CD
GitHub Actions runs on push to `main`:
1. Tests with coverage
2. Static analysis with `go vet`
3. Coverage upload to Codecov

## Current Development
Working on Hype Train v2 feature implementation (branch: `feat/jackmc/hype-train-v2`)