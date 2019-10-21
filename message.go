package camunda_client_go

import camunda_client_go "github.com/citilinkru/camunda-client-go"

// Message a client for Message API
type Message struct {
	client *Client
}

type CorrelationKeySet struct {
	Value interface{}
	Type  string
}

// ReqMessageCreate a request to message create
type ReqMessageCreate struct {
	MessageName       string                                 `json:"messageName"`
	BusinessKey       *string                                `json:"businessKey,omitempty"`
	TenantId          *string                                `json:"tenantId,omitempty"`
	WithoutTenantId   *bool                                  `json:"withoutTenantId,omitempty"`
	ProcessInstanceId *string                                `json:"processInstanceId,omitempty"`
	CorrelationKeys   *map[string]CorrelationKeySet          `json:"correlationKeys,omitempty"`
	ProcessVariables  *map[string]camunda_client_go.Variable `json:"processVariables,omitempty"`
	All               *bool                                  `json:"all,omitempty"`
}

func (m *Message) Create(reqMessageCreate ReqMessageCreate) error {
	_, err := m.client.doPostJson("/message", map[string]string{}, reqMessageCreate)

	return err
}
