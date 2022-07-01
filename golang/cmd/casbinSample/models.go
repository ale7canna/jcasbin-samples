package utils

type CustomSubject struct {
	Name    string
	IsAdmin bool
}

func (cs *CustomSubject) GetCacheKey() string {
	return cs.Name
}
