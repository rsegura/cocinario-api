package foods

type Usecase interface {
	GetById(id string) (interface{}, error)
	Fetch() (interface{}, error)
	//GetRandom(n int16) ([]interface{}, error)
}
