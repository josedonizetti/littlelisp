package littlelisp

import (
  "unicode"
  "unicode/utf8"
  "fmt"
  "strconv"
  "strings"
)

type tokenType int

const(
  tokenLeft tokenType = iota
  tokenRight
  tokenQuote
  tokenString
  tokenError
  tokenNumber
  tokenSymbol
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

func Lex(input string) *Lexer {
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

func (l *Lexer) NextToken() token {
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
      return lexLeft
    case isSpace(r):
      l.ignore()
    case isAlphaNumeric(r):
      l.backup()
      return lexIdentifier
    case r == '(':
      l.backup()
      return lexLeft
    case r == ')':
      l.emit(tokenRight)
      return lexText
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
  if r == eof {
    emitIdentifier(l)
    return lexText
  }

  switch {
  case isSpace(r):
    l.backup()
    emitIdentifier(l)
    l.ignore()
    return lexInsideList
  case r == ')':
    l.backup()
    emitIdentifier(l)
    return lexInsideList
  }

  return lexIdentifier
}

func emitIdentifier(l *Lexer) {
  input := l.input[l.start:l.pos]
  _, err := strconv.Atoi(input)

  if err == nil {
    l.emit(tokenNumber)
  } else if strings.HasPrefix(input,"\"") {
    l.emit(tokenString)
  } else {
    l.emit(tokenSymbol)
  }
}

func isSpace(r rune) bool {
  return r == ' '
}

func isAlphaNumeric(r rune) bool {
  return unicode.IsLetter(r) || unicode.IsDigit(r)
}
