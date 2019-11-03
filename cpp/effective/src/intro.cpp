#include "intro.h"

namespace ECPP_Intro {

Widget::Widget(int _x) : x{_x} {
    std::cout << "default constructor" << std::endl;
}

Widget::Widget(const Widget & rhs) {
    std::cout << "copy constructor" << std::endl;
    this->x = rhs.x;
}

Widget & Widget::operator=(const Widget & rhs) {
    this->x = rhs.x;
    std::cout << "copy assignment operator" << std::endl;
    return *this;
}

void funcUsingWidget(Widget w) {
    (void)w;
}

}
