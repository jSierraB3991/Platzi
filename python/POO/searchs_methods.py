import random
import time

def binari_search(list_elements, start, finish, objetive):
    if start > finish:
        return False

    medium = (start+finish) // 2
    if list_elements[medium] == objetive:
        return True
    elif list_elements[medium] < objetive:
        return binari_search(list_elements, medium+1, finish, objetive)
    else:
        return binari_search(list_elements, start, medium - 1, objetive)
 

def lineal_search( list_elements, objetive ):
    response=False

    for element in list_elements:
        if element == objetive:
            response=True
            break

    return response



if __name__ == '__main__':
    lenght_of_list=(int(input("lenght of list: ")))
    objetive=(int(input("value to search: ")))

    list_elements = [random.randint(0, 100) for i in range(lenght_of_list)]

    startingTime = time.time()
    encontrado=lineal_search(list_elements, objetive)
    print(f'The value {objetive} {"here" if encontrado else "not here"} in the list')
    endTime = time.time()
    print(f"LINEAL SEARCH \t{endTime - startingTime}");

    startingTime = time.time()
    encontrado=binari_search(sorted(list_elements), 0, len(list_elements), objetive)
    print(f'The value {objetive} {"here" if encontrado else "not here"} in the list')
    endTime = time.time()
    print(f"BINARI SEARCH \t{endTime - startingTime}");
