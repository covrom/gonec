package core

// VMChan - канал для передачи любого типа вирт. машины
type VMChan chan VMValuer

func (x VMChan) vmval() {}

func (x VMChan) Interface() interface{} {
	return x
}

func (x VMChan) Send(v VMValuer) {
	x <- v
}

func (x VMChan) Recv() (VMValuer, bool) {
	rv, ok := <-x
	return rv, ok
}

func (x VMChan) TrySend(v VMValuer) (ok bool) {
	select {
	case x <- v:
		ok = true
	default:
		ok = false
	}
	return
}

func (x VMChan) TryRecv() (v VMValuer, ok bool, notready bool) {
	select {
	case v, ok = <-x:
		notready = false
	default:
		ok = false
		notready = true
	}
	return
}

func (x VMChan) Close() { close(x) }

func (x VMChan) Закрыть() { close(x) }
