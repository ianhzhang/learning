package container_api

import(
)

type List []interface{}
type Element interface{}

func (list *List) Append(elem Element) {
	*list = append(*list, elem)
}

func (list *List) GetByIndex(index int) Element{
	return (*list)[index]
}

