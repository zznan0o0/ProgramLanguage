#include <iostream>

using namespace std;


int main(){
	int a = 1;
	int b = 2;

	int &c = a;
	a = 3;
	cout << c << endl;

	// cout << *a << endl;
	cout << &a << endl;
	cout << *&a << endl;


	return 0;
}
