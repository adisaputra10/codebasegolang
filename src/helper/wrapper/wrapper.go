package wrapper

// Result property
type Result struct {
	Err         error
	Data        interface{}
	Message     string
	PartialMeta *PartialMeta
}

// PartialMeta property
type PartialMeta struct {
	TotalDataOnAppear int `json:"totalDataOnPage" xml:"totalDataOnPage"`
	TotalDataAtAll    int `json:"totalData" xml:"totalData"`
	TotalPart         int `json:"totalPage" xml:"totalPage"`
	OnPart            int `json:"page" xml:"page"`
}

// HTTPResult property
type HTTPResult struct {
	Success bool         `json:"success" xml:"success"`
	Data    interface{}  `json:"data" xml:"data"`
	Message string       `json:"message" xml:"message"`
	Code    int          `json:"code" xml:"code"`
	Meta    *PartialMeta `json:"meta,omitempty" xml:"meta,omitempty"`
}
