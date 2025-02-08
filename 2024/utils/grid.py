from dataclasses import dataclass
from enum import Enum


class Direction(Enum):
    UP = (-1, 0)
    RIGHT = (0, 1)
    DOWN = (1, 0)
    LEFT = (0, -1)

    def __getitem__(self, item) -> int:
        return self.value[item]

    def rotate_left(self):
        if self == Direction.UP:
            return Direction.LEFT
        elif self == Direction.RIGHT:
            return Direction.UP
        elif self == Direction.DOWN:
            return Direction.RIGHT
        elif self == Direction.LEFT:
            return Direction.DOWN

    def rotate_right(self):
        if self == Direction.UP:
            return Direction.RIGHT
        elif self == Direction.RIGHT:
            return Direction.DOWN
        elif self == Direction.DOWN:
            return Direction.LEFT
        elif self == Direction.LEFT:
            return Direction.UP

    def opposite(self):
        if self == Direction.UP:
            return Direction.DOWN
        elif self == Direction.RIGHT:
            return Direction.LEFT
        elif self == Direction.DOWN:
            return Direction.UP
        elif self == Direction.LEFT:
            return Direction.RIGHT


@dataclass(frozen=True)
class Position:
    row: int = 0
    column: int = 0

    def __add__(self, other):
        if isinstance(other, Position):
            return Position(self.row + other.row, self.column + other.column)
        elif isinstance(other, Direction):
            return Position(self.row + other[0], self.column + other[1])
        else:
            raise NotImplementedError

    def __str__(self):
        return f"({self.row}, {self.column})"