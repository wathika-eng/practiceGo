import random

"""
define the function (the list and the target value) 
position of the target value in the list, if found, or -1 if not found
It will compare each item in the list, hence a loop is needed
create a range which will start at 0 and end at the length of the list
len(list) will be a linear time complexity
to get the value of the list, we will use list[i]
then now compare the value of the list to the target value
and terminate the loop if the target value is found else continue the loop and return -1 if the target value is not found
"""


def linear_search(list: list, target: int):
    for i in range(0, len(list)):
        """
        the for loop goes one by one through the list, hence a linear time complexity
        """
        if list[i] == target:
            return i
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
list: int = [x for x in range(1, 21)]
target = random.randint(1, 25)
# print(list)
# print(target)
verify(linear_search(list, target))
