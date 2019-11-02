#include "triangle.h"

std::ostream &operator<<(std::ostream &os, Point const &p) {
    return os << "(" << p.x << ", " << p.y << ")";
}

std::ostream& operator<<(std::ostream &os, Triangle const &t) {
    return os << t.points[0] << "\t" << t.points[1] << "\t" << t.points[2];
}

//std::ostream& operator<<(std::ostream &os, std::unique_ptr<const Point> & p) {
//    return os << *(p.get());
//}
