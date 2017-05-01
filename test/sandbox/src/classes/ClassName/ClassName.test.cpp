#ifndef CLASSNAME_CLASS_CPP
    #define CLASSNAME__CPP

    //ClassName class unit test

    #include <iostream>
    #include "ClassName.class.hpp"
    #include "../../config/UnitTests.hpp"

    int main(int argc, char const *argv[])
    {
        unsigned int err = 0;
        ClassName a;

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
