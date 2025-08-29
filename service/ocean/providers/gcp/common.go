package gcp

import "github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"

// region Tag

type Tag struct {
	Key   *string `json:"tagKey,omitempty"`
	Value *string `json:"tagValue,omitempty"`

	forceSendFields []string
	nullFields      []string
}

func (o Tag) MarshalJSON() ([]byte, error) {
	type noMethod Tag
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Tag) SetKey(v *string) *Tag {
	if o.Key = v; o.Key == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

func (o *Tag) SetValue(v *string) *Tag {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

type Filters struct {
	ExcludeFamilies []string `json:"excludeFamilies,omitempty"`
	IncludeFamilies []string `json:"includeFamilies,omitempty"`
	MaxMemoryGiB    *float64 `json:"maxMemoryGiB,omitempty"`
	MaxVcpu         *int     `json:"maxVcpu,omitempty"`
	MinMemoryGiB    *float64 `json:"minMemoryGiB,omitempty"`
	MinVcpu         *int     `json:"minVcpu,omitempty"`

	forceSendFields []string
	nullFields      []string
}

func (o Filters) MarshalJSON() ([]byte, error) {
	type noMethod Filters
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Filters) SetExcludeFamilies(v []string) *Filters {
	if o.ExcludeFamilies = v; o.ExcludeFamilies == nil {
		o.nullFields = append(o.nullFields, "ExcludeFamilies")
	}
	return o
}

func (o *Filters) SetIncludeFamilies(v []string) *Filters {
	if o.IncludeFamilies = v; o.IncludeFamilies == nil {
		o.nullFields = append(o.nullFields, "IncludeFamilies")
	}
	return o
}

func (o *Filters) SetMaxMemoryGiB(v *float64) *Filters {
	if o.MaxMemoryGiB = v; o.MaxMemoryGiB == nil {
		o.nullFields = append(o.nullFields, "MaxMemoryGiB")
	}
	return o
}

func (o *Filters) SetMaxVcpu(v *int) *Filters {
	if o.MaxVcpu = v; o.MaxVcpu == nil {
		o.nullFields = append(o.nullFields, "MaxVcpu")
	}
	return o
}

func (o *Filters) SetMinMemoryGiB(v *float64) *Filters {
	if o.MinMemoryGiB = v; o.MinMemoryGiB == nil {
		o.nullFields = append(o.nullFields, "MinMemoryGiB")
	}
	return o
}

func (o *Filters) SetMinVcpu(v *int) *Filters {
	if o.MinVcpu = v; o.MinVcpu == nil {
		o.nullFields = append(o.nullFields, "MinVcpu")
	}
	return o
}

// endregion
