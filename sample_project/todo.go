package sample_project

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	// "os"
	"time"
	"github.com/alexeyco/simpletable"
)


type item struct {
	Task string
	Done bool
	CreatedAt time.Time
	CompletedAt time.Time
}


type Todos []item


func (t *Todos) Add(task string) error {
	if len(task) == 0 {
		return errors.New("invalid string")
	}

	i := item {
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, i)
	return nil
}


func (t *Todos) Complete(index int) error {
	if index <= 0 || index > len(*t) {
		return errors.New("invalid values")
	}
	(*t)[index-1].Done = true
	(*t)[index-1].CompletedAt = time.Now()
	return nil
}


func (t *Todos) Delete(index int) error {
	if index <= 0 || index > len(*t) {
		return errors.New("invalid index")
	}
	*t = append((*t)[:index-1], (*t)[index:]...)
	return nil
}


func (t *Todos) Print() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			{Align: simpletable.AlignCenter, Text: "CompletedAt"},
		},
	}

	for i, item := range *t {
		taskName := blue(item.Task)
		done := "\u274C"
		if (item.Done) {
			taskName = green(item.Task)
			done = "\u2705"
		}

		r := []*simpletable.Cell {
			{Text: fmt.Sprintf("%d", i+1)},
			{Text: fmt.Sprintf("%v", taskName)},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%v", done)},
			{Text: fmt.Sprintf("%v", item.CreatedAt.Format("2006-01-02 15:04:05"))},
			{Text: fmt.Sprintf("%v", item.CompletedAt.Format("2006-01-02 15:04:05"))},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			// {Align: simpletable.AlignRight, Span: 4, Text: red("Task pending: ")},
			{Align: simpletable.AlignCenter, Span:5, Text: red(fmt.Sprintf("You have %d pending task!", (*t).countPending()))},
		},
	}


	table.SetStyle(simpletable.StyleUnicode)
	table.Println()

	fmt.Println()
}

func (t *Todos) Load(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return err
	}

	err = json.Unmarshal(data, t)
	if err != nil {
		return err
	}

	return nil
}


func (t *Todos) Store(fileName string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}


func (t *Todos) countPending() int {
	sum := 0
	for _, value := range *t {
		if (!value.Done) {
			sum ++
		}
	}
	return sum
}
