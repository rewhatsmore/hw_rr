package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if in == nil {
		panic("nil in")
	}

	out := in

	for _, stage := range stages {
		bi := make(Bi)
		go chanPipe(out, bi, done)
		out = stage(bi)
	}
	return out
}

func chanPipe(in In, out Bi, done In) {
	defer close(out)
	for {
		select {
		case v, ok := <-in:
			if !ok {
				return
			}
			out <- v
		case <-done:
			return
		}
	}
}
