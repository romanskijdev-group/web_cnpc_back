package typescore

type Response struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data,omitempty"`
	Error      *WEvent     `json:"error,omitempty"`
	Count      int         `json:"count,omitempty"`
	TotalCount *uint64     `json:"total_count,omitempty"`
}
