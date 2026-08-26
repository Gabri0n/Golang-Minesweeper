// User interface

package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func draw_board(table *tview.Table, matrix [][]Cell) {
	for x, row := range matrix {
		for y, val := range row {
			color := tcell.ColorWhite

			var text string
			switch {
			case val.flagged: // Flag
				text = "?"
				color = tcell.ColorPurple

			case val.bomb && !val.revealed: // Unrevealed Bomb
				text = "#"
				color = tcell.ColorWhite

			case val.bomb: // Revealed Bomb
				text = "*"
				color = tcell.ColorRed

			case val.revealed && val.adjacentBombs > 0: // Adjacent Bomb Count
				text = strconv.Itoa(val.adjacentBombs)
				color = tcell.ColorYellow

			case val.revealed: // Revealed Empty Cell
				text = " "
				color = tcell.ColorBlack

			default: // Default Cell
				text = "#"
				color = tcell.ColorWhite
			}

			table.SetCell(x, y, tview.NewTableCell(text).SetTextColor(color).SetAlign(tview.AlignCenter).SetMaxWidth(1))

		}
	}
}

func render_tui(matrix [][]Cell) {

	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(false).SetSeparator(' ')

	table.SetSelectedStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorRed))
	table.SetSelectable(true, true)

	draw_board(table, matrix)
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		x, y := table.GetSelection()

		switch event.Rune() {

		case 'f': // flag
			if !matrix[x][y].revealed {
				matrix[x][y].flagged = true
				draw_board(table, matrix)
			}
			return nil

		case ' ':
			if matrix[x][y].bomb {
				matrix = bomb_selected(matrix)
				draw_board(table, matrix)

			} else if !matrix[x][y].flagged && !matrix[x][y].revealed {
				matrix = flood_fill(matrix, x, y)
				draw_board(table, matrix)
			}
			return nil
		}

		if event.Key() == tcell.KeyEscape {
			app.Stop()
			return nil
		}

		return event // arrows, Esc, etc. fall through to the table
	})

	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
