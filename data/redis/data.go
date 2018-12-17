package redis

type Data struct{}

func New() *Data {
	return &Data{}
}

func (d *Data) UserAccountExist(account string) (bool, error) {
	return true, nil
}
