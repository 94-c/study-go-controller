package enums

// SortDirection represents sorting direction
type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

// String returns the string representation of SortDirection
func (d SortDirection) String() string {
	return string(d)
}

// IsValid checks if the sort direction is valid
func (d SortDirection) IsValid() bool {
	switch d {
	case SortDirectionAsc, SortDirectionDesc:
		return true
	default:
		return false
	}
}

// CommonStatus represents general status used across multiple domains
type CommonStatus string

const (
	CommonStatusActive   CommonStatus = "active"
	CommonStatusInactive CommonStatus = "inactive"
	CommonStatusDeleted  CommonStatus = "deleted"
)

// String returns the string representation of CommonStatus
func (s CommonStatus) String() string {
	return string(s)
}

// IsValid checks if the common status is valid
func (s CommonStatus) IsValid() bool {
	switch s {
	case CommonStatusActive, CommonStatusInactive, CommonStatusDeleted:
		return true
	default:
		return false
	}
}

// Priority represents priority levels used across domains
type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
	PriorityUrgent Priority = "urgent"
)

// String returns the string representation of Priority
func (p Priority) String() string {
	return string(p)
}

// IsValid checks if the priority is valid
func (p Priority) IsValid() bool {
	switch p {
	case PriorityLow, PriorityMedium, PriorityHigh, PriorityUrgent:
		return true
	default:
		return false
	}
}
