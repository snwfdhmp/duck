// Project duck [duck managed]
// Class Ta (src/classes/Ta/Ta.test.cpp)
#ifndef TA_TEST_CPP
#define TA_TEST_CPP

//Ta class unit test

#include <iostream>
#include "Ta.class.hpp"
#include "../../config/UnitTests.hpp"

int main(int argc, char const *argv[])
{
    unsigned int err = 0;
    Ta a;

    /*
        unit tests here
        use macro SHOULD_BE_TRUE(expression) and SHOULD_BE_FALSE(expression)
        to increment err when errors
    */
        
    if(err) {
        cout << 'Test failed with ' << err << ' errors.' << endl;
        return -1;
    }

    cout << 'Test executed successfully' << endl;
    return 0;
}

#endif
