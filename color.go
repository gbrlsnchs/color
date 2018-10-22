package color

import (
	"fmt"
	"strconv"
	"strings"
)

const escape = "\x1b"

// Color is a text formatter that colors
// text with ANSI escape sequences.
type Color struct {
	seq      string
	disabled bool
}

func New(codes ...Code) Color {
	var c Color
	c.Add(codes...)
	return c
}

func (c *Color) Add(codes ...Code) {
	if length := len(codes); length > 0 {
		seq := make([]string, 0, length+1)
		seq = append(seq, c.seq)
		for _, cd := range codes {
			seq = append(seq, strconv.Itoa(int(cd)))
		}
		c.seq = strings.Join(seq, ";")
	}
}

func (c *Color) SetDisabled(disabled bool) {
	c.disabled = disabled
}

func (c Color) Wrap(v ...interface{}) string {
	return c.text(fmt.Sprint(v...))
}

func (c Color) Wrapf(format string, v ...interface{}) string {
	return c.text(fmt.Sprintf(format, v...))
}

func (c Color) Wrapln(v ...interface{}) string {
	return c.text(fmt.Sprintln(v...))
}

func (c Color) text(s string) string {
	if c.disabled {
		return s
	}
	var bd strings.Builder
	fmt.Fprintf(&bd, "%s[%sm", escape, c.seq)
	bd.WriteString(s)
	fmt.Fprintf(&bd, "%s[%dm", escape, CodeReset)
	return bd.String()
}
