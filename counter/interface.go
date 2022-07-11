package counter

type Counter interface {
	Increment()
	Get() uint
}
