package request

const (
	// MimeTypeJSON is application/json type
	MimeTypeJSON = "application/json"

	// MimeTypeFormData is application/form-data
	MimeTypeFormData = "multipart/form-data"

	// MimeTypeFormURL is application/x-www-form-urlencoded type
	MimeTypeFormURL = "application/x-www-form-urlencoded"
)

type requestAPI string

const (
	GetAPI    requestAPI = "Get"
	PostAPI   requestAPI = "Post"
	DeleteAPI requestAPI = "Delete"
	PatchAPI  requestAPI = "Patch"
)

type httpHeader string

const (
	ContentTypeHeader   httpHeader = "Content-Type"
	UserAgentHeader     httpHeader = "User-Agent"
	AuthorizationHeader httpHeader = "Authorization"
)
