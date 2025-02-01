from collections import defaultdict, deque

from utils.api import Api


def get_tree(input: str) -> dict[str, list[str]]:
    return {
        line.strip().split(':')[0]: list(line.strip().split(':')[1].split(','))
        for line in input.strip().split('\n')
    }


def part1(input: str) -> str:
    tree = get_tree(input)

    fruit_paths_by_length = defaultdict(set)

    queue = deque()
    queue.append(('RR', "RR"))

    while queue:
        node = queue.popleft()

        if node[0] not in tree:
            continue

        for child in tree[node[0]]:
            if child == '@':
                fruit_path = node[1] + '@'
                fruit_paths_by_length[len(fruit_path)].add(fruit_path)
                continue

            queue.append((child, node[1] + child))

    for length, fruit_paths in fruit_paths_by_length.items():
        if len(fruit_paths) == 1:
            return next(iter(fruit_paths))


def part2(input: str) -> str:
    tree = get_tree(input)

    fruit_paths_by_length = defaultdict(set)

    queue = deque()
    queue.append(('RR', "RR"))

    while queue:
        node = queue.popleft()

        if node[0] not in tree:
            continue

        for child in tree[node[0]]:
            if child == '@':
                fruit_path = node[1] + ',@'
                fruit_paths_by_length[len(fruit_path)].add(fruit_path)
                continue

            queue.append((child, node[1] + ',' + child))

    for length, fruit_paths in fruit_paths_by_length.items():
        if len(fruit_paths) == 1:
            return ''.join(node[0] for node in next(iter(fruit_paths)).split(','))


def part3(input: str) -> str:
    tree = get_tree(input)

    fruit_paths_by_length = defaultdict(set)

    visited_nodes = set()

    queue = deque()
    queue.append(('RR', "RR"))

    while queue:
        node = queue.popleft()

        if node[0] not in tree or node[0] in visited_nodes:
            continue

        visited_nodes.add(node[0])

        for child in tree[node[0]]:
            if child == '@':
                fruit_path = node[1] + ',@'
                fruit_paths_by_length[len(fruit_path)].add(fruit_path)
                continue

            queue.append((child, node[1] + ',' + child))

    for length, fruit_paths in fruit_paths_by_length.items():
        if len(fruit_paths) == 1:
            return ''.join(node[0] for node in next(iter(fruit_paths)).split(','))


input_1, input_2, input_3 = Api().get_inputs()
if input_1:
    print(f"Part 1: {part1(input_1)}")
if input_2:
    print(f"Part 2: {part2(input_2)}")
if input_3:
    print(f"Part 3: {part3(input_3)}")
