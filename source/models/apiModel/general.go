package apiModel

type Role struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}

type BasicResp struct {
	Code         int    `bson:"code" json:"code"`                  // Mã lỗi. Code 0: Thành công, Code != 0 Thất bại
	Message      string `bson:"message" json:"message"`            // Nội dung thông báo lỗi
	ReturnIndex  int    `bson:"return_index" json:"returnIndex"`   // Thứ tự lỗi
	ReturnAction string `bson:"return_action" json:"returnAction"` // funcion nào đang bị lỗi
}

func (it *BasicResp) Update(errIndex int, errAction string, code int, err error) {
	it.Code = code
	it.Message = err.Error()
	it.ReturnIndex = errIndex
	it.ReturnAction = errAction
}
