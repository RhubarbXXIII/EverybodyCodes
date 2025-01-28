from utils.api import Api

MONSTER_MAP = {
    'x': 0,
    'A': 0,
    'B': 1,
    'C': 3,
    'D': 5,
}

def part1(input: str) -> int:
    return input.count('B') + 3 * input.count('C')


def part2(input: str) -> int:
    potion_count = 0
    for i in range(0, len(input), 2):
        monster_1, monster_2 = list(input[i:i + 2])
        potion_count += MONSTER_MAP[monster_1] + MONSTER_MAP[monster_2]
        if monster_1 == 'x' or monster_2 == 'x':
            continue

        potion_count += 2

    return potion_count


def part3(input: str) -> int:
    potion_count = 0
    for i in range(0, len(input), 3):
        monsters = list(input[i:i + 3])
        for monster in monsters:
            potion_count += MONSTER_MAP[monster]

        empty_count = monsters.count('x')
        if empty_count == 1:
            potion_count += 2
        elif empty_count == 0:
            potion_count += 6

    return potion_count


input_1, input_2, input_3 = Api().get_inputs()
if input_1:
    print(f"Part 1: {part1(input_1)}")
if input_2:
    print(f"Part 2: {part2(input_2)}")
if input_3:
    print(f"Part 3: {part3(input_3)}")