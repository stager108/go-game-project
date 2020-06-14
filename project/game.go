package main

import (
	"fmt"
    "strconv"
 //    "io/ioutil"
)

func (game *Game) Create() (map[string] interface{}) {

	games := make([]*Game, 0)
	err := GetDB().Table("games").Where("email = ?", game.Email).Find(&games).Error
	if err != nil {
        return Message(false, "Failed to create game")
	}
	
	game.Number = strconv.Itoa(len(games) + 1)
	GetDB().Create(game)
    
    if game.ID <= 0 {
		return Message(false, "Failed to create account, connection error.")
	}

	resp := Message(true, "success")
	resp["game"] = game
	return resp
}

func GetGame(email string, number int) (*Game) {
	game := &Game{}
	err := GetDB().Table("games").Where("email = ? and number=?", email, number).First(&game).Error
	if err != nil {
		return nil
	}	
	
	return game
}

func GetGames(email string) ([]*Game) {

	games := make([]*Game, 0)
	err := GetDB().Table("games").Where("email = ?", email).Find(&games).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return games
}
