#include <iostream>
// #include <string>
#include <map>

using namespace std;

int main()
{
    std::map<std::string, int> mymap;
    mymap["id"] = 1;
    // mymap["age"] = 11;
    // mymap["no"] = 55;

    std::cout << mymap["id"] << std::endl;
    return 0;
}