package nuonuo

type Content struct {
	Order Order `json:"order"`
}

type Order struct {
	BillInfoNo       string          `json:"billInfoNo,omitempty"`
	BuyerAccount     string          `json:"buyerAccount"` // 购买方银行及账号
	BuyerAddress     string          `json:"buyerAddress"` // 购买方地址
	BuyerName        string          `json:"buyerName"`    // Y 购买方名称
	BuyerPhone       string          `json:"buyerPhone"`   // 购方手机（pushMode为1或2时，此项为必填）
	BuyerTaxNum      string          `json:"buyerTaxNum"`  // 购买方税号
	BuyerTel         string          `json:"buyerTel"`     // 购买方电话
	CallBackURL      string          `json:"callBackUrl,omitempty"`
	Checker          string          `json:"checker"`
	Clerk            string          `json:"clerk,omitempty"` // Y 开票员
	ClerkID          string          `json:"clerkId,omitempty"`
	DepartmentID     string          `json:"departmentId,omitempty"`
	Email            string          `json:"email"` // 推送邮箱（pushMode为0或2时，此项为必填）
	ExtensionNumber  string          `json:"extensionNumber,omitempty"`
	InvoiceCode      string          `json:"invoiceCode,omitempty"`
	InvoiceDate      string          `json:"invoiceDate"` // Y 订单时间
	InvoiceDetail    []InvoiceDetail `json:"invoiceDetail"`
	InvoiceLine      string          `json:"invoiceLine,omitempty"`
	InvoiceNum       string          `json:"invoiceNum,omitempty"`
	InvoiceType      string          `json:"invoiceType"` // 开票类型：1:蓝票;2:红票
	ListFlag         string          `json:"listFlag,omitempty"`
	ListName         string          `json:"listName,omitempty"`
	MachineCode      string          `json:"machineCode,omitempty"`
	OrderNo          string          `json:"orderNo"` // Y 订单号（每个企业唯一）
	Payee            string          `json:"payee,omitempty"`
	ProductOilFlag   string          `json:"productOilFlag,omitempty"`
	ProxyInvoiceFlag string          `json:"proxyInvoiceFlag,omitempty"`
	PushMode         string          `json:"pushMode"`
	Remark           string          `json:"remark"`
	SalerAccount     string          `json:"salerAccount"` // 销方银行账号和开户行地址
	SalerAddress     string          `json:"salerAddress"` // Y 销方地址
	SalerTaxNum      string          `json:"salerTaxNum"`  // Y 销方税号（使用沙箱环境请求时消息体参数salerTaxNum和消息头参数userTax填写339901999999142）
	SalerTel         string          `json:"salerTel"`     // Y 销方电话
	TerminalNumber   string          `json:"terminalNumber,omitempty"`
	VehicleFlag      string          `json:"vehicleFlag,omitempty"`
	VehicleInfo      VehicleInfo     `json:"vehicleInfo,omitempty"`
}

type InvoiceDetail struct {
	Deduction           string `json:"deduction,omitempty"`
	FavouredPolicyFlag  string `json:"favouredPolicyFlag,omitempty"`
	FavouredPolicyName  string `json:"favouredPolicyName,omitempty"`
	GoodsCode           string `json:"goodsCode,omitempty"`
	GoodsName           string `json:"goodsName"` // Y 商品名称（如invoiceLineProperty =1，则此商品行为折扣行，折扣行不允许多行折扣，折扣行必须紧邻被折扣行，商品名称必须与被折扣行一致）
	InvoiceLineProperty string `json:"invoiceLineProperty,omitempty"`
	Num                 string `json:"num,omitempty"`
	Price               string `json:"price,omitempty"`
	SelfCode            string `json:"selfCode,omitempty"`
	SpecType            string `json:"specType,omitempty"`
	Tax                 string `json:"tax,omitempty"`
	TaxExcludedAmount   string `json:"taxExcludedAmount,omitempty"`
	TaxIncludedAmount   string `json:"taxIncludedAmount,omitempty"`
	TaxRate             string `json:"taxRate"` // Y 税率，注：纸票清单红票存在为null的情况
	Unit                string `json:"unit,omitempty"`
	WithTaxFlag         string `json:"withTaxFlag"` // Y 单价含税标志：0:不含税,1:含税
	ZeroRateFlag        string `json:"zeroRateFlag,omitempty"`
}

type VehicleInfo struct {
	BrandModel       string `json:"brandModel"`
	Certificate      string `json:"certificate"`
	EngineNum        string `json:"engineNum"`
	IDNumOrgCode     string `json:"idNumOrgCode"`
	ImportCerNum     string `json:"importCerNum"`
	InsOddNum        string `json:"insOddNum"`
	IntactCerNum     string `json:"intactCerNum"`
	ManufacturerName string `json:"manufacturerName"`
	MaxCapacity      string `json:"maxCapacity"`
	ProductOrigin    string `json:"productOrigin"`
	TaxOfficeCode    string `json:"taxOfficeCode"`
	TaxOfficeName    string `json:"taxOfficeName"`
	Tonnage          string `json:"tonnage"`
	VehicleCode      string `json:"vehicleCode"`
	VehicleType      string `json:"vehicleType"`
}
