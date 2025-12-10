package constants


// role
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

	UNAVAILABLE             = "UNAVAILABLE"
	PENDING                 = "PENDING"
	COMPLETED               = "COMPLETED"
	APPROVED                = "APPROVED"
	ASSIGNEDSLOT            = "ASSIGNEDSLOT"
	ACCEPTED                = "ACCEPTED"
	REJECTED                = "REJECTED"
	SERVICESTARTED          = "SERVICE STARTED"
	SERVICEENDED            = "SERVICE ENDED"
	DELIVERED               = "DELIVERED"
	WAITING                 = "WAITING"
	APPROVEDMSG             = "Your request has been approved. Our team will contact you soon."
	ACCEPTEDMSG             = "Your booking has been accepted. Stay tuned, our crew is on it!"
	REJECTEDMSG             = "Your request has been rejected. Please review and try again."
	PENDINGMSG              = "Your request is pending. Our team will update you shortly."
	UNAVAILABLEMSG          = "This service is currently unavailable. Please try again later."
	COMPLETEDMSG            = "The process is completed. Thank you for your patience!"
	ASSIGNEDSLOTMSG         = "A slot has been assigned for your service. Our team will keep you updated."
	SERVICESTARTEDMSG       = "Service has started — your vehicle is now being worked on."
	SERVICEENDEDMSG         = "Service is completed! Our team will contact you for the next steps."
	DELIVEREDMSG            = "Your vehicle/service has been delivered. Thank you for choosing us!"
	WAITINGMSG              = "Your request is in the queue. Our team will connect with you soon."
	FORBIDDENMSG            = "You don't have permission to access this resource."
	UNAUTHORIZEDMSG         = "Unauthorized request. Please authenticate to continue."
	BADREQUESTMSG           = "Something seems off with your request. Please check and try again."
	NOTFOUNDMSG             = "The resource you’re looking for could not be found."
	NOTIMPLEMENTEDMSG       = "This feature is not implemented yet. Coming soon!"
	ACCEPTORASSIGNEDBOOKING = "Request accepted you can see more details on booked"

	// New page constants
	PLAN_PAGE     = "PLAN_PAGE"
	INVOICE_PAGE  = "INVOICE_PAGE"
	CHECKOUT_PAGE = "CHECKOUT_PAGE"
	PLAN_STATIC   = 7

	
	// preload
	PRELOADUSER     = "Users"
	PRELOADSTAFF    = "Staff"
	PRELOADSLOT     = "Slot"
	PRELOADBOOKINGS = "Bookings"
	PRELOADBOOKED   = "Bookeds"
)

