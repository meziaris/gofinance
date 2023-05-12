package reason

var (
	UserAlreadyExist   = "user already exist"
	RegisterFailed     = "failed to register user"
	UserNotFound       = "user not exist"
	FailedLogin        = "failed to login, your email or password is incorrect"
	Unauthorized       = "unauthorized request"
	FailedRefreshToken = "failed to refresh token, please check your token" //nolint
	FailedLogout       = "failed logout"
)
