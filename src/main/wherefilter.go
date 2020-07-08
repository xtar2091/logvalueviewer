package main

import "strconv"

type WhereFilter interface {
	DoFilter(left, right string) bool
}

/*
创建where过滤器

@param op 操作，目前支持以下操作
= ==
!=
<
<=
>
>=
 */
func CreateWhereFilter(op string) WhereFilter {
	var filter WhereFilter = nil
	switch op {
	case "=":
		filter = WhereFilterEqual{}
		break
	case "==":
		filter = WhereFilterEqual{}
		break
	case "!=":
		filter = WhereFilterNotEqual{}
		break
	case "<":
		filter = WhereFilterLess{}
		break
	case "<=":
		filter = WhereFilterLessEqual{}
		break
	case ">":
		filter = WhereFilterMore{}
		break
	case ">=":
		filter = WhereFilterMoreEqual{}
		break
	}
	return filter
}

type WhereFilterEqual struct {

}

func (obj WhereFilterEqual) DoFilter(left, right string) bool {
	leftValue, err := strconv.ParseInt(left, 0, 64)
	if err != nil {
		return left == right
	}

	rightValue, err := strconv.ParseInt(right, 0, 64)
	if err != nil {
		return left == right
	}

	return leftValue == rightValue
}

type WhereFilterNotEqual struct {

}

func (obj WhereFilterNotEqual) DoFilter(left, right string) bool {
	leftValue, err := strconv.ParseInt(left, 0, 64)
	if err != nil {
		return left != right
	}

	rightValue, err := strconv.ParseInt(right, 0, 64)
	if err != nil {
		return left != right
	}

	return leftValue != rightValue
}

type WhereFilterLess struct {

}

func (obj WhereFilterLess) DoFilter(left, right string) bool {
	leftValue, err := strconv.ParseInt(left, 0, 64)
	if err != nil {
		return left < right
	}

	rightValue, err := strconv.ParseInt(right, 0, 64)
	if err != nil {
		return left < right
	}

	return leftValue < rightValue
}

type WhereFilterLessEqual struct {

}

func (obj WhereFilterLessEqual) DoFilter(left, right string) bool {
	leftValue, err := strconv.ParseInt(left, 0, 64)
	if err != nil {
		return left <= right
	}

	rightValue, err := strconv.ParseInt(right, 0, 64)
	if err != nil {
		return left <= right
	}

	return leftValue <= rightValue
}

type WhereFilterMore struct {

}

func (obj WhereFilterMore) DoFilter(left, right string) bool {
	leftValue, err := strconv.ParseInt(left, 0, 64)
	if err != nil {
		return left > right
	}

	rightValue, err := strconv.ParseInt(right, 0, 64)
	if err != nil {
		return left > right
	}

	return leftValue > rightValue
}

type WhereFilterMoreEqual struct {

}

func (obj WhereFilterMoreEqual) DoFilter(left, right string) bool {
	leftValue, err := strconv.ParseInt(left, 0, 64)
	if err != nil {
		return left >= right
	}

	rightValue, err := strconv.ParseInt(right, 0, 64)
	if err != nil {
		return left >= right
	}

	return leftValue >= rightValue
}
