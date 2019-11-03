#pragma once
#include <iostream>

// Constructors declared explicit are usually preferable to non-explicit ones
// because they prevent compilers from performing unexpected type conversions

// The copy constructor is used to initialize an object with a different object
// of the same type. The copy assignment operator is used to copy the value from
// one object to another of the same type.

// The copy constructor defines how an object is passed by value. Thus,
// funcUsingWidget will invoke the copy constructor.

// (Note: it's generally a bad idea to pass user-defined types by value;
// pass-by-reference-to-const is typically better; see item 20)

// C++ does not offer interfaces, but Item 31 discusses how to approximate them.

namespace ECPP_Intro {

class Widget {
    private:
        int x;
    public:
        Widget(int _x=0);                       // default constructor
        Widget(const Widget & rhs);             // copy constructor
        Widget & operator=(const Widget & rhs); // copy assignment operator
        int getX() { return x; }
};

void funcUsingWidget(Widget w);

}
