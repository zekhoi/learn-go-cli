package utils

import (
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/zekhoi/learn-go-cli/pkg/entity"
)

func CreateTable(list []entity.Shorten) *simpletable.Table {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Original URL"},
			{Align: simpletable.AlignCenter, Text: "Shorten URL"},
			{Align: simpletable.AlignCenter, Text: "Date"},
		},
	}
	for _, row := range list {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: fmt.Sprint(row.ID)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf(row.OriginalUrl)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf(row.ShortenUrl)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf(row.CreatedAt.String())},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}
	return table
}
