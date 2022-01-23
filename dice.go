package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

type DiceResult struct {
	Result int
	Bold   bool
}

// Handle a discordgo roll command interaction
func RollInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	//dice, err := ParseDiceArguments(i.)
	fmt.Print(i)

	args := i.ApplicationCommandData().Options[0].StringValue()

	dice, err := ParseDiceArguments(args)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to run command: " + err.Error(),
			},
		})
		return
	}
	results, err := RollDice(dice)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to run command: " + err.Error(),
			},
		})
		return
	}

	var sum int = 0
	var resultString = ""
	for _, result := range results {
		sum += result.Result
		if len(resultString) != 0 {
			resultString += " + "
		}
		resultString += fmt.Sprint(result.Result)
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%v (%v)", sum, resultString),
		},
	})
}

//Convert a string of 'd's, numbers, and '+'s to a slice of dice to roll
// example args: "d6", "3d13 + d4" & "78d6"
func ParseDiceArguments(arg string) ([]int, error) {
	dice := make([]int, 0)

	countingDice := true
	numberOfDice := ""
	diceSize := ""

	addDice := func() error {
		numberOfDiceInt, err := strconv.ParseInt(numberOfDice, 10, 8)
		if err != nil {
			numberOfDiceInt = 1
		}
		diceSizeInt, err := strconv.ParseInt(diceSize, 10, 8)
		if err != nil {
			return err
		}

		for i := numberOfDiceInt; i > 0; i-- {
			dice = append(dice, int(diceSizeInt))
		}

		countingDice = true
		numberOfDice = ""
		diceSize = ""
		return nil
	}

	for _, c := range arg {
		switch c {
		case ' ':
			// noop
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			if countingDice {
				numberOfDice += string(c)
			} else {
				diceSize += string(c)
			}
		case 'd', 'D':
			countingDice = false
		case '+':
			addDice()
		default:
			return nil, fmt.Errorf("unexpected character %v", c)
		}
	}

	err := addDice()
	if err != nil {
		return nil, err
	} else {
		return dice, nil
	}
}

// Roll multiple dice
func RollDice(dice []int) ([]DiceResult, error) {
	var results = make([]DiceResult, len(dice))

	for i, die := range dice {
		result, err := Roll(die)
		if err != nil {
			return nil, err
		}
		results[i] = *result
	}

	return results, nil
}

// Roll a dice of diceSize and product a DiceResult or an error if the diceSize is less than 1
func Roll(diceSize int) (*DiceResult, error) {
	var result int
	if diceSize < 1 {
		return nil, fmt.Errorf("invalid dice size %v", diceSize)
	} else if diceSize == 1 {
		result = 1
	} else {
		result = rand.Intn(diceSize-1) + 1
	}

	return &DiceResult{
		Result: result,
		Bold:   result == diceSize,
	}, nil

}
