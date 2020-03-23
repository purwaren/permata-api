package permata

type GetBillRq struct {
    GetBillRq BillReq
}

type BillReq struct {
    InstCode    string  `json:"INSTCODE"`
    ViVaNumber  string  `json:"VI_VANUMBER"`
    ViTraceNo   string  `json:"VI_TRACENO"`
    ViTrnDate   string  `json:"VI_TRNDATE"`
    ViDelChnl   string  `json:"VI_DELCHANNEL"`
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

type PayBillReq struct {
    InstCode    string `json:"INSTCODE"`
    ViVaNumber  string  `json:"VI_VANUMBER"`
    ViTraceNo   string  `json:"VI_TRACENO"`
    ViTrnDate   string  `json:"VI_TRNDATE"`
    BillAmount  string  `json:"BILL_AMOUNT"`
    ViCcy       string  `json:"VI_CCY"`
    ViDelChnl   string  `json:"VI_DELCHANNEL"`
    RefInfo     []RefInfo `json:"RefInfo"`
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


    