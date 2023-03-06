package errors

type Error string

var (
	ErrEmptyMempoolTx Error = "mempool transaction is empty"
)

func (e Error) Error() string {
	return string(e)
}
