package permata

import validation "github.com/go-ozzo/ozzo-validation"

type GetBillRq struct {
    GetBillRq BillReq
}

func (p GetBillRq) Validate() error {
    return validation.ValidateStruct(&p,
        validation.Field(&p.GetBillRq, validation.Required),
    )
}

type BillReq struct {
    InstCode    string  `json:"INSTCODE"`
    ViVaNumber  string  `json:"VI_VANUMBER"`
    ViTraceNo   string  `json:"VI_TRACENO"`
    ViTrnDate   string  `json:"VI_TRNDATE"`
    ViDelChnl   string  `json:"VI_DELCHANNEL"`
}

func (p BillReq) Validate() error {
    return validation.ValidateStruct(&p,
        validation.Field(&p.InstCode, validation.Required),
        validation.Field(&p.ViVaNumber, validation.Required),
        validation.Field(&p.ViTraceNo, validation.Required),
        validation.Field(&p.ViTrnDate, validation.Required),
        validation.Field(&p.ViDelChnl, validation.Required),
    )
}

func (p BillReq) ValidateTrnDate() error {
    return validation.ValidateStruct(&p,
        validation.Field(&p.ViTrnDate, validation.Date("2006-01-02T15:04:05+07:00")),
    )
}

type GetBillRs struct {
    GetBillRs BillRes
}

type BillRes struct {
    CustName    string  `json:"CUSTNAME"`
    BillAmount  string  `json:"BILL_AMOUNT"`
    ViCcy       string  `json:"VI_CCY"`
    Status      string  `json:"STATUS"`
}

type PayBillRq struct {
    PayBillRq PayBillReq
}

func (p PayBillRq) Validate() error {
    return validation.ValidateStruct(&p,
        validation.Field(&p.PayBillRq, validation.Required),
    )
}

type PayBillReq struct {
    InstCode    string  `json:"INSTCODE"`
    ViVaNumber  string  `json:"VI_VANUMBER"`
    ViTraceNo   string  `json:"VI_TRACENO"`
    ViTrnDate   string  `json:"VI_TRNDATE"`
    BillAmount  string  `json:"BILL_AMOUNT"`
    ViCcy       string  `json:"VI_CCY"`
    ViDelChnl   string  `json:"VI_DELCHANNEL"`
    RefInfo     []RefInfo `json:"RefInfo"`
}

func (p PayBillReq) Validate() error {
    return validation.ValidateStruct(&p,
        validation.Field(&p.InstCode, validation.Required),
        validation.Field(&p.ViVaNumber, validation.Required),
        validation.Field(&p.ViTraceNo, validation.Required),
        validation.Field(&p.ViTrnDate, validation.Required),
        validation.Field(&p.BillAmount, validation.Required),
        validation.Field(&p.ViCcy, validation.Required),
        validation.Field(&p.ViDelChnl, validation.Required),
    )
}

func (p PayBillReq) ValidateTrnDate() error {
    return validation.ValidateStruct(&p,
        validation.Field(&p.ViTrnDate, validation.Date("2006-01-02T15:04:05+07:00")),
    )
}

type RefInfo struct {
    RefName     string  `json:"RefName"`
    RefValue    string  `json:"RefValue"`
}

type PayBillRs struct {
    PayBillRs PayBillRes
}

type PayBillRes struct {
    Status      string  `json:"STATUS"`
}

type RevBillRq struct {
    RevBillRq PayBillReq
}

type RevBillRs struct {
    RevBillRs PayBillRes
}


    