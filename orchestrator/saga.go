package orchestrator

type Rollback interface {
	Step(func(...interface{}), ...interface{})
	Flag()
	Run()
}

type Step struct {
	function func(...interface{})
	args     []interface{}
}

type RollbackEngine struct {
	steps []Step
	flag  bool
}

func NewRollbackEngine() *RollbackEngine {
	return &RollbackEngine{}
}

func (r *RollbackEngine) Step(f func(...interface{}), args ...interface{}) {
	step := Step{
		function: f,
		args:     args,
	}

	if r.steps == nil {
		r.steps = []Step{}
	}

	r.steps = append(r.steps, step)
}

func (r *RollbackEngine) Flag() {
	r.flag = true
}

func (r *RollbackEngine) Run() {
	if r.flag && len(r.steps) > 0 {
		for i := len(r.steps) - 1; i >= 0; i-- {
			r.steps[i].function(r.steps[i].args...)
		}
	}
}
