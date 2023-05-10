package nuonuo

type Content struct {
	Order Order `json:"order"`
}

// https://open.jss.com.cn/#/api-doc/common-api?id=100607
type Order struct {
	BuyerName               string          `json:"buyerName,omitempty"`               // Y 企业名称/个人 100 购方名称
	BuyerTaxNum             string          `json:"buyerTaxNum,omitempty"`             // N 339901999999198 20 购方税号（企业要填，个人可为空；数电专票、二手车销售统一发票时必填）
	BuyerTel                string          `json:"buyerTel,omitempty"`                // N 0571-88888888 50 购方电话（购方地址+电话总共不超100字符；二手车销售统一发票时必填）
	BuyerAddress            string          `json:"buyerAddress,omitempty"`            // N 杭州市 80 购方地址（购方地址+电话总共不超100字符；二手车销售统一发票时必填）
	BuyerAccount            string          `json:"buyerAccount,omitempty"`            // N 中国工商银行 111111111111 100 购方银行开户行及账号
	SalerTaxNum             string          `json:"salerTaxNum,omitempty"`             // Y 339901999999199 20 销方税号（使用沙箱环境请求时消息体参数salerTaxNum和消息头参数userTax填写339902999999789113）
	SalerTel                string          `json:"salerTel,omitempty"`                // Y 0571-77777777 20 销方电话（在诺税通saas工作台配置过的可以不传，以传入的为准）
	SalerAddress            string          `json:"salerAddress,omitempty"`            // Y  80 销方地址（在诺税通saas工作台配置过的可以不传，以传入的为准）
	SalerAccount            string          `json:"salerAccount,omitempty"`            // N  100 销方银行开户行及账号(二手车销售统一发票时必填)
	OrderNo                 string          `json:"orderNo,omitempty"`                 // Y 201701053332079312313 20 订单号（每个企业唯一）
	InvoiceDate             string          `json:"invoiceDate,omitempty"`             // Y 2022-01-13 12:30:00 20 订单时间
	InvoiceCode             string          `json:"invoiceCode,omitempty"`             // N  12 冲红时填写的对应蓝票发票代码（红票必填 10位或12 位， 11位的时候请左补 0）
	InvoiceNum              string          `json:"invoiceNum,omitempty"`              // N  8 冲红时填写的对应蓝票发票号码（红票必填，不满8位请左补0）
	RedReason               string          `json:"redReason,omitempty"`               // N 1 1 冲红原因：1:销货退回;2:开票有误;3:服务中止;4:发生销售折让(开具红票时且票种为p,c,e,f,r需要传--成品油发票除外；不传时默认为 1)
	BillInfoNo              string          `json:"billInfoNo,omitempty"`              // N 1403011904008472 24 红字信息表编号.专票冲红时此项必填，且必须在备注中注明“开具红字增值税专用发票信息表编号ZZZZZZZZZZZZZZZZ”字样，其 中“Z”为开具红字增值税专用发票所需要的长度为16位信息表编号（建议16位，最长可支持24位）。
	DepartmentId            string          `json:"departmentId,omitempty"`            // N 9F7E9439CA8B4C60A2FFF3EA3290B088 32 部门门店id（诺诺系统中的id）
	ClerkId                 string          `json:"clerkId,omitempty"`                 // N  32 开票员id（诺诺系统中的id）
	Remark                  string          `json:"remark,omitempty"`                  // N 备注信息 230 冲红时，在备注中注明“对应正数发票代码:XXXXXXXXX号码:YYYYYYYY”文案，其中“X”为发票代码，“Y”为发票号码，可以不填，接口会自动添加该文案；机动车发票蓝票时备注只能为空；数电票时最长为200字符
	Checker                 string          `json:"checker,omitempty"`                 // N 王五 20 复核人
	Payee                   string          `json:"payee,omitempty"`                   // N 李四 20 收款人
	Clerk                   string          `json:"clerk,omitempty"`                   // Y 张三 20 开票员（数电票时需要传入和开票登录账号对应的开票员姓名）
	ListFlag                string          `json:"listFlag,omitempty"`                // N 0 1 清单标志：非清单:0；清单:1，默认:0，电票固定为0
	ListName                string          `json:"listName,omitempty"`                // N 详见销货清单 92 清单项目名称：对应发票票面项目名称（listFlag为1时，必填，默认为“详见销货清单”）
	PushMode                string          `json:"pushMode,omitempty"`                // N 1 2 推送方式：-1,不推送;0,邮箱;1,手机（默认）;2,邮箱、手机
	BuyerPhone              string          `json:"buyerPhone,omitempty"`              // Y 15858585858 20 购方手机（pushMode为1或2时，此项为必填，同时受企业资质是否必填控制）
	Email                   string          `json:"email,omitempty"`                   // Y test@xx.com 50 推送邮箱（pushMode为0或2时，此项为必填，同时受企业资质是否必填控制）
	InvoiceType             string          `json:"invoiceType,omitempty"`             // Y 1 1 开票类型：1:蓝票;2:红票 （数电票冲红请对接数电快捷冲红接口）
	InvoiceLine             string          `json:"invoiceLine,omitempty"`             // N p 2 发票种类：p,普通发票(电票)(默认);c,普通发票(纸票);s,专用发票;e,收购发票(电票);f,收购发票(纸质);r,普通发票(卷式);b,增值税电子专用发票;j,机动车销售统一发票;u,二手车销售统一发票;bs:电子发票(增值税专用发票)-即数电专票(电子),pc:电子发票(普通发票)-即数电普票(电子),es:数电纸质发票(增值税专用发票)-即数电专票(纸质);ec:数电纸质发票(普通发票)-即数电普票(纸质)
	PaperInvoiceType        string          `json:"paperInvoiceType,omitempty"`        // N  12 数电纸票类型(数电纸票时才需要传)：（票种为ec时，默认04；票种为es时，默认为1130）; 04 2016版增值税普通发票（二联折叠票）, 05 2016版增值税普通发票（五联折叠票), 1130 增值税专用发票（中文三联无金额限制版）, 1140 增值税专用发票（中文四联无金额限制版）, 1160 增值税专用发票（中文六联无金额限制版）, 1170 增值税专用发票（中文七联无金额限制版）
	SpecificFactor          string          `json:"specificFactor,omitempty"`          // N 0 2 特定要素：0普通发票（默认）、1 成品油 、3 建筑服务、4 货物运输服务、5 不动产销售、6 不动产经营租赁服务、9 旅客运输服务、16 农产品收购、31 建安发票 、 32 房地产销售发票、33 二手车发票反向开具、 34 电子烟、 35 矿产品
	ForceFlag               string          `json:"forceFlag,omitempty"`               // N 0 2 是否强制开具标识：0 否、1 是 （发票种类为u，且特定要素为 33-二手车发票反向开具时才需要填； 默认为 0；若为1时，则不校验卖方自然人身份证号的合规性）
	ProxyInvoiceFlag        string          `json:"proxyInvoiceFlag,omitempty"`        // N 0 1 代开标志：0非代开;1代开。代开蓝票时备注要求填写文案：代开企业税号:***,代开企业名称:***；代开红票时备注要求填写文案：对应正数发票代码:***号码:***代开企业税号:***代开企业名称:***
	TaxRebateProxy          string          `json:"taxRebateProxy,omitempty"`          // N 0 1 代办退税标记：0否（默认），1是；仅代办退税资质企业可传1
	InvoiceDifferenceType   string          `json:"invoiceDifferenceType,omitempty"`   // N 02 2 数电发票差额征税开具方式：01 全额开票（暂不支持），02 差额开票；非数电发票开具差额时，不传
	CallBackUrl             string          `json:"callBackUrl,omitempty"`             // N http:127.0.0.1/invoice/callback/  回传发票信息地址（开票完成、开票失败）
	ExtensionNumber         string          `json:"extensionNumber,omitempty"`         // N 0 5 分机号（只能为空或者数字）
	TerminalNumber          string          `json:"terminalNumber,omitempty"`          // N  4 终端号（开票终端号，只能 为空或数字）
	MachineCode             string          `json:"machineCode,omitempty"`             // N 123456789123 12 机器编号（12位盘号）
	VehicleFlag             string          `json:"vehicleFlag,omitempty"`             // N 1 1 是否机动车类专票 0-否 1-是
	HiddenBmbbbh            string          `json:"hiddenBmbbbh,omitempty"`            // N 0 1 是否隐藏编码表版本号 0-否 1-是（默认0，在企业资质中也配置为是隐藏的时候，并且此字段传1的时候代开发票 税率显示***）
	NextInvoiceCode         string          `json:"nextInvoiceCode,omitempty"`         // N  12 指定发票代码（票种为c普纸、f收购纸票时允许指定卷开具） 非必填
	NextInvoiceNum          string          `json:"nextInvoiceNum,omitempty"`          // N  8 发票起始号码，当指定代码有值时，发票起始号码必填
	InvoiceNumEnd           string          `json:"invoiceNumEnd,omitempty"`           // N  8 发票终止号码，当指定代码有值时，发票终止号码必填
	SurveyAnswerType        string          `json:"surveyAnswerType,omitempty"`        // N  1 3%、1%税率开具理由（企业为小规模/点下户时才需要），对应值：1-开具发票为2022年3月31日前发生纳税义务的业务； 2-前期已开具相应征收率发票，发生销售折让、中止或者退回等情形需要开具红字发票，或者开票有误需要重新开具； 3-因为实际经营业务需要，放弃享受免征增值税政策
	BuyerManagerName        string          `json:"buyerManagerName,omitempty"`        // N 张三 16 购买方经办人姓名（数电票特有字段）
	ManagerCardType         string          `json:"managerCardType,omitempty"`         // N 201 40 经办人证件类型：101-组织机构代码证, 102-营业执照, 103-税务登记证, 199-其他单位证件, 201-居民身份证, 202-军官证, 203-武警警官证, 204-士兵证, 205-军队离退休干部证, 206-残疾人证, 207-残疾军人证（1-8级）, 208-外国护照, 210-港澳居民来往内地通行证, 212-中华人民共和国往来港澳通行证, 213-台湾居民来往大陆通行证, 214-大陆居民往来台湾通行证, 215-外国人居留证, 216-外交官证 299-其他个人证件(数电发票特有)
	ManagerCardNo           string          `json:"managerCardNo,omitempty"`           // N  20 经办人证件号码（数电票特有字段）
	BField1                 string          `json:"bField1,omitempty"`                 // N  255 业务方自定义字段1，本应用只作保存
	BField2                 string          `json:"bField2,omitempty"`                 // N  255 业务方自定义字段2，本应用只作保存
	BField3                 string          `json:"bField3,omitempty"`                 // N  255 业务方自定义字段3，本应用只作保存
	NaturalPersonFlag       string          `json:"naturalPersonFlag,omitempty"`       // N 0 1 购买方自然人标志：0-否（默认），1-是；仅在开具数电普票(电子)时使用，如受票方（发票抬头）为自然人，并要求能将发票归集在个人票夹中展示，需提供姓名及身份证号（自然人纳税人识别号），此参数传入1；如受票方（发票抬头）为个体工商户，需提供社会统一信用代码或纳税人识别号，此参数传入0
	TaxNumVerifyFlag        string          `json:"taxNumVerifyFlag,omitempty"`        // N  1 对购方税号校验（ 0-不校验 1-校验，仅对数电票有效，未传时则取企业配置的值；注：若开启校验，当购方税号未能在电子税局中找到时 则会开票失败）
	NaturalPersonVerifyFlag string          `json:"naturalPersonVerifyFlag,omitempty"` // N  1 对购方名称校验（ 0-不校验 1-校验，仅对数电普票（电子）有效，未传时则取企业配置的值；若开启校验，当开具非自然人标记的数电普票（电子）时，将限制对于“购买方名称长度小于等于4位”的发票的开具）
	AdditionalElementName   string          `json:"additionalElementName,omitempty"`   // N 测试模版 50 附加模版名称（数电电票特有字段，附加模版有值时需要添加附加要素信息列表对象，需要先在电子税局平台维护好模版）
	InvoiceDetail           []InvoiceDetail `json:"invoiceDetail,omitempty"`
	VehicleInfo             struct {
		VehicleType      string `json:"vehicleType,omitempty"`      // Y	轿车	40	车辆类型,同明细中商品名称，开具机动车发票时明细有且仅有一行，商品名称为车辆类型且不能为空
		BrandModel       string `json:"brandModel,omitempty"`       // Y	宝马3系	60	厂牌型号
		ProductOrigin    string `json:"productOrigin,omitempty"`    // Y	北京	32	原产地
		Certificate      string `json:"certificate,omitempty"`      // N	WDL042613263551	50	合格证号
		ImportCerNum     string `json:"importCerNum,omitempty"`     // N		36	进出口证明书号
		InsOddNum        string `json:"insOddNum,omitempty"`        // N		32	商检单号
		EngineNum        string `json:"engineNum,omitempty"`        // N	10111011111	50	发动机号码
		VehicleCode      string `json:"vehicleCode,omitempty"`      // Y	LHGK43284342384234	23	车辆识别号码/车架号
		IntactCerNum     string `json:"intactCerNum,omitempty"`     // N		32	完税证明号码
		Tonnage          string `json:"tonnage,omitempty"`          // N	2	8	吨位
		MaxCapacity      string `json:"maxCapacity,omitempty"`      // N	5	12	限乘人数
		DdNumOrgCode     string `json:"ddNumOrgCode,omitempty"`     // N	9114010034683511XD	30	其他证件号码；该字段为空则为2021新版常规机动车发票，此时购方税号必填（个人在购方税号中填身份证号）；该字段有值，则为2021新版其他证件号码的机动车发票（可以录入汉字、大写字母、数字、全角括号等，此时购方税号需要为空；注：仅用于港澳台、国外等特殊身份/税号开机动车票时使用；国内个人身份证号码开具请勿传入该字段，需要传入到购方税号字段中）
		ManufacturerName string `json:"manufacturerName,omitempty"` // N	华晨宝马汽车生产有限公司	80	生产厂家（A9开票服务器类型可支持200）
		TaxOfficeName    string `json:"taxOfficeName,omitempty"`    // N	杭州税务	80	主管税务机关名称（A9开票服务器类型必填）
		TaxOfficeCode    string `json:"taxOfficeCode,omitempty"`    // N	13399000	11	主管税务机关代码（A9开票服务器类型必填）
	} `json:"vehicleInfo,omitempty"` // N		1	机动车销售统一发票才需要传
	SecondHandCarInfo struct {
		OrganizeType          string `json:"organizeType,omitempty"`          // Y	1	1	开票方类型 1：经营单位 2：拍卖单位 3：二手车市场 （只有传1-经营单位时，才支持 特定要素为33的 二手车发票反向开具）
		VehicleType           string `json:"vehicleType,omitempty"`           // Y	轿车	40	车辆类型,同明细中商品名称，开具机动车发票时明细有且仅有一行，商品名称为车辆类型且不能为空
		BrandModel            string `json:"brandModel,omitempty"`            // Y	宝马3系	60	厂牌型号
		VehicleCode           string `json:"vehicleCode,omitempty"`           // Y	LHGK43284342384234	23	车辆识别号码/车架号
		IntactCerNum          string `json:"intactCerNum,omitempty"`          // N		32	完税证明号码
		LicenseNumber         string `json:"licenseNumber,omitempty"`         // Y	浙A12345	20	车牌照号
		RegisterCertNo        string `json:"registerCertNo,omitempty"`        // Y	330022123321	20	登记证号
		VehicleManagementName string `json:"vehicleManagementName,omitempty"` // Y	杭州	80	转入地车管所名称
		SellerName            string `json:"sellerName,omitempty"`            // Y	张三	80	卖方单位/个人名称（开票方类型为1时，必须与销方名称一致，若反向开具时则为对应自然人信息）
		SellerTaxnum          string `json:"sellerTaxnum,omitempty"`          // Y	330100199001010000	20	卖方单位代码/身份证号码（开票方类型为1时，必须与销方税号一致，若反向开具时则为对应自然人信息）
		SellerAddress         string `json:"sellerAddress,omitempty"`         // Y	杭州文一路888号	80	卖方单位/个人地址（开票方类型为1时，必须与销方地址一致，若反向开具时则为对应自然人信息）
		SellerPhone           string `json:"sellerPhone,omitempty"`           // Y	13888888888	20	卖方单位/个人电话（开票方类型为1时，必须与销方电话一致，若反向开具时则为对应自然人信息）
	} `json:"secondHandCarInfo,omitempty"` // 开具二手车销售统一发票才需要传
	InvoiceBuildingInfo struct {
		BuildingAddress string `json:"buildingAddress,omitempty"` // Y	浙江省杭州市西湖区		建筑服务发生地（传对应省市区中文名称--需与行政区划名称一致）
		DetailedAddress string `json:"detailedAddress,omitempty"` // N		120	详细地址（建筑服务发生地+详细地址 总长度最大120字符）
		LandVatItemNo   string `json:"landVatItemNo,omitempty"`   // N		16	土地增值税项目编号
		ItemName        string `json:"itemName,omitempty"`        // Y	宇宙城	80	建筑项目名称
		CrossCityFlag   string `json:"crossCityFlag,omitempty"`   // Y	0	1	跨地（市）标志（0-否 1-是）
	} `json:"invoiceBuildingInfo,omitempty"` //N 具建筑服务特定要素的数电票才需要传（specificFactor 为 3时）；注：数电建筑服务发票 只能有 一条明细 且 规格型号、单位、数量、单价 都不能有值

	// TODO
	// invoiceTravellerTransport	Array	N			实际变量名：invoiceTravellerTransportInfoList 开具旅客运输服务特定要素的数电票时才需要填（specificFactor 为 9时）最多2000行，至少1行
	// realPropertyRentInfo	Object	N			开具不动产经营租赁服务特定要素的数电票才需要传（specificFactor 为 6时）；注：此时 商品只能有 一条明细 且 规格型号、单位都不能有值
	// realPropertySellInfo	Object	N			开具不动产销售特定要素的数电票才需要传（specificFactor 为 5时）；注：此时 商品只能有 一条明细 且 规格型号、单位都不能有值
	// differenceVoucherInfoList	Array	N			数电发票差额扣除凭证列表，开具数电差额征税-差额开票时，必传
	// additionalElementName	string //	N	测试模版	50	附加模版名称（数电电票特有字段，附加模版有值时需要添加附加要素信息列表对象，需要先在电子税局平台维护好模版）
	// additionalElementList	Array	N		10	附加要素信息列表（数电电票特有字段，附加要素信息可以有多个，有值时需要附加模版名称也有值）

}

type InvoiceDetail struct {
	GoodsName           string `json:"goodsName,omitempty"`           //	Y	电脑	90	商品名称（如invoiceLineProperty =1，则此商品行为折扣行，折扣行不允许多行折扣，折扣行必须紧邻被折扣行，商品名称必须与被折扣行一致）
	GoodsCode           string `json:"goodsCode,omitempty"`           //	N	1090511030000000000	19	商品编码（商品税收分类编码开发者自行填写）
	SelfCode            string `json:"selfCode,omitempty"`            //	N	130005426000000000	16	自行编码（可不填）
	WithTaxFlag         string `json:"withTaxFlag,omitempty"`         //	Y	1	1	单价含税标志：0:不含税,1:含税
	Price               string `json:"price,omitempty"`               //	N		16	单价（精确到小数点后8位），当单价(price)为空时，数量(num)也必须为空；(price)为空时，含税金额(taxIncludedAmount)、不含税金额(taxExcludedAmount)、税额(tax)都不能为空
	Num                 string `json:"num,omitempty"`                 //	N		16	数量（精确到小数点后8位，开具红票时数量为负数）
	Unit                string `json:"unit,omitempty"`                //	N	台	20	单位
	SpecType            string `json:"specType,omitempty"`            //	N	y460	40	规格型号
	Tax                 string `json:"tax,omitempty"`                 //	N	0.12	16	税额（精确到小数点后2位），[不含税金额] * [税率] = [税额]；税额允许误差为 0.06。红票为负。不含税金额、税额、含税金额任何一个不传时，会根据传入的单价，数量进行计算，可能和实际数值存在误差，建议都传入
	TaxRate             string `json:"taxRate,omitempty"`             //	Y	0.13	10	税率，注：1、纸票清单红票存在为null的情况；2、二手车发票税率为null或者0
	TaxExcludedAmount   string `json:"taxExcludedAmount,omitempty"`   //	N	0.88	16	不含税金额（精确到小数点后2位）。红票为负。不含税金额、税额、含税金额任何一个不传时，会根据传入的单价，数量进行计算，可能和实际数值存在误差，建议都传入
	TaxIncludedAmount   string `json:"taxIncludedAmount,omitempty"`   //	N	1	16	含税金额（精确到小数点后2位），[不含税金额] + [税额] = [含税金额]，红票为负。不含税金额、税额、含税金额任何一个不传时，会根据传入的单价，数量进行计算，可能和实际数值存在误差，建议都传入
	InvoiceLineProperty string `json:"invoiceLineProperty,omitempty"` //	N	0	1	发票行性质：0,正常行;1,折扣行;2,被折扣行，红票只有正常行
	FavouredPolicyFlag  string `json:"favouredPolicyFlag,omitempty"`  //	N	0	2	优惠政策标识：0,不使用;1,使用; 数电票时： 01：简易征收 02：稀土产品 03：免税 04：不征税 05：先征后退 06：100%先征后退 07：50%先征后退 08：按3%简易征收 09：按5%简易征收 10：按5%简易征收减按1.5%计征 11：即征即退30% 12：即征即退50% 13：即征即退70% 14：即征即退100% 15：超税负3%即征即退 16：超税负8%即征即退 17：超税负12%即征即退 18：超税负6%即征即退
	FavouredPolicyName  string `json:"favouredPolicyName,omitempty"`  //	N	0	50	增值税特殊管理（优惠政策名称）,当favouredPolicyFlag为1时，此项必填 （数电票时为空）
	Deduction           string `json:"deduction,omitempty"`           //	N	0	16	扣除额，差额征收时填写，目前只支持填写一项。 注意：当传0、空或字段不传时，都表示非差额征税；传0.00才表示差额征税：0.00 （数电票暂不支持）
	ZeroRateFlag        string `json:"zeroRateFlag,omitempty"`        //	N	0	1	零税率标识：空,非零税率;1,免税;2,不征税;3,普通零税率；1、当税率为：0%，且增值税特殊管理：为“免税”， 零税率标识：需传“1” 2、当税率为：0%，且增值税特殊管理：为"不征税" 零税率标识：需传“2” 3、当税率为：0%，且增值税特殊管理：为空 零税率标识：需传“3” （数电票时为空）
	DField1             string `json:"dField1,omitempty"`             //	Y		255	业务方明细自定义字段1，本应用只作保存
	DField2             string `json:"dField2,omitempty"`             //	Y		255	业务方明细自定义字段2，本应用只作保存
	DField3             string `json:"dField3,omitempty"`             //	Y		255	业务方明细自定义字段3，本应用只作保存
	DField4             string `json:"dField4,omitempty"`             //	Y		255	业务方明细自定义字段4，本应用只作保存
	DField5             string `json:"dField5,omitempty"`             //	Y		255	业务方明细自定义字段5，本应用只作保存

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
