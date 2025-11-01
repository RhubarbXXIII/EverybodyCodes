from utils.api import Api


def get_shrine(input: str) -> list[list[str]]:
    return [list(line.strip()) for line in input.split('\n')]


def get_runic_word(shrine: list[list[str]]):
    runic_word = ""
    for row_index in range(2, len(shrine) - 2):
        for column_index in range(2, len(shrine[0]) - 2):
            symbols = set()

            for i in range(len(shrine)):
                if (cell := shrine[i][column_index]) != '.':
                    symbols.add(cell)

            for j in range(len(shrine[0])):
                if (cell := shrine[row_index][j]) != '.':
                    if cell in symbols:
                        runic_word += cell

                    symbols.add(cell)

    return runic_word


def get_power(runic_word: str):
    power = 0
    for index, symbol in enumerate(runic_word):
        power += (index + 1) * (ord(symbol) - 64)

    return power


def part1(input: str) -> str:
    shrine = [list(line.strip()) for line in input.split('\n')]
    return get_runic_word(shrine)


def part2(input: str) -> str:
    shrines = [list(line.strip()) for line in input.split('\n')]

    power = 0

    for shrine_row_index in range(len(shrines) // 9 + 1):
        for shrine_column_index in range(len(shrines[0]) // 9 + 1):
            row_index = shrine_row_index * 9
            column_index = shrine_column_index * 9

            shrine = [
                shrines[i][column_index:column_index + 8]
                for i in range(row_index, row_index + 8)
            ]
            runic_word = get_runic_word(shrine)

            power += get_power(runic_word)

    return str(power)


def part3(input: str) -> str:
    shrines = [list(line.strip()) for line in input.split('\n')]

    power = 0

    unsolved_shrine_indexes = {
        (shrine_row_index, shrine_column_index)
        for shrine_row_index in range(len(shrines) // 6)
        for shrine_column_index in range(len(shrines[0]) // 6)
    }

    solved_shrine_count = None
    while not (solved_shrine_count and solved_shrine_count == 0):
        solved_shrine_indexes = set()

        for (shrine_row_index, shrine_column_index) in unsolved_shrine_indexes:
            previous_runic_word = None
            current_runic_word = None

            while not (previous_runic_word and previous_runic_word == current_runic_word):
                previous_runic_word = current_runic_word
                current_runic_word = ""

                for row_offset in range(2, 6):
                    for column_offset in range(2, 6):
                        runic_symbol = shrines[shrine_row_index + row_offset][shrine_column_index + column_offset]
                        if runic_symbol != '.':
                            current_runic_word += runic_symbol
                            continue

                        row_symbols = {
                            shrines[shrine_row_index + row_offset][shrine_column_index + column_index]
                            for column_index in (0, 1, 6, 7)
                        }
                        column_symbols = {
                            shrines[shrine_row_index + row_index][shrine_column_index + column_offset]
                            for row_index in (0, 1, 6, 7)
                        }

                        if len(row_symbols) < 4 or len(column_symbols) < 4:
                            current_runic_word += '.'
                            continue

                        row_symbols.remove('?')
                        column_symbols.remove('?')

                        intersection_symbols = row_symbols.intersection(column_symbols)
                        if len(intersection_symbols) == 1:
                            runic_symbol = next(iter(intersection_symbols))

                            shrines[shrine_row_index + row_offset][shrine_column_index + column_offset] = runic_symbol
                            current_runic_word += runic_symbol
                            continue

                        runic_row_symbols = {
                            shrines[shrine_row_index + row_offset][shrine_column_index + column_index]
                            for column_index in range(2, 6)
                        }
                        runic_column_symbols = {
                            shrines[shrine_row_index + row_index][shrine_column_index + column_offset]
                            for row_index in range(2, 6)
                        }

                        runic_row_symbols.remove('.')
                        runic_column_symbols.remove('.')

                        runic_symbol = None

                        unused_row_symbols = row_symbols.difference(runic_row_symbols)
                        if len(unused_row_symbols) == 1:
                            runic_symbol = next(iter(unused_row_symbols))

                        unused_column_symbols = column_symbols.difference(runic_column_symbols)
                        if len(unused_column_symbols) == 1:
                            runic_symbol = next(iter(unused_column_symbols))

                        if not runic_symbol:
                            current_runic_word += '.'
                            continue

                        # replace question marks







    return ""


input_1, input_2, input_3 = Api().get_inputs()
if input_1:
    print(f"Part 1: {part1(input_1)}")
if input_2:
    print(f"Part 2: {part2(input_2)}")
if input_3:
    print(f"Part 3: {part3(input_3)}")
