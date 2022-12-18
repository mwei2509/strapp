package orchestrator

type Database struct {
	Name string
	Type string // e.g. postgres mysql etc
	Port int64
}
