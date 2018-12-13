package logs

type Logger interface {
	Info()
}

var TraceChan = make(chan Logger)

func init() {
	go func() {
		for v := range TraceChan {
			v.Info()
		}
	}()
}
