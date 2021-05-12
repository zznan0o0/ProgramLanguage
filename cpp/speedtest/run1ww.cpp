#include <iostream>
#include <time.h>
using namespace std;


void func(long range_num){
    clock_t start, finish;
    //clock_t为CPU时钟计时单元数
    start = clock();

    long num = 0;

    for(long i=0; i<range_num; i++){
        num += i;
    }

    finish = clock();
    cout << num << endl;
	cout <<endl<<"the time cost is:" << double(finish - start) / CLOCKS_PER_SEC<<endl;
}

int main(){
	func(100000000);

	return 0;
}