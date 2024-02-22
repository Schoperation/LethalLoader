package file

type SteamDao struct {
	path string
}

func NewSteamDao(path string) SteamDao {
	return SteamDao{
		path: path,
	}
}
