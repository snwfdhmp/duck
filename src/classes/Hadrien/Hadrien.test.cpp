// Project duck [duck managed]
// Class Hadrien (src/classes/Hadrien/Hadrien.test.cpp)
#ifndef HADRIEN_TEST_CPP
#define HADRIEN_TEST_CPP

//Hadrien class unit test

#include <iostream>
#include "Hadrien.class.hpp"
#include "../../config/UnitTests.hpp"

int main(int argc, char const *argv[])
{
    unsigned int err = 0;
    Hadrien a;

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
