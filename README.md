
# Golang Minesweeper

A TUI implementation of Minesweeper in golang using tview.

This project was for me to learn a bit of Go and its flow. Most of the project is hand written except for some parts of the TUI. I used claude to help me out when I got stuck implementing Tview. As simple as it is, developing it has been a lot of fun over the 3 days that it took!


## Run Locally

#### Clone the project:

```bash
  git clone https://github.com/Gabri0n/Golang-Minesweeper
```

#### Go to the project directory:

```bash
  cd Golang-Minesweeper
```

#### Run it:

```bash
go run . <Field Size> <Mine Count>

# Default is a 16x16 field with 30 mines.
```


| Control | Action  |   |   |   |
|---------|---------|---|---|---|
| Space   | Select  |   |   |   |
| F       | Flag    |   |   |   |
| Enter   | Restart |   |   |   |



## Future Work:

1. Add main menu screen
2. Add game win / lose screen
3. Add sound
4. Make it pretty
5. Fix interface.go to not be so awful. 
