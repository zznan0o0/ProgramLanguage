#include <iostream>
using namespace std;

template <typename T1, typename T2, typename T3>

T1 sum(T2 a, T3 b)
{
    return a + b;
}



int main()
{
    cout << sum<float, float, int>(1.1, 2) << endl;
    return 0;
}