package payload

const (
	StatusContinue           = 100 // RFC 7231, 6.2.1
	StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1

	StatusOK                   = 200 // RFC 7231, 6.3.1
	StatusCreated              = 201 // RFC 7231, 6.3.2
	StatusAccepted             = 202 // RFC 7231, 6.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 7231, 6.3.4
	StatusNoContent            = 204 // RFC 7231, 6.3.5
	StatusResetContent         = 205 // RFC 7231, 6.3.6
	StatusPartialContent       = 206 // RFC 7233, 4.1
	StatusMultiStatus          = 207 // RFC 4918, 11.1
	StatusAlreadyReported      = 208 // RFC 5842, 7.1
	StatusIMUsed               = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices   = 300 // RFC 7231, 6.4.1
	StatusMovedPermanently  = 301 // RFC 7231, 6.4.2
	StatusFound             = 302 // RFC 7231, 6.4.3
	StatusSeeOther          = 303 // RFC 7231, 6.4.4
	StatusNotModified       = 304 // RFC 7232, 4.1
	StatusUseProxy          = 305 // RFC 7231, 6.4.5
	_                       = 306 // RFC 7231, 6.4.6 (Unused)
	StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7
	StatusPermanentRedirect = 308 // RFC 7538, 3

	StatusBadRequest                   = 400 // RFC 7231, 6.5.1
	StatusUnauthorized                 = 401 // RFC 7235, 3.1
	StatusPaymentRequired              = 402 // RFC 7231, 6.5.2
	StatusForbidden                    = 403 // RFC 7231, 6.5.3
	StatusNotFound                     = 404 // RFC 7231, 6.5.4
	StatusMethodNotAllowed             = 405 // RFC 7231, 6.5.5
	StatusNotAcceptable                = 406 // RFC 7231, 6.5.6
	StatusProxyAuthRequired            = 407 // RFC 7235, 3.2
	StatusRequestTimeout               = 408 // RFC 7231, 6.5.7
	StatusConflict                     = 409 // RFC 7231, 6.5.8
	StatusGone                         = 410 // RFC 7231, 6.5.9
	StatusLengthRequired               = 411 // RFC 7231, 6.5.10
	StatusPreconditionFailed           = 412 // RFC 7232, 4.2
	StatusRequestEntityTooLarge        = 413 // RFC 7231, 6.5.11
	StatusRequestURITooLong            = 414 // RFC 7231, 6.5.12
	StatusUnsupportedMediaType         = 415 // RFC 7231, 6.5.13
	StatusRequestedRangeNotSatisfiable = 416 // RFC 7233, 4.4
	StatusExpectationFailed            = 417 // RFC 7231, 6.5.14
	StatusTeapot                       = 418 // RFC 7168, 2.3.3
	StatusUnprocessableEntity          = 422 // RFC 4918, 11.2
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusUpgradeRequired              = 426 // RFC 7231, 6.5.15
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 7231, 6.6.1
	StatusNotImplemented                = 501 // RFC 7231, 6.6.2
	StatusBadGateway                    = 502 // RFC 7231, 6.6.3
	StatusServiceUnavailable            = 503 // RFC 7231, 6.6.4
	StatusGatewayTimeout                = 504 // RFC 7231, 6.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 7231, 6.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)

var statusText = map[int]*MossMessage{
	StatusContinue:           {Code: StatusContinue, Msg: "Continue"},
	StatusSwitchingProtocols: {Code: StatusSwitchingProtocols, Msg: "Switching Protocols"},
	StatusProcessing:         {Code: StatusProcessing, Msg: "Processing"},

	StatusOK:                   {Code: StatusOK, Msg: "Ok"},
	StatusCreated:              {Code: StatusCreated, Msg: "Created"},
	StatusAccepted:             {Code: StatusAccepted, Msg: "Accepted"},
	StatusNonAuthoritativeInfo: {Code: StatusNonAuthoritativeInfo, Msg: "Non-Authoritative Information"},
	StatusNoContent:            {Code: StatusNoContent, Msg: "No Content"},
	StatusResetContent:         {Code: StatusResetContent, Msg: "Reset Content"},
	StatusPartialContent:       {Code: StatusPartialContent, Msg: "Partial Content"},
	StatusMultiStatus:          {Code: StatusMultiStatus, Msg: "Multi-Status"},
	StatusAlreadyReported:      {Code: StatusAlreadyReported, Msg: "Already Reported"},
	StatusIMUsed:               {Code: StatusIMUsed, Msg: "IM Used"},

	StatusMultipleChoices:   {Code: StatusMultipleChoices, Msg: "Multiple Choices"},
	StatusMovedPermanently:  {Code: StatusMovedPermanently, Msg: "Moved Permanently"},
	StatusFound:             {Code: StatusFound, Msg: "Found"},
	StatusSeeOther:          {Code: StatusSeeOther, Msg: "See Other"},
	StatusNotModified:       {Code: StatusNotModified, Msg: "Not Modified"},
	StatusUseProxy:          {Code: StatusUseProxy, Msg: "Use Proxy"},
	StatusTemporaryRedirect: {Code: StatusTemporaryRedirect, Msg: "Temporary Redirect"},
	StatusPermanentRedirect: {Code: StatusPermanentRedirect, Msg: "Permanent Redirect"},

	StatusBadRequest:                   {Code: StatusBadRequest, Msg: "Bad Request"},
	StatusUnauthorized:                 {Code: StatusUnauthorized, Msg: "Unauthorized"},
	StatusPaymentRequired:              {Code: StatusPaymentRequired, Msg: "Payment Required"},
	StatusForbidden:                    {Code: StatusForbidden, Msg: "Forbidden"},
	StatusNotFound:                     {Code: StatusNotFound, Msg: "Not Found"},
	StatusMethodNotAllowed:             {Code: StatusMethodNotAllowed, Msg: "Method Not Allowed"},
	StatusNotAcceptable:                {Code: StatusNotAcceptable, Msg: "Not Acceptable"},
	StatusProxyAuthRequired:            {Code: StatusProxyAuthRequired, Msg: "Proxy Authentication Required"},
	StatusRequestTimeout:               {Code: StatusRequestTimeout, Msg: "Request Timeout"},
	StatusConflict:                     {Code: StatusConflict, Msg: "Conflict"},
	StatusGone:                         {Code: StatusGone, Msg: "Gone"},
	StatusLengthRequired:               {Code: StatusLengthRequired, Msg: "Length Required"},
	StatusPreconditionFailed:           {Code: StatusPreconditionFailed, Msg: "Precondition Failed"},
	StatusRequestEntityTooLarge:        {Code: StatusRequestEntityTooLarge, Msg: "Request Entity Too Large"},
	StatusRequestURITooLong:            {Code: StatusRequestURITooLong, Msg: "Request URI Too Long"},
	StatusUnsupportedMediaType:         {Code: StatusUnsupportedMediaType, Msg: "Unsupported Media Type"},
	StatusRequestedRangeNotSatisfiable: {Code: StatusRequestedRangeNotSatisfiable, Msg: "Requested Range Not Satisfiable"},
	StatusExpectationFailed:            {Code: StatusExpectationFailed, Msg: "Expectation Failed"},
	StatusTeapot:                       {Code: StatusTeapot, Msg: "I'm a teapot"},
	StatusUnprocessableEntity:          {Code: StatusUnprocessableEntity, Msg: "Unprocessable Entity"},
	StatusLocked:                       {Code: StatusLocked, Msg: "Locked"},
	StatusFailedDependency:             {Code: StatusFailedDependency, Msg: "Failed Dependency"},
	StatusUpgradeRequired:              {Code: StatusUpgradeRequired, Msg: "Upgrade Required"},
	StatusPreconditionRequired:         {Code: StatusPreconditionRequired, Msg: "Precondition Required"},
	StatusTooManyRequests:              {Code: StatusTooManyRequests, Msg: "Too Many Requests"},
	StatusRequestHeaderFieldsTooLarge:  {Code: StatusRequestHeaderFieldsTooLarge, Msg: "Request Header Fields Too Large"},
	StatusUnavailableForLegalReasons:   {Code: StatusUnavailableForLegalReasons, Msg: "Unavailable For Legal Reasons"},

	StatusInternalServerError:           {Code: StatusInternalServerError, Msg: "Internal Server Error"},
	StatusNotImplemented:                {Code: StatusNotImplemented, Msg: "Internal Server Error"},
	StatusBadGateway:                    {Code: StatusBadGateway, Msg: "Bad Gateway"},
	StatusServiceUnavailable:            {Code: StatusServiceUnavailable, Msg: "Service Unavailable"},
	StatusGatewayTimeout:                {Code: StatusGatewayTimeout, Msg: "Gateway Timeout"},
	StatusHTTPVersionNotSupported:       {Code: StatusHTTPVersionNotSupported, Msg: "HTTP Version Not Supported"},
	StatusVariantAlsoNegotiates:         {Code: StatusVariantAlsoNegotiates, Msg: "Variant Also Negotiates"},
	StatusInsufficientStorage:           {Code: StatusInsufficientStorage, Msg: "Insufficient Storage"},
	StatusLoopDetected:                  {Code: StatusLoopDetected, Msg: "Loop Detected"},
	StatusNotExtended:                   {Code: StatusNotExtended, Msg: "Not Extended"},
	StatusNetworkAuthenticationRequired: {Code: StatusNetworkAuthenticationRequired, Msg: "Network Authentication Required"},
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) *MossMessage {
	return statusText[code]
}
