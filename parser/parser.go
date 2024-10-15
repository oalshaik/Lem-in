package parser
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct{
	name string
	x int
	y int
}

func parseFile(filename string) (int, map[string]Room, [][2]string, error) {
	file , err
}