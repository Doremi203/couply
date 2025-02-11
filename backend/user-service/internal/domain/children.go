package domain

import desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"

type Children int

const (
	ChildrenUnspecified Children = iota
	ChildrenNo
	ChildrenNotYet
	ChildrenYes
)

func PBToChildren(children desc.Children) Children {
	switch children {
	case desc.Children_CHILDREN_UNSPECIFIED:
		return ChildrenUnspecified
	case desc.Children_CHILDREN_NO:
		return ChildrenNo
	case desc.Children_CHILDREN_NOT_YET:
		return ChildrenNotYet
	case desc.Children_CHILDREN_YES:
		return ChildrenYes
	default:
		return Children(0)
	}
}

func ChildrenToPB(children Children) desc.Children {
	switch children {
	case ChildrenUnspecified:
		return desc.Children_CHILDREN_UNSPECIFIED
	case ChildrenNo:
		return desc.Children_CHILDREN_NO
	case ChildrenNotYet:
		return desc.Children_CHILDREN_NOT_YET
	case ChildrenYes:
		return desc.Children_CHILDREN_YES
	default:
		return desc.Children(0)
	}
}
