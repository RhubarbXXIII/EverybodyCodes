from utils.api import Api


def get_heights(input: str) -> list[int]:
    return [int(height) for height in input.strip().split('\n')]


def part1(input: str) -> int:
    heights = get_heights(input)
    minimum_height = min(heights)

    return sum(height - minimum_height for height in heights)


def part2(input: str) -> int:
    heights = get_heights(input)
    minimum_height = min(heights)

    return sum(height - minimum_height for height in heights)


def part3(input: str) -> int:
    heights = get_heights(input)

    def get_hit_count(target_height: int):
        nonlocal heights
        return sum(abs(height - target_height) for height in heights)

    current_minimum_height = min(heights)
    current_maximum_height = max(heights)

    while current_minimum_height < current_maximum_height:
        current_height = (current_minimum_height + current_maximum_height) // 2

        hit_count_lower = get_hit_count(current_height - 1)
        hit_count = get_hit_count(current_height)
        hit_count_higher = get_hit_count(current_height + 1)

        if hit_count <= hit_count_lower and hit_count <= hit_count_higher:
            return hit_count
        elif hit_count_lower < hit_count < hit_count_higher:
            current_maximum_height = current_height
        elif hit_count_lower > hit_count > hit_count_higher:
            current_minimum_height = current_height
        else:
            raise ValueError("Unable to continue search.")

    raise ValueError("No minimum found.")


input_1, input_2, input_3 = Api().get_inputs()
if input_1:
    print(f"Part 1: {part1(input_1)}")
if input_2:
    print(f"Part 2: {part2(input_2)}")
if input_3:
    print(f"Part 3: {part3(input_3)}")
