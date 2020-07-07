package color

import "fmt"

//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

//  fmt.Printf(" %c[%d;%d;%dm%s%c[0m ", 0x1B, d, b, f, "阿巴阿巴", f, b, d, 0x1B)
//  %c 对应01xb头尾都有，标记符，没有意义（应该）
//  %d 高亮，就是上面【代码】那段
//  %d 背景
//  %d 前景，字的颜色
//  0m,  0好像是默认的
type Color  int

const (
	WHITE Color = iota + 30
	RED
	GREEN
	YELLOW
	BLUE
	PURPLE
	Magenta
	GRAY
	LIGHT_GRAY
	_
	BLACK
)

const (
	WHITE_BACK = iota +40
	RED_BACK
	GREEN_BACK
	YELLOW_BACK
	BLUE_BACK
	PURPLE_BACK
	Magenta_BACK
	GRAY_BACK
)

func ColorPrint(txt string, c Color) string {
	return  fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, 30, c, txt, 0x1B)
}

func ColorPrint_Back(txt string, c Color, backColor Color) string {
	return  fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, backColor, c, txt, 0x1B)
}
