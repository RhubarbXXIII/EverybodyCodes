from utils.api import Api


def parse_mine(input: str):
    return [
        [int(cell) for cell in list(row.strip().replace('.', '0').replace('#', '1'))]
        for row in input.strip().split('\n')
    ]


def part1(input: str) -> int:
    current_mine = parse_mine(input)
    current_depth = 1

    current_blocks_removed = {
        (row_index, column_index)
        for row_index, row in enumerate(current_mine)
        for column_index, cell in enumerate(row)
        if cell == 1
    }
    total_blocks_removed_count = len(current_blocks_removed)

    while current_blocks_removed:
        next_mine = [[cell for cell in row] for row in current_mine]

        current_blocks_removed = set()

        for row_index, row in enumerate(current_mine):
            for column_index, cell in enumerate(row):
                if cell != current_depth:
                    continue

                # Based on input, no need to check bounds.
                if (
                    current_mine[row_index - 1][column_index] == current_depth
                    and current_mine[row_index][column_index + 1] == current_depth
                    and current_mine[row_index + 1][column_index] == current_depth
                    and current_mine[row_index][column_index - 1] == current_depth
                ):
                    next_mine[row_index][column_index] += 1

                    current_blocks_removed.add((row_index, column_index))

        total_blocks_removed_count += len(current_blocks_removed)

        current_mine = next_mine
        current_depth += 1

    return total_blocks_removed_count


def part2(input: str) -> int:
    current_mine = parse_mine(input)
    current_depth = 1

    current_blocks = {
        (row_index, column_index)
        for row_index, row in enumerate(current_mine)
        for column_index, cell in enumerate(row)
        if cell == 1
    }
    total_block_count = len(current_blocks)

    while current_blocks:
        next_mine = [[cell for cell in row] for row in current_mine]
        next_blocks = set()

        for block in current_blocks:
            row_index, column_index = block

            # Based on input, no need to check bounds.
            if (
                current_mine[row_index - 1][column_index] == current_depth
                and current_mine[row_index][column_index + 1] == current_depth
                and current_mine[row_index + 1][column_index] == current_depth
                and current_mine[row_index][column_index - 1] == current_depth
            ):
                next_mine[row_index][column_index] += 1

                next_blocks.add((row_index, column_index))

        total_block_count += len(next_blocks)

        current_mine = next_mine
        current_blocks = next_blocks
        current_depth += 1

    return total_block_count


def part3(input: str) -> int:
    current_mine = parse_mine(input)
    current_depth = 1

    current_blocks = {
        (row_index, column_index)
        for row_index, row in enumerate(current_mine)
        for column_index, cell in enumerate(row)
        if cell == 1
    }
    total_block_count = len(current_blocks)

    while current_blocks:
        next_mine = [[cell for cell in row] for row in current_mine]
        next_blocks = set()

        for block in current_blocks:
            row_index, column_index = block

            if (
                row_index < 1
                or row_index >= len(current_mine) - 1
                or column_index < 1
                or column_index >= len(current_mine[0]) - 1
            ):
                continue

            if all(
                current_mine[row_index + i][column_index + j] == current_depth
                for i in range(-1, 2)
                for j in range(-1, 2)
                if i != 0 or j != 0
            ):
                next_mine[row_index][column_index] += 1

                next_blocks.add((row_index, column_index))

        total_block_count += len(next_blocks)

        current_mine = next_mine
        current_blocks = next_blocks
        current_depth += 1

    return total_block_count


input_1, input_2, input_3 = Api().get_inputs()
if input_1:
    print(f"Part 1: {part1(input_1)}")
if input_2:
    print(f"Part 2: {part2(input_2)}")
if input_3:
    print(f"Part 3: {part3(input_3)}")