package base

type ResponseEntity struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
