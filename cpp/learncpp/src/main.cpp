#include "triangle.h"
#include "register.h"

#include <iostream>
#include <string>
#include <memory>
#include <map>

void use_triangle_class()
{
    Point f{5, 3};
    Point g{2, 2};
    Point h{0, 0};
    Triangle t0{};
    Triangle t1{f,g,h};

    std::unique_ptr<const Point> t0_centroid_uptr = t0.centroid();
    std::cout << *t0_centroid_uptr << std::endl;

    t0_centroid_uptr.reset(nullptr);

    double d = f.distance(g);

    std::cout << f << std::endl;
    std::cout << g << std::endl;
    std::cout << h << std::endl;
    std::cout << d << std::endl;
    std::cout << t0 << std::endl;
    std::cout << t1 << std::endl;
}

void use_register_class()
{
    Register r{};
}

int main(int argc, char * argv[])
{
    (void)argc;
    (void)argv;

    use_register_class();
}
