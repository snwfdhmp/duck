// Project duck [duck managed]
// Class Window (src/classes/Window/Window.test.cpp)
#ifndef WINDOW_TEST_CPP
#define WINDOW_TEST_CPP

//Window class unit test

#include <iostream>
#include "Window.class.hpp"
#include "../../config/UnitTests.hpp"

int main(int argc, char const *argv[])
{
    unsigned int err = 0;
    Window a;

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
