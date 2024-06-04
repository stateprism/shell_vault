package providers

import (
	"go.uber.org/zap"
	"reflect"
	"time"
)

type EventType string

const (
	EVENT_TYPE_CREATE EventType = "CREATE"
	EVENT_TYPE_UPDATE EventType = "UPDATE"
	EVENT_TYPE_DELETE EventType = "DELETE"
)

type EventEntity string

const (
	EVENT_ENTITY_USER                   EventEntity = "USER"
	EVENT_ENTITY_CERTIFICATE            EventEntity = "CERTIFICATE"
	EVENT_ENTITY_CERTIFICATE_REQUEST    EventEntity = "CERTIFICATE_REQUEST"
	EVENT_ENTITY_CERTIFICATE_REVOCATION EventEntity = "CERTIFICATE_REVOCATION"
)

type EventTimestampResolution string

const (
	EVENT_TS_RES_SECOND EventEntity = "SECOND"
	EVENT_TS_RES_MS     EventEntity = "MILLIS"
	EVENT_TS_RES_NANO   EventEntity = "NANO"
)

func (EventTimestampResolution) String() string {

}

type AuditEvent struct {
	EventType
	EventEntity
	EventData map[string]any
	Timestap  int64
}

type AuditEvents struct {
	logger *zap.Logger
	Events []AuditEvent
}

// New creates a new instance of AuditEvent with the provided EventType, EventEntity,
// and EventData. The EventData is copied to the new AuditEvent instance, ensuring that
// any pointers are dereferenced.
//
// Note: No pointers are allowed as values of the EventData map, any pointers in that map will be dereferenced after nil checking
//
// Example usage:
//
//	eventType := EventType("login")
//	eventEntity := EventEntity("user")
//	eventData := map[string]any{
//	  "username": "john_doe",
//	  "email": "john@example.com",
//	}
//	auditEvent := AuditEvent{}.New(eventType, eventEntity, eventData)
//
//	// Use the created auditEvent instance...
//
//	fmt.Println(auditEvent.EventType)   // Output: login
//	fmt.Println(auditEvent.EventEntity) // Output: user
//	fmt.Println(auditEvent.EventData)   // Output: map[username:john_doe email:john@example.com]
func (a *AuditEvents) New(eType EventType, eEntity EventEntity, eData map[string]any) error {
	e := &AuditEvent{
		EventType:   eType,
		EventEntity: eEntity,
		EventData:   make(map[string]any),
		Timestap:    time.Now().UnixMilli(),
	}
	i := 0
	for k, v := range eData {
		if isPointer(v) {
			if v, ok := v.(*any); !ok || v == nil {
				a.logger.Fatal("A pointer that cannot be dereferenced was passed into an audit event event data map",
					zap.Any("context", e),
				)
				panic("integrity violation detected on auditing, killing offending request")
			} else {
				e.EventData[k] = *v
			}
		} else {
			e.EventData[k] = v
		}
		i += 1
	}

}

func isPointer(v any) bool {
	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr {
		return true
	} else {
		return false
	}
}
