package container

import (
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

// RouteInfo holds information about a route
type RouteInfo struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
	Middleware  []gin.HandlerFunc
}

// AutoRouter handles automatic route registration
type AutoRouter struct {
	routes []RouteInfo
}

// NewAutoRouter creates a new auto router
func NewAutoRouter() *AutoRouter {
	return &AutoRouter{
		routes: make([]RouteInfo, 0),
	}
}

// RegisterHandler automatically registers all routes for a handler
func (ar *AutoRouter) RegisterHandler(basePath string, handler interface{}) {
	handlerType := reflect.TypeOf(handler)
	handlerValue := reflect.ValueOf(handler)

	// Iterate through all methods of the handler
	for i := 0; i < handlerType.NumMethod(); i++ {
		method := handlerType.Method(i)
		methodValue := handlerValue.Method(i)

		// Check if method has route tag
		if route := ar.parseMethodName(method.Name); route != nil {
			// Create gin.HandlerFunc wrapper
			handlerFunc := ar.createHandlerFunc(methodValue)

			route.Path = basePath + route.Path
			route.HandlerFunc = handlerFunc

			ar.routes = append(ar.routes, *route)
			log.Printf("🔗 Auto-registered route: %s %s -> %s.%s",
				route.Method, route.Path, handlerType.Elem().Name(), method.Name)
		}
	}
}

// parseMethodName converts method name to route info using convention
func (ar *AutoRouter) parseMethodName(methodName string) *RouteInfo {
	// Convention-based routing:
	// CreateUser -> POST /
	// GetUser -> GET /:id
	// GetAllUsers -> GET /
	// UpdateUser -> PUT /:id
	// DeleteUser -> DELETE /:id
	// GetUserProfile -> GET /:id/profile
	// ChangePassword -> PUT /:id/password

	route := &RouteInfo{}

	switch {
	case strings.HasPrefix(methodName, "Create"):
		route.Method = "POST"
		route.Path = ""

	case strings.HasPrefix(methodName, "GetAll"):
		route.Method = "GET"
		route.Path = ""

	case strings.HasPrefix(methodName, "Get") && strings.Contains(methodName, "Profile"):
		route.Method = "GET"
		route.Path = "/:id/profile"

	case strings.HasPrefix(methodName, "Get"):
		route.Method = "GET"
		route.Path = "/:id"

	case strings.HasPrefix(methodName, "Update"):
		route.Method = "PUT"
		route.Path = "/:id"

	case strings.HasPrefix(methodName, "Delete"):
		route.Method = "DELETE"
		route.Path = "/:id"

	case strings.HasPrefix(methodName, "Change") && strings.Contains(methodName, "Password"):
		route.Method = "PUT"
		route.Path = "/:id/password"

	default:
		// Not a route method
		return nil
	}

	return route
}

// createHandlerFunc creates a gin.HandlerFunc from a method
func (ar *AutoRouter) createHandlerFunc(methodValue reflect.Value) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Call the handler method with gin.Context
		methodValue.Call([]reflect.Value{reflect.ValueOf(c)})
	}
}

// RegisterRoutes registers all collected routes to a router group
func (ar *AutoRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	for _, route := range ar.routes {
		switch route.Method {
		case "GET":
			routerGroup.GET(route.Path, route.HandlerFunc)
		case "POST":
			routerGroup.POST(route.Path, route.HandlerFunc)
		case "PUT":
			routerGroup.PUT(route.Path, route.HandlerFunc)
		case "DELETE":
			routerGroup.DELETE(route.Path, route.HandlerFunc)
		case "PATCH":
			routerGroup.PATCH(route.Path, route.HandlerFunc)
		}
	}
}

// GetRoutes returns all registered routes
func (ar *AutoRouter) GetRoutes() []RouteInfo {
	return ar.routes
}
