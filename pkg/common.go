package common

type QueueType int

const (
	QueueTypePod            QueueType = 0
	QueueTypeNode           QueueType = 1
	QueueTypePodMetrics     QueueType = 2
	QueueTypeNodeMetrics    QueueType = 3
)

type AgentQueueItem struct {
	queueType QueueType
	dataItem  interface{}
}