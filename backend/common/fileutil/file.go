package fileutil

type FileInfo struct {
	Name      string
	Extension *string
}

func NewFileInfo(name string, extension *string) FileInfo {
	return FileInfo{
		Name:      name,
		Extension: extension,
	}
}

func (fi FileInfo) GetCombined() string {
	if fi.Extension == nil {
		return fi.Name
	}

	return fi.Name + *fi.Extension
}
