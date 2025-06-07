# Complete Gin Framework Tutorial - From Scratch to Production

## Table of Contents
1. [Introduction to Gin](#introduction-to-gin)
2. [Setup and Installation](#setup-and-installation)
3. [Basic Gin Application](#basic-gin-application)
4. [Routing Fundamentals](#routing-fundamentals)
5. [Parameters and Query Strings](#parameters-and-query-strings)
6. [JSON Binding and Validation](#json-binding-and-validation)
7. [Middleware](#middleware)
8. [Error Handling](#error-handling)
9. [Route Groups](#route-groups)
10. [File Uploads](#file-uploads)
11. [Database Integration](#database-integration)
12. [Configuration Management](#configuration-management)
13. [Authentication with JWT](#authentication-with-jwt)
14. [Deployment](#deployment)

## Introduction to Gin

### What is Gin?
Gin is a high-performance HTTP web framework written in Go. It features a martini-like API with much better performance.

### Gin vs net/http
- **net/http**: Go's standard library for HTTP servers
  - More verbose
  - Requires manual routing implementation
  - No built-in middleware support
  - Lower-level control

- **Gin**: Third-party framework built on top of net/http
  - Simpler, more expressive API
  - Built-in routing with pattern matching
  - Middleware support out of the box
  - JSON binding and validation
  - Better performance than many alternatives

### When to use Gin?
- Building REST APIs
- Need fast development with clean code
- Want middleware support
- Require JSON handling
- Building microservices
- Need good performance with simplicity

## Setup and Installation

### Prerequisites
```bash
# Ensure you have Go installed (1.16+)
go version

# Initialize a new Go module
mkdir gin-tutorial
cd gin-tutorial
go mod init gin-tutorial

# Install Gin
go get github.com/gin-gonic/gin
```

## Basic Gin Application

Let's start with the simplest possible Gin application:

```go
package main

import (
    // Import the Gin framework
    "github.com/gin-gonic/gin"
    // Import net/http for HTTP status codes
    "net/http"
)

func main() {
    // Create a Gin router with default middleware (logger and recovery)
    // gin.Default() sets up a router with Logger and Recovery middleware attached
    router := gin.Default()
    
    // Define a GET route for the root path "/"
    // c *gin.Context contains all information about the HTTP request and response
    router.GET("/", func(c *gin.Context) {
        // Send a JSON response with HTTP status 200 (OK)
        // gin.H is a shortcut for map[string]interface{}
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World!",
            "status":  "success",
        })
    })
    
    // Start the HTTP server on port 8080
    // This is a blocking call - the program will wait here for HTTP requests
    router.Run(":8080") // Default port is :8080 if not specified
}
```

**Key Concepts Explained:**
- `gin.Default()`: Creates a router with default middleware (logging and panic recovery)
- `gin.Context`: Contains request/response information and helper methods
- `gin.H`: Type alias for `map[string]interface{}` - convenient for JSON responses
- `router.Run()`: Starts the HTTP server

## Routing Fundamentals

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // Create router instance
    router := gin.Default()
    
    // GET route - retrieve data
    router.GET("/users", func(c *gin.Context) {
        // Simulate getting all users
        users := []gin.H{
            {"id": 1, "name": "John Doe", "email": "john@example.com"},
            {"id": 2, "name": "Jane Smith", "email": "jane@example.com"},
        }
        
        // Return JSON response with 200 status code
        c.JSON(http.StatusOK, gin.H{
            "data":    users,
            "message": "Users retrieved successfully",
        })
    })
    
    // POST route - create new data
    router.POST("/users", func(c *gin.Context) {
        // In a real app, you'd process the request body here
        c.JSON(http.StatusCreated, gin.H{
            "message": "User created successfully",
            "id":      123, // Simulated new user ID
        })
    })
    
    // PUT route - update existing data (full update)
    router.PUT("/users/:id", func(c *gin.Context) {
        // Extract the ID from URL path parameter
        userID := c.Param("id")
        
        c.JSON(http.StatusOK, gin.H{
            "message": "User updated successfully",
            "id":      userID,
        })
    })
    
    // PATCH route - partial update
    router.PATCH("/users/:id", func(c *gin.Context) {
        userID := c.Param("id")
        
        c.JSON(http.StatusOK, gin.H{
            "message": "User partially updated",
            "id":      userID,
        })
    })
    
    // DELETE route - remove data
    router.DELETE("/users/:id", func(c *gin.Context) {
        userID := c.Param("id")
        
        c.JSON(http.StatusOK, gin.H{
            "message": "User deleted successfully",
            "id":      userID,
        })
    })
    
    // Start server
    router.Run(":8080")
}
```

## Parameters and Query Strings

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv" // For string to integer conversion
)

func main() {
    router := gin.Default()
    
    // Path Parameters Example
    // URL: /users/123
    router.GET("/users/:id", func(c *gin.Context) {
        // Extract path parameter using c.Param()
        // :id in the route corresponds to the parameter name
        userID := c.Param("id")
        
        // Convert string to integer
        id, err := strconv.Atoi(userID)
        if err != nil {
            // If conversion fails, return error
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid user ID format",
            })
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "user_id": id,
            "message": "User found",
        })
    })
    
    // Multiple Path Parameters
    // URL: /users/123/posts/456
    router.GET("/users/:userID/posts/:postID", func(c *gin.Context) {
        userID := c.Param("userID")   // Extract first parameter
        postID := c.Param("postID")   // Extract second parameter
        
        c.JSON(http.StatusOK, gin.H{
            "user_id": userID,
            "post_id": postID,
            "message": "Post found for user",
        })
    })
    
    // Query Parameters Example
    // URL: /search?name=john&age=25&active=true
    router.GET("/search", func(c *gin.Context) {
        // Get query parameter by key
        name := c.Query("name")           // Returns string, empty if not found
        ageStr := c.Query("age")          // Get age as string
        active := c.Query("active")       // Get active status
        
        // Get query parameter with default value
        page := c.DefaultQuery("page", "1")     // Default to "1" if not provided
        limit := c.DefaultQuery("limit", "10")  // Default to "10" if not provided
        
        // Convert string parameters to appropriate types
        age, err := strconv.Atoi(ageStr)
        if err != nil {
            age = 0 // Set default if conversion fails
        }
        
        // Convert string to boolean
        isActive := active == "true"
        
        c.JSON(http.StatusOK, gin.H{
            "search_params": gin.H{
                "name":   name,
                "age":    age,
                "active": isActive,
                "page":   page,
                "limit":  limit,
            },
            "message": "Search parameters received",
        })
    })
    
    // Query Arrays Example
    // URL: /filter?tags=go&tags=web&tags=api
    router.GET("/filter", func(c *gin.Context) {
        // Get array of query parameters with same key
        tags := c.QueryArray("tags")
        
        // Alternative: Get query parameter and split by comma
        // URL: /filter?categories=tech,programming,golang
        categories := c.Query("categories")
        
        c.JSON(http.StatusOK, gin.H{
            "tags":       tags,       // ["go", "web", "api"]
            "categories": categories, // "tech,programming,golang"
            "message":    "Filters applied",
        })
    })
    
    // Combined Path and Query Parameters
    // URL: /users/123/posts?status=published&sort=date
    router.GET("/users/:id/posts", func(c *gin.Context) {
        // Path parameter
        userID := c.Param("id")
        
        // Query parameters
        status := c.DefaultQuery("status", "all")
        sortBy := c.DefaultQuery("sort", "created_at")
        
        c.JSON(http.StatusOK, gin.H{
            "user_id": userID,
            "filters": gin.H{
                "status":  status,
                "sort_by": sortBy,
            },
            "message": "Posts filtered for user",
        })
    })
    
    router.Run(":8080")
}
```

## JSON Binding and Validation

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Define a struct for user data
// Struct tags provide JSON field names and validation rules
type User struct {
    // `json:"name"` - JSON field name when marshaling/unmarshaling
    // `binding:"required"` - Field is required during binding
    Name     string `json:"name" binding:"required"`
    
    // `binding:"required,email"` - Required field that must be valid email
    Email    string `json:"email" binding:"required,email"`
    
    // `binding:"required,min=18,max=100"` - Required integer between 18-100
    Age      int    `json:"age" binding:"required,min=18,max=100"`
    
    // Optional field (no required tag)
    Phone    string `json:"phone"`
    
    // `binding:"oneof=admin user guest"` - Must be one of specified values
    Role     string `json:"role" binding:"required,oneof=admin user guest"`
}

// Struct for updating user (partial data allowed)
type UserUpdate struct {
    Name  string `json:"name"`
    Email string `json:"email" binding:"omitempty,email"` // Validate email only if provided
    Age   int    `json:"age" binding:"omitempty,min=18,max=100"`
    Phone string `json:"phone"`
}

func main() {
    router := gin.Default()
    
    // POST route with JSON binding and validation
    router.POST("/users", func(c *gin.Context) {
        var user User
        
        // ShouldBindJSON attempts to bind request body JSON to struct
        // It also validates the data according to struct tags
        if err := c.ShouldBindJSON(&user); err != nil {
            // If binding/validation fails, return error details
            c.JSON(http.StatusBadRequest, gin.H{
                "error":   "Validation failed",
                "details": err.Error(), // Contains specific validation errors
            })
            return // Exit early if validation fails
        }
        
        // If we reach here, validation passed
        // In real app, you'd save user to database here
        
        c.JSON(http.StatusCreated, gin.H{
            "message": "User created successfully",
            "user":    user, // Echo back the validated user data
        })
    })
    
    // PUT route for updating user
    router.PUT("/users/:id", func(c *gin.Context) {
        userID := c.Param("id")
        var userUpdate UserUpdate
        
        // Bind and validate the update data
        if err := c.ShouldBindJSON(&userUpdate); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error":   "Invalid update data",
                "details": err.Error(),
            })
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "message":    "User updated successfully",
            "user_id":    userID,
            "updated_data": userUpdate,
        })
    })
    
    // Example with query parameter binding
    // URL: /users?page=1&limit=10&sort=name
    type QueryParams struct {
        Page  int    `form:"page" binding:"omitempty,min=1"`    // form tag for query params
        Limit int    `form:"limit" binding:"omitempty,min=1,max=100"`
        Sort  string `form:"sort" binding:"omitempty,oneof=name email age"`
    }
    
    router.GET("/users", func(c *gin.Context) {
        var params QueryParams
        
        // ShouldBindQuery binds query parameters to struct
        if err := c.ShouldBindQuery(&params); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error":   "Invalid query parameters",
                "details": err.Error(),
            })
            return
        }
        
        // Set defaults for zero values
        if params.Page == 0 {
            params.Page = 1
        }
        if params.Limit == 0 {
            params.Limit = 10
        }
        if params.Sort == "" {
            params.Sort = "name"
        }
        
        // Simulate fetching users with pagination
        users := []gin.H{
            {"id": 1, "name": "John", "email": "john@example.com"},
            {"id": 2, "name": "Jane", "email": "jane@example.com"},
        }
        
        c.JSON(http.StatusOK, gin.H{
            "data": users,
            "pagination": gin.H{
                "page":  params.Page,
                "limit": params.Limit,
                "sort":  params.Sort,
            },
        })
    })
    
    // Example with custom validation
    router.POST("/custom-validation", func(c *gin.Context) {
        var user User
        
        // First, do the standard binding and validation
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }
        
        // Then, add custom business logic validation
        if user.Name == "admin" && user.Role != "admin" {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "User named 'admin' must have admin role",
            })
            return
        }
        
        c.JSON(http.StatusCreated, gin.H{
            "message": "User created with custom validation",
            "user":    user,
        })
    })
    
    router.Run(":8080")
}
```

**Key Validation Tags:**
- `required` - Field must be present
- `email` - Must be valid email format
- `min=n` - Minimum value for numbers or length for strings
- `max=n` - Maximum value for numbers or length for strings
- `oneof=a b c` - Must be one of the specified values
- `omitempty` - Skip validation if field is empty

## Middleware

```go
package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

// Custom middleware function
// Middleware functions take gin.Context and return gin.HandlerFunc
func CustomLogger() gin.HandlerFunc {
    // This function is called once when middleware is registered
    return func(c *gin.Context) {
        // This function is called for each request
        
        // Record start time
        startTime := time.Now()
        
        // Get request information
        method := c.Request.Method
        path := c.Request.URL.Path
        clientIP := c.ClientIP()
        
        fmt.Printf("[CUSTOM] Started %s %s from %s\n", method, path, clientIP)
        
        // Call c.Next() to proceed to the next middleware or handler
        // This is crucial - without it, the request stops here
        c.Next()
        
        // Code here runs after the request is processed
        
        // Calculate processing time
        duration := time.Since(startTime)
        statusCode := c.Writer.Status()
        
        fmt.Printf("[CUSTOM] Completed %s %s (%d) in %v\n", 
            method, path, statusCode, duration)
    }
}

// Authentication middleware
func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get authorization header
        token := c.GetHeader("Authorization")
        
        // Simple token check (in real app, verify JWT or session)
        if token == "" {
            // If no token, return unauthorized and stop processing
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Authorization token required",
            })
            c.Abort() // Stop processing, don't call c.Next()
            return
        }
        
        // Simple validation (replace with real token validation)
        if token != "Bearer valid-token" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid authorization token",
            })
            c.Abort()
            return
        }
        
        // Set user information in context for use by handlers
        c.Set("user_id", "123")
        c.Set("user_role", "admin")
        
        // Continue to next middleware or handler
        c.Next()
    }
}

// Rate limiting middleware (simplified)
func RateLimit() gin.HandlerFunc {
    // In-memory map to track requests (use Redis in production)
    requestCounts := make(map[string]int)
    
    return func(c *gin.Context) {
        clientIP := c.ClientIP()
        
        // Check current request count for this IP
        if requestCounts[clientIP] >= 100 { // Allow 100 requests
            c.JSON(http.StatusTooManyRequests, gin.H{
                "error": "Rate limit exceeded",
            })
            c.Abort()
            return
        }
        
        // Increment request count
        requestCounts[clientIP]++
        
        c.Next()
    }
}

// CORS middleware
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Set CORS headers
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
        
        // Handle preflight requests
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    }
}

func main() {
    // Create router without default middleware
    router := gin.New()
    
    // Add built-in middleware manually
    router.Use(gin.Logger())   // Request logging middleware
    router.Use(gin.Recovery()) // Panic recovery middleware
    
    // Add custom middleware globally (applies to all routes)
    router.Use(CustomLogger())  // Our custom logging middleware
    router.Use(CORSMiddleware()) // CORS handling
    
    // Public routes (no authentication required)
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Public endpoint - no auth required",
        })
    })
    
    router.POST("/login", func(c *gin.Context) {
        // Simulate login logic
        c.JSON(http.StatusOK, gin.H{
            "token": "Bearer valid-token",
            "message": "Login successful",
        })
    })
    
    // Protected routes (authentication required)
    // Apply middleware to specific routes
    router.GET("/profile", AuthRequired(), func(c *gin.Context) {
        // Get user information set by middleware
        userID, _ := c.Get("user_id")
        userRole, _ := c.Get("user_role")
        
        c.JSON(http.StatusOK, gin.H{
            "user_id":   userID,
            "user_role": userRole,
            "message":   "Protected profile data",
        })
    })
    
    // Multiple middleware on single route
    router.GET("/admin", AuthRequired(), RateLimit(), func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Admin only endpoint with rate limiting",
        })
    })
    
    // Middleware with parameters
    router.GET("/limited", func(c *gin.Context) {
        // Inline middleware
        clientIP := c.ClientIP()
        
        // Simple rate check
        if clientIP == "127.0.0.1" {
            c.JSON(http.StatusOK, gin.H{
                "message": "Access allowed for localhost",
            })
        } else {
            c.JSON(http.StatusForbidden, gin.H{
                "error": "Access denied",
            })
        }
    })
    
    router.Run(":8080")
}
```

**Middleware Key Points:**
- `c.Next()` - Proceeds to next middleware/handler
- `c.Abort()` - Stops processing chain
- `c.Set()/c.Get()` - Store/retrieve data in request context
- Middleware can run before and after handlers
- Order matters - middleware runs in the order it's registered

## Error Handling

```go
package main

import (
    "errors"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

// Custom error types
type APIError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Details string `json:"details,omitempty"`
}

// Implement the error interface
func (e APIError) Error() string {
    return e.Message
}

// Error handling middleware
func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Execute the request
        c.Next()
        
        // Check if any errors occurred during request processing
        if len(c.Errors) > 0 {
            // Get the last error
            err := c.Errors.Last()
            
            // Check error type and respond accordingly
            switch e := err.Err.(type) {
            case APIError:
                // Custom API error
                c.JSON(e.Code, gin.H{
                    "error":   e.Message,
                    "details": e.Details,
                })
            default:
                // Generic error
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error": "Internal server error",
                })
            }
            
            // Prevent further error processing
            c.Abort()
        }
    }
}

// Simulated database operations
func getUserFromDB(id string) (gin.H, error) {
    if id == "999" {
        return nil, APIError{
            Code:    http.StatusNotFound,
            Message: "User not found",
            Details: fmt.Sprintf("No user exists with ID: %s", id),
        }
    }
    
    if id == "500" {
        return nil, errors.New("database connection failed")
    }
    
    return gin.H{
        "id":    id,
        "name":  "John Doe",
        "email": "john@example.com",
    }, nil
}

func createUserInDB(user gin.H) error {
    if user["name"] == "error" {
        return APIError{
            Code:    http.StatusBadRequest,
            Message: "Invalid user data",
            Details: "Username 'error' is not allowed",
        }
    }
    return nil
}

func main() {
    router := gin.Default()
    
    // Add error handling middleware
    router.Use(ErrorHandler())
    
    // Route with error handling using middleware
    router.GET("/users/:id", func(c *gin.Context) {
        userID := c.Param("id")
        
        // Call function that might return an error
        user, err := getUserFromDB(userID)
        if err != nil {
            // Add error to context - middleware will handle it
            c.Error(err)
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "data":    user,
            "message": "User retrieved successfully",
        })
    })
    
    // Direct error handling (without middleware)
    router.GET("/direct-error/:id", func(c *gin.Context) {
        userID := c.Param("id")
        
        user, err := getUserFromDB(userID)
        if err != nil {
            // Handle error directly in handler
            if apiErr, ok := err.(APIError); ok {
                // Custom API error
                c.JSON(apiErr.Code, gin.H{
                    "error":   apiErr.Message,
                    "details": apiErr.Details,
                })
            } else {
                // Generic error
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error": "Internal server error",
                })
            }
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "data": user,
        })
    })
    
    // POST route with validation and error handling
    router.POST("/users", func(c *gin.Context) {
        var user gin.H
        
        // Bind JSON and handle binding errors
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error":   "Invalid JSON data",
                "details": err.Error(),
            })
            return
        }
        
        // Validate required fields
        if user["name"] == nil || user["name"] == "" {
            c.JSON(http.StatusBadRequest, gin.H{
                "error":   "Validation failed",
                "details": "Name is required",
            })
            return
        }
        
        // Try to create user
        if err := createUserInDB(user); err != nil {
            c.Error(err) // Let middleware handle it
            return
        }
        
        c.JSON(http.StatusCreated, gin.H{
            "message": "User created successfully",
            "data":    user,
        })
    })
    
    // Panic recovery example
    router.GET("/panic", func(c *gin.Context) {
        // This will cause a panic, but Gin's Recovery middleware will catch it
        panic("Something went wrong!")
    })
    
    // Custom error response format
    router.GET("/custom-error", func(c *gin.Context) {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error": gin.H{
                "code":    "INVALID_REQUEST",
                "message": "The request contains invalid parameters",
                "timestamp": "2024-01-15T10:30:00Z",
            },
        })
    })
    
    // Multiple error scenarios
    router.GET("/errors/:type", func(c *gin.Context) {
        errorType := c.Param("type")
        
        switch errorType {
        case "not-found":
            c.JSON(http.StatusNotFound, gin.H{
                "error": "Resource not found",
            })
        case "unauthorized":
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Authentication required",
            })
        case "forbidden":
            c.JSON(http.StatusForbidden, gin.H{
                "error": "Access denied",
            })
        case "validation":
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Validation failed",
                "details": []string{
                    "Name is required",
                    "Email must be valid",
                    "Age must be between 18 and 100",
                },
            })
        case "server":
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Internal server error",
            })
        default:
            c.JSON(http.StatusOK, gin.H{
                "message": "No error triggered",
            })
        }
    })
    
    router.Run(":8080")
}
```

**Error Handling Best Practices:**
- Use consistent error response format
- Log errors for debugging
- Never expose internal error details to clients
- Use appropriate HTTP status codes
- Handle validation errors separately
- Use middleware for centralized error handling

## Route Groups

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Middleware for API versioning
func APIVersionMiddleware(version string) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Add version to response headers
        c.Header("API-Version", version)
        c.Next()
    }
}

// Authentication middleware for admin routes
func AdminAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Simple admin check (replace with real authentication)
        token := c.GetHeader("Admin-Token")
        if token != "admin-secret" {
            c.JSON(http.StatusForbidden, gin.H{
                "error": "Admin access required",
            })
            c.Abort()
            return
        }
        c.Next()
    }
}

func main() {
    router := gin.Default()
    
    // Root level routes
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Welcome to the API",
        })
    })
    
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "healthy",
        })
    })
    
    // API v1 group
    // Group allows you to organize related routes under a common prefix
    v1 := router.Group("/api/v1")
    {
        // Apply middleware to entire group
        v1.Use(APIVersionMiddleware("1.0"))
        
        // User routes within v1 group
        // These routes will have prefix /api/v1
        v1.GET("/users", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "users":   []string{"user1", "user2"},
                "version": "v1",
            })
        })
        
        v1.POST("/users", func(c *gin.Context) {
            c.JSON(http.StatusCreated, gin.H{
                "message": "User created in v1",
                "version": "v1",
            })
        })
        
        v1.GET("/users/:id", func(c *gin.Context) {
            userID := c.Param("id")
            c.JSON(http.StatusOK, gin.H{
                "user_id": userID,
                "version":  "v1",
            })
        })
    }
    
    // API v2 group with different functionality
    v2 := router.Group("/api/v2")
    {
        v2.Use(APIVersionMiddleware("2.0"))
        
        // Enhanced user routes in v2
        v2.GET("/users", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "users": []gin.H{
                    {"id": 1, "name": "John", "email": "john@example.com"},
                    {"id": 2, "name": "Jane", "email": "jane@example.com"},
                },
                "version": "v2",
                "features": []string{"pagination", "filtering", "sorting"},
            })
        })
        
        v2.POST("/users", func(c *gin.Context) {
            c.JSON(http.StatusCreated, gin.H{
                "message": "User created in v2 with enhanced validation",
                "version": "v2",
            })
        })
    }
    
    // Admin routes group
    admin := router.Group("/admin")
    {
        // Apply admin authentication to all routes in this group
        admin.Use(AdminAuth())
        
        // Admin dashboard
        admin.GET("/dashboard", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "message": "Admin dashboard data",
                "stats": gin.H{
                    "total_users": 1250,
                    "active_sessions": 89,
                    "server_uptime": "24h 15m",
                },
            })
        })
        
        // User management for admins
        admin.GET("/users", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "message": "All users with admin details",
                "users": []gin.H{
                    {"id": 1, "name": "John", "status": "active", "last_login": "2024-01-15"},
                    {"id": 2, "name": "Jane", "status": "inactive", "last_login": "2024-01-10"},
                },
            })
        })
        
        admin.DELETE("/users/:id", func(c *gin.Context) {
            userID := c.Param("id")
            c.JSON(http.StatusOK, gin.H{
                "message": "User deleted by admin",
                "user_id": userID,
            })
        })
    }
    
    // Nested groups example
    api := router.Group("/api")
    {
        // Products group within API group
        products := api.Group("/products")
        {
            products.GET("/", func(c *gin.Context) {
                c.JSON(http.StatusOK, gin.H{
                    "products": []string{"laptop", "phone", "tablet"},
                })
            })
            
            products.POST("/", func(c *gin.Context) {
                c.JSON(http.StatusCreated, gin.H{
                    "message": "Product created",
                })
            })
            
            // Category routes within products
            categories := products.Group("/:id/categories")
            {
                categories.GET("/", func(c *gin.Context) {
                    productID := c.Param("id")
                    c.JSON(http.StatusOK, gin.H{
                        "product_id": productID,
                        "categories": []string{"electronics", "computers"},
                    })
                })
            }
        }
        
        // Orders group within API group
        orders := api.Group("/orders")
        {
            orders.GET("/", func(c *gin.Context) {
                c.JSON(http.StatusOK, gin.H{
                    "orders": []gin.H{
                        {"id": 1, "status": "pending", "total": 99.99},
                        {"id": 2, "status": "completed", "total": 149.99},
                    },
                })
            })
            
            orders.GET("/:id", func(c *gin.Context) {
                orderID := c.Param("id")
                c.JSON(http.StatusOK, gin.H{
                    "order_id": orderID,
                    "status": "pending",
                    "items": []string{"laptop", "mouse"},
                })
            })
        }
    }
    
    // Group with multiple middleware
    protected := router.Group("/protected")
    {
        // Chain multiple middleware functions
        protected.Use(func(c *gin.Context) {
            // Custom logging middleware
            fmt.Printf("Accessing protected route: %s\n", c.Request.URL.Path)
            c.Next()
        })
        
        protected.Use(func(c *gin.Context) {
            // Simple authentication check
            if c.GetHeader("Authorization") == "" {
                c.JSON(http.StatusUnauthorized, gin.H{
                    "error": "Authorization required",
                })
                c.Abort()
                return
            }
            c.Next()
        })
        
        protected.GET("/data", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "message": "Protected data",
                "data": "sensitive information",
            })
        })
    }
    
    router.Run(":8080")
}
```

**Route Groups Benefits:**
- Organize related routes logically
- Apply middleware to specific route groups
- Easy to version APIs
- Cleaner code structure
- Prefix management for different API sections

## File Uploads

```go
package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "io"
    "mime/multipart"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "time"
)

// Function to validate file type
func isValidFileType(filename string, allowedTypes []string) bool {
    ext := strings.ToLower(filepath.Ext(filename))
    for _, allowedType := range allowedTypes {
        if ext == allowedType {
            return true
        }
    }
    return false
}

// Function to generate unique filename
func generateUniqueFilename(originalName string) string {
    // Get file extension
    ext := filepath.Ext(originalName)
    // Get filename without extension
    name := strings.TrimSuffix(originalName, ext)
    // Add timestamp to make it unique
    timestamp := time.Now().Unix()
    return fmt.Sprintf("%s_%d%s", name, timestamp, ext)
}

// Function to save uploaded file
func saveUploadedFile(file *multipart.FileHeader, dst string) error {
    // Open the uploaded file
    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close() // Always close the source file
    
    // Create the destination file
    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close() // Always close the destination file
    
    // Copy file contents
    _, err = io.Copy(out, src)
    return err
}

func main() {
    router := gin.Default()
    
    // Set maximum memory for multipart forms (default is 32 MiB)
    router.MaxMultipartMemory = 8 << 20 // 8 MiB
    
    // Create uploads directory if it doesn't exist
    if err := os.MkdirAll("uploads", 0755); err != nil {
        fmt.Printf("Failed to create uploads directory: %v\n", err)
    }
    
    // Single file upload
    router.POST("/upload", func(c *gin.Context) {
        // Get the uploaded file from form field named 'file'
        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "No file uploaded or invalid file",
                "details": err.Error(),
            })
            return
        }
        
        // Validate file size (5MB limit)
        const maxFileSize = 5 << 20 // 5 MB
        if file.Size > maxFileSize {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "File too large",
                "details": fmt.Sprintf("File size: %d bytes, max allowed: %d bytes", 
                    file.Size, maxFileSize),
            })
            return
        }
        
        // Validate file type
        allowedTypes := []string{".jpg", ".jpeg", ".png", ".gif", ".pdf", ".txt"}
        if !isValidFileType(file.Filename, allowedTypes) {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid file type",
                "details": fmt.Sprintf("Allowed types: %v", allowedTypes),
            })
            return
        }
        
        // Generate unique filename to avoid conflicts
        uniqueFilename := generateUniqueFilename(file.Filename)
        filePath := filepath.Join("uploads", uniqueFilename)
        
        // Save the file
        if err := c.SaveUploadedFile(file, filePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to save file",
                "details": err.Error(),
            })
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "message": "File uploaded successfully",
            "file_info": gin.H{
                "original_name": file.Filename,
                "saved_name":    uniqueFilename,
                "size":          file.Size,
                "path":          filePath,
            },
        })
    })
    
    // Multiple files upload
    router.POST("/upload-multiple", func(c *gin.Context) {
        // Parse multipart form
        form, err := c.MultipartForm()
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Failed to parse multipart form",
                "details": err.Error(),
            })
            return
        }
        
        // Get files from form field named 'files'
        files := form.File["files"]
        if len(files) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "No files uploaded",
            })
            return
        }
        
        // Limit number of files
        const maxFiles = 5
        if len(files) > maxFiles {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": fmt.Sprintf("Too many files. Maximum allowed: %d", maxFiles),
            })
            return
        }
        
        var uploadedFiles []gin.H
        var errors []string
        
        // Process each file
        for _, file := range files {
            // Validate each file
            if file.Size > (5 << 20) { // 5MB limit
                errors = append(errors, fmt.Sprintf("File %s is too large", file.Filename))
                continue
            }
            
            allowedTypes := []string{".jpg", ".jpeg", ".png", ".gif"}
            if !isValidFileType(file.Filename, allowedTypes) {
                errors = append(errors, fmt.Sprintf("File %s has invalid type", file.Filename))
                continue
            }
            
            // Generate unique filename and save
            uniqueFilename := generateUniqueFilename(file.Filename)
            filePath := filepath.Join("uploads", uniqueFilename)
            
            if err := c.SaveUploadedFile(file, filePath); err != nil {
                errors = append(errors, fmt.Sprintf("Failed to save %s: %v", file.Filename, err))
                continue
            }
            
            // Add to successful uploads
            uploadedFiles = append(uploadedFiles, gin.H{
                "original_name": file.Filename,
                "saved_name":    uniqueFilename,
                "size":          file.Size,
                "path":          filePath,
            })
        }
        
        response := gin.H{
            "message": fmt.Sprintf("Processed %d files", len(files)),
            "uploaded_files": uploadedFiles,
            "uploaded_count": len(uploadedFiles),
        }
        
        if len(errors) > 0 {
            response["errors"] = errors
            response["error_count"] = len(errors)
        }
        
        c.JSON(http.StatusOK, response)
    })
    
    // Upload with additional form data
    router.POST("/upload-with-data", func(c *gin.Context) {
        // Get form data
        title := c.PostForm("title")
        description := c.PostForm("description")
        category := c.PostForm("category")
        
        // Validate required fields
        if title == "" {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Title is required",
            })
            return
        }
        
        // Get uploaded file
        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "File is required",
            })
            return
        }
        
        // Save file
        uniqueFilename := generateUniqueFilename(file.Filename)
        filePath := filepath.Join("uploads", uniqueFilename)
        
        if err := c.SaveUploadedFile(file, filePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to save file",
            })
            return
        }
        
        // In a real app, you'd save this data to a database
        c.JSON(http.StatusOK, gin.H{
            "message": "File and data uploaded successfully",
            "data": gin.H{
                "title":       title,
                "description": description,
                "category":    category,
                "file": gin.H{
                    "original_name": file.Filename,
                    "saved_name":    uniqueFilename,
                    "size":          file.Size,
                    "path":          filePath,
                },
            },
        })
    })
    
    // Serve uploaded files
    router.GET("/files/:filename", func(c *gin.Context) {
        filename := c.Param("filename")
        filePath := filepath.Join("uploads", filename)
        
        // Check if file exists
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "File not found",
            })
            return
        }
        
        // Serve the file
        c.File(filePath)
    })
    
    // Get file information
    router.GET("/files/:filename/info", func(c *gin.Context) {
        filename := c.Param("filename")
        filePath := filepath.Join("uploads", filename)
        
        // Get file stats
        fileInfo, err := os.Stat(filePath)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "File not found",
            })
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "filename": filename,
            "size":     fileInfo.Size(),
            "modified": fileInfo.ModTime(),
            "is_dir":   fileInfo.IsDir(),
        })
    })
    
    // Delete uploaded file
    router.DELETE("/files/:filename", func(c *gin.Context) {
        filename := c.Param("filename")
        filePath := filepath.Join("uploads", filename)
        
        // Check if file exists
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "File not found",
            })
            return
        }
        
        // Delete the file
        if err := os.Remove(filePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to delete file",
                "details": err.Error(),
            })
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "message": "File deleted successfully",
            "filename": filename,
        })
    })
    
    // Upload progress simulation (for demonstration)
    router.POST("/upload-progress", func(c *gin.Context) {
        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
            return
        }
        
        // Simulate processing time based on file size
        processingTime := time.Duration(file.Size/1000000) * time.Second // 1s per MB
        if processingTime > 10*time.Second {
            processingTime = 10 * time.Second // Cap at 10 seconds
        }
        
        // In a real application, you'd implement actual progress tracking
        // using websockets or server-sent events
        time.Sleep(processingTime)
        
        uniqueFilename := generateUniqueFilename(file.Filename)
        filePath := filepath.Join("uploads", uniqueFilename)
        
        if err := c.SaveUploadedFile(file, filePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "message": "File uploaded with simulated progress",
            "processing_time": processingTime.String(),
            "file_info": gin.H{
                "original_name": file.Filename,
                "saved_name":    uniqueFilename,
                "size":          file.Size,
            },
        })
    })
    
    router.Run(":8080")
}
```

**File Upload Best Practices:**
- Always validate file types and sizes
- Generate unique filenames to avoid conflicts
- Store files outside the web root for security
- Implement virus scanning for production
- Use cloud storage (AWS S3, Google Cloud) for scalability
- Implement progress tracking for large files
- Add rate limiting to prevent abuse

## Database Integration

```go
package main

import (
    "database/sql"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "time"
    
    // PostgreSQL driver
    _ "github.com/lib/pq"
    
    // For MongoDB (commented out - choose one database)
    // "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"
)

// User struct for database operations
type User struct {
    ID        int       `json:"id" db:"id"`
    Name      string    `json:"name" db:"name"`
    Email     string    `json:"email" db:"email"`
    Age       int       `json:"age" db:"age"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Database connection global variable
var db *sql.DB

// Initialize database connection
func initDB() error {
    var err error
    
    // PostgreSQL connection string
    // Format: "postgres://username:password@localhost/dbname?sslmode=disable"
    connStr := "postgres://username:password@localhost/testdb?sslmode=disable"
    
    // Open database connection
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        return fmt.Errorf("failed to open database: %v", err)
    }
    
    // Test the connection
    if err = db.Ping(); err != nil {
        return fmt.Errorf("failed to ping database: %v", err)
    }
    
    // Set connection pool settings
    db.SetMaxOpenConns(25)                 // Maximum number of open connections
    db.SetMaxIdleConns(5)                  // Maximum number of idle connections
    db.SetConnMaxLifetime(5 * time.Minute) // Maximum lifetime of a connection
    
    fmt.Println("Database connection established")
    return nil
}

// Create users table
func createUsersTable() error {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,                    -- Auto-incrementing primary key
        name VARCHAR(100) NOT NULL,               -- User name (required)
        email VARCHAR(255) UNIQUE NOT NULL,       -- Email (unique and required)
        age INTEGER CHECK (age >= 0 AND age <= 150), -- Age with constraints
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Update timestamp
    )`
    
    // Execute the query
    _, err := db.Exec(query)
    if err != nil {
        return fmt.Errorf("failed to create users table: %v", err)
    }
    
    fmt.Println("Users table created or already exists")
    return nil
}

// Database operations

// Get all users with pagination
func getUsers(page, limit int) ([]User, error) {
    // Calculate offset for pagination
    offset := (page - 1) * limit
    
    // SQL query with LIMIT and OFFSET for pagination
    query := `
        SELECT id, name, email, age, created_at, updated_at 
        FROM users 
        ORDER BY created_at DESC 
        LIMIT $1 OFFSET $2`
    
    // Execute query with parameters ($1 = limit, $2 = offset)
    rows, err := db.Query(query, limit, offset)
    if err != nil {
        return nil, fmt.Errorf("failed to query users: %v", err)
    }
    defer rows.Close() // Always close rows when done
    
    var users []User
    
    // Iterate through result rows
    for rows.Next() {
        var user User
        
        // Scan row data into user struct
        err := rows.Scan(
            &user.ID,
            &user.Name, 
            &user.Email, 
            &user.Age, 
            &user.CreatedAt, 
            &user.UpdatedAt,
        )
        if err != nil {
            return nil, fmt.Errorf("failed to scan user row: %v", err)
        }
        
        users = append(users, user)
    }
    
    // Check for any error that occurred during iteration
    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %v", err)
    }
    
    return users, nil
}

// Get user by ID
func getUserByID(id int) (*User, error) {
    var user User
    
    // SQL query to get single user by ID
    query := `
        SELECT id, name, email, age, created_at, updated_at 
        FROM users 
        WHERE id = $1`
    
    // QueryRow returns at most one row
    row := db.QueryRow(query, id)
    
    // Scan the row into user struct
    err := row.Scan(
        &user.ID,
        &user.Name,
        &user.Email,
        &user.Age,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("failed to get user: %v", err)
    }
    
    return &user, nil
}

// Create new user
func createUser(user User) (*User, error) {
    // SQL query to insert new user and return the created record
    query := `
        INSERT INTO users (name, email, age, created_at, updated_at) 
        VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) 
        RETURNING id, created_at, updated_at`
    
    // Execute query and scan returned values
    err := db.QueryRow(query, user.Name, user.Email, user.Age).Scan(
        &user.ID,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    
    if err != nil {
        return nil, fmt.Errorf("failed to create user: %v", err)
    }
    
    return &user, nil
}

// Update user
func updateUser(id int, user User) (*User, error) {
    // SQL query to update user and return updated record
    query := `
        UPDATE users 
        SET name = $1, email = $2, age = $3, updated_at = CURRENT_TIMESTAMP 
        WHERE id = $4 
        RETURNING id, name, email, age, created_at, updated_at`
    
    var updatedUser User
    
    // Execute update and scan returned values
    err := db.QueryRow(query, user.Name, user.Email, user.Age, id).Scan(
        &updatedUser.ID,
        &updatedUser.Name,
        &updatedUser.Email,
        &updatedUser.Age,
        &updatedUser.CreatedAt,
        &updatedUser.UpdatedAt,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("failed to update user: %v", err)
    }
    
    return &updatedUser, nil
}

// Delete user
func deleteUser(id int) error {
    // SQL query to delete user
    query := `DELETE FROM users WHERE id = $1`
    
    // Execute delete query
    result, err := db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("failed to delete user: %v", err)
    }
    
    // Check if any rows were affected
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to get rows affected: %v", err)
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("user not found")
    }
    
    return nil
}

// Get total user count (for pagination metadata)
func getUserCount() (int, error) {
    var count int
    
    query := `SELECT COUNT(*) FROM users`
    err := db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, fmt.Errorf("failed to get user count: %v", err)
    }
    
    return count, nil
}

func main() {
    // Initialize database connection
    if err := initDB(); err != nil {
        fmt.Printf("Database initialization failed: %v\n", err)
        return
    }
    defer db.Close() // Close database connection when main exits
    
    // Create tables
    if err := createUsersTable(); err != nil {
        fmt.Printf("Table creation failed: %v\n", err)
        return
    }
    
    // Create Gin router
    router := gin.Default()
    
    // API routes
    
    // GET /users - Get all users with pagination
    router.GET("/users", func(c *gin.Context) {
        // Get pagination parameters from query string
        pageStr := c.DefaultQuery("page", "1")
        limitStr := c.DefaultQuery("limit", "10")
        
        // Convert string parameters to integers
        page, err := strconv.Atoi(pageStr)
        if err != nil || page < 1 {
            page = 1
        }
        
        limit, err := strconv.Atoi(limitStr)
        if err != nil || limit < 1 || limit > 100 {
            limit = 10 // Default limit, max 100
        }
        
        // Get users from database
        users, err := getUsers(page, limit)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to retrieve users",
                "details": err.Error(),
            })
            return
        }
        
        // Get total count for pagination metadata
        totalCount, err := getUserCount()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to get user count",
            })
            return
        }
        
        // Calculate pagination metadata
        totalPages := (totalCount + limit - 1) / limit // Ceiling division
        
        c.JSON(http.StatusOK, gin.H{
            "data": users,
            "pagination": gin.H{
                "page":         page,
                "limit":        limit,
                "total_count":  totalCount,
                "total_pages":  totalPages,
                "has_next":     page < totalPages,
                "has_prev":     page > 1,
            },
        })
    })
    
    // GET /users/:id - Get user by ID
    router.GET("/users/:id", func(c *gin.Context) {
        // Get ID parameter and convert to integer
        idStr := c.Param("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid user ID",
            })
            return
        }
        
        // Get user from database
        user, err := getUserByID(id)
        if err != nil {
            if err.Error() == "user not found" {
                c.JSON(http.StatusNotFound, gin.H{
                    "error": "User not found",
                })
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error": "Failed to retrieve user",
                    "details": err.Error(),
                })
            }
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "data": user,
        })
    })
    
    // POST /users - Create new user
    router.POST("/users", func(c *gin.Context) {
        var user User
        
        // Bind JSON request body to user struct
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid request data",
                "details": err.Error(),
            })
            return
        }
        
        // Validate required fields
        if user.Name == "" || user.Email == "" {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Name and email are required",
            })
            return
        }
        
        // Create user in database
        createdUser, err := createUser(user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to create user",
                "details": err.Error(),
            })
            return
        }
        
        c.JSON(http.StatusCreated, gin.H{