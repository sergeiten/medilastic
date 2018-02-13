package medilastic

// PermitStatus is a structure used for serialize/deserialize data in ElasticSearch
type PermitStatus struct {
	ID             int    `json:"id"`
	Prduct         string `json:"prduct"`
	Entrps         string `json:"entrps"`
	PrductPrmisnNo string `json:"prduct_prmisn_no"`
	MeaClassNo     string `json:"mea_class_no"`
	TypeName       string `json:"type_name"`
	UsePurps       string `json:"use_purps"`
}

type Fda struct {
	ID                int    `json:"id"`
	BrandName         string `json:"brand_name"`
	CompanyName       string `json:"company_name"`
	DeviceDescription string `json:"device_description"`
	GmdnPtName        string `json:"gmdn_pt_name"`
	GmdnPtDefinition  string `json:"gmdn_pt_definition"`
	ProductCode       string `json:"product_code"`
	ProductCodeName   string `json:"product_code_name"`
}
