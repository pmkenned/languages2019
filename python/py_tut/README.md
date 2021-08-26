# The Python Tutorial

[Source](https://docs.python.org/3/tutorial/index.html)

## [2. Using the Python Interpreter](https://docs.python.org/3/tutorial/interpreter.html)

- Default encoding: UTF-8
  - Use `# -*- coding: encoding -*-` to specify another encoding 
  - This must go on the first line (or first line after shebang line if present)

## [3. An Informal Introduction to Python](https://docs.python.org/3/tutorial/introduction.html)

- `/` always returns a float
- `//` does floor division
- In interactive mode, the last printed expression is assigned to the variable `_`
  - This variable is read-only for the user
- In addition to `int` and `float`, there is `Decimal` and `Fraction`
- Complex numbers: `3+5j` (or `J`)
- Raw-strings: `r'C:\some\name'`
- """...""" or '''...''' can include newlines
  - To prevent newlines from being included in the string, use `\`
- Adjacent string literals are automatically concatenated
- Out of bounds indexes generate an `IndexError` exception, but slices do not
- Lists can contain items of different types
- Lists are mutable
- Slice operations return a new list containing the requested elements
  - `myList[:]` returns a shallow copy of the list
- Can use `+` operator to concatenate lists
- Assignment to slices is possible and can change the size of the list
  - Can be use to insert, replace, or remove elements
- `end` parameter on `print` can be used to prevent/replace the newline

## [4. More Control Flow Tools](https://docs.python.org/3/tutorial/controlflow.html)

- Examples for how to iterate over a copy of a sequence
- `range` and `enumerate`, see [looping techniques](https://docs.python.org/3/tutorial/datastructures.html#tut-loopidioms)
- `range(10)` returns a range object which is an [iterable](https://docs.python.org/3/glossary.html#term-iterable) object that behaves like a list
- `sum` is a function which takes an iterable
- Loop may have an `else` clause which is executed when the loop terminates other than due to a `break`
- The first statement of a function body can optionally be a [docstring](https://docs.python.org/3/tutorial/controlflow.html#tut-docstrings)
- Variable references first look in the local symbol table, then in the local symbol tables of enclosing functions, then in the global symbol table, and finally in the table of built-in names
- Thus, global variables and variables of enclosing functions cannot be directly assigned a value within a function (unless, for global variables, named in a global statement, or, for variables of enclosing functions, named in a nonlocal statement), although they may be referenced
- Arguments are passed using call by value (where the value is always an object reference, not the value of the object)
- Functions return `None` by default
- Default parameters
  - Default parameter values are evaluated once! See example of using mutable value as default parameter
- Keyword arguments
  - `def parrot(voltage, state='a stiff', action='voom', type='Norwegian Blue'):`
  - `def cheeseshop(kind, *args, **kwargs):`
    - `args` is a tuple containing the positional arguments
    - `kwargs` is a dictionary containing the keyword arguments except for those corresponding to a formal parameter
      - order is preserved somehow
- Special parameters: `/` and `*`
  - If neither are present, arguments are all positional-or-keyword
  - If a `/` is present, any parameters before it are positional-only, parameters after it may be either positional-or-keyword or keyword-only
  - Any parameters after a `*` are keyword-only
  - See example of how this resolves an ambiguous case
- `*`-operator unpacks arguments out of a list or tuple: `range(*args)`
- `**`-operator unpacks arguments out of a dict: `range(**kwargs)`
- Lambda expressions: `lambda x: x + n`
- Docstrings:
  - Line should begin with a capital letter, end with a period
  - If there are more lines in the documentation string, the second line should be blank
- Example of function annotation: `def f(ham: str, eggs: str = 'eggs') -> str:`
- Style guide: [PEP8](https://www.python.org/dev/peps/pep-0008/)
  - `UpperCamelCase` for classes and `lowercase_with_underscores` for functions and methods
  - UTF-8 is Python's default encoding

## [5. Data Structures]()

- `sort()` can only be used on lists containing elements which can be pair-wise compared
  - This excludes `None` (which cannot be compared with values of any other type), `str` with `int`, and complex numbers
- Lists can be used as queues, but `collections.deque` is faster
- List comprehension example: `squares = list(map(lambda x: x**2, range(10)))`
  - Equivalent to: `squares = [x**2 for x in range(10)]`

```python
[(x, y) for x in [1,2,3] for y in [3,1,4] if x != y]
# the above is equivalent to:
for x in [1,2,3]:
    for y in [3,1,4]:
        if x != y:
            combs.append((x, y))
```

- The initial expression in a list comprehension can be any arbitrary expression, including another list comprehension
- Example of nested list comprehensions which transpose a matrix: `[[row[i] for row in matrix] for i in range(4)]`
- `del` keyword: can be used to remove elements or slices from a list, or entire variables
- Tuples: similar to lists, but are immutable, usually contain elements of different types and usually accessed by indexing or unpacking
  - For tuples with zero elements: `()`, for tuples with one element: `(x,)`
  - Unpacking: `x, y, z = t`
- Sets: unordered collection with no duplicates
  - Created using curly braces (though with zero elements, this would create a dictionary) or `set()`
  - `&`: set intersection
  - '|': set union
  - `-`: set difference
  - `^`: elements in either but not both
  - Set comprehension example: `a = {x for x in 'abracadabra' if x not in 'abc'}`
- Dictionaries
  - Keys can be any immutable type, such as `str`, `int`, or tuples containing only immutable types
  - Lists cannot be used as keys since they are mutable
  - `list(d)` returns a list of all the keys in dictionary `d` in insertion order
  - Dictionary comprehension example: `{x: x**2 for x in (2, 4, 6)}`
- Looping techniques
  - Iterating a `dict` using `items()`: `for k, v in knights.items():`
  - Example using `enumerate`: `for i, v in enumerate(['tic', 'tac', 'toe']):`
  - Example using `zip`: `for q, a in zip(questions, answers):`
- Conditions
  - `in`, `not in`; `is`, `is not`
  - `a < b == c`
  - Boolean operators short-circuit, and when used as a value, equal the last evaluated argument
  - Assignment inside expressions must use the walrus operator `:=`
- Comparing sequences
  - If one sequence is an initial sub-sequence of the other, the shorter sequence is the smaller (lesser) one

## [6. Modules]()

- Within a module, the module’s name (as a string) is available as the value of the global variable `__name__`
- Using `import`
  - `import fibo`
  - `from fibo import fib, fib2`
  - `from fibo import *` imports all symbols except those beginning with `_`
  - `import fibo as fib`
  - `from fibo import fib as fibonacci`
  - To reload a module during an interactive session: `import importlib; importlib.reload(modulename)`
- `if __name__ == "__main__":`

Left off at [6.1.2. The Module Search Path]()

## [7. Input and Output]()
## [8. Errors and Exceptions]()
## [9. Classes]()
## [10. Brief Tour of the Standard Library]()
## [11. Brief Tour of the Standard Library — Part II]()
## [12. Virtual Environments and Packages]()
## [13. What Now?]()
## [14. Interactive Input Editing and History Substitution]()
## [15. Floating Point Arithmetic: Issues and Limitations]()
