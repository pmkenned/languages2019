#pragma once

#include <iostream>
#include <vector>

namespace ECPP_Chap1 {

// Item 2: Prefer consts, enums, and inlines to #defines
//
// Instead of
//      #define ASPECT_RATIO 1.653
// Use:
//      const double AspectRatio = 1.653;
//
// When replacing #defines with constants, two special cases:
// * Conceringin pointers: because constant definitions are typically put in
//   header files, it's important that the pointer be declared const, usually
//   in addition to what the pointer points to:
//      const char * const authorName = "Scott Meyers";
//   Even better:
//      const std::string authorName("Scott Meyers");
// * Concerning class-specific constants: to limit the scope of a constant to a
//   class, you must make it a member, and to ensure there's at most one copy
//   of the constant, you must make it a static member
//      static const int NumTurns = 5;
//   Note: that is a declaration of NumTurns, not a definition. Class-specific
//   constants that are static and of integral type (integers, chars, bools) are
//   an exception. As long as you don't take their address, you can declare them
//   and use them without providing a definition. If you do take the address of
//   a class constant, you must provide a definition (do this in an
//   implementation file). No initial value is permitted at the point of
//   definition. Example:
//      const int GamePlayer::NumTurns;
//   Note: there's no such thing as a "private" #define, so this is another
//   reason why constants are preferable: encapsulation.
//   Note: see also enum-hack for old compilers that don't allow in-class
//   initialization (a few other reasons to know about it).
//
// Instead of
//      #define CALL_WITH_MAX(a,b) f((a) > (b) ? (a) : (b))
// Use:
//      template<typename T>
//      inline void callWithMax(const T& a, const T& b)
//      {
//          f(a > b ? a : b);
//      }

template<typename T>
inline int getMax(const T& a, const T& b)
{
    return (a > b ? a : b);
}

// Item 3: Use const whenever possible
//
// If the word const appears to the left of the asterisk, what's pointed to is
// constant; if the word const appears to the right of the asterisk, the pointer
// itself is constant
// Note: const Widget * pw and Widget const * pw are the same
//
// STL iterators are modeled on pointers, so an iterator acts much like a T*
// pointer. Declaring an iterator const is like declaring a pointer const: the
// iterator isn't allowed to point to something different, but the thing it
// points to may be modified. If you want an iterator that points to something
// that can't be modified, you want a const_iterator.
//
// const Member Functions
// The purpose of const member functions is to identify which member functions
// may be invoked on const objects.
// One of the fundamental ways to improve a C++ program's performance is to pass
// objects by reference-to-const.

void const_it();

}
