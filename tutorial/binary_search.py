import random

"""
define a function which will take a list and a target value
binary search splits the list into two halves and compares the target value to the middle value
we start at 0 index
"""


def binary_search(sort_list: list, target: int):
    first_val: int = 0
    last_val: int = len(sort_list) - 1
    """
    use a while loop to iterate through the list
    a mid point is calculated by adding the first and last value and dividing by 2
    mid point is important because it is the value we will compare the target value to
    two // are used to round the value to the nearest whole number
    """
    while first_val <= last_val:
        mid_point = (first_val + last_val) // 2
        print(mid_point)
        # if the mid point is the target value, return the mid point
        if sort_list[mid_point] == target:
            return mid_point
        # else if the mid point is less than the target value, the first value is updated to the mid point
        elif sort_list[mid_point] < target:
            # the first value is updated to the mid point + 1
            # this takes time and space complexity to O(log n)
            first_val = mid_point + 1
        else:
            # the last value is updated to the mid point - 1
            last_val = mid_point - 1
    return None


def verify(index: int):
    """
    verify if an index is returned or not
    """
    if index != None:
        print(f"Number {target} found at index: ", index)
    else:
        print(f"Number {target} not found in the list")


"""
A list of numbers from 1 to 20
A random target value between 1 and 25
"""
int_list = [9, 3, 14, 1]
sort_list = sorted(int_list)
target = 14
# print(sort_list)
# print(target)
verify(binary_search(sort_list, target))
