package recipes

type Usecase interface {
	Fetch() (interface{}, error)
	//GetRandom(n int16) ([]interface{}, error)
}
