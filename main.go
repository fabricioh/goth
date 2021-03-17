package main

import (
  "fmt"
  // "strings"
  "os"
  "strconv"
  "regexp"
)

func main() {
  if len(os.Args) == 1 || os.Args[1] == "--help" {
    fmt.Println("goth v0.1 - fabricio h")
    fmt.Println("\nPasse como primeiro argumento uma expressão matemática.")
    fmt.Print("\nExemplos:\n\tgoth \"5 * (2 + 3)\"\n\tgoth \"10 / 10\"\n\tgoth 4+2\n\n")
    os.Exit(0)
  }

  // fmt.Printf("result: \t%v", solve(os.Args[1]))
  fmt.Printf("\nexpressão: %s\n", os.Args[1])

  result := solve(os.Args[1])
  fmt.Printf("\nresultado: %v\n\n", result)
}

func solve(exp string) int64 {
  parens := 0
  var numbers []int64
  var operator rune
  var operands []string
  var operand string
  var result int64 = 404

  // Tira os parenteses em volta da expressão
  // (caso strings.Trim fosse usado, apagariam-se
  // mais parenteses que o necessário em algumas situações)
  if ok, _ := regexp.MatchString("^\\(.+\\)$", exp); ok {
    exp = exp[1:][:len(exp)-2]
  }
  
  // Detecta os elementos da expressão
  for _, i := range exp {
    if i == ' ' {continue}

    if (i == '+' || i == '-' || i == '*' || i == '/') && parens == 0 {
      operator = i
      // fmt.Printf("\tfound operator: %c\n", operator)
      operands = append(operands, operand)
      operand = ""
      continue
    }
    
    if i == '(' {
      parens++
      // fmt.Printf("\tparens: %v\n", parens)
    }
    
    if i == ')' {
      parens--
      // fmt.Printf("\tparens: %v\n", parens)
    }

    operand += string(i)
  }

  operands = append(operands, operand)

  // Se a expressão for composta de apenas um número,
  // retorna ele
  if len(operands) == 1 {
    result, err := strconv.ParseInt(operands[0], 10, 64)
    
    if err != nil {
      fmt.Printf("\n-- ERRO: Não foi possível processar o termo '%v' --\n\n", operands[0])
      os.Exit(1)
    }

    // fmt.Printf("\n%v\n", result)
    // fmt.Print("----------------------------\n")

    return result
  }

  // Resolve expressões recursivamente
  for i, op := range operands {
    if ok, _ := regexp.MatchString("\\(.+\\)", op); ok {
      operands[i] = strconv.Itoa(int(solve(op)))
    }
  }

  // LOG
  // fmt.Printf("\nexpression: %s %c %s\n", operands[0], operator, operands[1])
  // fmt.Printf("operand: \t%s\n", operands[0])
  // fmt.Printf("operator: \t%c\n", operator)
  // fmt.Printf("operand: \t%s\n", operands[1])

  // Converte os termos para números
  for _, op := range operands {
    result, err := strconv.ParseInt(op, 10, 64)
    
    if err != nil {
      fmt.Printf("\n-- ERRO: Não foi possível processar o termo '%v' --\n\n", op)
      os.Exit(1)
    }

    numbers = append(numbers, result)
  }

  // Retorna o resultado
  switch operator {
  case '+':
    result = numbers[0] + numbers[1]
  case '-':
    result = numbers[0] - numbers[1]
  case '*':
    result = numbers[0] * numbers[1]
  case '/':
    result = numbers[0] / numbers[1]
  }

  fmt.Printf("\n%s %c %s = %v\n", operands[0], operator, operands[1], result)
  fmt.Print("----------------------------\n")
  // fmt.Printf("result: %v\n", result)

  return result
}
