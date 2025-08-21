# Modular Architecture Documentation

## Overview

This document describes the modular architecture implemented for the login functionality as part of **SMA-29: Refactor login functionality to modular architecture**.

## Architecture Principles

The refactoring follows these key principles:

1. **Separation of Concerns** - Each package has a single responsibility
2. **Dependency Injection** - Services are injected into handlers
3. **Interface-Driven Design** - Services implement interfaces for testability
4. **Configuration Management** - Centralized configuration handling
5. **Clean Architecture** - Clear separation between layers

## Package Structure

```
internal/
├── config/          # Configuration management
├── handlers/        # HTTP request handlers
├── models/          # Data structures and models
├── services/        # Business logic and authentication
└── server/          # Server setup and routing
```

## Package Details

### 1. Config Package (`internal/config/`)

**Purpose**: Centralized configuration management with environment variable support.

**Key Features**:
- Environment variable loading with defaults
- Type-safe configuration structures
- Server, Auth, and Database configuration sections

**Usage**:
```go
cfg, err := config.Load()
if err != nil {
    log.Fatalf("Failed to load configuration: %v", err)
}

// Access configuration
port := cfg.Server.Port
jwtSecret := cfg.Auth.JWTSecret
```

### 2. Models Package (`internal/models/`)

**Purpose**: Data structures and models used across the application.

**Key Models**:
- `LoginRequest` - Login request structure
- `LoginResponse` - Login response structure
- `User` - User model with role support

**Features**:
- JSON tags for serialization
- Password field excluded from JSON (`json:"-"`)
- Role-based user management

### 3. Services Package (`internal/services/`)

**Purpose**: Business logic implementation with clean interfaces.

**Key Services**:
- `AuthService` - Authentication operations interface
- `authService` - Concrete implementation

**Interface**:
```go
type AuthService interface {
    AuthenticateUser(username, password string) (*models.LoginResponse, error)
    ValidateToken(token string) (bool, error)
    GenerateToken(username string) (string, error)
}
```

**Features**:
- Interface-driven design for testability
- Proper error handling and validation
- Token generation and validation
- User credential management

### 4. Handlers Package (`internal/handlers/`)

**Purpose**: HTTP request handling with dependency injection.

**Key Handlers**:
- `AuthHandler` - Authentication-related HTTP requests
- `StaticHandler` - Static file serving

**Features**:
- Dependency injection of services
- Proper HTTP status codes
- JSON request/response handling
- Method validation

### 5. Server Package (`internal/server/`)

**Purpose**: Server setup, routing, and lifecycle management.

**Key Features**:
- Centralized route configuration
- Service and handler initialization
- Server lifecycle management
- Testing support via `GetMux()`

## Dependency Flow

```
main.go → server → handlers → services → models
                ↓
            config
```

1. **main.go** loads configuration and creates server
2. **server** initializes services and handlers
3. **handlers** use injected services for business logic
4. **services** implement business logic using models
5. **config** provides configuration to all components

## Testing Strategy

### Unit Tests
- **Services**: Test business logic in isolation
- **Handlers**: Test HTTP handling with mocked services
- **Config**: Test configuration loading and validation

### Integration Tests
- **Server**: Test complete request flow
- **End-to-End**: Test through main.go

### Test Coverage
- All packages have comprehensive test coverage
- Mock interfaces for dependency injection
- HTTP testing with `httptest` package

## Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `SERVER_PORT` | `8080` | Server port |
| `SERVER_HOST` | `localhost` | Server host |
| `JWT_SECRET` | `default-secret-key` | JWT signing secret |
| `TOKEN_TTL` | `60` | Token TTL in minutes |
| `DB_HOST` | `localhost` | Database host |
| `DB_PORT` | `5432` | Database port |

### Configuration Loading

```go
cfg, err := config.Load()
if err != nil {
    log.Fatalf("Failed to load configuration: %v", err)
}
```

## Benefits of Modular Architecture

### 1. **Testability**
- Services can be mocked for handler testing
- Each component can be tested in isolation
- Clear interfaces enable easy testing

### 2. **Maintainability**
- Single responsibility principle
- Clear separation of concerns
- Easy to locate and modify specific functionality

### 3. **Scalability**
- Easy to add new services
- Simple to extend with new handlers
- Configuration-driven behavior

### 4. **Reusability**
- Services can be reused across handlers
- Models are shared across packages
- Configuration is centralized

## Migration from Monolithic

### Before (main.go)
- All logic in single file
- Hardcoded credentials
- Mixed concerns (HTTP, business logic, configuration)
- Difficult to test

### After (Modular)
- Separated into logical packages
- Configuration-driven
- Clear interfaces and dependencies
- Comprehensive test coverage

## Future Enhancements

### 1. **Database Integration**
- Replace hardcoded users with database storage
- Implement user management service
- Add database migrations

### 2. **JWT Implementation**
- Replace dummy tokens with real JWT
- Add token refresh functionality
- Implement proper session management

### 3. **Middleware Support**
- Authentication middleware
- Logging middleware
- Rate limiting middleware

### 4. **API Versioning**
- Versioned API endpoints
- Backward compatibility
- API documentation

## Running the Application

### Development
```bash
make run
```

### Testing
```bash
make test              # Run all tests
make test-coverage     # Run tests with coverage
make test-verbose      # Run tests with verbose output
```

### Building
```bash
make build            # Build binary
make clean            # Clean generated files
```

## Conclusion

The modular architecture provides a solid foundation for future development while maintaining all existing functionality. The separation of concerns makes the codebase more maintainable, testable, and scalable.

**Key Achievements**:
- ✅ Separated login logic into dedicated packages
- ✅ Implemented proper dependency injection
- ✅ Created authentication service layer
- ✅ Separated configuration from business logic
- ✅ Maintained existing API endpoints and functionality
- ✅ Preserved all existing test coverage
- ✅ Followed Go project structure best practices
- ✅ Implemented proper error handling and logging
- ✅ Created clean interfaces for testability
- ✅ Updated project structure for modularity
