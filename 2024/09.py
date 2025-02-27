import sys
from collections import defaultdict

from utils.api import Api


def get_brightnesses(input: str) -> list[int]:
    return [int(line.strip()) for line in input.split('\n')]


def get_beetles(brightness: int, stamps: list[int]) -> dict[int, int]:
    dot_increments = defaultdict(lambda: 0)
    for dot_increment in stamps:
        while brightness > 0:
            if dot_increment > brightness:
                break

            brightness -= dot_increment
            dot_increments[dot_increment] += 1

    return dot_increments


def part1(input: str) -> str:
    stamps: list[int] = sorted([1, 3, 5, 10], reverse=True)

    return str(sum(
        sum(v for k, v in get_beetles(brightness, stamps).items())
        for brightness in get_brightnesses(input)
    ))


def part2(input: str) -> str:
    stamps = [1, 3, 5, 10, 15, 16, 20, 24, 25, 30]
    brightnesses = get_brightnesses(input)

    beetle_counts = {}
    for brightness in range(1, max(brightnesses) + 1):
        best_beetle_count = sys.maxsize

        for stamp in stamps:
            if brightness - stamp in beetle_counts:
                best_beetle_count = min(beetle_counts[brightness - stamp] + 1, best_beetle_count)
            elif stamp == brightness:
                best_beetle_count = 1

        if best_beetle_count < sys.maxsize:
            beetle_counts[brightness] = best_beetle_count

    return str(sum(beetle_counts[brightness] for brightness in brightnesses))


def part3(input: str) -> str:
    stamps = [1, 3, 5, 10, 15, 16, 20, 24, 25, 30, 37, 38, 49, 50, 74, 75, 100, 101]
    brightnesses = get_brightnesses(input)

    beetle_counts = {}
    for brightness in range(1, max(brightnesses) + 1):
        best_beetle_count = sys.maxsize

        for stamp in stamps:
            if brightness - stamp in beetle_counts:
                best_beetle_count = min(beetle_counts[brightness - stamp] + 1, best_beetle_count)
            elif stamp == brightness:
                best_beetle_count = 1

        if best_beetle_count < sys.maxsize:
            beetle_counts[brightness] = best_beetle_count

    total_beetle_count = 0
    for brightness in brightnesses:
        left = brightness // 2
        right = brightness - left

        best_beetle_count = sys.maxsize
        while right - left < 100:
            if left in beetle_counts and right in beetle_counts:
                best_beetle_count = min(beetle_counts[left] + beetle_counts[right], best_beetle_count)

            left -= 1
            right += 1

        total_beetle_count += best_beetle_count

    return str(total_beetle_count)


input_1, input_2, input_3 = Api().get_inputs()
if input_1:
    print(f"Part 1: {part1(input_1)}")
if input_2:
    print(f"Part 2: {part2(input_2)}")
if input_3:
    print(f"Part 3: {part3(input_3)}")
