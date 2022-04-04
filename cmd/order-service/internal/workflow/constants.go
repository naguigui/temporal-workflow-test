package workflow

type Queue string

const (
	CreateOrderTaskQueue Queue = "CREATE_ORDER_TASK_QUEUE"
)

func (q Queue) String() string {
	return string(q)
}
