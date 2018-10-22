package color

type Code int

const (
	codeFg       Code = 30
	codeBg       Code = 40
	codeBrightFg Code = 90
	codeBrightBg Code = 100
)

const (
	CodeReset Code = iota
	CodeBold
	CodeFaint
	CodeItalic
	CodeUnderline
	CodeBlinkSlow
	CodeBlinkRapid
	CodeReverseVideo
	CodeConcealed
	CodeCrossedOut
)

const (
	CodeBgBlack = iota + codeBg
	CodeBgRed
	CodeBgGreen
	CodeBgYellow
	CodeBgBlue
	CodeBgMagenta
	CodeBgCyan
	CodeBgWhite
)

const (
	CodeFgBlack = iota + codeFg
	CodeFgRed
	CodeFgGreen
	CodeFgYellow
	CodeFgBlue
	CodeFgMagenta
	CodeFgCyan
	CodeFgWhite
)

const (
	CodeBgBrightBlack = iota + codeBrightBg
	CodeBgBrightRed
	CodeBgBrightGreen
	CodeBgBrightYellow
	CodeBgBrightBlue
	CodeBgBrightMagenta
	CodeBgBrightCyan
	CodeBgBrightWhite
)

const (
	CodeFgBrightBlack = iota + codeBrightFg
	CodeFgBrightRed
	CodeFgBrightGreen
	CodeFgBrightYellow
	CodeFgBrightBlue
	CodeFgBrightMagenta
	CodeFgBrightCyan
	CodeFgBrightWhite
)
