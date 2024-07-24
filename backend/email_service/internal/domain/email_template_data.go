package domain

type EmailType int

const (
	ConfirmLink EmailType = iota
)

func (e EmailType) String() string {
	return [...]string{"ConfirmLink"}[e]
}

type EmailTemplateData struct {
	Type EmailType
	Data map[string]interface{}
}
