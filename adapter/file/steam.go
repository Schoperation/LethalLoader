package file

type SteamDao struct {
	path string
}

func NewSteamDao() SteamDao {
	return SteamDao{}
}

func (dao *SteamDao) SetPath(path string) {
	dao.path = path
}
