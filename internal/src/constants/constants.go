package constants

//role
const (
	User  = "user"
	Admin = "admin"
	Staff = "staff"
)

var (
	SUCESS               = 200
	BADREQUEST           = 400
	UNAUTHORIZED         = 401
	FORBIDDEN            = 403
	NOTFOUND             = 404
	METHODNOTALLOWED     = 405
	INTERNALSERVERERROR  = 500
	NOTIMPLEMENTED       = 501
	BADGATEWAY           = 502
	SERVICEUNAVAILABLE   = 503
	GATEWAYTIMEOUT       = 504
	UNSUPPORTEDMEDIATYPE = 415
	UNPROCESSABLEENTITY  = 422
	PENDING              = "PENDING"
	COMPLETED            = "COMPLETED"
	APPROVED             = "APPROVED"
	GENERATE             = "GENERATE"
	ACCEPTED             = "ACCEPTED"
	REJECTED             = "REJECTED"

	// New page constants
	PLAN_PAGE     = "PLAN_PAGE"
	INVOICE_PAGE  = "INVOICE_PAGE"
	CHECKOUT_PAGE = "CHECKOUT_PAGE"
	PLAN_STATIC   = 7

	// User role IDs
	PLATFORM_USER_ID = 1
	ADMIN_USER_ID    = 2
)