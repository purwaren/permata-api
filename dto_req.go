package permata

// === API req ===

// APIReq wraps all requests as defined by spec
type APIReq struct {
	BalanceInquiryReq BalanceInquiryReq `json:"BalInqRq,omitempty"`
}

// === Balance Inquiry req ===

// BalanceInquiryReq is request for Balance Inquiry
type BalanceInquiryReq struct {
	MsgReqHeader   MsgReqHeader   `json:"MsgRqHdr,omitempty"`
	InquiryInfoReq InquiryInfoReq `json:"InqInfo,omitempty"`
}

// InquiryInfoReq is embedded info request for Balance Inquiry
type InquiryInfoReq struct {
	AccountNumber string
}

// === Common Req ===

// MsgReqHeader is embedded in every request
type MsgReqHeader struct {
	RequestTimestamp string
	CustRefID        string
}
