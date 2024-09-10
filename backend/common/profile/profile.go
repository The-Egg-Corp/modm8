package profile

type Profile struct {
	Name       string   `json:"name" mapstructure:"name"`
	Favourited bool     `json:"favourited" mapstructure:"favourited"`
	Mods       ModStore `json:"mods" mapstructure:"mods"`
}

func NewProfile(name string) Profile {
	return Profile{
		Name:       name,
		Favourited: false,
		Mods:       ModStore{},
	}
}

func (p *Profile) Favourite() {
	p.Favourited = true
}

func (p *Profile) AddMod(name string, mod IMod) {
	p.Mods[name] = mod
}
