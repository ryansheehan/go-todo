package shared

type AppState struct {
	Count uint64
	Name  string
}

func InitAppState(count uint64, name string) *AppState {
	return &AppState{
		Count: count,
		Name:  name,
	}
}
