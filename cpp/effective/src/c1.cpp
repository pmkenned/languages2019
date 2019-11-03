#include "c1.h"

namespace ECPP_Chap1 {

void const_it() {
    std::vector<int> vec;
    vec.push_back(5);
    vec.push_back(6);
    const std::vector<int>::iterator iter = vec.begin();
    std::cout << "*iter: " << *iter << std::endl;
    *iter = 10;
    std::vector<int>::const_iterator cIter = vec.begin();
    std::cout << "*cIter: " << *cIter << std::endl;
    ++cIter;
    std::cout << "*cIter: " << *cIter << std::endl;
}

}
