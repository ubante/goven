package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"gopkg.in/inconshreveable/log15.v2"
	"os"
	"sort"
	"time"
)

/*
Trying to stick to the definitions in https://www.pokernews.com/pokerterms/reraise.htm
and the rules in http://www.pokercoach.us/RobsPkrRules11.doc

In addition to those definitions:
  - hand: the five cards a player constructs
  - round: a betting stage; there are four per game: preflop/flop/turn/river
  - game: what happens between the shuffling of the deck
  - tournament: from the first game until there's only one player left
*/

type Card struct {
	suit         string
	numericaRank int
	rank         string
}

func NewCard(s string, nr int) Card {
	var c Card
	c.suit = s
	c.numericaRank = nr

	if nr == 14 {
		c.rank = "A" // Aces are aces.
	}

	switch nr {
	case 14:
		c.rank = "A" // Aces are aces.
	case 13:
		c.rank = "K"
	case 12:
		c.rank = "Q"
	case 11:
		c.rank = "J"
	case 10:
		c.rank = "T"
	default:
		c.rank = strconv.Itoa(nr)
	}

	return c
}

func (c *Card) toString() string {
	return fmt.Sprintf("%s%s", c.suit, c.rank)
}

func (c Card) String() string {
	return c.toString()
}

func (c Card) isSuited(c2 Card) bool {
	if c.suit == c2.suit {
		return true
	}

	return false
}

func (c Card) isPaired(c2 Card) bool {
	if c.numericaRank == c2.numericaRank {
		return true
	}

	return false
}

type CardSet struct {
	cards         []*Card
	bestHand      *CardSet
	bestEval      *Evaluation
	possibleHands []*CardSet
}

func NewCardSet(cards ...Card) CardSet {
	var cs CardSet

	for _, card := range cards {
		cs.add(card)
	}

	return cs
}

// I should rename toString to String and explicitly call
// SomeType.String() when I want a string outside of Println()
func (cs CardSet) toString() string {
	var toString string

	for _, c := range cs.cards {
		if toString == "" {
			toString = c.toString()
			continue
		}

		toString += fmt.Sprintf(" %s", c.toString())
	}

	return toString
}

func (cs CardSet) String() string {
	return cs.toString()
}

func (cs *CardSet) add(c Card) {
	cs.cards = append(cs.cards, &c)
}

/*
This will accept a cardset and return the union of it with this.
*/
func (cs *CardSet) combine(cs2 CardSet) CardSet {
	var combined CardSet

	for _, card := range cs.cards {
		combined.add(*card)
	}
	for _, card := range cs2.cards {
		combined.add(*card)
	}

	return combined
}

// Find all the possible combinations.
// I like it gross.
func (cs *CardSet) setPossibleHands() {
	cards := cs.cards
	for a := 0; a < len(cards)-4; a++ {
		for b := a + 1; b < len(cards)-3; b++ {
			for c := b + 1; c < len(cards)-2; c++ {
				for d := c + 1; d < len(cards)-1; d++ {
					for e := d + 1; e < len(cards); e++ {
						possibleHand := NewCardSet(*cards[a], *cards[b], *cards[c], *cards[d], *cards[e])
						cs.possibleHands = append(cs.possibleHands, &possibleHand)
					}
				}
			}
		}
	}
}

func (cs *CardSet) findBestHand() {
	cs.setPossibleHands()
	for _, ph := range cs.possibleHands {
		eval := NewEvaluation(*ph)
		//fmt.Println(eval)

		if cs.bestHand == nil {
			cs.bestHand = ph
			cs.bestEval = eval
			//fmt.Println("Here's the initial bestHand:", cs.bestHand)
			continue
		}

		// Maybe stash this eval so no need to recompute it?
		currentEval := NewEvaluation(*cs.bestHand)
		//if currentEval.compare(*eval) == -1 {
		//	cs.bestHand = ph
		//	cs.bestEval = eval
		//	//fmt.Println("This is the new bestHand:", cs.bestHand)
		//	//fmt.Println("This is the new bestHand:", eval)
		//}

		if eval.isBetterThan(*currentEval) {
			cs.bestHand = ph
			cs.bestEval = eval
		}
		//switch currentEval.compare(*eval) {
		//case -1:
		//	cs.bestHand = ph
		//case 1
		//}
	}

	//for i, ph := range cs.possibleHands {
	//	fmt.Println(i, "evalling:", ph)
	//	eval := NewEvaluation(*ph)
	//	fmt.Println(eval)
	//}
}

/*
I couldn't find a pop function - weird.
https://groups.google.com/forum/#!topic/Golang-nuts/obZI4uyZTe0
*/
func (cs *CardSet) pop() *Card {
	card := cs.cards[0]
	copy(cs.cards, cs.cards[1:])
	cs.cards = cs.cards[:len(cs.cards)-1]

	return card
}

func (cs *CardSet) shuffle() {
	rand.Shuffle(cs.length(), func(i, j int) {
		cs.cards[i], cs.cards[j] = cs.cards[j], cs.cards[i]
	})
}

func (cs CardSet) length() int {
	return len(cs.cards)
}

func (cs CardSet) getReverseOrderedNumericRanks() []int {
	var orderedRanks []int
	for _, card := range cs.cards {
		orderedRanks = append(orderedRanks, card.numericaRank)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(orderedRanks)))

	return orderedRanks
}

/*
Primary ranks are:
  9: Straight flush
  8: Four of a kind
  7: Full house
  6: Flush
  5: Straight
  4: Three of a kind
  3: Two pair
  2: One pair
  1: High card

Secondary ranks vary with the primary rank.  For example,  for
a straight, the secondary rank will be the highest card.  For
a full house, the secondary rank will be the numeric value of
the three of a kind.

Tertiary rank also vary's with the primary rank.  For a flush,
this would be the second highest card.  For a full house, this
would be the numeric value of the pair.

This just goes on and on.  For a flush and high card, it is
possible to have a rank for each of the five cards.  For the
other types of hands, these will remain at zero.

It is possible for different hands to have the same evaluation.
*/
type Evaluation struct {
	cardSet        *CardSet
	humanEval      string
	allRanks       [6]*int
	primaryRank    int
	secondaryRank  int
	tertiaryRank   int
	quaternaryRank int
	quinaryRank    int
	senaryRank     int
	flattenedScore int // A simple way to score the hand.  Higher is better.
}

func NewEvaluation(cardSet CardSet) *Evaluation {
	var eval Evaluation
	eval.cardSet = &cardSet
	eval.evaluate()

	// How to do the below in one line?
	eval.allRanks[0] = &eval.primaryRank
	eval.allRanks[1] = &eval.secondaryRank
	eval.allRanks[2] = &eval.tertiaryRank
	eval.allRanks[3] = &eval.quaternaryRank
	eval.allRanks[4] = &eval.quinaryRank
	eval.allRanks[5] = &eval.senaryRank
	eval.flattenScore()

	return &eval
}

func (e Evaluation) String() string {
	toString := fmt.Sprintf("%s: %s with ranks: [%d %d %d %d %d %d]", e.cardSet, e.humanEval,
		e.primaryRank, e.secondaryRank, e.tertiaryRank, e.quaternaryRank, e.quinaryRank, e.senaryRank)

	return toString
}

func (e Evaluation) isFlush() bool {
	var suit string
	for _, card := range e.cardSet.cards {
		if suit == "" {
			suit = card.suit
		} else {
			if card.suit != suit {
				return false
			}
		}
	}

	return true
}

func (e Evaluation) isStraight() bool {
	// This needs to handle a wheel straight, ie A2345.

	orderedRanks := e.cardSet.getReverseOrderedNumericRanks()
	previous := 0
	for _, rank := range orderedRanks {
		if previous == 0 {
			previous = rank
		} else {
			if previous-rank != 1 {
				return false
			}
			previous = rank
		}
	}

	return true
}

// This will return a map of ints -> ints.
// The key will be the size of the match, ie 4 means quads and 1 means
// unmatched cards.  The value will be the numeric rank(s) of that size.
// For the numeric rank 2, there could be two values, ie two pairs.
//
// If cards have ranks 13 13 11 11 7, then the returned map will be:
//   2 -> [13, 11]  // a pair of kings and jacks
//   1 -> [7]       // a lone 7
func (e Evaluation) hasMatches() map[int][]int {
	frequency := make(map[int]int)
	for _, card := range e.cardSet.cards {
		frequency[card.numericaRank]++
	}

	// We want the keys to the frequency map to be added to the matches
	// map in descending order.
	var orderedKeys []int
	for k := range frequency {
		orderedKeys = append(orderedKeys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(orderedKeys)))

	matches := make(map[int][]int)
	for _, rank := range orderedKeys {
		matches[frequency[rank]] = append(matches[frequency[rank]], rank)
	}

	return matches
}

// This will flatten the allRanks array into an int preserving
// ordinality. This could be more compact by using powers of 13 but then
// I'd have to cast to float64 to use Pow.  smh
func (e *Evaluation) flattenScore() {
	e.flattenedScore = 0
	//fmt.Println(e, "Starting with e.flattenedScore of", e.flattenedScore)
	for _, subScore := range e.allRanks {
		e.flattenedScore *= 100
		//fmt.Println("at subscore of", *subScore, "e.flattenedScore is now", e.flattenedScore)
		e.flattenedScore += *subScore // this line panics
	}

	//fmt.Println("Flattened", e.allRanks, "to", e.flattenedScore)
}

// This is because defer can call functions but cannot call methods.
// Yeah, I did it.  Wot?
//func notAMethodFlattenScore(e *Evaluation) {
//	fmt.Println("flattening:", e)
//	e.flattenScore()
//}

func (e *Evaluation) evaluate() {
	//defer notAMethodFlattenScore(e)
	e.humanEval = "TBDeval"

	if e.isStraight() && e.isFlush() {
		e.humanEval = "staight flush"
		e.primaryRank = 9
		highestRank := e.cardSet.getReverseOrderedNumericRanks()[0]
		if highestRank == 14 {
			e.humanEval = "royal flush"
		}
		e.secondaryRank = highestRank
		return
	}

	allMatches := e.hasMatches()
	// https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
	if rank, ok := allMatches[4]; ok {
		e.humanEval = "quads"
		e.primaryRank = 8
		e.secondaryRank = rank[0]
		e.tertiaryRank = allMatches[1][0]
		return
	}

	// Guess golang if's only allow one initialization statement?
	//if rank3, ok3 := allMatches[3]; rank2, ok2 := allMatches[2]; ok2 {
	var trips, firstPair, secondPair int
	var ok3, ok2 bool
	if _, ok := allMatches[3]; ok {
		ok3 = true
		trips = allMatches[3][0]
	}
	if _, ok := allMatches[2]; ok {
		ok2 = true
		firstPair = allMatches[2][0]
		if len(allMatches[2]) == 2 {
			secondPair = allMatches[2][1]
		}
	}

	if ok3 && ok2 {
		e.humanEval = "full house"
		e.primaryRank = 7
		e.secondaryRank = trips
		e.tertiaryRank = firstPair
		return
	}

	if e.isFlush() {
		e.humanEval = "flush"
		e.primaryRank = 6
		orderedRanks := e.cardSet.getReverseOrderedNumericRanks()
		e.secondaryRank = orderedRanks[0]
		e.tertiaryRank = orderedRanks[1]
		e.quaternaryRank = orderedRanks[2]
		e.quinaryRank = orderedRanks[3]
		e.senaryRank = orderedRanks[4]
		return
	}

	if e.isStraight() {
		e.humanEval = "straight"
		e.primaryRank = 5
		e.secondaryRank = e.cardSet.getReverseOrderedNumericRanks()[0]
		return
	}

	// At this point, we're guaranteed to have unmatched cards.
	singles := allMatches[1]
	sort.Sort(sort.Reverse(sort.IntSlice(singles)))

	if trips != 0 {
		e.humanEval = "three of a kind"
		e.primaryRank = 4
		e.secondaryRank = trips
		e.tertiaryRank = singles[0]
		e.quinaryRank = singles[1]
		return
	}

	if firstPair*secondPair != 0 {
		e.humanEval = "two pairs"
		e.primaryRank = 3
		e.secondaryRank = firstPair
		e.tertiaryRank = secondPair
		e.quaternaryRank = singles[0]
		return
	}

	if firstPair != 0 {
		e.humanEval = "pair"
		e.primaryRank = 2
		e.secondaryRank = firstPair
		e.tertiaryRank = singles[0]
		e.quaternaryRank = singles[1]
		e.quinaryRank = singles[2]
		return
	}

	// Don't laugh.
	e.humanEval = "high card"
	e.primaryRank = 1
	e.secondaryRank = singles[0]
	e.tertiaryRank = singles[1]
	e.quaternaryRank = singles[2]
	e.quinaryRank = singles[3]
	e.senaryRank = singles[4]
}

// This will return 1 if this Evaluation is greater than the given
// Evaluation and will return -1 if this Evaluation is lesser than
// the other Evaluation.  If they are even, then this will return 0.
// Maybe use an enum?
//
// Greater values in rank means betterness.
func (e Evaluation) compare(otherEval Evaluation) int {

	for i := 0; i < 6; i++ {
		//fmt.Println(i, "this:", e.allRanks[i], ", other:", otherEval.allRanks[i])
		if *e.allRanks[i] < *otherEval.allRanks[i] {
			return -1
		}
		if *e.allRanks[i] > *otherEval.allRanks[i] {
			return 1
		}
	}

	return 0
}

// Courtesy method.  Note that comparing two Evaluations has three possible
// results.
func (e Evaluation) isBetterThan(otherEval Evaluation) bool {
	results := e.compare(otherEval)

	if results == 1 {
		return true
	}

	return false
}

type HoleCards struct {
	cardSet *CardSet
}

func (hc *HoleCards) toString() string {
	if hc.cardSet == nil {
		return ""
	}

	return hc.cardSet.toString()
}

func (hc HoleCards) String() string {
	return hc.toString()
}

func (hc *HoleCards) add(c Card) {
	hc.cardSet.add(c)
}

func (hc *HoleCards) empty() {
	ecs := NewCardSet()
	hc.cardSet = &ecs
}

func (hc *HoleCards) toss() {
	hc.empty()
}

type Deck struct {
	cardset *CardSet // Deck is not a CardSet.
}

func NewDeck() *Deck {
	var d Deck
	ecs := NewCardSet()
	d.cardset = &ecs

	for _, suit := range []string{"S", "H", "D", "C"} {
		// https://stackoverflow.com/questions/21950244/is-there-a-way-to-iterate-over-a-range-of-integers-in-golang
		for numericRank := range [13]int{} {
			newCard := NewCard(suit, numericRank+2)
			d.cardset.add(newCard)
		}
	}

	return &d
}

func (d *Deck) getStatus() string {
	status := fmt.Sprintf("Deck has %d cards:\n", len(d.cardset.cards))
	for i, card := range d.cardset.cards {
		if i != 0 && i%13 == 0 {
			status += "\n"
		}
		status += card.toString()
		status += " "
	}

	return status
}

func (d Deck) length() int {
	return d.cardset.length()
}

func (d *Deck) shuffle() {
	d.cardset.shuffle()
}

func (d *Deck) getCard() *Card {
	return d.cardset.pop()
}

// Maybe I can get away with just using a CardSet instead of this
// struct.
type Community struct {
	cards *CardSet
}

func NewCommunity() Community {
	ecs := NewCardSet()

	var c Community
	c.cards = &ecs

	return c
}

// Will eventually replace all the getStatus() methods.
func (c Community) String() string {
	return c.cards.String()
	//return c.cards.getStatus()
}

func (c *Community) add(card Card) {
	c.cards.add(card)
}

type SubPot struct {
	amount              int
	contributingPlayers []*Player
}

func (sp SubPot) String() string {
	toString := fmt.Sprintf("$%d: ", sp.amount)
	for _, cp := range sp.contributingPlayers {
		cpp := *cp
		toString += fmt.Sprintf("%s, ", cpp.getName())
	}
	return toString
}

func (sp SubPot) contains(player *Player) bool {
	// This could have been a map but then I'd have to convert the slice
	// in NewPot.
	for _, cp := range sp.contributingPlayers {
		if cp == player {
			return true
		}
	}
	return false
}

func (sp *SubPot) deposit(player *Player, amount int) {
	pp := *player
	if !sp.contains(player) {
		sp.contributingPlayers = append(sp.contributingPlayers, player)
	}
	sp.amount += amount
	pp.setBet(pp.getBet() - amount)
}

type Pot struct {
	subPots     []*SubPot // The first SubPot is the main pot.
	subPotIndex int
}

func (pot Pot) String() string {
	toString := fmt.Sprintf("POT total is $%d", pot.getValue())
	for i, sp := range pot.subPots {
		toString += fmt.Sprintf("\npot #%d, %s", i, *sp)
	}
	return toString
}

func NewPot(players []*Player) Pot {
	var pot Pot

	aSubPot := SubPot{contributingPlayers: players, amount: 0}
	pot.subPots = append(pot.subPots, &aSubPot)
	pot.subPotIndex = 0

	return pot
}

func (pot *Pot) getValue() int {
	value := 0
	for _, sp := range pot.subPots {
		value += sp.amount
	}

	return value
}

func (pot *Pot) recordRoundBets(players []*Player) {
	// First find all-in players with bets.
	var allInBetValues []int
	for _, p := range players {
		pp := *p
		if pp.checkIsAllIn() && pp.getBet() > 0 {
			allInBetValues = append(allInBetValues, pp.getBet())
		}
	}

	// Loop through them in ascending order by Player.bet.
	sort.Sort(sort.IntSlice(allInBetValues))
	for _, bet := range allInBetValues {
		// Loop through all the players with bets.
		for _, p := range players {
			pp := *p
			if pp.getBet() == 0 {
				continue
			}

			// Deposit their bets (up to the current all-in player's
			// bet) into the current subPot.
			if pp.getBet() > bet {
				pot.subPots[pot.subPotIndex].deposit(p, bet)
			} else {
				pot.subPots[pot.subPotIndex].deposit(p, pp.getBet())
			}
		}

		// Create a new subPot and move the index.
		//fmt.Printf("Pot #%d is closed at $%d\n", pot.subPotIndex, pot.subPots[pot.subPotIndex].amount)
		fmt.Printf("Pot #%d is closed at $%d\n", pot.subPotIndex, bet)
		pot.subPots = append(pot.subPots, &SubPot{contributingPlayers: nil, amount: 0})
		pot.subPotIndex++
	}

	// Catch the rest of the bets by looping over all players with bets.
	for _, p := range players {
		pp := *p
		if pp.getBet() == 0 {
			continue
		}

		// Deposit their bets in its entirety into the current subPot.
		pot.subPots[pot.subPotIndex].deposit(p, pp.getBet())
	}
}

func (pot *Pot) getWinnings(playerScores map[*Player]int) map[*Player]int {
	winnings := make(map[*Player]int)

	// Loop through all the subPots.
	for i, sp := range pot.subPots {
		// It is possible to have an empty subPot.  eg, some goes all
		// in, someone calls and no one raises until the end of the
		// game.
		if sp.amount == 0 {
			continue
		}

		// Find the contributingPlayer(s) with the highest score.
		highestScore := 0
		var subPotWinners []*Player

		for _, cp := range sp.contributingPlayers {
			if playerScores[cp] > highestScore {
				highestScore = playerScores[cp]
				subPotWinners = nil
				subPotWinners = append(subPotWinners, cp)
			} else if playerScores[cp] == highestScore {
				subPotWinners = append(subPotWinners, cp)
			}
		}
		fmt.Printf("We have %d winners of a subPot worth $%d (each $%d):\n", len(subPotWinners), sp.amount,
			sp.amount/len(subPotWinners))
		fmt.Printf("SubPot #%d was won by ", i)
		for _, spw := range subPotWinners {
			spwp := *spw
			fmt.Printf("%s, ", spwp.getName())

			// Tally their winnings.
			if _, ok := winnings[spw]; ok {
				winnings[spw] += sp.amount / len(subPotWinners)
			} else {
				winnings[spw] = sp.amount / len(subPotWinners)
			}
		}
		fmt.Println()

	}

	return winnings
}

// This is most-likely an anti-pattern.
type Player interface {
	fold()
	allIn()
	check(t *Table)
	preset()
	setNextPlayer(p *Player)
	getNextPlayer() *Player
	setPreviousPlayer(p *Player)
	getPreviousPlayer() *Player
	setName(n string)
	getName() string
	raise(raiseAmount int)
	setBet(newBet int)
	getBet() int
	addToStack(payout int)
	getStack() int
	addHoleCard(c Card)
	getHoleCards() CardSet
	payBlind(blindAmount int)
	checkHasFolded() bool
	checkIsAllIn() bool
	chooseAction(t *Table)
}

type GenericPlayer struct {
	name           string
	nextPlayer     *Player
	previousPlayer *Player

	// The below get preset after each game.
	holeCards HoleCards
	stack     int
	bet       int
	hasFolded bool
	isAllIn   bool
}

// GenericPlayer constructor
// http://www.golangpatterns.info/object-oriented/constructors
// Maybe this could be replaced with new() and some helper lines.
func NewGenericPlayer(name string) GenericPlayer {
	ecs := NewCardSet()
	hc := HoleCards{cardSet: &ecs}
	initialStack := 1000 // dollars
	newPlayer := GenericPlayer{name, nil, nil, hc, initialStack, 0,
		false, false}
	return newPlayer
}

func (gp *GenericPlayer) getStatus() string {
	status := ""
	status = fmt.Sprintf("%s: [%s] stack=%d bet=%d", gp.name, gp.holeCards.toString(), gp.stack, gp.bet)

	return status
}

func (gp GenericPlayer) String() string {
	return fmt.Sprintf("%s: [%s] $%d/$%d", gp.name, gp.holeCards.toString(), gp.bet, gp.stack)
}

/*
Preset before each game.

Maybe use "prepare()" instead of "preset()" because the latter implies
something you do afterwards.
*/
func (gp *GenericPlayer) preset() {
	// Maybe NewGenericPlayer can call this?
	ecs := NewCardSet()
	gp.holeCards = HoleCards{cardSet: &ecs}
	gp.bet = 0
	gp.hasFolded = false
	gp.isAllIn = false
}

func (gp *GenericPlayer) setNextPlayer(p *Player) {
	gp.nextPlayer = p
}

func (gp *GenericPlayer) getNextPlayer() *Player {
	return gp.nextPlayer
}

func (gp *GenericPlayer) setPreviousPlayer(p *Player) {
	gp.previousPlayer = p
}

func (gp *GenericPlayer) getPreviousPlayer() *Player {
	return gp.previousPlayer
}

func (gp *GenericPlayer) setName(n string) {
	gp.name = n
}

func (gp *GenericPlayer) getName() string {
	return gp.name
}

func (gp *GenericPlayer) setBet(newBet int) {
	gp.bet = newBet
}

func (gp *GenericPlayer) getBet() int {
	return gp.bet
}

func (gp *GenericPlayer) addToStack(payout int) {
	gp.stack += payout
}

func (gp *GenericPlayer) getStack() int {
	return gp.stack
}

func (gp *GenericPlayer) addHoleCard(c Card) {
	gp.holeCards.add(c)
}

func (gp *GenericPlayer) getHoleCards() CardSet {
	return *gp.holeCards.cardSet
}

func (gp *GenericPlayer) payBlind(blindAmount int) {
	if gp.stack > blindAmount {
		gp.stack -= blindAmount
		gp.bet = blindAmount
	} else {
		gp.allIn()
	}
}

func (gp *GenericPlayer) check(t *Table) {
	// This may one day do something.

	return
}

func (gp *GenericPlayer) call(t *Table) {
	maxBet := t.getMaxBet()
	increase := maxBet - gp.bet
	stackBefore := gp.stack
	betBefore := gp.bet

	if increase >= gp.stack {
		gp.allIn()
	} else {
		gp.bet = maxBet
		gp.stack -= increase
	}

	fmt.Printf("calling - max bet:%d, increase:%d, bet before/after:%d/%d, stack before/after:%d/%d\n",
		maxBet, increase, betBefore, gp.bet, stackBefore, gp.stack)
}

func (gp *GenericPlayer) fold() {
	gp.hasFolded = true
	gp.holeCards.toss()
}

// I know there is a difference between bet and raise but that seems
// like irrelevant semantics here.
func (gp *GenericPlayer) raise(raiseAmount int) {
	if gp.stack < raiseAmount {
		fmt.Println("Player tried to raise more than its stack.  This is a fatal error.")
		fmt.Println(gp)
		os.Exit(9)
	}

	gp.bet += raiseAmount
	gp.stack -= raiseAmount
}

func (gp *GenericPlayer) allIn() {
	gp.raise(gp.stack)
	//gp.bet += gp.stack
	//gp.stack = 0
	gp.isAllIn = true
}

func (gp *GenericPlayer) checkOrCall(t *Table) {
	if t.getMaxBet() == 0 {
		gp.check(t)
		return
	}

	gp.call(t)

	return
}

func (gp *GenericPlayer) checkOrFold(t *Table) {
	if t.getMaxBet() > gp.getBet() {
		fmt.Printf("%s thinks the pot ($%d) is too rich - folding.\n", gp.getName(), t.getMaxBet())
		gp.fold()
		return
	}

	gp.check(t)
	return
}

func (gp *GenericPlayer) chooseActionPreflop(t *Table) {
	fmt.Println("Using the default preflop action of check-calling")
	gp.checkOrCall(t)

	return
}

func (gp *GenericPlayer) chooseActionFlop(t *Table) {
	fmt.Println("Using the default flop action of check-folding")
	gp.checkOrFold(t)

	return
}

func (gp *GenericPlayer) chooseActionTurn(t *Table) {
	fmt.Println("Using the default turn action of check-folding")
	gp.checkOrFold(t)

	return
}

func (gp *GenericPlayer) chooseActionRiver(t *Table) {
	fmt.Println("Using the default river action of check-folding")
	gp.checkOrFold(t)

	return
}

func (gp *GenericPlayer) chooseAction(t *Table) {

	// This is handled by Table() but be redundant for clearness.
	if gp.hasFolded {
		fmt.Println("I,", gp.name, "have already checkHasFolded so no action.")
		return
	}
	if gp.isAllIn {
		fmt.Println("I,", gp.name, "am all in so no action.")
		return
	}

	fmt.Println(gp.getStatus())

	switch t.bettingRound {
	case "PREFLOP":
		gp.chooseActionPreflop(t)
	case "FLOP":
		gp.chooseActionFlop(t)
	case "TURN":
		gp.chooseActionTurn(t)
	case "RIVER":
		gp.chooseActionRiver(t)
	}
	return
}

func (gp *GenericPlayer) checkHasFolded() bool {
	return gp.hasFolded
}

func (gp *GenericPlayer) checkIsAllIn() bool {
	return gp.isAllIn
}

// This Player will always call.
type CallingStationPlayer struct {
	GenericPlayer
}

// Repeating the constructor is kinda lame.
func NewCallingStationPlayer(name string) CallingStationPlayer {
	ecs := NewCardSet()
	hc := HoleCards{cardSet: &ecs}
	initialStack := 1000 // dollars

	newPlayer := new(CallingStationPlayer)
	newPlayer.name = name
	newPlayer.holeCards = hc
	newPlayer.stack = initialStack

	return *newPlayer
}

func (csp *CallingStationPlayer) chooseAction(t *Table) {
	fmt.Println("This is CSP overriding chooseAction")
	currentTableBet := t.getMaxBet()

	if currentTableBet > csp.bet {
		fmt.Printf("My bet ($%d) is below the table's bet ($%d) so calling.\n", csp.bet, currentTableBet)
		csp.call(t)
	} else {
		fmt.Printf("My bet ($%d) is >= the table bet ($%d) so just checking.\n", csp.bet, currentTableBet)
		csp.check(t)
	}
}

// This Player likes to see a specific street.  But then check/folds.
type StreetTestPlayer struct {
	GenericPlayer
	foldingStreet string
}

func NewStreetTestPlayer(name string, street string) StreetTestPlayer {
	ecs := NewCardSet()
	hc := HoleCards{cardSet: &ecs}
	initialStack := 1000 // dollars

	newPlayer := new(StreetTestPlayer)
	newPlayer.name = name
	newPlayer.holeCards = hc
	newPlayer.stack = initialStack
	newPlayer.foldingStreet = street

	return *newPlayer
}

func (stp *StreetTestPlayer) chooseAction(t *Table) {
	if stp.hasFolded {
		fmt.Println("I have already folded so no action.  How did this codepath happen btw?.")
		return
	}

	if t.bettingRound == stp.foldingStreet {
		fmt.Printf("(%s) @%s: I've seen my street so check-folding\n", stp.getName(), stp.foldingStreet)
		stp.checkOrFold(t)
	} else {
		fmt.Printf("(%s) @%s: check-calling\n", stp.getName(), stp.foldingStreet)
		stp.checkOrCall(t)
	}
}

type MinRaisingPlayer struct {
	GenericPlayer
}

// This Player makes sure that no street goes unbet.
func NewMinRaisingPlayer(name string) MinRaisingPlayer {
	ecs := NewCardSet()
	hc := HoleCards{cardSet: &ecs}
	initialStack := 1000 // dollars

	newPlayer := new(MinRaisingPlayer)
	newPlayer.name = name
	newPlayer.holeCards = hc
	newPlayer.stack = initialStack

	return *newPlayer
}

func (mrp *MinRaisingPlayer) chooseAction(t *Table) {
	currentTableBet := t.getMaxBet()
	if currentTableBet > mrp.bet {
		fmt.Printf("My bet ($%d) is below the table's bet ($%d) so calling.\n", mrp.bet, currentTableBet)
		mrp.call(t)
	} else {
		if mrp.bet == 0 {
			fmt.Printf("To keep the bluffers nervous, min-raising $%d\n", t.bigBlindValue)
			mrp.raise(t.bigBlindValue)
		} else {
			fmt.Println("I have bet but I'm now behind the bet.  Calling")
			mrp.call(t)
		}
	}
}

// This Player will check if able, otherwise will fold.  This models an
// online player who has disconnected.  If the present players are
// aggressive enough, it is common to see a FoldingPlayer make it to the
// money.
type FoldingPlayer struct {
	GenericPlayer
}

func NewFoldingPlayer(name string) FoldingPlayer {
	ecs := NewCardSet()
	hc := HoleCards{cardSet: &ecs}
	initialStack := 1000 // dollars

	newPlayer := new(FoldingPlayer)
	newPlayer.name = name
	newPlayer.holeCards = hc
	newPlayer.stack = initialStack

	return *newPlayer
}

func (fp *FoldingPlayer) chooseAction(t *Table) {
	fmt.Println("This is FP overriding chooseAction")
	currentTableBet := t.getMaxBet()

	if currentTableBet > fp.bet {
		fmt.Printf("My bet ($%d) is below the table's bet ($%d) so folding like a champ.\n", fp.bet, currentTableBet)
		fp.fold()
	} else {
		fmt.Printf("My bet ($%d) is >= the table bet ($%d) so I'm still in it.\n", fp.bet, currentTableBet)
		fp.check(t)
	}
}

// This Player is pure aggro.  He will always shove.
type AllInAlwaysPlayer struct {
	GenericPlayer
}

func NewAllInAlwaysPlayer(name string) AllInAlwaysPlayer {
	ecs := NewCardSet()
	hc := HoleCards{cardSet: &ecs}
	initialStack := 1000 // dollars

	newPlayer := new(AllInAlwaysPlayer)
	newPlayer.name = name
	newPlayer.holeCards = hc
	newPlayer.stack = initialStack

	return *newPlayer
}

func (ap *AllInAlwaysPlayer) chooseAction(t *Table) {
	fmt.Println("I am always all-in always.")
	ap.allIn()
}

// This Player is a calling station unless the current bet is >=5x his
// current non-zero bet.  In that case, he will fold.
type CallToFivePlayer struct {
	GenericPlayer
}

func NewCallToFivePlayer(name string) CallToFivePlayer {
	ecs := NewCardSet()
	hc := HoleCards{cardSet: &ecs}
	initialStack := 1000 // dollars

	newPlayer := new(CallToFivePlayer)
	newPlayer.name = name
	newPlayer.holeCards = hc
	newPlayer.stack = initialStack

	return *newPlayer
}

func (ct5p *CallToFivePlayer) chooseAction(t *Table) {
	fmt.Println("This is CT5 overriding chooseAction")
	currentTableBet := t.getMaxBet()

	// First check or call if I haven't bet yet.
	if ct5p.bet == 0 {
		ct5p.checkOrCall(t)
	}

	// Then call up to 5x my bet.
	if currentTableBet >= 5*ct5p.bet {
		fmt.Printf("My bet ($%d) is less than a fifth of the table's bet ($%d) so folding.\n",
			ct5p.bet, currentTableBet)
		ct5p.fold()
	} else if currentTableBet > ct5p.bet {
		fmt.Printf("My bet ($%d) is below the table's bet ($%d) so calling.\n", ct5p.bet, currentTableBet)
		ct5p.call(t)
	} else {
		fmt.Printf("My bet ($%d) is >= the table bet ($%d) so just checking.\n", ct5p.bet, currentTableBet)
		ct5p.check(t)
	}

}

// These are future players.

/**
This breaks my brain.
*/
type Table struct {
	players []*Player
	gameCtr int

	// The below get preset before each game.
	community        Community
	button           Player
	smallBlindValue  int
	smallBlindPlayer *Player
	bigBlindValue    int
	bigBlindPlayer   *Player
	bettingRound     string
	deck             *Deck
	pot              Pot
}

// getStatus is more verbose than toString.
func (t *Table) getStatus() string {
	status := "------\n"
	status += fmt.Sprintf("%s -- %d players\n", t.bettingRound, len(t.players))

	betTotal := 0
	stackTotal := 0
	for _, p := range t.players {
		pp := *p
		status += fmt.Sprintf("%s\n", *p)
		betTotal += pp.getBet()
		stackTotal += pp.getStack()
	}

	status += fmt.Sprintf("Pot: %d\n", t.pot.getValue())
	status += fmt.Sprintf("Community: %s\n", t.community)
	status += fmt.Sprintf("Bet totals: %d\n", betTotal)
	status += fmt.Sprintf("Stack totals: %d\n", stackTotal)
	status += "------\n"
	return status
}

/*
This happens at the start of tournaments.
*/
func (t *Table) initialize() {
	t.gameCtr = 0

	// https://stackoverflow.com/questions/33994677/pick-a-random-value-from-a-go-slice
	rand.Seed(time.Now().Unix())
}

/*
This happens at the start of games.
*/
func (t *Table) preset() {
	t.gameCtr++

	t.deck = NewDeck()
	t.deck.shuffle()
	t.pot = NewPot(t.players)
	t.community = NewCommunity()

	for _, p := range t.players {
		player := *p
		player.preset()
	}
}

//func (t *Table) addPlayerPointerVersion(player *Player) {
//	if len(t.players) == 0 {
//		t.players = append(t.players, player)
//		return
//	}
//
//	playerDerefd := *player
//
//	initialPlayer := *t.players[0]
//	lastPlayerPtr := t.players[len(t.players)-1]
//	lastPlayer := *lastPlayerPtr
//
//	lastPlayer.setNextPlayer(player)
//	playerDerefd.setPreviousPlayer(&lastPlayer)
//	playerDerefd.setNextPlayer(t.players[0])
//	initialPlayer.setPreviousPlayer(player)
//
//	t.players = append(t.players, player)
//
//	//fmt.Printf("CURRENT PLAYER: %s > %s > %s\n",
//	//	lastPlayer.getName(), playerDerefd.getName(), initialPlayer.getName())
//	return
//}

func (t *Table) addPlayer(player Player) {
	if len(t.players) == 0 {
		t.players = append(t.players, &player)
		return
	}

	initialPlayer := *t.players[0]
	lastPlayerPtr := t.players[len(t.players)-1]
	lastPlayer := *lastPlayerPtr
	lastPlayer.setNextPlayer(&player)
	player.setPreviousPlayer(lastPlayerPtr)
	player.setNextPlayer(t.players[0])
	initialPlayer.setPreviousPlayer(&player)
	t.players = append(t.players, &player)

	//fmt.Printf("CURRENT PLAYER: %s > %s > %s\n",
	//	lastPlayer.getName(), player.getName(), initialPlayer.getName())
	return
}

func (t Table) printPlayerList() {
	if len(t.players) == 1 {
		lonePlayer := *t.players[0]
		fmt.Println("There is only one player at the table:", lonePlayer.getName())
		return
	}

	fmt.Println("Players: ")
	for _, p := range t.players {
		player := *p // Needed because t.players is a slice of *Player.
		np := player.getNextPlayer()
		nextPlayer := *np
		nnp := nextPlayer.getNextPlayer()
		nextNextPlayer := *nnp
		fmt.Println(player.getName(), nextPlayer.getName(), nextNextPlayer.getName())
	}
}

func (t Table) printLinkList(reverse bool, p *Player) {
	// the zero value of a bool is false

	if p == nil {
		p = t.players[0]
	}
	player := *p

	if reverse == false {
		if player.getNextPlayer() == t.players[0] {
			fmt.Println(player.getName())
			return
		}
	} else {
		if player.getPreviousPlayer() == t.players[0] {
			fmt.Println(player.getName())
			return
		}
	}

	if reverse == false {
		fmt.Printf("%s -> ", player.getName())
	} else {
		fmt.Printf("%s <- ", player.getName())
	}

	if reverse == false {
		t.printLinkList(reverse, player.getNextPlayer()) // error (fixed)
	} else {
		t.printLinkList(reverse, player.getPreviousPlayer())
	}
	return
}

func (t *Table) assignInitialButtonAndBlinds() {
	n := rand.Int() % len(t.players)
	t.button = *t.players[n]

	fmt.Println("Assigning the button to:", t.button.getName())
	t.smallBlindPlayer = t.button.getNextPlayer()
	smallBlindDerefd := *t.smallBlindPlayer
	fmt.Println("Assigning SB to:", smallBlindDerefd.getName())
	t.bigBlindPlayer = smallBlindDerefd.getNextPlayer()
	bigBlindPlayerDerefd := *t.bigBlindPlayer
	fmt.Println("Assigning BB to:", bigBlindPlayerDerefd.getName())
}

func (t *Table) defineBlinds(sb int) {
	t.smallBlindValue = sb
	t.bigBlindValue = sb * 2
}

func (t *Table) postBlinds() (table Table) {
	bbp := *t.bigBlindPlayer
	bbp.payBlind(t.bigBlindValue)
	fmt.Println(bbp.getName(), "just paid the blind of $", t.bigBlindValue, "and has $",
		bbp.getStack(), "left.")
	sbp := *t.smallBlindPlayer
	sbp.payBlind(t.smallBlindValue)
	fmt.Println(sbp.getName(), "just paid the blind of $", t.smallBlindValue, "and has $",
		sbp.getStack(), "left.")

	return
}

func (t *Table) dealHoleCards() {
	for _, p := range t.players {
		player := *p
		player.addHoleCard(*t.deck.getCard())
		player.addHoleCard(*t.deck.getCard())
	}
}

func (t *Table) getMaxBet() int {
	var maxBet int
	if t.bettingRound == "PREFLOP" {
		maxBet = t.bigBlindValue
	} else {
		maxBet = 0
	}

	for _, p := range t.players {
		player := *p
		if maxBet < player.getBet() {
			maxBet = player.getBet()
		}
	}

	return maxBet
}

func (t *Table) getPlayerAction(playerPtr *Player) {
	player := *playerPtr
	if player.checkHasFolded() {
		fmt.Println(player.getName(), "has folded so no action.")
		return
	}

	if player.checkIsAllIn() {
		fmt.Println(player.getName(), "is all-in so no action.")
		return
	}

	fmt.Println(player.getName(), "has action - finding it.")
	player.chooseAction(t)

	return
}

/*
Return false unless all non folded players either have the same bet or
are all-in.
*/
func (t *Table) checkBetParity() bool {
	maxBet := t.getMaxBet()
	for _, p := range t.players {
		player := *p
		if player.checkHasFolded() || player.checkIsAllIn() {
			continue
		}

		if player.getBet() != maxBet {
			return false
		}
	}

	return true
}

func (t *Table) genericBet(firstBetter *Player) {
	firstBetterDerefd := *firstBetter
	fmt.Println("The first better is", firstBetterDerefd.getName())
	log15.Info("The first better is", firstBetterDerefd.getName())

	better := firstBetterDerefd
	t.getPlayerAction(&better)
	better = *better.getNextPlayer()

	// First go around the table.
	for better != firstBetterDerefd {
		fmt.Println(better.getName(), "is the better.")

		if t.checkForOnePlayer() {
			fmt.Println("We are down to one player.")
			return
		}

		t.getPlayerAction(&better)
		better = *better.getNextPlayer()

	}

	fmt.Println("After going around the table once, we have:")
	fmt.Println(t.getStatus())

	// There may be raises and re-raises so handle that.
	for {
		if t.checkForOnePlayer() {
			fmt.Println("There is only one player left with action.")
			break
		}

		if t.checkBetParity() {
			fmt.Println("Everyone had a chance to bet and everyone is all-in, has checkHasFolded or has called.")
			break
		}

		// These players have no action.
		if better.checkHasFolded() || better.checkIsAllIn() {
			fmt.Println(better.getName(), "has no action.")
			better = *better.getNextPlayer()
			continue
		}

		t.getPlayerAction(&better)
		better = *better.getNextPlayer()
	}
}

func (t *Table) preFlopBet() {
	bigBlindPlayer := *t.bigBlindPlayer
	firstBetter := bigBlindPlayer.getNextPlayer()
	t.genericBet(firstBetter)
}

func (t *Table) postPreFlopBet() {
	t.genericBet(t.smallBlindPlayer)
}

func (t *Table) countFoldedPlayers() {

}

func (t *Table) checkForOnePlayer() bool {
	remaingPlayerCount := len(t.players)
	if remaingPlayerCount > 1 {
		return false
	}

	return true
}

func (t *Table) moveBetsToPot() {
	fmt.Println("Moving bets to pot.")
	t.pot.recordRoundBets(t.players)
	fmt.Println("The pot is now:", t.pot)
}

func (t *Table) dealFlop() {
	t.bettingRound = "FLOP"
	for i := 1; i <= 3; i++ {
		card := t.deck.getCard()
		t.community.add(*card)
	}
}

func (t *Table) dealTurn() {
	t.bettingRound = "TURN"
	card := t.deck.getCard()
	t.community.add(*card)
}

func (t *Table) dealRiver() {
	t.bettingRound = "RIVER"
	card := t.deck.getCard()
	t.community.add(*card)
}

func (t *Table) payWinners() {
	fmt.Println("The pot:", t.pot)

	// Find all the players still in it and find their hand strength.
	fmt.Println("Finding the still active players.")
	var activePlayers []*Player
	playerScores := make(map[*Player]int)
	for _, p := range t.players {
		player := *p
		if player.checkHasFolded() {
			continue
		}
		activePlayers = append(activePlayers, p)

		// Evaluate the players' hand strengths.
		fmt.Println("Evaluating the hand of", player.getName())
		hc := player.getHoleCards()
		combinedCardSet := hc.combine(*t.community.cards) // 7 cards.
		combinedCardSet.findBestHand()
		fmt.Printf("%s's best hand is: %s\n", player.getName(), combinedCardSet.bestEval)
		playerScores[p] = combinedCardSet.bestEval.flattenedScore
	}

	// Send all the hand strengths to the pot to find payouts.
	payments := t.pot.getWinnings(playerScores)

	// Pay the players.
	fmt.Println("Paying the winners finally.")
	for p := range payments {
		pp := *p
		fmt.Println(pp.getName(), "gets", payments[p])
		pp.addToStack(payments[p])
	}
}

func (t *Table) removePlayer(player *Player) {
	for i, p := range t.players {
		pp := *p
		if p == player {
			fmt.Println("Removing", pp.getName(), "from position", i, "on the table.")
			t.players = append(t.players[:i], t.players[i+1:]...)
			return
		}
	}

	playerDerefd := *player
	fmt.Println("Could not find", playerDerefd.getName(), "so exiting program.")
	os.Exit(11)
}

func (t *Table) removeBustedPlayers() {
	var bustedPlayers []*Player  // We need this temp slice so not to affect the below range.
	for _, p := range t.players {
		pp := *p
		if pp.getStack() > 0 {
			continue
		}
		bustedPlayers = append(bustedPlayers, p)
	}

	for _, bp := range bustedPlayers {
		bpp := *bp

		// Remove the busted player.
		fmt.Println(bpp.getName(), "has busted out so removing from the table.")
		fmt.Printf("From: ")
		t.printLinkList(false, nil)
		fmt.Printf("      ")
		t.printLinkList(true, nil)
		previousPlayer := bpp.getPreviousPlayer()
		ppp := *previousPlayer
		nextPlayer := bpp.getNextPlayer()
		npp := *nextPlayer

		ppp.setNextPlayer(nextPlayer)
		npp.setPreviousPlayer(previousPlayer)

		t.removePlayer(bp)
		fmt.Printf("To: ")
		t.printLinkList(false, nil)
		fmt.Printf("    ")
		t.printLinkList(true, nil)
	}
}

// Is this still needed?
func (t *Table) payWinnersForSegment(segmentValue int, players []*Player) {
	// It is possible that a segment has no valid players because they
	// have all folded.  For example, if the small blind folds, he may
	// be the only player in the segment equal to t.smallBlindValue.
	// Another example is a single person folds after the flop.  The
	// preflop segment would only have one person and he folded so would
	// be invalid.
	// TODO: Handle these cases.

	fmt.Println("Finding the winner of the", segmentValue, "dollar segment.")

	var activePlayers []Player
	for _, p := range players {
		player := *p

		// You can't win if you don't play.
		if player.checkHasFolded() {
			continue
		}
		activePlayers = append(activePlayers, player)
	}

	fmt.Println("There are", len(activePlayers), "active players in this segment.")
	var segmentWinningPlayers []Player // Ties happen.
	var segmentWinningEvaluation Evaluation
	for _, ap := range activePlayers {
		fmt.Println("-", ap.getName())

		// I will pay the cost of reevaluating these hands so I don't
		// have to add more methods to the Player interface.
		aphc := ap.getHoleCards()
		combinedCardset := aphc.combine(*t.community.cards)
		combinedCardset.findBestHand()
		fmt.Printf("%s's best hand is: %s\n", ap.getName(), combinedCardset.bestEval)
		thisEval := *combinedCardset.bestEval

		if len(segmentWinningPlayers) == 0 {
			segmentWinningPlayers = []Player{ap}
			segmentWinningEvaluation = thisEval
			fmt.Println("Initialing the best player with", ap.getName())
			continue
		}
		switch segmentWinningEvaluation.compare(thisEval) {
		case -1:
			thisPlayer := ap
			segmentWinningPlayers = []Player{thisPlayer}
			segmentWinningEvaluation = thisEval
			fmt.Println("YYYYY: We have a new best player:", ap.getName())
		case 0:
			segmentWinningPlayers = append(segmentWinningPlayers, ap)
			fmt.Println("OOOOO:", ap.getName(), "has tied the best hand.")
		default:
			fmt.Println("NNNNNN: We do not have a new best player - the reign continues.")
			fmt.Printf("(%s) remains the best\n  over %s (%s)\n", segmentWinningEvaluation,
				ap.getName(), thisEval)
		}
	}

	fmt.Println("\nThe community cards were:", t.community)
	fmt.Println("The winning hand:", segmentWinningEvaluation)
	payout := segmentValue / len(segmentWinningPlayers)
	fmt.Printf("The winners of the $%d segment, each winning $%d:\n", segmentValue, payout)
	for i, p := range segmentWinningPlayers {
		fmt.Printf("%d: %v\n", i, p)
		p.addToStack(payout)
		fmt.Printf("   %v\n", p)
		//fmt.Println("  ", segmentWinningEvaluation)
	}
	os.Exit(4)
}

func runTournament() {
	var table Table
	table.initialize()

	temp := NewAllInAlwaysPlayer("Adam")
	temp.stack = 100 // To test Pot, we need an all-in short stack player.
	table.addPlayer(&temp)
	temp1 := NewAllInAlwaysPlayer("Alma")
	temp1.stack = 125 // To test Pot, we need an all-in short stack player.
	//temp1.stack = 0 // To test Pot, we need an all-in short stack player.
	table.addPlayer(&temp1)
	temp2 := NewGenericPlayer("Bert")
	table.addPlayer(&temp2)
	tempCSP := NewCallingStationPlayer("Cali")
	table.addPlayer(&tempCSP)
	temp4 := NewGenericPlayer("Dale")
	table.addPlayer(&temp4)
	temp5 := NewGenericPlayer("Eyor")
	table.addPlayer(&temp5)
	temp6 := NewFoldingPlayer("Fred")
	table.addPlayer(&temp6)
	temp7 := NewCallToFivePlayer("Greg")
	table.addPlayer(&temp7)
	temp8 := NewGenericPlayer("Hill")
	table.addPlayer(&temp8)
	temp9 := NewGenericPlayer("Igor")
	table.addPlayer(&temp9)
	temp10 := NewStreetTestPlayer("Flow", "FLOP")
	table.addPlayer(&temp10)
	temp11 := NewStreetTestPlayer("Turk", "TURN")
	table.addPlayer(&temp11)
	temp12 := NewStreetTestPlayer("Rivv", "RIVER")
	table.addPlayer(&temp12)
	temp13 := NewMinRaisingPlayer("Ming")
	table.addPlayer(&temp13)

	//table.printLinkList(false, nil)
	//table.printLinkList(true, nil)
	table.removeBustedPlayers()
	//table.printLinkList(false, nil)
	//table.printLinkList(true, nil)

	//os.Exit(7)


	fmt.Print("\n\n")

	// Set an initial small blind value.
	table.defineBlinds(25)

	for i := 1; i <= 2; i++ {
		fmt.Println("============================")
		table.preset()
		fmt.Printf("This is game #%d.\n", table.gameCtr)
		table.assignInitialButtonAndBlinds()
		table.bettingRound = "PREFLOP"

		//if i == 2 {
		//	fmt.Println("Figure out how to pay winners then we'll continue.")
		//	fmt.Println(table.getStatus())
		//	os.Exit(3)
		//}

		table.postBlinds()
		fmt.Println(table.getStatus())
		table.dealHoleCards()
		table.preFlopBet()
		table.moveBetsToPot()
		fmt.Println(table.getStatus())

		table.bettingRound = "FLOP"
		fmt.Println("Dealing the flop.")
		table.dealFlop()
		table.postPreFlopBet()
		table.moveBetsToPot()
		fmt.Println(table.getStatus())

		table.bettingRound = "TURN"
		fmt.Println("Dealing the turn.")
		table.dealTurn()
		table.postPreFlopBet()
		table.moveBetsToPot()
		fmt.Println(table.getStatus())

		table.bettingRound = "RIVER"
		fmt.Println("Dealing the river.")
		table.dealRiver()

		// Mock the community cards for testing Evaluation()
		//mockedCardSet := NewCommunity()
		//mockedCardSet.add(NewCard("H", 11))
		//mockedCardSet.add(NewCard("H", 6))
		//mockedCardSet.add(NewCard("C", 11))
		//mockedCardSet.add(NewCard("S", 11))
		//mockedCardSet.add(NewCard("C", 6))
		//table.community = mockedCardSet

		table.postPreFlopBet()
		table.moveBetsToPot()
		fmt.Println(table.getStatus())

		fmt.Println("Finding and paying the winners.")
		table.payWinners()
		table.removeBustedPlayers()

		fmt.Println("At the end of the game, the table looks like this:")
		fmt.Println(table.getStatus())
	}

}

/*
We will run multiple tournaments to find the best player type.
Each tournament has multiple poker _games_.
Each game has multiple betting _rounds_.
Each winner of each game has the best _hand_ - multiple winners are possible.
*/

// This is just to cast a GenericPlayer to Player.
//func castToPlayer(player Player) Player {
//	return player
//}

func main() {
	//card := NewCard("H", 12)
	//cardSet := NewCardSet(card)
	//fmt.Println(cardSet)
	//card2 := NewCard("S", 4)
	//cardSet.add(card2)
	//fmt.Println(cardSet)
	//fmt.Println("cardset1 =", cardSet)
	//
	////ecs := getEmptyCardSet()
	//ecs := NewCardSet()
	//fmt.Println("ECS:", ecs)
	//
	//cs2 := NewCardSet()
	//cs2.add(NewCard("C", 13))
	//cs2.add(NewCard("C", 12))
	//fmt.Println("cardset2 =", cs2)
	//
	//cs3 := cs2.combine(cardSet)
	//fmt.Println("cardset3 =", cs3)
	//
	//cs4 := NewCardSet(card, card2)
	//fmt.Println("cardset4 =", cs4)
	//
	//cs5 := NewCardSet(card, card2, NewCard("C", 4))
	//fmt.Println("using type:", reflect.TypeOf(card2))
	//fmt.Println("cardset5 =", cs5)
	//for i, c := range cs5.cards {
	//	fmt.Println(i, c)
	//}
	//
	//os.Exit(3)

	//for numericRank := range [13]int{} {
	//	newCard := NewCard("D", numericRank+2)
	//	//fmt.Println("New card:", newCard)
	//	cardSet.add(newCard)
	//}
	//fmt.Println(cardSet.getStatus())
	//cardSet.shuffle()
	//fmt.Println(cardSet.getStatus())
	/*
		H12 S4 D2 D3 D4 D5 D6 D7 D8 D9 D10 D11 D12 D13 D14
		D11 D10 D14 D3 D2 D7 D12 S4 H12 D6 D4 D5 D8 D13 D9
	*/

	// This will test the new pot.
	//pot := NewOldPot()
	//
	//temp6a := NewGenericPlayer("Fred")
	//temp6a.stack = 250
	//temp6i := castToPlayer(&temp6a)
	//temp7a := NewGenericPlayer("Ginger")
	//temp7i := castToPlayer(&temp7a)
	//temp8a := NewGenericPlayer("Harold")
	//temp8i := castToPlayer(&temp8a)
	//temp9a := NewCallingStationPlayer("Kyle")
	//temp9i := castToPlayer(&temp9a)
	//players := []*Player{&temp6i, &temp7i, &temp8i, &temp9i}
	//for _, p := range players {
	//	pp := *p
	//	pp.raise(100)
	//	fmt.Println(pp)
	//}
	//fmt.Println(pot)
	//for _, p := range players {
	//	pp := *p
	//	pot.addEquity(pp.getBet(), p)
	//}
	//fmt.Println(pot)
	//temp6i.allIn()
	//temp7i.fold()
	//temp8i.raise(200)
	//temp9i.raise(200)
	//for _, p := range players {
	//	pp := *p
	//	pot.addEquity(pp.getBet(), p)
	//	fmt.Println(pp)
	//}
	//fmt.Println("\nFred is all in; Ginger has folded.")
	//fmt.Println(pot)
	//
	//os.Exit(7)

	for i := 1; i <= 2; i++ {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Printf("Starting tournament #%d\n", i)
		runTournament()
	}

}
