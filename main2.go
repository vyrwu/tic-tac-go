package main

import "fmt"

func main() {
	var i int
	fmt.Scanf("%d", &i)
	fmt.Println(i)
}
//
//type Board struct {
//	Width int
//	Spaces int
//	Blanks int
//	Occupied int
//}
//
//func (b *Board) blank() bool {
//	return b.Blanks == b.Spaces
//}
//
//
//type Game struct {
//	Score      int
//	ActiveTurn string
//}
//
//func (g *Game) over() bool {
//	return true
//}
//
//func (g *Game) getAvailableMoves() [][]int {
//
//}
//
//func (g *Game) getNewState(move []int) *Game {
//
//}
//
//func (g *Game) win(player string) bool {
//
//}
//
//func score(g *Game) int {
//	if g.win(g.ActiveTurn) {
//		return 10
//	}
//}
//
//func minimax(g *Game) ([]int, int) {
//	if g.over() {
//		return nil, score(g)
//	}
//	var scores []int
//	var moves [][]int
//	for _, v := range g.getAvailableMoves() {
//		possibleGame := g.getNewState(v)
//		_, b := minimax(possibleGame)
//		scores = append(scores, b)
//		moves = append(moves, v)
//	}
//
//	minI, maxI := minMaxOf(scores)
//	if g.ActiveTurn == "Player" {
//		return moves[maxI], scores[maxI]
//	}
//	return moves[minI], scores[minI]
//}
//
//// This is faster than dispaching individual min/max methods
//func minMaxOf(slice []int) (int, int) {
//	if len(slice) == 0 {
//		panic("minMaxOf: Slice cannot be empty")
//	}
//	var maxScoreIndex, minScoreIndex int
//	var maxScore, minScore int
//	for i, v := range slice {
//		if i == 0 {
//			maxScore = v
//			minScore = v
//		} else {
//			if v > maxScore {
//				maxScoreIndex = i
//			}
//			if v < minScore {
//				minScoreIndex = i
//			}
//		}
//	}
//	return minScoreIndex, maxScoreIndex
//}
