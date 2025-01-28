import re

from utils.api import Api

def get_runic_words(input: str) -> set[str]:
    return set(input.split('\n')[0].strip().split(':')[1].strip().split(','))

def part1(input: str) -> int:
    runic_words = get_runic_words(input)
    inscription = input.strip().split('\n')[-1]

    return sum(inscription.count(runic_word) for runic_word in runic_words)


def part2(input: str) -> int:
    runic_words = get_runic_words(input)
    runic_words = runic_words.union({runic_word[::-1] for runic_word in runic_words})
    inscriptions = input.strip().split('\n')[2:]

    runic_symbol_count = 0
    for inscription in inscriptions:
        runic_symbol_indices = set()

        for runic_word in runic_words:
            for runic_word_index in (
                match.start() for match in re.finditer(f"(?=({runic_word}))", inscription)
            ):
                runic_symbol_indices = runic_symbol_indices.union(
                    range(runic_word_index, runic_word_index + len(runic_word))
                )

        runic_symbol_count += len(runic_symbol_indices)

    return runic_symbol_count


def part3(input: str) -> int:
    runic_words = get_runic_words(input)
    runic_words = runic_words.union({runic_word[::-1] for runic_word in runic_words})
    armor = [list(row.strip()) for row in input.strip().split('\n')[2:]]

    runic_word_lengths = {len(runic_word) for runic_word in runic_words}
    runic_symbols = set()

    for row_index, row in enumerate(armor):
        for column_index, cell in enumerate(row):
            for runic_word_length in runic_word_lengths:
                horizontal_word = ""
                vertical_word = ""

                for i in range(runic_word_length):
                    horizontal_word += armor[row_index][(column_index + i) % len(armor[0])]
                    vertical_word += armor[(row_index + i) % len(armor)][column_index]

                if horizontal_word in runic_words:
                    runic_symbols = runic_symbols.union(
                        {
                            (row_index, (column_index + i) % len(armor[0]))
                            for i in range(runic_word_length)
                        }
                    )
                if vertical_word in runic_words:
                    if row_index + runic_word_length > len(armor):
                        continue

                    runic_symbols = runic_symbols.union(
                        {(row_index + i, column_index) for i in range(runic_word_length)}
                    )

    return len(runic_symbols)


input_1, input_2, input_3 = Api().get_inputs()
if input_1:
    print(f"Part 1: {part1(input_1)}")
if input_2:
    print(f"Part 2: {part2(input_2)}")
if input_3:
    print(f"Part 3: {part3(input_3)}")