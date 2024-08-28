package profile

var Profiles = make(ProfileStore)

type ProfileStore map[string]Profile
type ModStore map[string]any

func (store *ProfileStore) Delete(p Profile) {
	delete(*store, p.Name)
}
