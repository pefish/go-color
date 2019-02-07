package console_color

import (
	"fmt"
	"testing"
)

func TestConsoleColorClass_Blue(t *testing.T) {
	fmt.Println(ConsoleColor.Blue(`123124`))
}
