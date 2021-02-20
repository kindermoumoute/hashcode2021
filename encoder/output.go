package encoder

import (
	"fmt"
	"github.com/kindermoumoute/hashcode2021/model"
)

type Output struct {
	Libraries []*model.Library
}

func EncodeOutput(o Output) string {
	output := fmt.Sprintf("%d\n", len(o.Libraries))

	for _, lib := range o.Libraries {
		output += fmt.Sprintf("%d %d\n", lib.ID, len(lib.GetAssignedBooks()))
		output += fmt.Sprintf("%s\n", lib.GetScannedBooksID())
	}

	return output
}
