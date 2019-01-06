# Logfile

## Day 1

* Input: row of digits.
* Structure: **circular buffer**, implemented with slice.
* Process: conditional sum.

## Day 2

* Input: list of rows of numbers.
* Process: sum of MinMax pairs, then sum of conditional pairs.

## Day 3

* Input: number.
* Structure: **2d map growing in all directions**, implemented as separate class.
  1. common: map of maps.
  2. **slightly optimized:** 4 growing slices for 4 segments!
* Process: iterate position with direction changes.

## Day 4

* Input: list of rows of words.
* Structure: set of strings, implemented as a map of strings to booleans.
* Process 1: search for duplicate words
* Process 2: search for duplicate sets of characters.

## Day 5

* Input: list of numbers.
* Structure: array of numbers.
* Process: execution.

## Day 6

* Input: row of numbers.
* Structure: **circular buffer**, implemented with slice.
* Process: iterate over circular buffer.

## Day 7

* Input: list of node definitions.
* Structure: tree.
* Process: weights calculation, search in tree.

## Day 8

* Input: list of dual instructions.
* Structure: registers.
* Process: dual execution.

## Day 9

* Input: string.
* Structure: counters.
* Process: parsing, switching between 2 modes.

## Day 10

* Input: row of numbers, interpreted as a string for part 2.
* Structure: **circular buffer**, implemented with fixed size array.
* Process: reversing chunks of circular buffer, hex encoding.

## Day 11

* Input: row of words (directions).
* Structure: **hex grid coordinates**, implemented as cube coordinates!
* Process: execution, manhattan distance in hex grid.

## Day 12

* Input: list of node links (rows of numbers).
* Structure: **unions!** OMG, they are so useful! Implemented as common union.New().
* Process: merging unions.

## Day 13

* Input: list of number pairs.
* Process:
  1. brute force: iterate each step (with direction changes) x each distance.
  2. **optimized:** iterate distances with **calculated scanner positions with mod.**
  3. not implemented: single distance calculation with gcd + mod.

## Day 14

* Input: string.
* Process: hash func from Day 10.
* Data structure and Process options:
  1. full iteration:
      * Structure: bitmap 128x128.
      * Process: repeated flood fill (BFS or DFS).
  2. **optimized:**
      * Structure: bitmap 2x128
      * Process: **find unions, by checking only 2 directions: left and up.**

## Day 15

* Input: 2 numbers.
* Structure: well, 2 numbers.
* Process: generator iteration, %mod.

## Day 16

* Input: row of instructions: letters and numbers.
* Structure: **circular buffer**, implemented as slice with %mod.
* Process: execution, with loop detection of visited states, calculating remaining steps.

## Day 17

* Input: number.
* Structure: **circular buffer (again!)**, now implemented as common circular.NewList.
* Process:
  1. brute force: iterate insertion over circular list.
  2. **optimized:** iterate single position check and ignore content of the list.

## Day 18

* Input: list of instructions.
* Structure: state(register map, ip, queue) x2
* Process:
  1. initial try: goroutines.
  2. optimized: execute single program, until locked by input.
  3. **simple, reliable:** execute ticks for both programs in every iteration.

## Day 19

* Input: 2d map drawn in text.
* Structure: 2d field, used field.Slice.
* Process: walking on map.
  1. custom walk.
  2. field.Walk.

## Day 20

* Input: list of particles (rows with integers).
* Structure: 3 vec3, implemented as single slice of length 9.
* Process: iteration, mark & sweep.
* Bugs:
  1. manhattan distance was calculating sum of squares, instead of sum of abs.
  2. part1 wants particle closest **in the long term** (== in distant future), not closest in all steps.
  3. part2 collision removal was checking particles slice from part1.

## Day 21

* Input: list of rules (2-part strings).
* **Option 1, slightly optimized:**
  * Structure: rule map, 2d state field.
  * Process: generation of rotated rules, iteration, with expanding field.
* Option 2, highly optimized, not implemented:
  * Structure: rule map converted to 3x3->9x9 rules.
  * Process: iteration, with counting field types. No real field storage required.
* Bugs:
  1. rotation was implemented wrong at first. Was swapping flipping 4 sectors of the square.
  2. missing rule. Rule pregeneration was missing adding key after flip, but before rotation.

## Day 22

* Input: text map (lines of chars).
* Structure: growing 2d map, used field.Slice.
* Process: walking on map and changing it, used field.Walk.

## Day 23

* Input: list of instructions (assembly ops)
* Structure: registers.
* Process: part 1, execute instructions.
* Process: part 2, just execute hand-written function.

## Day 24

* Input: list of number pairs.
* Structure: graph!
* Process: find heaviest path in graph, find longest path in graph.
* Option 1, not well thought, buggy custom implementation:
  * Structure: slice of built nodes, slice of flips, map of used nodes.
  * Process: DFS.
* Option 2, short and reliable:
  * Structure: common Graph.
  * Process: DFS, done with graph.IteratePathsFrom.

## Day 25

* Input: list of multiline instructions for each state of Turing machine.
* Structure: tape, implemented as:
  * Option 1, quick write: vanilla map - slow in general.
  * Option 2, optimized for the task: double slice, left and right from 0, implemented as common turing.Tape.
* Process: execute state instructions, inc/dec tape position, in the end: count 1's on tape.

## In total

* New common structures:
  * field.Field - 2d field interface.
    * field.Map - faster for sparce data.
    * field.Slice - 2d slice field, growing in all directions. Faster for compact/filled data.
  * union.New() - find unions of linked nodes.
  * circular.Buffer - looped buffer interface.
    * circular.NewList() - faster for random insertions.
    * circular.NewSlice() - faster for fixed size chunks.
  * graph.NewGraph() - graph with DFS with callback, with arbitrary data storage for every link and node.
  * turing.Tape - tape for Turing machine.

## The end of 2017 puzzles
