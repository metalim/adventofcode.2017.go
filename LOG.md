# Logfile

## Day 1

* Input: row of digits.
* Data structure: **circular buffer**, implemented with slice.
* Process: conditional sum.

## Day 2

* Input: list of rows of numbers.
* Process: sum of MinMax pairs, then sum of conditional pairs.

## Day 3

* Input: number.
* Data structure: **2d map growing in all directions**, implemented as separate class.
  * Option 1, common:
    * map of maps.
  * **Option 2, slightly optimized:**
    * 4 growing slices for 4 segments!
* Process: iterate position with direction changes.

## Day 4

* Input: list of rows of words.
* Data structure: set of strings, implemented as a map of strings to booleans.
* Process 1: search for duplicate words
* Process 2: search for duplicate sets of characters.

## Day 5

* Input: list of numbers.
* Data structure: array of numbers.
* Process: execution.

## Day 6

* Input: row of numbers.
* Data structure: **circular buffer**, implemented with slice.
* Process: iterate over circular buffer.

## Day 7

* Input: list of node definitions.
* Data structure: tree.
* Process: weights calculation, search in tree.

## Day 8

* Input: list of dual instructions.
* Data structure: registers.
* Process: dual execution.

## Day 9

* Input: string.
* Data structure: counters.
* Process: parsing, switching between 2 modes.

## Day 10

* Input: row of numbers, interpreted as a string for part 2.
* Data structure: **circular buffer**, implemented with fixed size array.
* Process: reversing chunks of circular buffer, hex encoding.

## Day 11

* Input: row of words (directions).
* Data structure: **hex grid coordinates**, implemented as cube coordinates!
* Process: execution, manhattan distance in hex grid.

## Day 12

* Input: list of node links (rows of numbers).
* Data structure: **unions!** OMG, they are so useful!
* Process: merging unions.

## Day 13

* Input: list of number pairs.
* Option 1, brute force:
  * Process: iterate each step (with direction changes) x each distance.
* **Option 2, optimized:**
  * Process: iterate distances with **calculated scanner positions with mod.**
* Option 3, not implemented:
  * Process: single distance calculation with gcd + mod.

## Day 14

* Input: string.
* Process: hash func from Day 10.
* Option 1, full iteration:
  * Data structure: bitmap 128x128.
  * Process: repeated flood fill (BFS or DFS).
* **Option 2, optimized:**
  * Data structure: bitmap 2x128
  * Process: **find unions, by checking only 2 directions: left and up.**

## Day 15

* Input: 2 numbers.
* Data structure: well, 2 numbers.
* Process: generator iteration, %mod.

## Day 16

* Input: row of instructions: letters and numbers.
* Data structure: **circular buffer**, implemented as slice with %mod.
* Process: execution, with loop detection of visited states, calculating remaining steps.

## Day 17

* Input: number.
* Data structure: **circular buffer (again!)**, now implemented as common circular.NewList.
* Option 1, brute force:
  * Process: iterate insertion over circular list.
* **Option 2, optimized:**
  * Process: iterate single position check and ignore content of the list.

## Day

* Input:
* Data structure:
* Process:
