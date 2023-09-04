package beerprotov1

func (x *ExpressionTree) Evaluate(variables map[string]float64) float64 {
	return x.Expression.evaluate(variables)
}

func (x *BinaryExpression) evaluate(variables map[string]float64) float64 {

	var left, right float64

	switch x.Left.(type) {
	case *BinaryExpression_BinaryLeft:
		exp := x.GetBinaryLeft()
		left = exp.evaluate(variables)
	case *BinaryExpression_UnaryLeft:
		exp := x.GetUnaryLeft()
		left = exp.evaluate(variables)
	case *BinaryExpression_ParameterLeft:
		exp := x.GetParameterLeft()
		left = exp.evaluate(variables)

	}

	switch x.Right.(type) {
	case *BinaryExpression_BinaryRight:
		exp := x.GetBinaryRight()
		right = exp.evaluate(variables)
	case *BinaryExpression_UnaryRight:
		exp := x.GetUnaryRight()
		right = exp.evaluate(variables)
	case *BinaryExpression_ParameterRight:
		exp := x.GetParameterRight()
		right = exp.evaluate(variables)
	}

	switch x.Operator {
	case BinaryArithmetic_BINARY_ARITHMETIC_ADDITION:
		return left + right
	case BinaryArithmetic_BINARY_ARITHMETIC_SUBTRACTION:
		return left - right
	case BinaryArithmetic_BINARY_ARITHMETIC_MULTIPLICATION:
		return left * right
	case BinaryArithmetic_BINARY_ARITHMETIC_DIVISION:
		return left / right
	}

	return 0
}

func (x *UnaryExpression) evaluate(variables map[string]float64) float64 {
	return x.Double
}

func (x *ParameterExpression) evaluate(variables map[string]float64) float64 {
	if v, ok := variables[x.Parameter]; ok {
		return v
	}
	return 0
}
