// Project duck [duck managed]
// Class test (src/classes/test/test.test.cpp)
#ifndef TEST_TEST_CPP
#define TEST_TEST_CPP

//test class unit test

#include <iostream>
#include "test.class.hpp"
#include "../../config/UnitTests.hpp"

int main(int argc, char const *argv[])
{
    unsigned int err = 0;
    test a;

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
