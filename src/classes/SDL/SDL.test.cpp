// Project duck [duck managed]
// Class SDL (src/classes/SDL/SDL.test.cpp)
#ifndef SDL_TEST_CPP
#define SDL_TEST_CPP

//SDL class unit test

#include <iostream>
#include "SDL.class.hpp"
#include "../../config/UnitTests.hpp"

int main(int argc, char const *argv[])
{
    unsigned int err = 0;
    SDL a;

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
