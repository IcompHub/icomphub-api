package enums

type (
	StatusEnum string
)

const (
	StatusActive          StatusEnum = "active"
	StatusInactive        StatusEnum = "inactive"
	StatusWaitingApproval StatusEnum = "waiting_approval"
)
