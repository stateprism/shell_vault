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

func (e EventType) String() string {
	switch e {
	case EVENT_TYPE_CREATE:
		return "CREATE"
	case EVENT_TYPE_UPDATE:
		return "UPDATE"
	case EVENT_TYPE_DELETE:
		return "DELETE"
	default:
		return "UNKNOWN"
	}

}

type EventEntity string

const (
	EVENT_ENTITY_USER                   EventEntity = "USER"
	EVENT_ENTITY_CERTIFICATE            EventEntity = "CERTIFICATE"
	EVENT_ENTITY_CERTIFICATE_REQUEST    EventEntity = "CERTIFICATE_REQUEST"
	EVENT_ENTITY_CERTIFICATE_REVOCATION EventEntity = "CERTIFICATE_REVOCATION"
)

func (e EventEntity) String() string {
	switch e {
	case EVENT_ENTITY_USER:
		return "USER"
	case EVENT_ENTITY_CERTIFICATE:
		return "CERTIFICATE"
	case EVENT_ENTITY_CERTIFICATE_REQUEST:
		return "CERTIFICATE_REQUEST"
	case EVENT_ENTITY_CERTIFICATE_REVOCATION:
		return "CERTIFICATE_REVOCATION"
	default:
		return "UNKNOWN"
	}
}

type EventTimestampResolution string

const (
	EVENT_TS_RES_SECOND EventTimestampResolution = "SECOND"
	EVENT_TS_RES_MS     EventTimestampResolution = "MILLIS"
	EVENT_TS_RES_NANO   EventTimestampResolution = "NANO"
)

func (e EventTimestampResolution) String() string {
	switch e {
	case EVENT_TS_RES_SECOND:
		return "SECOND"
	case EVENT_TS_RES_MS:
		return "MILLIS"
	case EVENT_TS_RES_NANO:
		return "NANO"
	default:
		return "UNKNOWN"
	}
}

type AuditEvent struct {
	EventType
	EventEntity
	EventData    map[string]any
	Timestamp    int64
	TsResolution EventTimestampResolution
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
		EventType:    eType,
		EventEntity:  eEntity,
		EventData:    make(map[string]any),
		Timestamp:    time.Now().UnixMilli(),
		TsResolution: EVENT_TS_RES_MS,
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

	a.Events = append(a.Events, *e)
	return nil
}

func isPointer(v any) bool {
	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr {
		return true
	} else {
		return false
	}
}
