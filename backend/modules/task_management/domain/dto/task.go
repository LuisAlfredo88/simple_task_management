package dto

type TaskRecord struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Color       string `json:"color"`
	CreatedBy   string `json:"createdBy"`
	CreatedById string `json:"createdById"`
	AssignedTo  string `json:"assignedTo"`
}
