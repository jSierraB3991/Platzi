import random
import time

def bumblee_method(list_elements):
    n = len(list_elements)

    for i in range(n):
        for j in range(0, n - i-1):
            if  list_elements[j] > list_elements[j+1]:
                list_elements[j], list_elements[j+1] = list_elements[j+1], list_elements[j]
    return list_elements

def merge_method( list_elements):
    if len(list_elements) > 1:
        medium = len(list_elements) //2
        left = list_elements[:medium]
        rigth = list_elements[medium:]
        print(left, '*' * 5, rigth)
        
        merge_method(rigth)
        merge_method(left)

        i=0
        j=0
        k=0

        while i < len(left) and j < len(rigth):
            if left[i] < rigth[j]:
                list_elements[k] = left[i]
                i+=1
            else:
                list_elements[k] = rigth[j]
                j+=1
            k+=1

        while i < len(left):
            list_elements[k] = left[i]
            i+=1
            k+=1

        while j < len(rigth):
            list_elements[k] = rigth[j]
            j+=1
            k+=1
        print(f'left {left}, rigth {rigth}')
        print(list_elements)
        print('-' * 50)

    return list_elements



if __name__ == '__main__':
    lenght_of_list=(int(input("lenght of list: ")))

    list_elements = [random.randint(0, 100) for i in range(lenght_of_list)]
    print(list_elements)
    print(bumblee_method(list_elements))

    print("")
    list_elements = [random.randint(0, 100) for i in range(lenght_of_list)]
    print(list_elements)
    print(merge_method(list_elements))
