package mysql

type UserModel struct {
}

func UserExist(uid int64, account string) (bool, error) {
	return true, nil
}
