package application

type Command interface {
	CommandName() string
}

type CommandHandler[T Command] interface {
	Handle(ctx CommandContext, cmd T) error
}

type CommandContext interface {
	// Add context methods as needed (e.g., UserID(), TraceID(), etc.)
}

type BaseCommandContext struct {
}

var _ CommandContext = (*BaseCommandContext)(nil)
