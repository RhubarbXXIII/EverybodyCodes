from collections import defaultdict

from utils.api import Api


def get_initial_columns(input: str) -> list[list[int]]:
    rows = [list(row.strip().split(' ')) for row in input.strip().split('\n')]
    return [[int(row[i]) for row in rows] for i in range(4)]


def step_columns(current_column: list[int], next_column: list[int]) -> [list[int], list[int]]:
    clapper = (current_column[0] - 1) % (2 * len(next_column)) + 1
    clapper_index = clapper - 1 if clapper <= len(next_column) else 2 * len(next_column) + 1 - clapper
    return current_column[1:], next_column[:clapper_index] + [current_column[0]] + next_column[clapper_index:]


def part1(input: str) -> int:
    columns = get_initial_columns(input)

    for i in range(10):
        columns[i % 4], columns[(i + 1) % 4] = step_columns(columns[i % 4], columns[(i + 1) % 4])

    return int(''.join(str(column[0]) for column in columns))


def part2(input: str) -> int:
    columns = get_initial_columns(input)

    round_count = 0
    shout_counts = defaultdict(lambda: 0)

    while True:
        i = round_count
        round_count += 1

        columns[i % 4], columns[(i + 1) % 4] = step_columns(columns[i % 4], columns[(i + 1) % 4])

        shout = int(''.join(str(column[0]) for column in columns))
        shout_counts[shout] += 1

        if shout_counts[shout] == 2024:
            return shout * round_count


def part3(input: str) -> int:
    columns = get_initial_columns(input)
    column_history = set()

    shouts = set()

    round_count = 0
    while True:
        column_record = tuple(tuple(column) for column in columns)
        if column_record in column_history:
            return max(shouts)

        column_history.add(column_record)

        round_count += 1
        i = round_count - 1

        columns[i % 4], columns[(i + 1) % 4] = step_columns(columns[i % 4], columns[(i + 1) % 4])

        shout = int(''.join(str(column[0]) for column in columns))
        shouts.add(shout)


input_1, input_2, input_3 = Api().get_inputs()
if input_1:
    print(f"Part 1: {part1(input_1)}")
if input_2:
    print(f"Part 2: {part2(input_2)}")
if input_3:
    print(f"Part 3: {part3(input_3)}")
