import time

start = time.time()



def addnum(long range_num):
    cdef long number = 0
    cdef long i
    for i in range(range_num):
        number += i

addnum(100000000)

print(f"Ellapsed time: {time.time() - start} s")
