package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"nelhage.com/tak/ai"
	"nelhage.com/tak/playtak"
	"nelhage.com/tak/ptn"
	"nelhage.com/tak/tak"
)

var (
	server = flag.String("server", "playtak.com:10000", "playtak.com server to connect to")
	depth  = flag.Int("depth", 5, "minimax depth")
	user   = flag.String("user", "", "username for login")
	pass   = flag.String("pass", "", "password for login")
)

const Client = "Takker AI"

func main() {
	flag.Parse()
	client := &client{
		debug: true,
	}
	err := client.Connect(*server)
	if err != nil {
		log.Fatal(err)
	}
	client.SendClient(Client)
	if *user != "" {
		err = client.Login(*user, *pass)
	} else {
		err = client.LoginGuest()
	}
	if err != nil {
		log.Fatal("login: ", err)
	}

	client.sendCommand("Seek", "5", "1200")
	for line := range client.recv {
		if strings.HasPrefix(line, "Game Start") {
			playGame(client, line)
		}
	}
}

func playGame(c *client, line string) {
	ai := ai.NewMinimax(*depth)
	p := tak.New(tak.Config{Size: 5})
	bits := strings.Split(line, " ")
	log.Printf(line)
	gameStr := fmt.Sprintf("Game#%s", bits[2])
	var color tak.Color
	switch bits[7] {
	case "white":
		color = tak.White
	case "black":
		color = tak.Black
	default:
		panic(fmt.Sprintf("bad color: %s", bits[7]))
	}
	for {
		if ok, _ := p.GameOver(); ok {
			break
		}
		if color == p.ToMove() {
			move := ai.GetMove(p)
			next, err := p.Move(*move)
			if err != nil {
				log.Printf("ai returned bad move: %s: %s",
					ptn.FormatMove(move), err)
				continue
			}
			p = next
			c.sendCommand(gameStr, playtak.FormatServer(move))
		} else {
		theirMove:
			for line := range c.recv {
				if !strings.HasPrefix(line, gameStr) {
					continue
				}
				bits = strings.Split(line, " ")
				switch bits[1] {
				case "P", "M":
					move, err := playtak.ParseServer(strings.Join(bits[1:], " "))
					if err != nil {
						panic(err)
					}
					p, err = p.Move(move)
					if err != nil {
						panic(err)
					}
					break theirMove
				case "Abandoned", "Over":
					return
				}
			}
		}
	}
}