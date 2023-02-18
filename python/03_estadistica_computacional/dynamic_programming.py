#import time
import sys

def fibonnaci_recurse(n):
    if n == 0 or n == 1:
        return 1
    return fibonnaci_recurse(n - 1) + fibonnaci_recurse(n -2)

def fibonnaci(n, mem = {}):
    if n == 0 or n == 1:
        return 1
    try:
        return mem[n]
    except KeyError:
        result = fibonnaci(n - 1, mem) + fibonnaci(n -2, mem)
        mem[n] = result
        return result

if __name__ == "__main__":
    n = int(input('number for search finnaci: '))

    sys.setrecursionlimit(10002)
    result = fibonnaci(n)
    #startingTime = time.time()
    #result = fibonnaci_recurse(n)
    print(f'fibonacci of {n} is : {result}')
    #endTime = time.time()
    #print(f"BINARI SEARCH \t{endTime - startingTime}");
