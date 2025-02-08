import functools


@functools.cache
def permute_string(items: str) -> set[str]:
    if len(items) == 1:
        return {items}

    permutations = set()
    for i, item in enumerate(items):
        for permutation in permute_string(items[:i] + items[i + 1:]):
            permutations.add(item + permutation)

    return permutations