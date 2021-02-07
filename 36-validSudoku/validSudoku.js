const EMPTY = '.'

function isValidSudoku(board) {
  const rowVisited = {}
  const columnVisited = {}
  const boxVisited = {}

  for (const [i, row] of board.entries()) {
    for (const [j, cell] of row.entries()) {
      if (cell === EMPTY) {
        continue
      }

      if (rowVisited[j]) {
        if (rowVisited[j][cell]) {
          return false
        }

        rowVisited[j][cell] = true
      } else {
        rowVisited[j] = { [cell]: true }
      }

      if (columnVisited[i]) {
        if (columnVisited[i][cell]) {
          return false
        }

        columnVisited[i][cell] = true
      } else {
        columnVisited[i] = { [cell]: true }
      }

      const boxIndex = (Math.floor(i / 3) * 3) + Math.floor(j / 3)
      if (boxVisited[boxIndex]) {
        if (boxVisited[boxIndex][cell]) {
          return false
        }

        boxVisited[boxIndex][cell] = true
      } else {
        boxVisited[boxIndex] = { [cell]: true }
      }
    }
  }

  return true
}

const board = [
  ['5', '3', '.', '.', '7', '.', '.', '.', '.'],
  ['6', '.', '.', '1', '9', '5', '.', '.', '.'],
  ['.', '9', '8', '.', '.', '.', '.', '6', '.'],
  ['8', '.', '.', '.', '6', '.', '.', '.', '3'],
  ['4', '.', '.', '8', '.', '3', '.', '.', '1'],
  ['7', '.', '.', '.', '2', '.', '.', '.', '6'],
  ['.', '6', '.', '.', '.', '.', '2', '8', '.'],
  ['.', '.', '.', '4', '1', '9', '.', '.', '5'],
  ['.', '.', '.', '.', '8', '.', '.', '7', '9'],
]

console.log(isValidSudoku(board))
