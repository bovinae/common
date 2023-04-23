package util

import (
	"context"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWriteCsvFile(t *testing.T) {
	cli := NewCsvClient()

	Convey("TestWriteCsvFile", t, func() {
		Convey("TestWriteCsvFile", func() {
			fileName := "./test.csv"
			values := [][]any{{"07974", "EH4 1DT", "01202", "EH32 1DT", "W1B 1JH", "W1Y 1JH"}, {"07974", "EH4 1DT", "01202", "EH32 1DT", "W1B 1JH", "W1Y 1JH"}}
			err := cli.WriteCsvFile(context.Background(), fileName, values)
			So(err, ShouldEqual, nil)
			ExecOsCommond(`C:\Program Files\Git\git-bash.exe`, "-c", fmt.Sprintf("rm -rf %v", fileName))
		})
	})
}
