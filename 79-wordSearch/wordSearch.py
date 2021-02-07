from typing import List


def makeGraph(board, word):
    graph = {}
    startList = []

    for i, row in enumerate(board):  # m
        for j, char in enumerate(row):  # n
            graph[(i, j)] = []

            if(j+1 < len(row)):
                graph[(i, j)].append((i, j+1))
            if(j-1 >= 0):
                graph[(i, j)].append((i, j-1))
            if(i+1 < len(board)):
                graph[(i, j)].append((i+1, j))
            if(i-1 >= 0):
                graph[(i, j)].append((i-1, j))

            if(word[0] == char):
                startList.append((i, j))

    return graph, startList


def exist(board: List[List[str]], word: str) -> bool:
    graph, startList = makeGraph(board, word) #mn

    # initial
    visited = []
    charIndex = 0
    currentChar = word[charIndex]
    stack = []

    if(len(startList) == 0):
        return False

    for node in startList:
        stack.append((node, charIndex))

    print(graph)
    # dfs
    while len(stack) > 0 and charIndex < len(word):
        print(stack)
        node, charIndex = stack.pop()
        visited = visited[:charIndex]
        currentChar = word[charIndex]
        adjList = graph[node]
        print(currentChar, charIndex)

        if(charIndex == len(word)-1):
            return True

        for adj in adjList:
            char = board[adj[0]][adj[1]]
            nextChar = word[charIndex+1]

            if(adj in visited or nextChar != char):
                continue

            visited.append(node)
            stack.append((adj, charIndex+1))

    return False


board = [["A", "B", "C", "E"], ["S", "F", "C", "S"], ["A", "D", "E", "E"]]
word = "ABCCED"
# true

# board = [["A", "B", "C", "E"], ["S", "F", "C", "S"], ["A", "D", "E", "E"]]
# word = "ABCB"
# false

# board = [["a"]]
# word = "a"
# true

# board = [["A", "B", "C", "E"], ["S", "F", "E", "S"], ["A", "D", "E", "E"]]
# word = "ABCESEEEFS"
# true

print(exist(board, word))
