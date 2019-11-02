#pragma once

#include <iostream>
#include <vector>
#include <cmath>
#include <memory>

class Triangle;

class Point {
    double x;
    double y;

    public:
        Point(const Point &other) : x{other.x}, y{other.y} {
            //std::cout << "point copy constructor:" << this << std::endl;
        }

        Point(double _x={}, double _y={}) : x(_x), y(_y) {
            //std::cout << "point constructor:" << this << std::endl;
        }

        ~Point() {
            std::cout << "point destructor:" << this << std::endl;
        }

        double distance(const Point & other) {
            const double dx = this->x - other.x;
            const double dy = this->y - other.y;
            return std::sqrt(dx*dx + dy*dy);
        }

        friend std::ostream &operator<<(std::ostream &os, Point const &p);
};

class Triangle {
    std::vector<Point> points;

    public:
        Triangle(Point const & p0 = Point{0,0}, Point const & p1 = Point{1,0}, Point const & p2 = Point{0,1}): points{p0, p1, p2}
        {
        }

        double perimeter()
        {
            return 0.0;
        }

        double area()
        {
            return 0.0;
        }

        std::unique_ptr<const Point> centroid()
        {
            double x = 1.1;
            double y = 1.2;
            return std::make_unique<Point>(x,y);
        }

        const Point * circumcenter()
        {
            return 0;
        }

        const Point * orthocenter()
        {
            return 0;
        }

        friend std::ostream& operator<<(std::ostream &os, Triangle const &t);
};


