package gatewayModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type DTOSendInfoServiceReq struct {
	TransactionId     primitive.ObjectID `json:"transactionId"`
	ServiceProviderId primitive.ObjectID `json:"serviceProviderId"` // id nguoi cung cap dich vu
}

type DTOSendInfoServiceResp struct {
	Code    int    `json:"code"`    // Mã lỗi. Code 0: Thành công, Code != 0 Thất bại
	Message string `json:"message"` // Nội dung thông báo lỗi
}
