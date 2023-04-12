package main

type OperatorFactory interface {
	Create() MathOperator
}

type MathOperator interface {
	SetOperandA(int)
	SetOperandB(int)
	ComputeResult() int
}

func (pf *PlusOperatorFactory) Create() MathOperator {
	return &PlusOperator{}
}

type MultiOperatorFactory struct {
}

type BaseOperator struct {
	operatorA, operatorB int
}

func (o *BaseOperator) SetOperatorA(operand int) {
	o.operatorA = operand
}

func (o *BaseOperator) SetOperatorB(operand int) {
	o.operatorB = operand
}

type PlusOperatorFactory struct {
}

func (pf *PlusOperator) ComputeResult() int {
	return pf
}
