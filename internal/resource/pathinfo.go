package resource

type PathInfo struct {
	Name  string
	IsDir bool
	Size  string
	Path  string
}

type PathInfoList struct {
	CurrentPath string
	Paths       []PathInfo
}
