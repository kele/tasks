package tasks

type Service interface {
	List(ListRequest) ListResponse

	BatchGet(BatchGetRequest) BatchGetResponse
}

type ListRequest struct {
	Category string
}

type ListResponse struct {
	IDs []int
}

type BatchGetRequest struct {
	IDs []int
}

type BatchGetResponse struct {
	Tasks []Task
}

type Task struct {
	ID       int
	Category string
	Content  string
}
