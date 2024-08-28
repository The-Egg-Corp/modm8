package profile

type Profile struct {
	Name       string
	Favourited bool
	Mods       ModStore
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

func (p *Profile) AddMod(name string, mod any) {
	p.Mods[name] = mod
}

func (p Profile) Delete() {
	Profiles.Delete(p)
}
