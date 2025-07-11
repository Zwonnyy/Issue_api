# 🛠️ 이슈 관리 API (Issue Management API)

Go 언어와 Gin 프레임워크를 활용해 구현한 간단한 이슈 관리 RESTful API입니다.  
담당자 지정, 상태 변경, 이슈 수정 등의 기능을 포함하며, 메모리 기반으로 동작합니다.

---

## 📌 프로젝트 개요

- **언어**: Go 1.20+
- **프레임워크**: Gin
- **실행 포트**: `8080`
- **데이터 저장소**: 메모리 (슬라이스 기반)
- **목적**: RESTful API 설계 및 비즈니스 로직 구현 과제

---

## 🧱 프로젝트 구조

```
Issue-api/
├── main.go
├── go.mod
├── go.sum
├── README.md
├── data/
│ └── seed.go
├── models/
│ └── models.go
├── controllers/
│ └── issue_controller.go
├── router/
│ └── router.go
```

seed.go: 초기 사용자 및 이슈 데이터

models.go: User, Issue 구조체 정의

issue_controller.go: API 핸들러 함수

router.go : 라우팅 설정

---

## 🚀 실행 방법

### 1. Go 모듈 초기화 및 의존성 설치

    bash
    go mod tidy


## 🚀 서버 실행
    go run main.go

    기본 포트는 8080입니다.
    서버가 정상적으로 실행되면 Listening and serving HTTP on :8080 메시지가 출력됩니다.

🧪 API 테스트 방법
    Postman, curl, 또는 브라우저를 통해 아래 API들을 테스트할 수 있습니다.

📌 1. 이슈 생성
    POST /issue

    담당자(userId)가 있으면 상태는 IN_PROGRESS, 없으면 PENDING으로 생성됨

        예시 요청:
            {
            "title": "버그 수정 필요",
            "description": "로그인 오류 발생",
            "userId": 1
            }
        Bash
            curl -X POST http://localhost:8080/issue \
            -H "Content-Type: application/json" \
            -d '{"title":"버그 수정 필요", "description":"로그인 오류", "userId":1}'
📌 2. 이슈 목록 조회
        -GET /issues

        -쿼리 파라미터 status로 필터링 가능: PENDING, IN_PROGRESS, COMPLETED, CANCELLED

        예시:
            bash
    
            curl http://localhost:8080/issues
            curl http://localhost:8080/issues?status=IN_PROGRESS
📌 3. 이슈 상세 조회
        -GET /issue/:id

        예시:
            bash
            curl http://localhost:8080/issue/1
📌 4. 이슈 수정
        -PATCH /issue/:id

        일부 필드만 변경 가능

        COMPLETED 또는 CANCELLED 상태의 이슈는 수정이 불가 합니다.

        예시:
            {
            "title": "로그인 버그 수정",
            "status": "IN_PROGRESS",
            "userId": 2
            }

            bash
                curl -X PATCH http://localhost:8080/issue/1 \
                -H "Content-Type: application/json" \
                -d '{"title":"로그인 버그 수정", "status":"IN_PROGRESS", "userId":2}'

❗ 에러 응답 형식
        모든 잘못된 요청은 다음 형식으로 JSON 에러 메시지를 반환합니다.

        json
        {
        "error": "에러 메시지",
        "code": 400
        }// 400: 잘못된 요청 , 404: 리소스를 찾을 수 없음, 405: 잘못된 메서드 요청


🧼 주의사항
    이 프로젝트는 메모리 기반 저장이므로 서버 재시작 시 모든 이슈 데이터가 초기화됩니다.

    실 서비스 환경에선 DB와 영속 저장소를 도입해야 합니다.

    go run main.go 실행 전 반드시 go mod tidy 로 의존성을 설치하세요.



📚 기술 스택
        Go 1.20+

        Gin Web Framework

        RESTful API 설계

        JSON 기반 데이터 처리

        단순한 슬라이스 기반 상태 관리
