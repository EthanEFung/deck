package main

type Card struct {
	Suit
	Rank
	Type
}
type Rank int
type Suit int
type Type int

const (
	Suitless Suit = iota
	Diamonds
	Clubs
	Hearts
	Spades
)

const (
	Rankless Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	Numerical Type = iota
	Face
	High
	Joker
)

var Ranks = []Rank{
	Ace,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Ten,
	Jack,
	Queen,
	King,
}

var Suits = []Suit{
	Diamonds,
	Clubs,
	Hearts,
	Spades,
}

func (s Suit) String() string {
	switch s {
	case Diamonds:
		return "Diamonds"
	case Clubs:
		return "Clubs"
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	default:
		return ""
	}
}

func (r Rank) String() string {
	switch r {
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	case Ace:
		return "Ace"
	default:
		return ""
	}
}

func (t Type) String() string {
	switch t {
	case Joker:
		return "Joker"
	case High:
		return "High"
	case Face:
		return "Face"
	case Numerical:
		return "Numerical"
	default:
		return ""
	}
}

func (c Card) Name() string {
	if c.Type == Joker {
		return "Joker"
	}
	return c.Rank.String() + " of " + c.Suit.String()
}
