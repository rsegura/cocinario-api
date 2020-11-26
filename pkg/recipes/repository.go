package recipes

type Repository interface {
	Fetch() (interface{}, error)
	//GetRandom(n int16) ([]interface{}, error)
}
