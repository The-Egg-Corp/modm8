package appcore

type AppCore struct {
	Settings    *AppSettings
	Persistence *Persistence
	Utils       *Utils
}

func New() *AppCore {
	core := &AppCore{
		Settings:    NewSettings(),
		Persistence: NewPersistence(),
		Utils:       NewUtils(),
	}

	core.Settings.Apply()
	return core
}
