package enums

import ()

type TodoStatus string
type TodoType string

const (
	TodoStatusDone       TodoStatus = "done"
	TodoStatusNotDone    TodoStatus = "not done"
	TodoTypePersonal     TodoType   = "Personal"
	TodoTypeProfessional TodoType   = "Professional"
)

var statusName = map[TodoStatus]string{
	TodoStatusDone:    "done",
	TodoStatusNotDone: "not done",
}

var typeName = map[TodoType]string{
	TodoTypePersonal:     "Personal",
	TodoTypeProfessional: "Professional",
}

func ParseTodoStatus(s string) (TodoStatus, bool) {
	for ts, name := range statusName {
		if name == s {
			return ts, true
		}
	}
	return "", false
}

func ParseTodoType(s string) (TodoType, bool) {
	for tt, name := range typeName {
		if name == s {
			return tt, true
		}
	}
	return "", false
}

func (ts TodoStatus) String() string {
	return statusName[ts]
}

func (tt TodoType) String() string {
	return typeName[tt]
}
