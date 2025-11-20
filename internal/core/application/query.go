package application

type Query interface {
	QueryName() string
}

type QueryHandler[T Query, R any] interface {
	Handle(ctx QueryContext, query T) (R, error)
}

type QueryContext interface {
	// Add context methods as needed (e.g., UserID(), TraceID(), etc.)
}

type BaseQueryContext struct {
}

var _ QueryContext = (*BaseQueryContext)(nil)
