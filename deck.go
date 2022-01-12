package deck

import (
	"math/rand"
	"sort"
)

/*
NewOptions is a struct passed to the New function detailing the configurations
for the generated deck
*/
type NewOptions struct {
	/*
	   Sort is a less predicate function that determines the order of the deck
	   the second and third parameters are the indices of the respective cards in
	   the initial order.
	*/
	Sort func([]Card, int, int) bool
	/*
	   Filter is a predicate function that determines which cards are included.
	   Jokers are added to the deck after the filters are applied
	*/
	Filter func([]Card, int) bool
	/*
	   Jokers is a number that represents how many
	*/
	Jokers int
	/*
	   Decks is a multiple of the number of standard decks the generated deck contains excluding jokers
	   (52 * options.Decks)
	*/
	Decks int
	/*
	   Shuffle will psuedo randomize the order of the cards from the number. Pass the number of times
	   the randomizer should run
	*/
	Shuffle int
	/*
	   Seed is the int64 number passed to the algorithm running the shuffle of cards
	*/
	Seed int64
}

/*
New generates a standard deck of cards
*/
func New(options NewOptions) []Card {
	deck := createStandard()
	deck = multiply(deck, options.Decks)
	sortDeck(deck, options.Sort)
	deck = filter(deck, options.Filter)
	deck = addJokers(deck, options.Jokers)
	shuffle(deck, options.Shuffle, options.Seed)
	return deck
}

func createStandard() []Card {
	deck := []Card{}
	for _, suit := range Suits {
		for _, rank := range Ranks {
			var cardType Type = Numerical
			if rank == Jack || rank == Queen || rank == King {
				cardType = Face
			} else if rank == Ace {
				cardType = High
			}
			deck = append(deck, Card{Suit: suit, Rank: rank, Type: cardType})
		}
	}
	return deck
}

func sortDeck(deck []Card, callback func(deck []Card, i, j int) bool) {
	if callback == nil {
		return
	}
	sort.Slice(deck, func(i, j int) bool {
		return callback(deck, i, j)
	})
}

func addJokers(deck []Card, n int) []Card {
	for i := 0; i < n; i++ {
		deck = append(deck, Card{Suit: Suitless, Rank: Rankless, Type: Joker})
	}
	return deck
}

func shuffle(deck []Card, n int, seed int64) {
	r := rand.New(rand.NewSource(seed))
	for i := 1; i <= n; i++ {
		r.Shuffle(len(deck), func(j, k int) {
			deck[j], deck[k] = deck[k], deck[j]
		})
	}
}

func filter(deck []Card, callback func([]Card, int) bool) []Card {
	if callback == nil {
		return deck
	}
	res := []Card{}
	for i, card := range deck {
		if callback(deck, i) {
			res = append(res, card)
		}
	}
	return res
}

func multiply(deck []Card, n int) []Card {
	temp := make([]Card, len(deck))
	copy(temp, deck)
	for i := 2; i <= n; i++ {
		deck = append(deck, temp...)
	}
	return deck
}

// func main() {
// 	deck := New(NewOptions{
// 		// Sort: func(deck []Card, i, j int) bool {
// 		// return deck[i].Rank > deck[j].Rank && deck[i].Suit > deck[j].Suit
// 		// },
// 		Filter: func(deck []Card, i int) bool {
// 			return deck[i].Type == Face
// 		},
// 		Decks:   1,
// 		Jokers:  2,
// 		Shuffle: 1,
// 		Seed:    time.Now().UnixNano(),
// 	})
// 	fmt.Println(deck)
// }
