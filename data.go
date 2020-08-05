package diff

import "fmt"

const (
	EditAdd = "add"
	EditDel = "del"
	EditEq  = "eq"
)

var Tags = map[string]string{
	EditEq:  " ",
	EditDel: "-",
	EditAdd: "+",
}

type EditPoint struct {
	PrevX int
	PrevY int
	X     int
	Y     int
}

type EditDiff struct {
	OpType string
	Old    string
	New    string
}

func (d EditDiff) Text() string {
	switch d.OpType {
	case EditAdd:
		return d.New
	case EditDel:
		return d.Old
	default:
		return d.Old
	}
}

func (d EditDiff) String() string {
	tag := Tags[d.OpType]
	return fmt.Sprintf("%v %v", tag, d.Text())
}
