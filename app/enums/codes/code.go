package codes

const (
	Success         = "OK"
	ParamError      = "400"
	AuthorizedError = "401"
	NotFoundError   = "404"
	SystemError     = "500"
	BusinessError   = "501"

	AuthCheckTokenFail    = "20001"
	AuthCheckTokenTimeout = "20002"
)

var ErrorMap = map[string]string{
	Success:               "success",
	ParamError:            "params error",
	AuthorizedError:       "unauthorized",
	NotFoundError:         "not found error",
	SystemError:           "system error",
	BusinessError:         "business fail",
	AuthCheckTokenFail:    "auth: check token fail",
	AuthCheckTokenTimeout: "auth: check token timeout",
}
