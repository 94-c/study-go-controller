# 📦 패키지 의존성 가이드

이 문서는 프로젝트 내 각 패키지가 어떻게 상호작용하는지 상세히 설명합니다.

## 🔗 의존성 계층 구조

```
                cmd/server/main.go
                       |
            ┌─────────────────────────┐
            |                         |
      pkg/database              internal/domain/
            |                         |
    ┌───────┴───────┐         ┌───────┴───────┐
    |               |         |               |
  user/entity   post/entity   user/         post/
                              domain        domain
```

## 📋 상세 의존성 매트릭스

### 🎯 **cmd/server/main.go**
```go
직접 의존성:
├── pkg/database               (데이터베이스 초기화)
├── internal/domain/user/      (User 도메인 모든 계층)
├── internal/domain/post/      (Post 도메인 모든 계층)
├── github.com/gin-gonic/gin   (HTTP 프레임워크)
└── github.com/joho/godotenv   (환경변수)

역할:
- 애플리케이션 부트스트랩
- 의존성 주입 (DI Container 역할)
- 라우터 설정 및 서버 시작
```

### 🔷 **pkg/database/**
```go
내부 구조:
└── database.go

의존성:
├── internal/domain/user/entity     (User 모델)
├── internal/domain/post/entity     (Post 모델)
├── gorm.io/gorm                   (ORM)
├── gorm.io/driver/mysql           (MySQL 드라이버)
└── os                             (환경변수)

사용처:
└── cmd/server/main.go

제공 기능:
- NewDatabase() *Database
- AutoMigrate() error
- Close() error
- getEnv(key, fallback) string
```

### 🔷 **pkg/response/**
```go
내부 구조:
└── response.go

의존성:
├── net/http                       (HTTP 상태 코드)
└── github.com/gin-gonic/gin       (Gin Context)

사용처:
├── internal/domain/user/handler/
└── internal/domain/post/handler/

제공 기능:
- SuccessResponse(c, statusCode, message, data)
- ErrorResponse(c, statusCode, message)
- ValidationErrorResponse(c, err)

타입:
- APIResponse struct
```

### 🔷 **pkg/models/**
```go
내부 구조:
└── common.go

의존성:
├── time                           (시간 타입)
└── gorm.io/gorm                   (GORM 타입)

사용처:
├── internal/domain/user/entity/   (BaseModel 임베딩)
├── internal/domain/post/entity/   (BaseModel 임베딩)
└── internal/domain/*/handler/     (페이지네이션)

제공 기능:
- BaseModel (ID, CreatedAt, UpdatedAt, DeletedAt)
- PaginationRequest
- PaginationResponse
- SortRequest
```

### 🔷 **pkg/enums/**
```go
내부 구조:
└── common_enums.go

의존성:
└── 없음 (순수 Go 타입)

사용처:
├── internal/domain/user/service/
├── internal/domain/post/service/
├── internal/domain/*/handler/     (정렬, 우선순위)
└── internal/domain/*/dto/

제공 열거형:
- SortDirection (asc, desc)
- CommonStatus (active, inactive, deleted)
- Priority (low, medium, high, urgent)

메서드:
- String() string
- IsValid() bool
```

### 🔷 **pkg/utils/**
```go
내부 구조:
└── string_utils.go

의존성:
├── crypto/rand                    (랜덤 생성)
├── regexp                         (정규표현식)
├── strings                        (문자열 처리)
└── unicode                        (유니코드 처리)

사용처:
├── internal/domain/user/service/  (패스워드 검증, 슬러그 생성)
├── internal/domain/post/service/  (제목 슬러그 생성)
└── pkg/validator/                 (검증 로직)

제공 기능:
- IsValidEmail(email) bool
- IsValidPassword(password) bool
- GenerateRandomString(length) string
- SlugifyString(input) string
- TruncateString(input, maxLength) string
- ContainsString(slice, item) bool
- RemoveString(slice, item) []string
```

### 🔷 **pkg/validator/**
```go
내부 구조:
└── custom_validators.go

의존성:
├── regexp                         (정규표현식)
└── unicode                        (유니코드 처리)

사용처:
├── internal/domain/user/service/  (사용자 검증)
├── internal/domain/post/service/  (포스트 검증)
└── internal/domain/*/handler/     (요청 검증)

제공 기능:
- ValidateUsername(username) error
- ValidatePassword(password) error
- ValidateEmail(email) error
- NewValidationError(message) *ValidationError

타입:
- ValidationError struct
```

## 🏗️ **internal/domain/user/**

### 📁 **entity/user.go**
```go
의존성:
├── time                           (시간 타입)
└── gorm.io/gorm                   (GORM 타입)

사용처:
├── repository/user_repository.go  (CRUD 작업)
├── service/user_service.go        (비즈니스 로직)
├── dto/user_dto.go               (DTO 변환)
├── pkg/database/database.go       (마이그레이션)
└── internal/domain/post/entity/   (Author 관계)

제공:
- User struct (ID, Username, Email, Password, Name, 타임스탬프)
- TableName() string
```

### 📁 **enums/user_enums.go**
```go
의존성:
└── 없음 (순수 Go 타입)

사용처:
├── entity/user.go                 (상태 필드)
├── service/user_service.go        (비즈니스 규칙)
├── dto/user_dto.go               (API 응답)
└── handler/user_handler.go        (요청 처리)

제공 열거형:
- UserStatus (active, inactive, suspended, deleted)
- UserRole (admin, moderator, user, guest)

메서드:
- String() string
- IsValid() bool
- GetAllUserStatuses() []UserStatus
```

### 📁 **dto/user_dto.go**
```go
의존성:
├── time                           (시간 타입)
└── internal/domain/user/entity    (User 엔티티)

사용처:
├── handler/user_handler.go        (HTTP 요청/응답)
└── internal/domain/post/dto/      (Author 정보)

제공:
- CreateUserRequest struct
- UpdateUserRequest struct
- UserResponse struct
- LoginRequest struct
- ChangePasswordRequest struct
- ToUserResponse(*entity.User) *UserResponse
- ToUserResponseList([]*entity.User) []*UserResponse
```

### 📁 **repository/user_repository.go**
```go
의존성:
├── internal/domain/user/entity    (User 엔티티)
└── gorm.io/gorm                   (ORM)

사용처:
└── service/user_service.go        (데이터 액세스)

제공:
- UserRepository interface
- userRepository struct
- NewUserRepository(db) UserRepository

메서드:
- Create(user) error
- GetByID(id) (*User, error)
- GetByEmail(email) (*User, error)
- GetByUsername(username) (*User, error)
- Update(user) error
- Delete(id) error
- GetAll() ([]*User, error)
```

### 📁 **service/user_service.go**
```go
의존성:
├── errors                         (에러 생성)
├── internal/domain/user/entity    (User 엔티티)
├── internal/domain/user/repository (데이터 액세스)
└── golang.org/x/crypto/bcrypt     (암호화)

사용처:
└── handler/user_handler.go        (HTTP 핸들러)

제공:
- UserService interface
- userService struct
- NewUserService(userRepo) UserService

메서드:
- CreateUser(username, email, password, name) (*User, error)
- GetUserByID(id) (*User, error)
- GetUserByEmail(email) (*User, error)
- GetUserByUsername(username) (*User, error)
- UpdateUser(id, username, email, name) (*User, error)
- DeleteUser(id) error
- GetAllUsers() ([]*User, error)
- ValidatePassword(password, hashedPassword) bool
```

### 📁 **handler/user_handler.go**
```go
의존성:
├── net/http                       (HTTP 상태 코드)
├── strconv                        (문자열 변환)
├── internal/domain/user/dto       (DTO)
├── internal/domain/user/service   (비즈니스 로직)
├── pkg/response                   (표준 응답)
└── github.com/gin-gonic/gin       (HTTP 프레임워크)

사용처:
└── routes/user_routes.go          (라우팅)

제공:
- UserHandler struct
- NewUserHandler(userService) *UserHandler

메서드:
- CreateUser(c *gin.Context)
- GetUser(c *gin.Context)
- GetAllUsers(c *gin.Context)
- UpdateUser(c *gin.Context)
- DeleteUser(c *gin.Context)
```

### 📁 **routes/user_routes.go**
```go
의존성:
├── internal/domain/user/handler   (HTTP 핸들러)
└── github.com/gin-gonic/gin       (라우팅)

사용처:
└── cmd/server/main.go             (라우터 등록)

제공:
- SetupUserRoutes(router, userHandler)

라우트:
- POST   /users
- GET    /users
- GET    /users/:id
- PUT    /users/:id
- DELETE /users/:id
```

## 🏗️ **internal/domain/post/**

### 📁 **entity/post.go**
```go
의존성:
├── time                           (시간 타입)
├── gorm.io/gorm                   (GORM 타입)
└── internal/domain/user/entity    (Author 관계)

사용처:
├── repository/post_repository.go  (CRUD 작업)
├── service/post_service.go        (비즈니스 로직)
├── dto/post_dto.go               (DTO 변환)
└── pkg/database/database.go       (마이그레이션)

제공:
- Post struct (ID, Title, Content, AuthorID, Author, 타임스탬프)
- TableName() string
```

### 📁 **enums/post_enums.go**
```go
의존성:
└── 없음 (순수 Go 타입)

사용처:
├── entity/post.go                 (상태 필드)
├── service/post_service.go        (비즈니스 규칙)
├── dto/post_dto.go               (API 응답)
└── handler/post_handler.go        (요청 처리)

제공 열거형:
- PostStatus (draft, published, archived, deleted)
- PostCategory (tech, lifestyle, news, review)

메서드:
- String() string
- IsValid() bool
- GetAllPostStatuses() []PostStatus
```

### 📁 **dto/post_dto.go**
```go
의존성:
├── time                           (시간 타입)
├── internal/domain/post/entity    (Post 엔티티)
└── internal/domain/user/dto       (Author 정보)

사용처:
└── handler/post_handler.go        (HTTP 요청/응답)

제공:
- CreatePostRequest struct
- UpdatePostRequest struct
- PostResponse struct
- PostListResponse struct
- ToPostResponse(*entity.Post) *PostResponse
- ToPostListResponse(*entity.Post) *PostListResponse
- ToPostResponseList([]*entity.Post) []*PostResponse
- ToPostListResponseList([]*entity.Post) []*PostListResponse
```

## 💡 **의존성 규칙**

### ✅ **허용되는 의존성 방향**
```
cmd → pkg → external libraries
cmd → internal → pkg → external libraries
internal/domain/A → internal/domain/B (DTO만)
internal/domain/*/layer → internal/domain/*/lower-layer
```

### ❌ **금지되는 의존성 방향**
```
pkg → internal (pkg는 internal을 참조할 수 없음)
internal/domain/A/service → internal/domain/B/service (서비스 간 직접 호출 금지)
lower-layer → upper-layer (계층 역전 금지)
```

### 🎯 **계층별 의존성 규칙**
```
Handler → Service + DTO + pkg/response
Service → Repository + Entity + pkg/utils + pkg/validator
Repository → Entity + GORM
Entity → pkg/models (BaseModel)
DTO → Entity + other domain DTO (read-only)
```

## 🔄 **도메인 간 통신 패턴**

### ✅ **올바른 패턴**
```go
// Post DTO에서 User DTO 참조 (읽기 전용)
type PostResponse struct {
    Author *userDto.UserResponse `json:"author,omitempty"`
}

// 서비스 계층에서 다른 도메인 호출 (이벤트 기반 권장)
func (s *postService) CreatePost(authorID uint) {
    // Direct call (간단한 경우)
    user, err := s.userService.GetUserByID(authorID)
    
    // Event-driven (복잡한 경우 권장)
    s.eventBus.Publish("post.created", event)
}
```

### ❌ **피해야 할 패턴**
```go
// Repository에서 다른 도메인 직접 접근
func (r *postRepository) GetWithAuthor() {
    // ❌ Bad: Repository가 다른 도메인 접근
    userRepo.GetByID(post.AuthorID)
}

// Entity에서 다른 도메인 비즈니스 로직 호출
func (p *Post) ValidateAuthor() {
    // ❌ Bad: Entity에 비즈니스 로직
    userService.IsActive(p.AuthorID)
}
```

## 📊 **패키지 크기 및 복잡도 가이드라인**

### 🎯 **권장 크기**
```
Entity:     50-150 lines    (도메인 모델만)
Repository: 100-300 lines   (CRUD + 복잡 쿼리)
Service:    200-500 lines   (비즈니스 로직)
Handler:    150-400 lines   (HTTP 처리)
DTO:        100-250 lines   (변환 함수 포함)
```

### 🔍 **복잡도 신호**
```
❌ 파일이 500줄 초과 → 분리 고려
❌ 의존성이 5개 초과 → 추상화 고려
❌ 순환 의존성 발생 → 아키텍처 재설계
❌ 테스트 어려움 → 의존성 주입 개선
```

이 의존성 가이드를 따르면 유지보수가 쉽고 확장 가능한 코드를 작성할 수 있습니다! 🚀 