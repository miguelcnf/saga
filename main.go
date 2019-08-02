package main

import (
	"errors"
	"log"

	"saga/orchestrator"
)

// successful action
func action1() (string, error) {
	return "action 1", nil
}

// successful action
func action2() (string, error) {
	return "action 2", nil
}

// failed action
func action3() (string, error) {
	return "", errors.New("error: action 3")
}

func main() {
	// init rollback engine for the current function
	rollbackEngine := orchestrator.NewRollbackEngine()
	// defer any rollback execution
	defer rollbackEngine.Run()

	// run command
	res1, err := action1()
	// handle error idiomatically
	if err != nil {
		// if error flag the orchestrator to rollback
		rollbackEngine.Flag()

		log.Print(err.Error())
		return
	}
	// if successful register the rollback step
	rollbackEngine.Step(func(args ...interface{}) {
		val := args[0].(int)
		log.Print("rollback: action ", val)
	}, 1)

	log.Print(res1)

	res2, err := action2()
	if err != nil {
		rollbackEngine.Flag()

		log.Print(err.Error())
		return
	}
	rollbackEngine.Step(func(args ...interface{}) {
		val := args[0].(int)
		log.Print("rollback: action ", val)
	}, 2)

	log.Print(res2)

	res3, err := action3()
	if err != nil {
		rollbackEngine.Flag()

		log.Print(err.Error())
		return
	}
	rollbackEngine.Step(func(args ...interface{}) {
		val := args[0].(int)
		log.Print("rollback: action ", val)
	}, 3)

	log.Print(res3)
}
