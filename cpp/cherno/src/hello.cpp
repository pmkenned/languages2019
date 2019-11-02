#include <iostream>
#include <ios>

int main()
{
    uint32_t y = 5;
    y = ~0;

    std::ios_base::fmtflags f(std::cout.flags());
    std::cout << std::hex;
    std::cout << y << std::endl;
    std::cout << sizeof(y) << std::endl;
    std::cout.flags(f);
}
