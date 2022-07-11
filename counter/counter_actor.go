package counter

type incrementReq struct{}

type getReq struct {
	result chan uint
}

type counterActor struct {
	state        counter
	requestQueue chan interface{}
}

func (a *counterActor) run() {
	for reqUntyped := range a.requestQueue {
		switch req := reqUntyped.(type) {
		case *incrementReq:
			a.state.Increment()
		case *getReq:
			req.result <- a.state.Get()
		}
	}
}

func (a *counterActor) Increment() {
	a.requestQueue <- &incrementReq{}
}

func (a *counterActor) Get() uint {
	ch := make(chan uint)
	a.requestQueue <- &getReq{ch}

	return <-ch
}

func NewActorCounter() Counter {
	actor := &counterActor{
		state:        newCounter(),
		requestQueue: make(chan interface{}, 256),
	}
	go actor.run()

	return actor
}
