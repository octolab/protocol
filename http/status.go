package http

const (
	StatusAuthenticationTimeout = 419
	StatusRetryWith             = 449
	StatusClientClosedRequest   = 499

	StatusBandwidthLimitExceeded = 509
	StatusUnknownError           = 520
	StatusWebServerIsDown        = 521
	StatusConnectionTimedOut     = 522
	StatusOriginIsUnreachable    = 523
	StatusATimeoutOccurred       = 524
	StatusSSLHandshakeFailed     = 525
	StatusInvalidSSLCertificate  = 526
)
