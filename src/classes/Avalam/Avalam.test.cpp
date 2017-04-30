// Project duck [duck managed]
// Class Avalam (src/classes/Avalam/Avalam.test.cpp)
#ifndef AVALAM_TEST_CPP
#define AVALAM_TEST_CPP

//Avalam class unit test

#include <iostream>
#include "Avalam.class.hpp"
#include "../../config/UnitTests.hpp"

int main(int argc, char const *argv[])
{
    unsigned int err = 0;
    Avalam a;

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
