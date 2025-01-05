package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	ErrInvalidToken     = 30001
	// Register code
	ErrCodeUserHasExists = 50001
)

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "token is invalid",
	ErrCodeUserHasExists: "User has already registered",
}
