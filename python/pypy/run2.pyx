import time

start = time.time()
cdef long number = 0
cdef long i

for i in range(100000000):
    number += i
    
print(f"Ellapsed time: {time.time() - start} s")
print(number)