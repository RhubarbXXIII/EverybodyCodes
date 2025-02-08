from utils.api import Api
from utils.grid import Direction, Position
from utils.numbers import permute_string


RACETRACK = """S-=++=-==++=++=-=+=-=+=+=--=-=++=-==++=-+=-=+=-=+=+=++=-+==++=++=-=-=--
-                                                                     -
=                                                                     =
+                                                                     +
=                                                                     +
+                                                                     =
=                                                                     =
-                                                                     -
--==++++==+=+++-=+=-=+=-+-=+-=+-=+=-=+=--=+++=++=+++==++==--=+=++==+++-"""
RACETRACK_DUEL = """S+= +=-== +=++=     =+=+=--=    =-= ++=     +=-  =+=++=-+==+ =++=-=-=--
- + +   + =   =     =      =   == = - -     - =  =         =-=        -
= + + +-- =-= ==-==-= --++ +  == == = +     - =  =    ==++=    =++=-=++
+ + + =     +         =  + + == == ++ =     = =  ==   =   = =++=
= = + + +== +==     =++ == =+=  =  +  +==-=++ =   =++ --= + =
+ ==- = + =   = =+= =   =       ++--          +     =   = = =--= ==++==
=     ==- ==+-- = = = ++= +=--      ==+ ==--= +--+=-= ==- ==   =+=    =
-               = = = =   +  +  ==+ = = +   =        ++    =          -
-               = + + =   +  -  = + = = +   =        +     =          -
--==++++==+=+++-= =-= =-+-=  =+-= =-= =--   +=++=+++==     -=+=++==+++-"""


def get_paths(input: str) -> dict[str, list[str]]:
    return {
        line.strip().split(':')[0]: list(line.strip().split(':')[1].split(','))
        for line in input.strip().split('\n')
    }


def get_track(track: str) -> list[str]:
    track_line_length = len(track.split('\n')[0])
    track_lines = list(line.ljust(track_line_length, ' ') for line in track.split('\n'))

    position = Position(0, 1)
    direction = Direction.RIGHT

    track_path = []
    while not track_path or track_path[-1] != 'S':
        track_path.append(track_lines[position.row][position.column])

        for next_direction in (direction, direction.rotate_right(), direction.rotate_left()):
            next_position = position + next_direction

            if (
                0 <= next_position.row < len(track_lines)
                and 0 <= next_position.column < len(track_lines[0])
                and track_lines[next_position.row][next_position.column] != ' '
            ):
                direction = next_direction
                break

        position = next_position

    return track_path


def score_path(path: list[str]) -> int:
    score = 0
    power = 10

    for segment in path:
        if segment == '+':
            power += 1
        elif segment == '-':
            power -= 1

        score += power

    return score


def score_path_on_track(path: list[str], track: list[str], loops: int = 1) -> int:
    score = 0
    power = 10

    for i in range(loops * len(track)):
        if track[i % len(track)] == '+':
            power += 1
        elif track[i % len(track)] == '-':
            power -= 1
        elif path[i % len(path)] == '+':
            power += 1
        elif path[i % len(path)] == '-':
            power -= 1

        score += power

    return score


def part1(input: str) -> str:
    paths = get_paths(input)
    path_scores = {key: score_path(path) for key, path in paths.items()}
    return ''.join(k for k in sorted(path_scores.keys(), key=lambda k: -path_scores[k]))


def part2(input: str) -> str:
    paths = get_paths(input)
    track = get_track(RACETRACK)

    path_scores = {key: score_path_on_track(path, track, 10) for key, path in paths.items()}
    return ''.join(k for k in sorted(path_scores.keys(), key=lambda k: -path_scores[k]))


def part3(input: str) -> str:
    paths = permute_string(5 * '+' + 3 * '-' + 3 * '=')
    path_opponent = next(iter(get_paths(input).values()))
    track = get_track(RACETRACK_DUEL)

    loops = len(path_opponent) + 2024 % len(path_opponent)

    path_score_opponent = score_path_on_track(path_opponent, track, loops)
    path_scores = [score_path_on_track(list(path), track, loops) for path in paths]
    return str(len([path_score for path_score in path_scores if path_score > path_score_opponent]))


input_1, input_2, input_3 = Api().get_inputs()
if input_1:
    print(f"Part 1: {part1(input_1)}")
if input_2:
    print(f"Part 2: {part2(input_2)}")
if input_3:
    print(f"Part 3: {part3(input_3)}")
