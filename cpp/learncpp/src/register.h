#pragma once

#include <map>
#include <cstdint>

// Sequence containers:
// * array
// * vector
// * deque
// * forward_list
// * list

// Container adaptors:
// * stack
// * queue
// * priority_queue

// Associative containers:
// * set
// * multiset
// * map
// * multimap

// Unordered associative containers:
// * unordered_set
// * unordered_multiset
// * unordered_map
// * unordered_multimap

class RegisterDefinition
{

    // fields
    std::map<std::pair<int, int>, std::string> fields;

    public:
        RegisterDefinition()
        {
        }
};

class MultiBitValue
{
    uint_fast8_t width;

    public:
        MultiBitValue()
        {
        }
};

class Register
{
    uint32_t width;

    RegisterDefinition regDef;
    MultiBitValue value;

    //std::map<std::string, int> myMap;
    //myMap["a"] = 5;
    //myMap["b"] = 6;

    //for(auto const & it : myMap) {
    //    std::cout << it.first << ": " << it.second << std::endl;
    //}

    // std::vector<uint_fast8_t> bits;

    public:

        Register() : width{1} {
        }

        uint32_t getWidth() { return width; }
};
