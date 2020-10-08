package parameter

type Request struct {
	Url      string                 `json:"url"`
	Data     map[string]interface{} `json:"data"`
	Method   string                 `json:"method"`
	Duration int64                  `json:"duration"`
}
