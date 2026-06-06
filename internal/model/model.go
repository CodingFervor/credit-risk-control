package model

import "time"

type LoanApplication struct {
	ID           int64     `json:"id" db:"id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	Amount       float64   `json:"amount" db:"amount"`
	Term         int       `json:"term" db:"term"` // months
	Purpose      string    `json:"purpose" db:"purpose"`
	Status       string    `json:"status" db:"status"` // pending, approved, rejected, disbursed
	CreditScore  int       `json:"credit_score" db:"credit_score"`
	RiskLevel    string    `json:"risk_level" db:"risk_level"` // low, medium, high, reject
	ApproverID   *int64    `json:"approver_id" db:"approver_id"`
	RejectReason string    `json:"reject_reason" db:"reject_reason"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type CreditAccount struct {
	ID            int64   `json:"id" db:"id"`
	UserID        int64   `json:"user_id" db:"user_id"`
	TotalLimit    float64 `json:"total_limit" db:"total_limit"`
	UsedLimit     float64 `json:"used_limit" db:"used_limit"`
	AvailableLimit float64 `json:"available_limit" db:"available_limit"`
	OverdueCount  int     `json:"overdue_count" db:"overdue_count"`
	Status        string  `json:"status" db:"status"` // active, frozen, closed
}

type RepaymentPlan struct {
	ID          int64     `json:"id" db:"id"`
	LoanID      int64     `json:"loan_id" db:"loan_id"`
	Period      int       `json:"period" db:"period"`
	DueDate     time.Time `json:"due_date" db:"due_date"`
	Principal   float64   `json:"principal" db:"principal"`
	Interest    float64   `json:"interest" db:"interest"`
	Total       float64   `json:"total" db:"total"`
	Status      string    `json:"status" db:"status"` // pending, paid, overdue
	PaidAt      *time.Time `json:"paid_at" db:"paid_at"`
}

type RiskRule struct {
	ID          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Category    string  `json:"category" db:"category"` // identity, credit, behavior, device
	Condition   string  `json:"condition" db:"condition"` // JSON rule expression
	ScoreWeight float64 `json:"score_weight" db:"score_weight"`
	IsEnabled   bool    `json:"is_enabled" db:"is_enabled"`
}

type BlacklistEntry struct {
	ID        int64     `json:"id" db:"id"`
	UserID    *int64    `json:"user_id" db:"user_id"`
	Phone     string    `json:"phone" db:"phone"`
	IDCard    string    `json:"-" db:"id_card"`
	DeviceID  string    `json:"device_id" db:"device_id"`
	Reason    string    `json:"reason" db:"reason"`
	Source    string    `json:"source" db:"source"` // manual, auto, external
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CollectionTask struct {
	ID         int64     `json:"id" db:"id"`
	LoanID     int64     `json:"loan_id" db:"loan_id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	OverdueDays int      `json:"overdue_days" db:"overdue_days"`
	Amount     float64   `json:"amount" db:"amount"`
	AssigneeID *int64    `json:"assignee_id" db:"assignee_id"`
	Status     string    `json:"status" db:"status"` // pending, calling, promised, paid, escalated
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type RiskEvent struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	EventType string    `json:"event_type" db:"event_type"`
	Detail    string    `json:"detail" db:"detail"`
	RiskScore int       `json:"risk_score" db:"risk_score"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
