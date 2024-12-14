package process

type Symbol struct {
	symbol   string
	position *Position
	state    SymbolState
}

func NewSymbol(symbol string, position *Position) *Symbol {
	return &Symbol{symbol, position, SymbolStateInitial}
}

func (s *Symbol) GetSymbol() string {
	return s.symbol
}

func (s *Symbol) GetPosition() *Position {
	return s.position
}

type SymbolState int

const (
	SymbolStateInitial SymbolState = iota
	SymbolStateToBeUpdated
	SymbolStateChecking
	SymbolStateChecked
	SymbolStateMatched
	SymbolStateNotMatched
)
