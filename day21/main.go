package day21

import "fmt"

func Part1() {
	roll := 1

	player1 := 8
	player1Score := 0
	player2 := 7
	player2Score := 0

	turn := func(player *int, score *int) {
		*player += roll
		roll++
		*player += roll
		roll++
		*player += roll
		roll++

		if *player > 10 {
			*player = *player % 10
			if *player == 0 {
				*player = 10
			}
		}

		*score += *player
	}

	for {
		turn(&player1, &player1Score)
		if player1Score >= 1000 {
			fmt.Printf("player1: %d\n", player1Score)
			fmt.Printf("player2: %d\n", player2Score)
			fmt.Printf("roll: %d\n", roll-1)

			fmt.Printf("Answer: %d\n", player2Score*(roll-1))
			break
		}
		turn(&player2, &player2Score)
		if player2Score >= 1000 {
			fmt.Printf("player1: %d\n", player1Score)
			fmt.Printf("player2: %d\n", player2Score)
			fmt.Printf("roll: %d\n", roll-1)

			fmt.Printf("Answer: %d\n", player1Score*(roll-1))
			break
		}
	}
}

type GamestateMap map[string]*Gamestate

func (ths GamestateMap) PutGamestates(states []*Gamestate) {
	for _, state := range states {
		if _, ok := ths[state.Hash()]; ok {
			ths[state.Hash()].num += state.num
		} else {
			ths[state.Hash()] = state
		}
	}
}

type Gamestate struct {
	player1      int
	player1Score int
	player2      int
	player2Score int

	num int
}

func increment(val int, inc int) int {
	val += inc
	if val > 10 {
		val = val % 10
	}
	return val
}

func (ths *Gamestate) DiceRollStates(isPlayer1Turn bool) []*Gamestate {
	if isPlayer1Turn {
		return []*Gamestate{
			{
				player1:      increment(ths.player1, 1),
				player1Score: ths.player1Score,
				player2:      ths.player2,
				player2Score: ths.player2Score,
				num:          ths.num,
			},
			{
				player1:      increment(ths.player1, 2),
				player1Score: ths.player1Score,
				player2:      ths.player2,
				player2Score: ths.player2Score,
				num:          ths.num,
			},
			{
				player1:      increment(ths.player1, 3),
				player1Score: ths.player1Score,
				player2:      ths.player2,
				player2Score: ths.player2Score,
				num:          ths.num,
			},
		}
	} else {
		return []*Gamestate{
			{
				player1:      ths.player1,
				player1Score: ths.player1Score,
				player2:      increment(ths.player2, 1),
				player2Score: ths.player2Score,
				num:          ths.num,
			},
			{
				player1:      ths.player1,
				player1Score: ths.player1Score,
				player2:      increment(ths.player2, 2),
				player2Score: ths.player2Score,
				num:          ths.num,
			},
			{
				player1:      ths.player1,
				player1Score: ths.player1Score,
				player2:      increment(ths.player2, 3),
				player2Score: ths.player2Score,
				num:          ths.num,
			},
		}
	}
}

func (ths Gamestate) Hash() string {
	return fmt.Sprintf("%d/%d|%d/%d", ths.player1, ths.player1Score, ths.player2, ths.player2Score)
}

func Part2() {

	initial := &Gamestate{
		player1: 8,
		player2: 7,
		num:     1,
	}

	states := GamestateMap{
		initial.Hash(): initial,
	}

	player1States := GamestateMap{}
	player2States := GamestateMap{}

	player1Turn := true

	turn := 1
	for len(states) > 0 {
		for i := 0; i < 3; i++ {
			newStates := GamestateMap{}
			for _, state := range states {
				newStates.PutGamestates(state.DiceRollStates(player1Turn))
			}
			states = newStates
		}
		for hash, state := range states {
			if player1Turn {
				state.player1Score += state.player1
				if state.player1Score >= 21 {
					player1States.PutGamestates([]*Gamestate{state})
					delete(states, hash)
				}
			} else {
				state.player2Score += state.player2
				if state.player2Score >= 21 {
					player2States.PutGamestates([]*Gamestate{state})
					delete(states, hash)
				}
			}
		}
		player1Turn = !player1Turn
		fmt.Printf("Turn %d: len states: %d, p1 states: %d, p2 states: %d\n", turn, len(states), len(player1States), len(player2States))
		turn++
	}

	p1Sum := 0
	for _, state := range player1States {
		p1Sum += state.num
	}
	println(p1Sum)

	p2Sum := 0
	for _, state := range player2States {
		p2Sum += state.num
	}
	println(p2Sum)
}
