package transactionDto

type CreateTransactionRequest struct {
	StartDate int    `json:"start_date"`
	DueDate   int    `json:"due_date"`
	UserID    int    `json:"user_id"`
	Status    string `json:"status" gorm:"type: VARCHAR(25)"`
}

type UpdateTransactionRequest struct {
	StartDate int    `json:"start_date" form:"start_date"`
	DueDate   int    `json:"due_date" form:"due_date"`
	UserID    int    `json:"user_id" form:"user_id"`
	Status    string `json:"status" form:"status" gorm:"type: VARCHAR(25)"`
}
