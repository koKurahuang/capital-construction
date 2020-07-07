package color

import (
	"fmt"
	"strconv"
	"testing"
)
func TestColorPrint_WHITE(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint(A, WHITE))
}

func TestColorPrint_RED(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint(A, RED))
}

func TestColorPrint_GREEN(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint(A, GREEN))
}

func TestColorPrint_YELLOW(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint(A, YELLOW))
}
func TestColorPrint_BLUE(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint(A, BLUE))
}
func TestColorPrint_PURPLE(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint(A, PURPLE))
}
func TestColorPrint_Magenta(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint(A, Magenta))
}
func TestColorPrint_GRAY(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint(A, GRAY))
}
func TestColorPrint_LIGHT_GTAY(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint(A, LIGHT_GRAY))
}

func TestColorPrint_BLACK(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint(A, BLACK))
}


func TestColorPrint_Back_WHITE(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint_Back(A, WHITE, WHITE_BACK))
}

func TestColorPrint_Back_RED(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint_Back(A, WHITE, RED_BACK))
}

func TestColorPrint_Back_GREEN(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint_Back(A, WHITE, GREEN_BACK))
}

func TestColorPrint_Back_YELLOW(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint_Back(A, WHITE, YELLOW_BACK))
}

func TestColorPrint_Back_BLUE(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint_Back(A, WHITE, BLUE_BACK))
}

func TestColorPrint_Back_PURPLE(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint_Back(A, WHITE, PURPLE_BACK))
}

func TestColorPrint_Back_MagentaK(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint_Back(A, WHITE, Magenta_BACK))
}

func TestColorPrint_Back_GRAY(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint_Back(A, WHITE, GRAY_BACK))
}

func TestColorPrint_BLACK_Back_WHITE(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint_Back(A, BLACK, WHITE_BACK))
}

func TestAll(t *testing.T) {
	for i:= 0 ; i<= 50 ; i++{
		for j:=0; j< 50 ; j ++ {
			s := strconv.Itoa(i) + " " +strconv.Itoa(j)
			fmt.Print(fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, i, j, s, 0x1B))
		}
		fmt.Println()
	}
}

func TestTwo(t *testing.T) {
	A := "阿巴阿巴"
	fmt.Println(ColorPrint_Back(A, BLACK, WHITE_BACK)+ColorPrint_Back(A, BLACK, WHITE_BACK))
}