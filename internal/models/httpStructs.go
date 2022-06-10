package models

type NewPaymentRequest struct {
	UserId    int    `json:"user_id" example:"5"`
	UserEmail string `json:"user_email" example:"secret@mail.ru"`
	Amount    int    `json:"amount" example:"6535"`
	Currency  string `json:"currency" example:"RUB"`
}

type NewPaymentResponse struct {
	Payment Payment `json:"payment"`
}

type ErrorResponse struct {
	Message string `json:"error" example:"incorrect user ID"`
}

type ChangeStatusRequest struct {
	Status string `json:"status" example:"SUCCESS"`
}

type PaymentStatusResponse struct {
	Status string `json:"status" example:"ERROR"`
}

type PaymentsResponse struct {
	Payments []Payment `json:"payments"`
}
