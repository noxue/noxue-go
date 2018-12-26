package model



type Blog struct {
	Article `bson:",inline"`
}

func (this *Blog) GetCName() string {
	return "blog"
}