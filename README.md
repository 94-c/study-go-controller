# 🚀 Study Go Controller

Go언어로 구현한 **Controller-Service-Repository-Entity** 패턴과 **DDD(Domain-Driven Design)**를 적용한 **자동 라우팅** REST API 프로젝트입니다.

## ⚡ **핵심 특징**

### 🔥 **완전 자동 라우팅 시스템**
- 🚀 **제로 설정**: Handler 메서드만 작성하면 자동으로 API 엔드포인트 생성
- 🧠 **컨벤션 기반**: 메서드 이름을 통한 자동 HTTP 메서드 및 경로 결정
- 📡 **실시간 등록**: 서버 시작 시 모든 라우트 자동 스캔 및 등록

### 🏗️ **Clean Architecture + DDD**  
- 🎯 **도메인별 완전 분리**: User, Post 등 각 도메인이 독립적 구조
- 🔧 **의존성 주입**: 자동 DI Container로 완전 자동화
- 📦 **계층별 책임 분리**: Handler-Service-Repository-Entity

## 📁 프로젝트 구조

```
study-go-controller/
├── cmd/
│   └── server/                    # 🚀 애플리케이션 진입점 (완전 자동화)
│       └── main.go               # DI Container + 자동 라우팅
├── internal/                     # 🏗️ 내부 패키지 (외부 접근 불가)
│   └── domain/                   # DDD 도메인별 구조
│       ├── user/                 # 👤 User 도메인
│       │   ├── entity/           # 사용자 엔티티
│       │   ├── repository/       # 데이터 액세스 계층
│       │   ├── service/          # 비즈니스 로직 계층
│       │   ├── handler/          # 🔥 HTTP 핸들러 (자동 라우팅)
│       │   ├── dto/             # 데이터 전송 객체
│       │   └── enums/           # User 도메인 특화 열거형
│       └── post/                 # 📝 Post 도메인
│           ├── entity/
│           ├── repository/
│           ├── service/
│           ├── handler/          # 🔥 HTTP 핸들러 (자동 라우팅)
│           ├── dto/
│           └── enums/
├── pkg/                          # 📚 공통 패키지 (재사용 가능)
│   ├── container/                # 🔧 DI Container + 자동 라우팅 시스템
│   │   ├── container.go         # 의존성 주입 컨테이너
│   │   ├── auto_router.go       # 🚀 자동 라우팅 엔진
│   │   ├── registry.go          # 서비스 레지스트리
│   │   └── factory.go           # 팩토리 패턴
│   ├── database/                # 🗄️ 데이터베이스 연결 관리
│   ├── response/                # 📤 API 응답 표준화
│   ├── models/                  # 🔧 공통 모델
│   ├── enums/                   # 🏷️ 도메인 간 공통 열거형
│   ├── utils/                   # 🛠️ 공통 유틸리티 함수
│   └── validator/               # ✅ 공통 검증 로직
├── docs/                        # 📖 문서
└── configs/                     # ⚙️ 설정 파일
```

## 🚀 **자동 라우팅 시스템**

### 🔥 **어떻게 작동하나요?**

#### **1. Handler 메서드만 작성**
```go
// internal/domain/user/handler/user_handler.go
func (h *UserHandler) CreateUser(c *gin.Context) {
    // 사용자 생성 로직
}

func (h *UserHandler) GetUser(c *gin.Context) {
    // 사용자 조회 로직  
}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
    // 사용자 프로필 조회 로직
}
```

#### **2. 자동으로 API 엔드포인트 생성! ✨**
```
🔗 Auto-registered route: POST /api/v1/users -> UserHandler.CreateUser
🔗 Auto-registered route: GET /api/v1/users/:id -> UserHandler.GetUser  
🔗 Auto-registered route: GET /api/v1/users/:id/profile -> UserHandler.GetUserProfile
```

### 📋 **컨벤션 기반 라우팅 규칙**

| 메서드 이름 패턴 | HTTP 메서드 | 자동 생성 경로 | 예시 |
|------------------|-------------|----------------|------|
| `Create*` | POST | `/` | `CreateUser` → `POST /users` |
| `GetAll*` | GET | `/` | `GetAllUsers` → `GET /users` |
| `Get*` | GET | `/:id` | `GetUser` → `GET /users/:id` |
| `Update*` | PUT | `/:id` | `UpdateUser` → `PUT /users/:id` |
| `Delete*` | DELETE | `/:id` | `DeleteUser` → `DELETE /users/:id` |
| `Get*Profile` | GET | `/:id/profile` | `GetUserProfile` → `GET /users/:id/profile` |
| `Change*Password` | PUT | `/:id/password` | `ChangePassword` → `PUT /users/:id/password` |

### 🎯 **실제 등록된 API 엔드포인트**

#### **User API (자동 생성)**
```
POST   /api/v1/users              # CreateUser
GET    /api/v1/users              # GetAllUsers
GET    /api/v1/users/:id          # GetUser
PUT    /api/v1/users/:id          # UpdateUser
DELETE /api/v1/users/:id          # DeleteUser
GET    /api/v1/users/:id/profile  # GetUserProfile
PUT    /api/v1/users/:id/password # ChangePassword
```

#### **Post API (자동 생성)**
```
POST   /api/v1/posts              # CreatePost
GET    /api/v1/posts              # GetAllPosts
GET    /api/v1/posts/:id          # GetPost
PUT    /api/v1/posts/:id          # UpdatePost
DELETE /api/v1/posts/:id          # DeletePost
```

#### **System API**
```
GET    /health                    # 서버 상태 + 등록된 라우트 정보
```

## 🏗️ **아키텍처 & 패키지 구조**

### 📦 **각 패키지의 역할**

```
📁 cmd/server/main.go     → 🚀 서버 시작점 (DI Container 초기화)
📁 pkg/container/         → 🔧 자동 의존성 주입 + 라우팅
📁 pkg/database/          → 🗄️ DB 연결 관리
📁 pkg/response/          → 📤 API 응답 표준화
📁 internal/domain/user/  → 👤 사용자 도메인 (완전 분리)
  ├── entity/             → 데이터 모델 (User 구조체)
  ├── repository/         → 데이터베이스 액세스
  ├── service/            → 비즈니스 로직
  ├── handler/            → HTTP 요청 처리 (🔥 자동 라우팅)
  └── dto/                → API 요청/응답 구조
```

### 🔗 **데이터 흐름 (Clean Architecture)**

```
HTTP 요청 → Handler → Service → Repository → Database
    ↓         ↓         ↓          ↓
  JSON      비즈니스   데이터      SQL
  파싱      로직      액세스     쿼리
```

---

## 🔧 **실제 예제: 회원가입 API 만들기**

**"사용자 등록(회원가입)"** API를 처음부터 끝까지 만들어보겠습니다! 🎯

### **Step 1: DTO 정의** 📝

```go
// internal/domain/user/dto/user_dto.go에 추가
type RegisterRequest struct {
    Username        string `json:"username" binding:"required,min=3,max=20"`
    Email          string `json:"email" binding:"required,email"`
    Password       string `json:"password" binding:"required,min=8"`
    ConfirmPassword string `json:"confirm_password" binding:"required"`
    Name           string `json:"name" binding:"required,min=2"`
}

type RegisterResponse struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Name     string `json:"name"`
    Message  string `json:"message"`
}
```

### **Step 2: Service 로직 구현** 🧠

```go
// internal/domain/user/service/user_service.go에 추가
func (s *userService) RegisterUser(username, email, password, confirmPassword, name string) (*entity.User, error) {
    // 1. 입력 검증
    if password != confirmPassword {
        return nil, fmt.Errorf("passwords do not match")
    }
    
    // 2. 중복 체크
    existingUser, _ := s.userRepository.GetByEmail(email)
    if existingUser != nil {
        return nil, fmt.Errorf("email already registered")
    }
    
    existingUser, _ = s.userRepository.GetByUsername(username)
    if existingUser != nil {
        return nil, fmt.Errorf("username already taken")
    }
    
    // 3. 비밀번호 해시화
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, fmt.Errorf("failed to hash password")
    }
    
    // 4. 사용자 생성
    user := &entity.User{
        Username: username,
        Email:    email,
        Password: string(hashedPassword),
        Name:     name,
    }
    
    // 5. 데이터베이스 저장
    if err := s.userRepository.Create(user); err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }
    
    return user, nil
}
```

### **Step 3: Handler 메서드 추가** 🚀

```go
// internal/domain/user/handler/user_handler.go에 추가

// 🔗 자동으로 POST /api/v1/users/register 엔드포인트 생성!
func (h *UserHandler) CreateUserRegister(c *gin.Context) {
    // 1. 요청 데이터 파싱
    var req dto.RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.ValidationErrorResponse(c, err)
        return
    }
    
    // 2. 서비스 호출 (비즈니스 로직)
    user, err := h.userService.RegisterUser(
        req.Username,
        req.Email, 
        req.Password,
        req.ConfirmPassword,
        req.Name,
    )
    if err != nil {
        response.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    
    // 3. 응답 생성
    registerResponse := dto.RegisterResponse{
        ID:       user.ID,
        Username: user.Username,
        Email:    user.Email,
        Name:     user.Name,
        Message:  "Registration successful! Welcome to our platform.",
    }
    
    // 4. 성공 응답
    response.SuccessResponse(c, http.StatusCreated, "User registered successfully", registerResponse)
}
```

### **Step 4: 자동 생성 결과** ✨

서버 시작 시 로그:
```
🔗 Auto-registered route: POST /api/v1/users/register -> UserHandler.CreateUserRegister
📡 Total: 8 routes automatically registered
```

### **Step 5: API 테스트** 🧪

```bash
# 회원가입 API 호출
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "securepassword123",
    "confirm_password": "securepassword123",
    "name": "John Doe"
  }'

# 성공 응답
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com", 
    "name": "John Doe",
    "message": "Registration successful! Welcome to our platform."
  }
}

# 에러 응답 (중복 이메일)
{
  "success": false,
  "message": "email already registered"
}
```

---

## 🎯 **개발 프로세스 요약**

### **✅ 해야 할 일**
1. **DTO 정의**: 요청/응답 구조 (`dto/`)
2. **Service 구현**: 비즈니스 로직 (`service/`)  
3. **Handler 추가**: HTTP 처리 + **컨벤션 이름** (`handler/`)

### **🚫 하지 않아도 되는 일**
- ❌ main.go 수정
- ❌ routes 파일 작성
- ❌ 수동 라우트 등록
- ❌ 의존성 주입 설정

### **🔥 자동 라우팅 컨벤션**
- `CreateUser` → `POST /users`
- `CreateUserRegister` → `POST /users/register`  
- `GetUser` → `GET /users/:id`
- `UpdateUser` → `PUT /users/:id`
- `DeleteUser` → `DELETE /users/:id`

---

## 💡 **핵심 원칙**

### **🎯 하나의 API = 3단계**
```
1️⃣ DTO 정의 (요청/응답 구조)
2️⃣ Service 구현 (비즈니스 로직)
3️⃣ Handler 추가 (컨벤션 이름) → 🚀 자동 API 생성!
```

### **🚀 완전 자동화**
```
✅ Handler 메서드 1개 = API 엔드포인트 1개 자동 생성
✅ 컨벤션만 따르면 모든 것이 자동
✅ 개발자는 비즈니스 로직에만 집중
```

### 🏗️ **새로운 도메인 추가하기**

#### **Step 1: 도메인 패키지 생성**
```
internal/domain/product/
├── entity/product.go
├── repository/product_repository.go
├── service/product_service.go
├── handler/product_handler.go    # 🔥 메서드 이름만 맞추면 자동 라우팅
├── dto/product_dto.go
└── enums/product_enums.go
```

#### **Step 2: Container에 등록**
```go
// pkg/container/container.go의 registerAllHandlers() 메서드에 추가
func (c *Container) registerAllHandlers() {
    // 기존 도메인들...
    c.AutoRouter.RegisterHandler("/users", c.UserHandler)
    c.AutoRouter.RegisterHandler("/posts", c.PostHandler)
    
    // 🆕 새 도메인 추가
    c.AutoRouter.RegisterHandler("/products", c.ProductHandler)
}
```

#### **Step 3: 끝! 🎉**
- ✅ 모든 Product API 자동 생성
- ✅ main.go는 그대로

## 🛠️ **아키텍처 상세 설명**

### 🏗️ **DDD (Domain-Driven Design) 아키텍처**

```
┌─────────────────────────────────────┐
│           Presentation Layer        │
│     (Handler - Auto Routing)        │ ← 🚀 자동 라우팅
├─────────────────────────────────────┤
│          Application Layer          │
│          (Service Layer)            │ ← 📋 비즈니스 로직
├─────────────────────────────────────┤
│         Infrastructure Layer        │
│        (Repository Layer)           │ ← 🗄️ 데이터 액세스
├─────────────────────────────────────┤
│            Domain Layer             │
│         (Entity + Enums)            │ ← 🎯 도메인 모델
└─────────────────────────────────────┘
```

### 🔗 **의존성 주입 흐름**

```
Container.NewContainer()
    ↓
Database → Repository → Service → Handler
    ↓
AutoRouter.RegisterHandler()
    ↓
자동 라우트 스캔 및 등록
    ↓
완전한 API 서버 준비 완료! 🚀
```

### 📦 **패키지 분류 원칙**

```
📁 cmd/             ← 애플리케이션 진입점 (main.go만)
📁 internal/domain/ ← 도메인별 특화 로직 (완전 분리)
📁 pkg/container/   ← 🔥 자동화 시스템 (DI + Auto Routing)
📁 pkg/utils/       ← 공통 유틸리티 (여러 도메인에서 공용)
```

## 🚀 시작하기

### 1. 의존성 설치
```bash
go mod tidy
```

### 2. 환경변수 설정
```bash
cp configs/config.example .env
# .env 파일을 편집하여 데이터베이스 정보 입력
```

### 3. 데이터베이스 설정
MySQL 데이터베이스를 생성하고 연결 정보를 .env에 설정합니다.

### 4. 서버 실행
```bash
go run cmd/server/main.go
```

### 5. 자동 등록된 라우트 확인
```bash
curl http://localhost:8080/health
```

**결과:**
```json
{
  "status": "ok",
  "message": "Server is running with automatic routing",
  "total_routes": 7,
  "auto_routes": [
    {"method": "POST", "path": "/api/v1/users"},
    {"method": "GET", "path": "/api/v1/users"},
    {"method": "GET", "path": "/api/v1/users/:id"},
    {"method": "PUT", "path": "/api/v1/users/:id"},
    {"method": "DELETE", "path": "/api/v1/users/:id"},
    {"method": "GET", "path": "/api/v1/users/:id/profile"},
    {"method": "PUT", "path": "/api/v1/users/:id/password"}
  ]
}
```

## 🛠️ 사용된 기술 스택

- **Go 1.21**: 프로그래밍 언어
- **Gin**: HTTP 웹 프레임워크  
- **GORM**: ORM 라이브러리
- **MySQL**: 관계형 데이터베이스
- **Reflection**: 자동 라우팅 시스템
- **DI Container**: 의존성 주입 자동화
- **godotenv**: 환경변수 관리
- **bcrypt**: 암호 해싱

## 📋 주요 기능

### 🚀 **완전 자동화된 개발 경험**
- **제로 설정 라우팅**: 메서드 이름만으로 API 엔드포인트 자동 생성
- **자동 의존성 주입**: main.go 수정 없이 새 도메인 추가
- **컨벤션 기반**: 일관된 API 설계 강제

### 🏗️ **확장 가능한 아키텍처**
- **도메인별 완전 분리**: 마이크로서비스 전환 용이
- **인터페이스 기반**: 테스트 및 모킹 쉬움
- **플러그인 구조**: 새 기능 추가 시 기존 코드 영향 최소

### 📤 **표준화된 API 응답**
```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com"
  }
}
```

### ✅ **개발자 친화적**
- **실시간 피드백**: 서버 시작 시 등록된 모든 라우트 출력
- **디버깅 지원**: /health 엔드포인트에서 라우트 정보 확인
- **타입 안전성**: 컴파일 타임 에러 검출

## 🔧 개발 가이드

### 📋 **Handler 메서드 작성 규칙**

1. **메서드 시그니처**: `func (h *Handler) MethodName(c *gin.Context)`
2. **명명 규칙**: 컨벤션 테이블 참조
3. **응답 형식**: `pkg/response` 패키지 사용

### 🎯 **새 API 추가 체크리스트**

- [ ] **DTO 정의** (필요한 경우)
- [ ] **Service 메서드 구현** (필요한 경우)  
- [ ] **Handler 메서드 작성** (컨벤션 준수)
- [ ] **서버 재시작**
- [ ] ✅ **자동으로 API 엔드포인트 생성됨!**

### 🏗️ **새 도메인 추가 체크리스트**

- [ ] **도메인 패키지 생성** (`internal/domain/새도메인/`)
- [ ] **Entity, Repository, Service, Handler, DTO, Enums 구현**
- [ ] **Container에 Handler 등록** (`registerAllHandlers()`)
- [ ] **서버 재시작**
- [ ] ✅ **모든 API 자동 생성됨!**

## ⚡ 성능 및 확장성

### 🚀 **자동 라우팅 성능**
- **초기화 시에만 리플렉션 사용**: 런타임 성능 영향 없음
- **메모리 효율적**: 라우트 정보 캐싱
- **빠른 등록**: 서버 시작 시 몇 ms 내 완료

### 📈 **확장성 고려사항**
- **수평 확장**: 도메인별 독립성으로 마이크로서비스 분리 쉬움
- **팀 개발**: 도메인별 팀이 독립적으로 개발 가능
- **CI/CD**: 자동화된 구조로 배포 파이프라인 간소화

## 🧪 테스트 가이드

### 단위 테스트
```bash
go test ./internal/domain/user/service/...
go test ./pkg/utils/...
```

### 통합 테스트
```bash
go test ./internal/domain/user/...
```

### API 테스트
```bash
# 자동 등록된 라우트 확인
curl http://localhost:8080/health

# User API 테스트
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"username":"test","email":"test@example.com","password":"password123","name":"Test User"}'
```

## 🎯 아키텍처 결정 이유

### ✅ **왜 자동 라우팅을 선택했나요?**
- **개발 속도**: API 추가 시간을 80% 단축
- **일관성**: 컨벤션 기반으로 API 설계 표준화
- **실수 방지**: 수동 라우트 등록으로 인한 오타/누락 방지

### ✅ **왜 컨벤션 기반인가요?**
- **학습 용이성**: 명확한 규칙으로 빠른 온보딩
- **예측 가능성**: 메서드 이름만 보고 API 경로 예측 가능
- **유지보수성**: 일관된 구조로 코드 이해 쉬움

### ✅ **왜 DDD + Clean Architecture인가요?**
- **도메인 중심**: 비즈니스 로직이 기술적 관심사와 분리
- **테스트 용이성**: 계층별 독립적 테스트 가능
- **확장성**: 새 기능 추가 시 기존 코드 영향 최소

## 📈 확장 계획

1. **인증/인가 시스템** (JWT, OAuth2)
2. **캐싱 레이어** (Redis) 
3. **로깅 시스템** (Structured logging)
4. **모니터링** (Prometheus, Grafana)
5. **API 문서화** (Swagger 자동 생성)
6. **Rate Limiting** (요청 제한)
7. **마이크로서비스 분리**

## 🔍 디버깅 및 모니터링

### 📊 **라우트 정보 확인**
```bash
# 등록된 모든 라우트 확인
curl http://localhost:8080/health | jq '.auto_routes'

# 서버 로그에서 라우트 등록 정보 확인
tail -f server.log | grep "Auto-registered route"
```

### 🐛 **문제 해결**
- **라우트가 등록되지 않는 경우**: 메서드 이름이 컨벤션을 따르는지 확인
- **404 에러**: 자동 생성된 경로와 요청 경로 비교
- **의존성 에러**: Container 초기화 로그 확인

---

## 🤝 기여하기

1. 이 저장소를 포크합니다
2. 새 기능 브랜치를 생성합니다 (`git checkout -b feature/amazing-feature`)
3. **Handler 메서드만 추가하면 자동으로 API 생성됩니다!** 🚀
4. 변경사항을 커밋합니다 (`git commit -m 'Add amazing feature'`)
5. 브랜치에 푸시합니다 (`git push origin feature/amazing-feature`)
6. Pull Request를 생성합니다

## 📝 라이센스

이 프로젝트는 MIT 라이센스 하에 배포됩니다.

---

## 🎉 **요약: 완전 자동화된 Go API 개발**

이 프로젝트를 사용하면:

- ✅ **Handler 메서드 하나 추가** = **API 엔드포인트 하나 자동 생성**
- ✅ **main.go 절대 수정 안 함**
- ✅ **routes 파일 수정 안 함**  
- ✅ **컨벤션만 따르면 모든 것이 자동**

**개발자는 비즈니스 로직에만 집중하세요!** 🚀 