package encoder

import (
	"fmt"
	"strconv"
	"strings"
)

func EncodeOutput(o Output) string {
	output := fmt.Sprintf("%d\n", len(o.Libraries))

	for _, lib := range o.Libraries {
		output += fmt.Sprintf("%d %d\n", lib.ID, len(lib.ScannedBookIDs))
		bookIDStrings := []string(nil)
		for _, bookID := range lib.ScannedBookIDs {
			bookIDStrings = append(bookIDStrings, strconv.Itoa(bookID))
		}
		output += strings.Join(bookIDStrings, " ")
	}

	return output
}

type Output struct {
	Libraries []*OutputLibrary
}

type OutputLibrary struct {
	ID             int
	ScannedBookIDs []int
}
