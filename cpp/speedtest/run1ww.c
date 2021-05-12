#include <stdio.h>
#include <time.h>

int main()
{
    int start, finish;
    //clock_t为CPU时钟计时单元数
    start = clock();
    long num = 0;
    for(long i=0; i<100000000; i++){
        num += i;
    }
    printf("%ld \n", num);
    finish = clock();
    printf("time: %f\n", double(finish - start) / CLOCKS_PER_SEC);
    return 0;
}