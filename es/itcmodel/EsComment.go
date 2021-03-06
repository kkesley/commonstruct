package itcmodel

import "time"

//EsComment holds the class in elasticsearch
type EsComment struct {
	ARN       *string `json:"arn,omitempty"`
	ClientARN *string `json:"client_arn,omitempty"`
	Code      *string `json:"code,omitempty"`

	Comment       *string  `json:"comment,omitempty"`
	ReplyTo       []string `json:"reply_to,omitempty"`
	Targets       []string `json:"targets,omitempty"`
	ServiceScopes []string `json:"service_scopes,omitempty"`
	Type          *string  `json:"type,omitempty"`

	Principal      *string    `json:"principal,omitempty"`
	PrincipalScope *string    `json:"principal_scope,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	UpdatedBy      *string    `json:"updated_by,omitempty"`
}
