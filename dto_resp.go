package permata

// === AUTH resp ===

// OAuth2Resp is OAuth2 response
type OAuth2Resp struct {
	AccessToken string `json:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	ExpiredIn   int64  `json:"expired_in,omitempty"`
	Scope       string `json:"scope,omitempty"`
}

// === API resp ===

// APIResp wraps all responses as defined by spec
type APIResp struct {
}

// === Balance Inquiry resp ===

// BalanceInquiryResp is response for Balance Inquiry
type BalanceInquiryResp struct {
	MsgRespHeader  MsgRespHeader  `json:"MsgRsHdr,omitempty"`
	InquiryInfoResp InquiryInfoResp `json:"InqInfo,omitempty"`
}

// InquiryInfoResp is embedded info response for Balance Inquiry
type InquiryInfoResp struct {
	AccountNumber        string
	AccountCurrency      string
	AccountBalanceAmount string
	BalanceType          string
}

// === Common resp ===

// MsgRespHeader is embedded in every response
type MsgRespHeader struct {
	RequestTimestamp string
	CustRefID        string
	StatusCode       string
	StatusDesc       string
}
