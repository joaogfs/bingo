NOTE: None of this is implemented yet

Every slice member should have one of the following tags:
- `delim`: the argument is a delimiter with hex notation `0x`
- `size`: the argument should be a previously declared member of type `int` or an action the returns `int`

Every member of a poiter type it treated as optional and should have the tag:
- `optional`: the argument is a previously declared member of type `bool` or an action the returns `int`

Member of `func` types are called "actions". They should be defined before calling the `parse` function. They will be executed after matching the member that comes before them in the struct
