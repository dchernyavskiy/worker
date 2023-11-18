package constants

const (
	Pending     string = "Pending"
	Accepted    string = "Accepted"
	InProgress  string = "In Progress"
	Completed   string = "Completed"
	Canceled    string = "Canceled"
	Rejected    string = "Rejected"
	Scheduled   string = "Scheduled"
	OnHold      string = "On Hold"
	Expired     string = "Expired"
	NeedsReview string = "Needs Review"
)

func AllStatuses() []string {
	return []string{Pending, Accepted, InProgress, Completed, Canceled, Rejected, Scheduled, OnHold, Expired, NeedsReview}
}
