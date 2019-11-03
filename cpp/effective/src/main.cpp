#include "intro.h"
#include "c1.h"
#include <iostream>

int main() {

    // introduction
    ECPP_Intro::Widget w1(5);   // invoke default constructor
    ECPP_Intro::Widget w2(w1);  // invoke copy constructor
    w1 = w2;                    // invoke copy assignment operator
    ECPP_Intro::Widget w3 = w2; // invoke copy constructor
    std::cout << w1.getX() << std::endl;
    std::cout << w2.getX() << std::endl;
    std::cout << w3.getX() << std::endl;
    ECPP_Intro::funcUsingWidget(w3); // invoke copy constructor

    // Chapter 1

    // item 2
    int x = 5, y = 7;
    int a = ECPP_Chap1::getMax(x++,y++);
    std::cout << a << std::endl;

    // item 3
    ECPP_Chap1::const_it();
}
