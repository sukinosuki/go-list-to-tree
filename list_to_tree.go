package tree

type Tree[T any] interface {
	GetId() uint
	GetPid() uint
	SetChildren(arr []T)
}

func ListToTree[T Tree[T]](list []T) []T {
	var tree []T
	idMap := make(map[uint][]T)

	for _, item := range list {
		if arr, ok := idMap[item.GetPid()]; ok {
			idMap[item.GetPid()] = append(arr, item)
		} else {
			idMap[item.GetPid()] = []T{item}
		}
	}

	for _, item := range list {
		if arr, ok := idMap[item.GetId()]; ok {
			item.SetChildren(arr)
		} else {
			//tree = append(tree, item)
		}

		if item.GetPid() == 0 {
			tree = append(tree, item)
		}
	}

	return tree
}
