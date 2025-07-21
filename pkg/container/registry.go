package container

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

// RouteRegistrar defines interface for route registration
type RouteRegistrar interface {
	RegisterRoutes(router *gin.RouterGroup)
}

// ServiceRegistry manages service registration
type ServiceRegistry struct {
	services    map[string]interface{}
	routeSetups []RouteSetupFunc
}

// RouteSetupFunc defines function signature for route setup
type RouteSetupFunc func(router *gin.RouterGroup, container *Container)

// NewServiceRegistry creates a new service registry
func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		services:    make(map[string]interface{}),
		routeSetups: make([]RouteSetupFunc, 0),
	}
}

// Register registers a service in the registry
func (r *ServiceRegistry) Register(name string, service interface{}) {
	r.services[name] = service
}

// Get retrieves a service from the registry
func (r *ServiceRegistry) Get(name string) (interface{}, bool) {
	service, exists := r.services[name]
	return service, exists
}

// RegisterRouteSetup registers a route setup function
func (r *ServiceRegistry) RegisterRouteSetup(setup RouteSetupFunc) {
	r.routeSetups = append(r.routeSetups, setup)
}

// SetupAllRoutes sets up all registered routes
func (r *ServiceRegistry) SetupAllRoutes(router *gin.RouterGroup, container *Container) {
	for _, setup := range r.routeSetups {
		setup(router, container)
	}
}

// AutoRegister automatically registers services based on reflection
func (r *ServiceRegistry) AutoRegister(services ...interface{}) {
	for _, service := range services {
		serviceName := reflect.TypeOf(service).Elem().Name()
		r.Register(serviceName, service)
	}
}

// Example usage:
// registry.AutoRegister(userService, postService, userHandler, postHandler)
