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

// Fda ...
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

// Kimes ...
type Kimes struct {
	ID            int    `json:"id"`
	Model         string `json:"model"`
	Country       string `json:"country"`
	Manufacture   string `json:"manufacture"`
	Specification string `json:"specification"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	Subcategory   string `json:"subcategory"`
}

// Medica ...
type Medica struct {
	ID                 int    `json:"id"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	CompanyTitle       string `json:"company_title"`
	CompanyDescription string `json:"company_description"`
}

// Pma ...
type Pma struct {
	ID          int    `json:"id"`
	Applicant   string `json:"applicant"`
	GenericName string `json:"generic_name"`
	TradeName   string `json:"trade_name"`
}

// Pas ...
type Pas struct {
	ID                     int    `json:"id"`
	ApplicationName        string `json:"application_name"`
	DeviceName             string `json:"device_name"`
	MedicalSpeciality      string `json:"medical_speciality"`
	StudyName              string `json:"study_name"`
	StudyDesignDescription string `json:"study_design_description"`
}
