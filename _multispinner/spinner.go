package _multispinner

import (
	"fmt"
	"log"
	"time"

	"foreplay/curse"
)

type Spinner struct {
	startTime time.Time
	ticker    *time.Ticker
	position  *curse.Position
	label     string
}

type Container struct {
	spinners []*Spinner
	cursor   *curse.Cursor
	updates  chan *Spinner
}

func New() *Container {
	cursor, err := curse.New()
	if err != nil {
		log.Fatal("new: ", err)
	}

	//fmt.Printf("starting=[%d,%d]\n", c.StartingPosition.X, c.StartingPosition.Y)

	return &Container{
		updates: make(chan *Spinner),
		cursor:  cursor,
	}
}

func (s *Spinner) Render() string {
	elapsed := time.Since(s.startTime)
	return fmt.Sprintf("%d %s", elapsed, s.label)
}

func (c *Container) AddSpinner(label string) {
	x, y, err := curse.GetCursorPosition()
	if err != nil {
		fmt.Println("add spinner: ", err)
	}
	spinner := &Spinner{
		startTime: time.Now(),
		ticker:    time.NewTicker(100 * time.Millisecond),
		position:  &curse.Position{X: x, Y: y},
		label:     label,
	}
	for _, v := range c.spinners {
		v.position.Y--
	}
	c.spinners = append(c.spinners, spinner)

	go func() {
		for {
			<-spinner.ticker.C
			c.updates <- spinner
		}
	}()
}

func (c *Container) Update() {
	for {
		s := <-c.updates
		x, y, err := curse.GetScreenDimensions()
		if err != nil {
			fmt.Println("update: ", err)
		}
		c.cursor.Move(s.position.X, s.position.Y)
		c.cursor.EraseCurrentLine()
		fmt.Print(s.Render())
		c.cursor.Move(x, y)
	}
}
