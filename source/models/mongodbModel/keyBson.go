package mongodbModel

const (
	/* ========= General ========= */
	KeyId        string = "_id"
	KeyIsRemoved string = "is_removed"
	KeyCreateAt  string = "create_at"
	KeyCode      string = "code"
	KeyStatus    string = "status"

	/* ========= Services ========= */
	KeyQuantifyIsUsed string = "quantity_is_used"

	/* ========= Account ========= */
	KeyPhone string = "phone"
	KeyRole  string = "role"

	/* ========= Transaction ========= */
	KeyTransactionId     string = "transaction_id"
	KeyIdServiceProvider string = "service_provider_id"
)
