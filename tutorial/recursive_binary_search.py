import random, time

"""
This will return True if value exists else false, unlike other binary search
we will start by checking if list is empty or not, if not then we will check
unlike previous search, we have a loop which ensured code always executed
"""


def recursive_binary_search(list, target: int):
    if len(list) == 0:
        return False
    else:
        # length of list is calculated by len(list) // 2
        midpoint = len(list) // 2
        if list[midpoint] == target:
            return True
        else:
            """
            if target is less than midpoint then we shall create a sublist which starts at midpoint + 1
            and ends at end of list
            call recursive_binary_search with new list, but same target
            """
            if list[midpoint] < target:
                # midpoint + 1 then end of list
                return recursive_binary_search(list[midpoint + 1 :], target)
            else:
                # if value is greater than midpoint then shall start at beginning of list and end at midpoint
                return recursive_binary_search(list[:midpoint], target)


def verify(result: bool):
    # will return True if value is found else False
    print("Target found: ", result)


list = [random.randint(0, 100) for i in range(10)]
list.sort()
print(list)
result = recursive_binary_search(list, 10)
verify(result)

 