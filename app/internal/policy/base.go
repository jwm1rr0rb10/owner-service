package policy

import (
	"time"
)

// Generation UUIDv4 -- from custom lib we can gen UUIDv7 more faster for index.
type Generator interface {
	GenerateID() string
}

// Clock interface for create time createAt and updateAt. use it in policy layer.
type Clock interface {
	Now() time.Time
}

// Struct for Constructor we can make naming NewController or NewConstructorController.
type BasePolicy struct {
	Generator
	Clock
}

// Constrcutor for initialize in app.go where we add all setups. To our application.
func NewConstructorBasePolicy(generator Generator, clock Clock) *BasePolicy {
	return &BasePolicy{Generator: generator, Clock: clock}
}