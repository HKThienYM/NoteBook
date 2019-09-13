package encoderesponse

//go:generate mockery -name=EncodeResponse -inpkg
// type Encoder interface {
// 	func NewEncoder(io.Writer) *EncodeResponse
// }

type EncodeResponse interface {
	Encode(v interface{}) error
}
