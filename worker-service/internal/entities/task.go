package entities

type Task struct {
	ID   string `json:"id"`
	Data int    `json:"data"`
}

type TaskResult struct {
	ID     string `json:"task_id"`
	Result int    `json:"result"`
}
