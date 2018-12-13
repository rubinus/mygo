package models

type FetchPoint struct {
	Kind       string `json:kind`
	UserId     string `json:user_id`
	Points     int    `json:points`
	GitCounts  int    `json:git_counts`
	FacePoster string `json:face_poster`
}
type ErrorFetchPoint struct {
	Error ErrFetchPointBody `json:"error"`
}
type ErrFetchPointBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (r FetchPoint) read() (interface{}, error) {
	return r, nil
}
func (r ErrorFetchPoint) read() (interface{}, error) {
	return r, nil
}
