package util

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/bovinae/common/comm_error"
	"github.com/pkg/errors"
)

type CondArgs []string

func (a *CondArgs) Set(val string) error {
	if len(val) == 0 {
		return nil
	}
	*a = CondArgs(strings.Split(val, ","))
	return nil
}

func (a *CondArgs) String() string {
	// *a = CondArgs(strings.Split("default is me", ","))
	return "It's none of my business"
}

type Args struct {
	TableName string
	ColName   string
	BeginId   int64
	EndId     int64
	Condition string
}

type Column struct {
	ColName string   `json:"colName"`
	Values  []string `json:"values"`
}

type CsvClient struct {
}

func NewCsvClient() *CsvClient {
	return &CsvClient{}
}

func (c *CsvClient) QueryCsv(ctx context.Context, args1, args2 Args) ([][]string, [][]string, error) {
	if args1.TableName == args2.TableName {
		data, err := c.QuerySameTable(ctx, args1, args2)
		return data, nil, err
	}

	data1, err := c.QueryOneTable(ctx, args1)
	if err != nil {
		return nil, nil, err
	}
	data2, err := c.QueryOneTable(ctx, args2)
	if err != nil {
		return nil, nil, err
	}
	return data1, data2, nil
}

func (c *CsvClient) QueryOneTable(ctx context.Context, args Args) ([][]string, error) {
	data, err := c.ReadCsvFile(ctx, args.TableName)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.Wrap(comm_error.ErrNoData, "QueryOneTable")
	}

	// filter column
	index, err := c.getColumnIndex(ctx, args.ColName, data[0])
	if err != nil {
		return nil, err
	}
	return c.filterData(ctx, data, index), nil
}

func (c *CsvClient) QuerySameTable(ctx context.Context, args1, args2 Args) ([][]string, error) {
	data, err := c.ReadCsvFile(ctx, args1.TableName)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.Wrap(comm_error.ErrNoData, "QuerySameTable")
	}

	if len(args1.ColName) == 0 && len(args2.ColName) == 0 {
		return data, nil
	}

	index1, err := c.getColumnIndex(ctx, args1.ColName, data[0])
	if err != nil {
		return nil, err
	}
	index2, err := c.getColumnIndex(ctx, args2.ColName, data[0])
	if err != nil {
		return nil, err
	}

	// index1 and index2:
	// case1: both >= 0
	if index1 >= 0 && index2 >= 0 {
		return c.filterData(ctx, data, index1, index2), nil
	}
	// case2: at least one index >= 0
	if index1 >= 0 {
		return c.filterData(ctx, data, index1), nil
	}
	return c.filterData(ctx, data, index2), nil
}

func (c *CsvClient) filterData(ctx context.Context, data [][]string, index ...int) [][]string {
	filtered := make([][]string, 0, len(data))
	for i := 0; i < len(data); i++ {
		line := make([]string, 0, len(index))
		for _, j := range index {
			line = append(line, data[i][j])
		}
		filtered = append(filtered, line)
	}
	return filtered
}

func (c *CsvClient) getColumnIndex(ctx context.Context, column string, header []string) (int, error) {
	if len(header) == 0 {
		return -1, errors.Wrap(comm_error.ErrNoData, "empty header when getColumnIndex")
	}
	if len(column) == 0 {
		return -1, nil
	}

	for i := 0; i < len(header); i++ {
		if column == header[i] {
			return i, nil
		}
	}

	return -1, errors.Wrap(comm_error.ErrGetColumnIndex, "get none index when getColumnIndex")
}

func (c *CsvClient) ReadCsvFile(ctx context.Context, fileName string) ([][]string, error) {
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var values [][]string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.Wrap(err, "read csv file failed")
		}
		values = append(values, line)
	}

	return values, nil
}

func (c *CsvClient) WriteCsvFile(ctx context.Context, fileName string, values [][]any) error {
	if len(values) == 0 {
		return nil
	}

	fileHandle, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return errors.Wrap(err, "os.OpenFile failed")
	}
	defer fileHandle.Close()

	// NewWriter 默认缓冲区大小是 4096
	buf := bufio.NewWriterSize(fileHandle, 4096*4096)

	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[i]); j++ {
			buf.WriteString(fmt.Sprintf(`"%v"`, values[i][j]))
			if j < len(values[i])-1 {
				buf.WriteString(",")
			}
		}
		buf.WriteString("\n")

		if i > 0 && i%10000 == 0 {
			if err := buf.Flush(); err != nil {
				return errors.Wrap(err, "flush csv file")
			}
		}
	}

	if err := buf.Flush(); err != nil {
		return errors.Wrap(err, "flush csv file")
	}
	return nil
}
