package littlelisp

import (
  "unicode"
  "unicode/utf8"
  "fmt"
)

type tokenType int

const(
  tokenLeft tokenType = iota
  tokenRight
  tokenQuote
  tokenString
  tokenError
  tokenNumber
  tokenEOF
)

const eof = -1

type stateFunc func(*Lexer) stateFunc

type token struct {
  typ tokenType
  val string
}

func (t token) String() string {
    switch t.typ {
    case tokenEOF:
        return "EOF"
    }
    if len(t.val) > 10 {
        return fmt.Sprintf("%.10q...", t.val)
    }
    return fmt.Sprintf("%q", t.val)
}


type Lexer struct {
  start int
  pos int
  input string
  tokens chan token
}

func lex(input string) *Lexer {
  lexer := &Lexer{
    input: input,
    tokens: make(chan token),
  }

  go lexer.run()

  return lexer
}

func (l *Lexer) run() {
  for state := lexText; state != nil; {
    state = state(l)
  }
  close(l.tokens)
}

func (l *Lexer) nextToken() token {
  token := <-l.tokens
  return token
}

func (l *Lexer) emit(typ tokenType) {
  l.tokens <- token{typ, l.input[l.start:l.pos]}
  l.start = l.pos
}

func (l *Lexer) next() rune {
  if l.pos >= len(l.input) {
    return eof
  }

  r, w := utf8.DecodeRuneInString(l.input[l.pos:])
  l.pos += w

  return r
}

func (l *Lexer) backup() {
  l.pos -= 1
}

func (l *Lexer) ignore() {
  l.start += 1
}

//state functions
func lexText(l *Lexer) stateFunc {
  for {
    r := l.next()
    if r == eof { break }
    switch {
    case r == '\'':
      l.emit(tokenQuote)
      return lexLeft(l)
    case isSpace(r):
      l.ignore()
    case isAlphaNumeric(r):
      l.backup()
      return lexIdentifier
    }
  }

  l.emit(tokenEOF)
  return nil
}


func lexLeft(l *Lexer) stateFunc {
  r := l.next()

  if r ==  '(' {
    l.emit(tokenLeft)
    return lexInsideList
  } else {
    l.emit(tokenError)
    return nil
  }
}

func lexInsideList(l *Lexer) stateFunc {
  r := l.next()

  switch  {
  case r == ')':
    l.emit(tokenRight)
    return lexText
  case isAlphaNumeric(r):
    l.backup()
    return lexIdentifier
  case r == '(':
    l.backup()
    return lexLeft
  }

  return lexInsideList
}

func lexIdentifier(l *Lexer) stateFunc {
  r := l.next()

  switch {
  case isSpace(r):
    l.backup()
    l.emit(tokenNumber)
    l.ignore()
    return lexInsideList
  case r == ')':
    l.backup()
    l.emit(tokenNumber)
    return lexInsideList
  }

  return lexIdentifier
}

func isSpace(r rune) bool {
  return r == ' '
}

func isAlphaNumeric(r rune) bool {
  return unicode.IsLetter(r) || unicode.IsDigit(r)
}
