# 🔧 의존성 주입 가이드

Go 프로젝트에서 의존성 주입(Dependency Injection)을 효과적으로 관리하는 방법을 설명합니다.

## 🚨 **현재 main.go의 문제점**

### ❌ **수동 의존성 관리의 한계**

```go
// 기존 main.go - 새로운 도메인 추가 시마다 수정 필요
func main() {
    // 데이터베이스 초기화
    db, err := database.NewDatabase()
    
    // 리포지토리 초기화 (새 도메인마다 추가)
    userRepository := userRepo.NewUserRepository(db.DB)
    postRepository := postRepo.NewPostRepository(db.DB)
    // ❌ 새로운 도메인 추가 시: productRepository := productRepo.NewProductRepository(db.DB)
    
    // 서비스 초기화 (새 도메인마다 추가)
    userSvc := userService.NewUserService(userRepository)
    postSvc := postService.NewPostService(postRepository)
    // ❌ 새로운 도메인 추가 시: productSvc := productService.NewProductService(productRepository)
    
    // 핸들러 초기화 (새 도메인마다 추가)
    userHdl := userHandler.NewUserHandler(userSvc)
    postHdl := handler.NewPostHandler(postSvc)
    // ❌ 새로운 도메인 추가 시: productHdl := productHandler.NewProductHandler(productSvc)
    
    // 라우터 설정 (새 도메인마다 추가)
    userRoutes.SetupUserRoutes(v1, userHdl)
    routes.SetupPostRoutes(v1, postHdl)
    // ❌ 새로운 도메인 추가 시: productRoutes.SetupProductRoutes(v1, productHdl)
}
```

### 🔥 **문제점들**

1. **🔄 반복적인 수정**: 새로운 도메인 추가 시마다 main.go 수정
2. **📈 복잡도 증가**: 도메인이 늘어날수록 main.go가 거대해짐
3. **🐛 실수 유발**: 수동 등록으로 인한 누락 또는 오타 가능성
4. **🧪 테스트 어려움**: 의존성이 하드코딩되어 모킹 어려움
5. **🔄 순환 의존성**: 복잡한 의존성 관계에서 순환 참조 위험

## ✅ **해결 방법들**

### 1. 🏗️ **DI Container 패턴**

```go
// pkg/container/container.go
type Container struct {
    DB *gorm.DB
    
    // Repositories
    UserRepo userRepo.UserRepository
    PostRepo postRepo.PostRepository
    
    // Services  
    UserService userService.UserService
    PostService postService.PostService
    
    // Handlers
    UserHandler *userHandler.UserHandler
    PostHandler *handler.PostHandler
}

func NewContainer() (*Container, error) {
    // 모든 의존성을 한 곳에서 초기화
    // ...
}
```

**장점:**
- ✅ 의존성을 한 곳에서 관리
- ✅ main.go 간소화
- ✅ 테스트 시 모킹 용이

**단점:**
- ⚠️ 여전히 새 도메인 추가 시 Container 수정 필요

### 2. 🏭 **Factory 패턴**

```go
// pkg/container/factory.go
type Factory struct {
    instances map[reflect.Type]interface{}
    database  *database.Database
}

func (f *Factory) Get(serviceType reflect.Type) (interface{}, error) {
    // 리플렉션을 사용한 자동 인스턴스 생성
    // 의존성 그래프 자동 해결
}
```

**장점:**
- ✅ 반자동 의존성 해결
- ✅ 싱글톤 인스턴스 관리
- ✅ 타입 안전성

**단점:**
- ⚠️ 리플렉션 사용으로 성능 오버헤드
- ⚠️ 컴파일 타임 에러 검출 어려움

### 3. 📋 **Registry 패턴**

```go
// pkg/container/registry.go
type ServiceRegistry struct {
    services    map[string]interface{}
    routeSetups []RouteSetupFunc
}

func (r *ServiceRegistry) AutoRegister(services ...interface{}) {
    for _, service := range services {
        serviceName := reflect.TypeOf(service).Elem().Name()
        r.Register(serviceName, service)
    }
}
```

**장점:**
- ✅ 서비스 자동 등록
- ✅ 동적 서비스 관리
- ✅ 플러그인 아키텍처 지원

### 4. 🔌 **Wire (Google) - 컴파일 타임 DI**

```go
//go:build wireinject

package main

import (
    "github.com/google/wire"
)

func InitializeContainer() (*Container, error) {
    wire.Build(
        database.NewDatabase,
        userRepo.NewUserRepository,
        userService.NewUserService,
        userHandler.NewUserHandler,
        // ... 다른 의존성들
        NewContainer,
    )
    return &Container{}, nil
}
```

**장점:**
- ✅ 컴파일 타임 의존성 해결
- ✅ 성능 오버헤드 없음
- ✅ 타입 안전성 보장
- ✅ 순환 의존성 컴파일 타임 검출

**단점:**
- ⚠️ 코드 생성 도구 필요
- ⚠️ 학습 곡선

## 🎯 **권장 접근법**

### 📊 **프로젝트 규모별 권장사항**

| 프로젝트 규모 | 권장 방법 | 이유 |
|--------------|-----------|------|
| **소규모** (1-3 도메인) | DI Container | 간단하고 이해하기 쉬움 |
| **중규모** (4-8 도메인) | Factory + Registry | 자동화와 유연성의 균형 |
| **대규모** (9+ 도메인) | Wire (Google) | 성능과 안정성 우선 |

### 🔄 **단계별 마이그레이션**

#### **1단계: DI Container 도입**
```go
// cmd/server/main.go - 간소화됨
func main() {
    container, err := container.NewContainer()
    if err != nil {
        log.Fatal("Failed to initialize container:", err)
    }
    
    router := gin.Default()
    container.RegisterRoutes(router)
    
    // 서버 시작...
}
```

#### **2단계: 자동 등록 추가**
```go
// 새 도메인 추가 시 자동 인식
func (c *Container) autoRegisterDomains() {
    // 리플렉션 또는 설정 파일 기반 자동 등록
}
```

#### **3단계: Wire 도입 (선택적)**
```go
// 성능이 중요한 경우 Wire로 전환
//go:generate wire
func InitializeApp() (*App, error) {
    // Wire가 자동 생성
}
```

## 🛠️ **실제 구현 예시**

### ✨ **개선된 main.go**

```go
package main

import (
    "log"
    "os"
    "study-go-controller/pkg/container"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // 환경변수 로드
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // ✨ DI Container로 모든 의존성 자동 초기화
    c, err := container.NewContainer()
    if err != nil {
        log.Fatal("Failed to initialize container:", err)
    }

    // 라우터 초기화
    router := gin.Default()

    // ✨ 모든 라우트 자동 등록
    c.RegisterRoutes(router)

    // 헬스체크
    router.GET("/health", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "status":  "ok",
            "message": "Server is running",
        })
    })

    // 서버 시작
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on port %s", port)
    if err := router.Run(":" + port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
```

### 🎯 **새로운 도메인 추가 프로세스**

#### **기존 방식 (❌)**
1. `internal/domain/product/` 패키지 생성
2. **main.go 수정** (필수)
3. 의존성 수동 추가
4. 라우터 수동 등록

#### **개선된 방식 (✅)**
1. `internal/domain/product/` 패키지 생성  
2. **Container에만 추가** (한 곳에서 관리)
3. 자동 의존성 해결
4. 자동 라우터 등록

### 📋 **의존성 추가 체크리스트**

#### **새 도메인 추가 시:**
- [ ] Entity 생성
- [ ] Repository 인터페이스 및 구현체 생성
- [ ] Service 인터페이스 및 구현체 생성
- [ ] Handler 생성
- [ ] DTO 생성
- [ ] Routes 생성
- [ ] **Container에 등록** (DI Container 사용 시)
- [ ] **main.go 수정** (수동 관리 시)

## 🧪 **테스트 개선**

### **기존 테스트 (어려움)**
```go
func TestUserService(t *testing.T) {
    // 모든 의존성을 수동으로 모킹
    mockRepo := &MockUserRepository{}
    service := userService.NewUserService(mockRepo)
    // ...
}
```

### **개선된 테스트 (쉬움)**
```go
func TestUserService(t *testing.T) {
    // Container에서 테스트용 의존성 주입
    container := container.NewTestContainer()
    container.UserRepo = &MockUserRepository{}
    
    service := container.UserService
    // ...
}
```

## 🎯 **결론**

### **현재 상황 개선 우선순위:**

1. **🥇 1순위**: DI Container 도입
   - 즉시 main.go 복잡도 감소
   - 테스트 용이성 향상

2. **🥈 2순위**: Factory 패턴 추가
   - 자동 의존성 해결
   - 더 적은 수동 작업

3. **🥉 3순위**: Wire 도입 고려
   - 성능 최적화 필요 시
   - 대규모 프로젝트 전환 시

### **핵심 원칙:**

- ✅ **Single Responsibility**: main.go는 애플리케이션 시작만 담당
- ✅ **Open/Closed**: 새 기능 추가 시 기존 코드 최소 변경
- ✅ **Dependency Inversion**: 구체 타입이 아닌 인터페이스에 의존
- ✅ **Testability**: 쉬운 테스트를 위한 의존성 주입

이렇게 개선하면 **새로운 도메인 추가 시 main.go를 거의 수정하지 않아도** 됩니다! 🚀 