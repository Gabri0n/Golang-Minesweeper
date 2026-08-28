// User interface

package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func DrawBoard(table *tview.Table, board *Board) {
	for x, row := range board.Cells {
		for y, val := range row {
			color := tcell.ColorWhite

			var text string
			switch {
			case val.isFlagged: // Flag
				text = "?"
				color = tcell.ColorPurple

			case val.isMine && !val.isRevealed: // Unrevealed Bomb
				text = "#"
				color = tcell.ColorWhite

			case val.isMine: // Revealed Bomb
				text = "*"
				color = tcell.ColorRed

			case val.isRevealed && val.adjacentBombs > 0: // Adjacent Bomb Count
				text = strconv.Itoa(val.adjacentBombs)
				color = tcell.ColorYellow

			case val.isRevealed: // Revealed Empty Cell
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

func RenderTui(board *Board) {

	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(false).SetSeparator(' ')

	table.SetSelectedStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorRed))
	table.SetSelectable(true, true)

	DrawBoard(table, board)
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		x, y := table.GetSelection()

		switch event.Rune() {

		case 'f': // flag
			board.Cells[x][y].ToggleFlagged()

			DrawBoard(table, board)

			return nil

		case ' ':
			board.Select(x, y)

			DrawBoard(table, board)

			return nil
		}

		if event.Key() == tcell.KeyEscape {
			app.Stop()
			return nil
		}

		return event
	})

	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
