package edvmodel

import "time"

//EsClass holds the class in elasticsearch
type EsClass struct {
	ARN        *string `json:"arn,omitempty"`
	ClientARN  *string `json:"client_arn,omitempty"`
	Code       *string `json:"code,omitempty"`
	CustomCode *string `json:"custom_code,omitempty"`
	Name       *string `json:"name,omitempty"`

	IsBoundedByCourse *bool `json:"is_bounded_by_course,omitempty"` //class.code == course.code

	IsRetired        *bool      `json:"is_retired,omitempty"`
	RetiredClassCode *string    `json:"retired_class_code,omitempty"`
	RetiredBatchTime *time.Time `json:"retired_batch_time,omitempty"`

	Courses  []string `json:"courses,omitempty"`
	Students []string `json:"students,omitempty"`

	Principal *string    `json:"principal,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
}