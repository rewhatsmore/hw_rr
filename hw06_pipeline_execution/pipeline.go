package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	var out Out = in
	for _, stage := range stages {
		bi := make(Bi)
		go func(in In) {
			defer close(bi)
			for v := range in {
				select {
				case bi <- v:
				case <-done:
					return
				}
			}
		}(out)
		out = stage(bi)
	}
	return out
}
