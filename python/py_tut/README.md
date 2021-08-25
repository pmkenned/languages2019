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

Left off at [4.7.7. Documentation Strings]()

## [5. Data Structures]()
## [6. Modules]()
## [7. Input and Output]()
## [8. Errors and Exceptions]()
## [9. Classes]()
## [10. Brief Tour of the Standard Library]()
## [11. Brief Tour of the Standard Library — Part II]()
## [12. Virtual Environments and Packages]()
## [13. What Now?]()
## [14. Interactive Input Editing and History Substitution]()
## [15. Floating Point Arithmetic: Issues and Limitations]()

## References

- [The Python Language Reference](https://docs.python.org/3/reference/index.html#reference-index)
- [The Python Standard Library](https://docs.python.org/3/library/index.html#library-index)
- [Extending and Embedding the Python Interpreter](https://docs.python.org/3/extending/index.html)
- [Python/C API Reference Manual](https://docs.python.org/3/c-api/index.html)
- [Python HOWTOs](https://docs.python.org/3/howto/index.html)
- [Python Frequently Asked Questions](https://docs.python.org/3/faq/index.html)
