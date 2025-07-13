package appctx

type AppCore struct {
	Settings    *AppSettings
	Persistence *Persistence
	Utils       *Utils
}

func NewAppCore() *AppCore {
	return &AppCore{
		Settings:    NewSettings(),
		Persistence: NewPersistence(),
		Utils:       NewUtils(),
	}
}
