from utils.api import Api


def part1(input: str) -> str:
    remaining_block_count = int(input) - 1
    width = 1

    while remaining_block_count > 0:
        width += 2

        necessary_block_count = width
        if necessary_block_count > remaining_block_count:
            return str(width * (necessary_block_count - remaining_block_count))

        remaining_block_count -= necessary_block_count


def part2(input: str) -> str:
    priest_count = int(input)
    acolyte_count = 1111
    block_count = 20240000

    width = 1
    thickness = 1

    remaining_block_count = block_count - 1
    while remaining_block_count > 0:
        width += 2
        thickness = (thickness * priest_count) % acolyte_count

        necessary_block_count = width * thickness
        if necessary_block_count > remaining_block_count:
            return str(width * (necessary_block_count - remaining_block_count))

        remaining_block_count -= necessary_block_count


def part3(input: str) -> str:
    priest_count = int(input)
    acolyte_count = 10
    block_count = 202400000

    thickness = 1
    columns = [1]

    necessary_block_count = 1
    while necessary_block_count < block_count:
        thickness = ((thickness * priest_count) % acolyte_count) + acolyte_count
        columns = [thickness] + [height + thickness for height in columns] + [thickness]

        necessary_block_count = sum(
            [columns[0]]
            + [height - ((priest_count * len(columns) * height) % acolyte_count) for height in columns[1:-1]]
            + [columns[-1]]
        )

    return str(necessary_block_count - block_count)


input_1, input_2, input_3 = Api().get_inputs()
if input_1:
    print(f"Part 1: {part1(input_1)}")
if input_2:
    print(f"Part 2: {part2(input_2)}")
if input_3:
    print(f"Part 3: {part3(input_3)}")
