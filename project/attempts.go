 
package main

import (
	"fmt"
    "strconv"
)

func (attempt *Attempt) Create() (map[string] interface{}) {
    
    attempt.Word = "test"
    //fmt.Printf( attempt.Word)
    
    attempts := make([]*Attempt, 0)
	err := GetDB().Table("attempts").Where("email = ? and game = ?", attempt.Email, attempt.Game).Find(&attempts).Error
	if err != nil {
		return Message(false, "Failed to create attempt")
	}
	
	attempt.Number = strconv.Itoa(len(attempts) + 1)  
    
	GetDB().Create(attempt)
	resp := Message(true, "New attempt is created!")
	resp["attempt"] = attempt
	return resp
}

func (attempt *AttemptAnswer) Validate() (map[string] interface{}) {
    
    fmt.Printf( attempt.Word)
	
	result := &Attempt{}
	err := GetDB().Table("attempts").Where("email = ? and game = ? and number = ?", 
                                           attempt.Email, attempt.Game, attempt.Number).First(&result).Error
	if err != nil {
		return Message(false, "Attempt is not recognized")
	}
	
    err = GetDB().Table("attempts").Where("number = ? and word = ?", attempt.Number, attempt.Word).First(&result).Error
    
    if err != nil {
		return Message(true, "Wrong answer!")
	}

	game := &Game{}
    err = GetDB().Table("games").Where("email = ? and number = ?", attempt.Email, attempt.Game).First(&game).Error
    if err != nil {
		return Message(true, "Your answer was right, but we can't process it!")
	}
		
	return Message(true, "success")
}

func GetAttempt(email string, game int, number int) (*Attempt) {

	attempt := &Attempt{}
	err := GetDB().Table("attempts").Where("email = ? and number = ? and game = ?", 
                                           email, number, game).First(&attempt).Error
	if err != nil {
		return nil
	}
	return attempt
}

func GetAttempts(email string, game int) ([]*Attempt) {

	attempts := make([]*Attempt, 0)
	err := GetDB().Table("attempts").Where("email = ? and game = ?", email, game).Find(&attempts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return attempts
}
