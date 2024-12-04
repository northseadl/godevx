package taskx

import (
	"strings"
)

type KeyManager struct {
	namespace string
}

func NewKeyManager(namespace string) *KeyManager {
	if namespace == "" {
		namespace = DefaultNamespace
	}
	return &KeyManager{namespace: namespace}
}

func (km *KeyManager) buildKey(parts ...string) string {
	elements := append([]string{km.namespace}, parts...)
	return strings.Join(elements, KeySeparator)
}

func (km *KeyManager) TaskLockKey(taskID string) string {
	return km.buildKey("locks", "tasks", taskID)
}

func (km *KeyManager) WorkerHeartbeatKey(workerID string) string {
	return km.buildKey("workers", "heartbeat", workerID)
}

func (km *KeyManager) TaskQueueKey() string {
	return km.buildKey("queues", "tasks")
}

func (km *KeyManager) TaskStatusKey(taskID string) string {
	return km.buildKey("status", "tasks", taskID)
}
