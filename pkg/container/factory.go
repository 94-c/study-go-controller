package container

import (
	"fmt"
	"reflect"
	"study-go-controller/internal/domain/post/handler"
	postRepo "study-go-controller/internal/domain/post/repository"
	postService "study-go-controller/internal/domain/post/service"
	userHandler "study-go-controller/internal/domain/user/handler"
	userRepo "study-go-controller/internal/domain/user/repository"
	userService "study-go-controller/internal/domain/user/service"
	"study-go-controller/pkg/database"
)

// Factory creates instances with automatic dependency injection
type Factory struct {
	instances map[reflect.Type]interface{}
	database  *database.Database
}

// NewFactory creates a new factory instance
func NewFactory() (*Factory, error) {
	db, err := database.NewDatabase()
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(); err != nil {
		return nil, err
	}

	return &Factory{
		instances: make(map[reflect.Type]interface{}),
		database:  db,
	}, nil
}

// Get retrieves or creates an instance of the specified type
func (f *Factory) Get(serviceType reflect.Type) (interface{}, error) {
	// Check if instance already exists
	if instance, exists := f.instances[serviceType]; exists {
		return instance, nil
	}

	// Create new instance based on type
	var instance interface{}
	var err error

	switch serviceType.String() {
	case "repository.UserRepository":
		instance = userRepo.NewUserRepository(f.database.DB)
	case "repository.PostRepository":
		instance = postRepo.NewPostRepository(f.database.DB)
	case "service.UserService":
		userRepoInstance, err := f.Get(reflect.TypeOf((*userRepo.UserRepository)(nil)).Elem())
		if err != nil {
			return nil, err
		}
		instance = userService.NewUserService(userRepoInstance.(userRepo.UserRepository))
	case "service.PostService":
		postRepoInstance, err := f.Get(reflect.TypeOf((*postRepo.PostRepository)(nil)).Elem())
		if err != nil {
			return nil, err
		}
		instance = postService.NewPostService(postRepoInstance.(postRepo.PostRepository))
	case "*handler.UserHandler":
		userSvcInstance, err := f.Get(reflect.TypeOf((*userService.UserService)(nil)).Elem())
		if err != nil {
			return nil, err
		}
		instance = userHandler.NewUserHandler(userSvcInstance.(userService.UserService))
	case "*handler.PostHandler":
		postSvcInstance, err := f.Get(reflect.TypeOf((*postService.PostService)(nil)).Elem())
		if err != nil {
			return nil, err
		}
		instance = handler.NewPostHandler(postSvcInstance.(postService.PostService))
	default:
		return nil, fmt.Errorf("unknown service type: %s", serviceType.String())
	}

	if err != nil {
		return nil, err
	}

	// Store instance for reuse
	f.instances[serviceType] = instance
	return instance, nil
}

// GetUserHandler convenience method for getting user handler
func (f *Factory) GetUserHandler() (*userHandler.UserHandler, error) {
	instance, err := f.Get(reflect.TypeOf((*userHandler.UserHandler)(nil)))
	if err != nil {
		return nil, err
	}
	return instance.(*userHandler.UserHandler), nil
}

// GetPostHandler convenience method for getting post handler
func (f *Factory) GetPostHandler() (*handler.PostHandler, error) {
	instance, err := f.Get(reflect.TypeOf((*handler.PostHandler)(nil)))
	if err != nil {
		return nil, err
	}
	return instance.(*handler.PostHandler), nil
}
