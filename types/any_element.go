package types

type AnyElement interface {
	SameType(elem AnyElement) bool
	TypeName() string
}
