package db

type DbError string

const (
	DbKeyNotFound   DbError = "Key not found"
	DbAlReadyExisti DbError = "Db Key AlReady Exist"
)

func (d DbError) Error() string {
	return string(d)
}
