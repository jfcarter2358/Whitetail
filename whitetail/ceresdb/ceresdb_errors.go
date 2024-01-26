package ceresdb

type CollectionExists struct{}
type DatabaseExists struct{}
type PermitExists struct{}
type RecordExists struct{}
type UserExists struct{}

func (e *CollectionExists) Error() string {
	return "collection exists"
}

func (e *DatabaseExists) Error() string {
	return "database exists"
}

func (e *PermitExists) Error() string {
	return "permit exists"
}

func (e *RecordExists) Error() string {
	return "record exists"
}

func (e *UserExists) Error() string {
	return "user exists"
}
