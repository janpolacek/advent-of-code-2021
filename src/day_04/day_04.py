import numpy as np
import copy

lines_list = [line for line in open('input.txt')]
numbers = np.fromstring(lines_list[0], dtype=int, sep=',')
boards = np.fromstring(''.join(lines_list[2:]), dtype=int, sep='\n').reshape(-1, 5, 5)


def run_best_bingo(numbers, boards):
    for number in numbers:
        for idx in np.ndindex(boards.shape):
            if boards[idx] == number:
                boards[idx] = 0

        for idx in np.ndindex(boards.shape):

            rows = np.sum(boards[idx[0]], axis=1)
            columns = np.sum(boards[idx[0]], axis=0)

            if not np.all(rows) or not np.all(columns):
                winning_index = idx[0]
                wining_board = boards[winning_index]
                winning_sum = np.sum(wining_board)

                print("Bingo: " + str(winning_index), str(winning_sum), str(winning_sum * number))
                return winning_sum * number


def run_worst_bingo(numbers, boards):
    won_boards = np.zeros(len(boards))
    for number in numbers:
        for idx in np.ndindex(boards.shape):
            if boards[idx] == number:
                boards[idx] = 0

        for idx in np.ndindex(boards.shape):
            rows = np.sum(boards[idx[0]], axis=1)
            columns = np.sum(boards[idx[0]], axis=0)

            if not np.all(rows) or not np.all(columns):
                winning_index = idx[0]
                wining_board = boards[winning_index]
                winning_sum = np.sum(wining_board)

                won_boards[winning_index] = 1
                if np.sum(won_boards) == len(won_boards):
                    print("Last Bingo: " + str(winning_index), str(winning_sum), str(winning_sum * number))
                    return winning_sum * number


print("part1: " + str(run_best_bingo(numbers, copy.deepcopy(boards))))
print("part1: " + str(run_worst_bingo(numbers, copy.deepcopy(boards))))
