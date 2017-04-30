// Project duck [duck managed]
// Class Test (src/classes/Test/Test.test.cpp)
#ifndef TEST_TEST_CPP
#define TEST_TEST_CPP

//Test class unit test

#include <iostream>
#include "Test.class.hpp"
#include "../../config/UnitTests.hpp"

int main(int argc, char const *argv[])
{
    unsigned int err = 0;
    Test a;

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
