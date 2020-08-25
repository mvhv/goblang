Goblang synax specification in Backus-Naur Form
```
<source-unit> ::= // source code document after processing

// need to define scope structure

<variable-declaration> ::= dec <whitespace> <type-specifier> <whitespace> <identifier> <line>

<function-definition> ::= def <whitespace> <identifier> <opt-whitespace> ( <function-parameter>* ) <opt-whitespace> { <body> }

<function-parameter> ::= <type-specifier> <whitespace> <identifier>

<statement> ::=

<expression> ::=

<type-specifier> ::= int
                   | int32
                   | int64
                   | uint32
                   | uint64
                   | float
                   | byte
                   | string


// whitespace fluff
<whitespace> ::= <whitespace-char>
               | <whitespace> <whitespace-char>

<whitespace-char> ::= " "
                    | "\n"
                    | "\t"

<opt-whitespace> ::= ""
                   | <whitespace>

<identifier> ::= 
```