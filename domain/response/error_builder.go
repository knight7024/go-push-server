package response

import "net/http"

var (
	UndefinedServerError = ErrorBuilder.New(http.StatusInternalServerError).
				ErrorCode("S001").
				Message("알 수 없는 서버 에러가 발생하였습니다.").
				Build()
	InternalFirebaseError = ErrorBuilder.New(http.StatusInternalServerError).
				ErrorCode("S002").
				Message("파이어베이스 관련 에러가 발생하였습니다.").
				Build()
	DatabaseServerError = ErrorBuilder.New(http.StatusInternalServerError).
				ErrorCode("S003").
				Message("데이터베이스 서버에서 에러가 발생하였습니다.").
				Build()
	RedisServerError = ErrorBuilder.New(http.StatusInternalServerError).
				ErrorCode("S004").
				Message("레디스 서버에서 에러가 발생하였습니다.").
				Build()

	BadRequestError = ErrorBuilder.New(http.StatusBadRequest).
			ErrorCode("R001").
			Message("유효하지 않은 요청입니다. 요청 메시지를 다시 확인해주세요.").
			Build()
	BadParametersError = ErrorBuilder.New(http.StatusBadRequest).
				ErrorCode("R002").
				Message("올바르지 않은 파라미터입니다.").
				Build()
	DataValidationFailedError = ErrorBuilder.New(http.StatusUnprocessableEntity).
					ErrorCode("R003").
					Message("조건에 맞지 않는 요청입니다. 요청 메시지를 다시 확인해주세요.").
					Build()

	AlreadyExistsUsernameError = ErrorBuilder.New(http.StatusConflict).
					ErrorCode("D001").
					Message("동일한 아이디가 이미 존재합니다. 변경 후 다시 시도해주세요.").
					Build()
	AlreadyExistsProjectIDError = ErrorBuilder.New(http.StatusConflict).
					ErrorCode("D002").
					Message("이미 등록된 프로젝트 ID입니다. 프로젝트는 하나만 등록할 수 있습니다.").
					Build()
	DataNotFoundError = ErrorBuilder.New(http.StatusNotFound).
				ErrorCode("D003").
				Message("존재하지 않는 데이터입니다.").
				Build()
	UserNotFoundError = ErrorBuilder.New(http.StatusNotFound).
				ErrorCode("D004").
				Message("아직 가입하지 않았거나 승인되지 않은 계정입니다.").
				Build()

	AuthenticationRequiredError = ErrorBuilder.New(http.StatusUnauthorized).
					ErrorCode("A001").
					Message("인증이 필요한 요청입니다. 로그인 후 다시 시도해주세요.").
					Build()
	InvalidTokenError = ErrorBuilder.New(http.StatusUnauthorized).
				ErrorCode("A002").
				Message("잘못된 접근입니다.").
				Build()
	ExpiredTokenError = ErrorBuilder.New(http.StatusUnauthorized).
				ErrorCode("A003").
				Message("만료된 토큰입니다. 다시 로그인해주세요.").
				Build()
	BeforeIssuedTokenError = ErrorBuilder.New(http.StatusUnauthorized).
				ErrorCode("A004").
				Message("아직 사용이 불가한 토큰입니다.").
				Build()
	PasswordNotMatchedError = ErrorBuilder.New(http.StatusUnauthorized).
				ErrorCode("A005").
				Message("비밀번호가 일치하지 않습니다.").
				Build()
	UsernameValidationFailedError = ErrorBuilder.New(http.StatusUnprocessableEntity).
					ErrorCode("A006").
					Message("아이디는 4 ~ 32자의 영문자와 숫자로 이루어져야 합니다.").
					Build()
	PasswordValidationFailedError = ErrorBuilder.New(http.StatusUnprocessableEntity).
					ErrorCode("A007").
					Message("비밀번호는 8 ~ 32자의 영문자와 숫자, 특수문자로 이루어져야 합니다.").
					Build()
	InvalidClientKeyError = ErrorBuilder.New(http.StatusForbidden).
				ErrorCode("A008").
				Message("올바르지 않거나 등록되지 않은 Client Key입니다.").
				Build()
)

type errorBuilder struct {
	code       string
	statusCode int
	message    string
	reason     string
	data       interface{}
}

type registeredErrorBuilder struct {
	code       string
	statusCode int
	message    string
	reason     string
	data       interface{}
}

type eb struct{}

var ErrorBuilder *eb

func (b *eb) New(statusCode int) *errorBuilder {
	return &errorBuilder{statusCode: statusCode}
}

func (b *eb) NewWithError(error *errorResponse) *registeredErrorBuilder {
	return &registeredErrorBuilder{
		code:       error.Code,
		statusCode: error.StatusCode,
		message:    error.Message,
		reason:     error.Reason,
		data:       error.Data,
	}
}

func (b *errorBuilder) ErrorCode(errorCode string) *errorBuilder {
	b.code = errorCode
	return b
}

func (b *errorBuilder) Message(message string) *errorBuilder {
	b.message = message
	return b
}

func (b *errorBuilder) Reason(reason string) *errorBuilder {
	b.reason = reason
	return b
}

func (b *errorBuilder) Data(data interface{}) *errorBuilder {
	b.data = data
	return b
}

func (rb *registeredErrorBuilder) Reason(reason string) *registeredErrorBuilder {
	rb.reason = reason
	return rb
}

func (rb *registeredErrorBuilder) Data(data interface{}) *registeredErrorBuilder {
	rb.data = data
	return rb
}

func (b *errorBuilder) Build() *errorResponse {
	return &errorResponse{
		Code:       b.code,
		StatusCode: b.statusCode,
		Message:    b.message,
		Reason:     b.reason,
		Data:       b.data,
	}
}

func (rb *registeredErrorBuilder) Build() *errorResponse {
	return &errorResponse{
		Code:       rb.code,
		StatusCode: rb.statusCode,
		Message:    rb.message,
		Reason:     rb.reason,
		Data:       rb.data,
	}
}
