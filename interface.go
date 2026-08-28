// User interface

package main

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func RenderTui(game *Game) {

	tview.Borders.Horizontal = tview.BoxDrawingsLightHorizontal
	tview.Borders.Vertical = tview.BoxDrawingsLightVertical
	tview.Borders.TopLeft = tview.BoxDrawingsLightArcDownAndRight
	tview.Borders.TopRight = tview.BoxDrawingsLightArcDownAndLeft
	tview.Borders.BottomLeft = tview.BoxDrawingsLightArcUpAndRight
	tview.Borders.BottomRight = tview.BoxDrawingsLightArcUpAndLeft

	tview.Borders.HorizontalFocus = tview.BoxDrawingsLightHorizontal
	tview.Borders.VerticalFocus = tview.BoxDrawingsLightVertical
	tview.Borders.TopLeftFocus = tview.BoxDrawingsLightArcDownAndRight
	tview.Borders.TopRightFocus = tview.BoxDrawingsLightArcDownAndLeft
	tview.Borders.BottomLeftFocus = tview.BoxDrawingsLightArcUpAndRight
	tview.Borders.BottomRightFocus = tview.BoxDrawingsLightArcUpAndLeft

	app := tview.NewApplication()

	// Table minefield

	Table := tview.NewTable().SetBorders(false).SetSeparator(' ')
	Table.SetSelectedStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorRed))
	Table.SetSelectable(true, true)

	// Header bar

	Header := tview.NewTextView().SetText("FLAGS HERE							TIMER HERE").SetTextColor(tcell.ColorWhite).SetTextAlign(tview.AlignCenter).SetScrollable(false)
	Header.ScrollToBeginning()
	Header.SetBorder(true).SetTitle(" Golang Minesweeper ")

	spacer := tview.NewBox().SetBackgroundColor(tcell.ColorBlack)

	HeaderFlex := tview.NewFlex().SetDirection(tview.FlexColumn)
	HeaderFlex.
		AddItem(spacer, 0, 1, false).
		AddItem(Header, 60, 0, false).
		AddItem(spacer, 0, 1, false)

	// Flex containing Minefield

	FieldWidth := 2*len(game.Board.Cells[0]) - 1
	FieldHeight := len(game.Board.Cells)

	TableFlex := tview.NewFlex()
	TableFlex.SetBorder(true).SetBorderColor(tcell.ColorWhite)
	TableFlex.AddItem(Table, 0, 1, true)

	InnerFieldFlex := tview.NewFlex().SetDirection(tview.FlexColumn)
	InnerFieldFlex.
		AddItem(spacer, 0, 1, false).
		AddItem(TableFlex, FieldWidth+2, 0, true).
		AddItem(spacer, 0, 1, false)

	OuterFieldFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	OuterFieldFlex.
		AddItem(spacer, 0, 1, false).
		AddItem(InnerFieldFlex, FieldHeight+2, 1, true).
		AddItem(spacer, 0, 1, false)

	boxContainer := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(HeaderFlex, 3, 0, false).AddItem(OuterFieldFlex, 0, 1, true)

	DrawBoard(Table, game)
	Table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		x, y := Table.GetSelection()

		switch event.Rune() {

		case 'f': // flag
			game.Board.FlagCell(x, y)
			DrawBoard(Table, game)
			DrawHeader(Header, game)
			return nil

		case ' ':
			game.Board.Select(x, y)
			DrawBoard(Table, game)
			DrawHeader(Header, game)
			return nil
		}

		if event.Key() == tcell.KeyEscape {
			app.Stop()
			return nil
		}

		return event
	})

	// This is completely AI here, had no clue how to fix it

	app.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
		w, h := screen.Size()
		availH := h - 3 // header is 3 rows tall

		tw := FieldWidth + 2 // +2 for the L/R border
		if tw > w {
			tw = w
		}
		th := FieldHeight + 2 // +2 for the T/B border
		if th > availH {
			th = availH
		}

		InnerFieldFlex.ResizeItem(TableFlex, tw, 0)      // WIDTH  → TableFlex, inside the FlexColumn
		OuterFieldFlex.ResizeItem(InnerFieldFlex, th, 0) // HEIGHT → InnerFieldFlex, inside the FlexRow
		return false
	})

	if err := app.SetRoot(boxContainer, true).SetFocus(Table).Run(); err != nil {
		panic(err)
	}
}

// Draw the Header

func DrawHeader(Header *tview.TextView, game *Game) {

	var HeaderFlags string = strconv.Itoa(game.Board.MineCount - game.Board.FlagCount)

	game.UpdateTimer()

	var HeaderText = fmt.Sprintf("Flags: %s							Time: %.1fs", HeaderFlags, game.ElapsedTime.Seconds())

	Header.SetText(HeaderText).SetTextColor(tcell.ColorWhite).SetTextAlign(tview.AlignCenter).SetScrollable(false)
	Header.ScrollToBeginning()

}

// Draw the Board

func DrawBoard(table *tview.Table, game *Game) {
	for x, row := range game.Board.Cells {
		for y, cell := range row {
			color := tcell.ColorWhite
			BackgroundColour := tcell.ColorDarkGrey

			var mineColors = map[int]tcell.Color{
				1: tcell.ColorBlue,
				2: tcell.ColorGreen,
				3: tcell.ColorRed,
				4: tcell.ColorDarkBlue,
				5: tcell.ColorDarkRed,
				6: tcell.ColorTeal,
				7: tcell.ColorBlack,
				8: tcell.ColorGray,
			}

			var text string
			switch {
			case cell.isFlagged: // Flag
				text = "?"
				color = tcell.ColorRed

			case cell.isMine && !cell.isRevealed: // Unrevealed Bomb
				text = "#"
				color = tcell.ColorWhite

			case cell.isMine: // Revealed Bomb
				text = "*"
				color = tcell.ColorBlack
				BackgroundColour = tcell.ColorRed

			case cell.isRevealed && cell.adjacentBombs > 0: // Adjacent Bomb Count
				text = strconv.Itoa(cell.adjacentBombs)
				color = mineColors[cell.adjacentBombs]

			case cell.isRevealed: // Revealed Empty Cell
				text = " "
				color = tcell.ColorBlack

			default: // Default Cell
				text = "#"
				color = tcell.ColorWhite

			}

			table.SetCell(x, y, tview.NewTableCell(text).SetTextColor(color).SetBackgroundColor(BackgroundColour).SetAlign(tview.AlignCenter).SetMaxWidth(1))

		}
	}
}
