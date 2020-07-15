package request

type CreateConfigEntryRequest struct {
	Key         string
	Value       string
	ConfigMapId uint
}
