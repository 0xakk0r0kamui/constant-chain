package transaction

const (
	// TxVersion is the current latest supported transaction version.
	TxVersion = 1
)

const (
	NoSort            = iota
	SortByAmount
	SortByCreatedTime
)
