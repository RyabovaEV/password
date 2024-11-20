package output

import (
	"github.com/fatih/color"
)

// PrintError  печать ошибок
func PrintError(value any) {
	/*intVal, ok := value.(int)
	if ok {
		color.Red("код ошибки: %d", intVal)
		return
	}
	strVal, ok := value.(string)
	if ok {
		color.Red(strVal)
		return
	}*/

	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("код ошибки: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Неизвестный тип ошибки")

	}
	//fmt.Println(value)
}

func sumGeneric[T int | float32](a, b T) T {
	return a + b
}

type List[T any] struct {
	//elements []T
}

func (l *List[T]) addElement() {
	//
}
