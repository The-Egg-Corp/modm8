package profile

type IMod interface {
	Identifier() string
}

type BaseMod struct {
	identifier string `mapstructure:"identifier"`
}

// Make BaseMod implement IMod interface
func (m BaseMod) Identifier() string {
	return m.identifier
}
