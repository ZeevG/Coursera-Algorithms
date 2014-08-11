
def merge_sort(unsorted):

    length = len(unsorted)

    # Base case where recursive calls stop
    if length <= 1:
        return unsorted

    left = unsorted[:length/2]
    right = unsorted[length/2:]

    left = merge_sort(left)
    right = merge_sort(right)

    return merge(left, right)


def merge(left, right):

    merged = []
    while len(left) != 0 and len(right) != 0:

        if left[0] < right[0]:
            merged.append(left.pop(0))
        else:
            merged.append(right.pop(0))

    # Empty any remaining elements from left/right
    # onto the sorted array
    while len(left) != 0:
        merged.append(left.pop(0))

    while len(right) != 0:
        merged.append(right.pop(0))

    return merged


if __name__ == "__main__":

    unsorted = [
        34,
        23,
        423,
        356,
        43,
        4,
        2,
        32,
        41,
        2154,
        5,
        63,
        43,
        123,
        42,
        34,
        52,
        35,
    ]

    print unsorted

    array = merge_sort(unsorted)

    print array
