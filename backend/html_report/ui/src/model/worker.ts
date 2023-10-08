const operations =  ["/", "*", "/", "+", "-"]
let keep = ["(", ")", ...operations]


export const SetEq = (eq: string): string => {
    let operationString = "";
    for (let i = 0; i < eq.length; i++) {
        const c = eq[i]
    
        if (keep.includes(c)) {
            operationString += c
        }
  }

  if (operationString.length === 0) {
    return "a"
  }
  // console.log(operationString)
  // then if there is an operation, we should fill it with the letter a on either side
  let newOperationString = "";
  for (let i = 0; i < operationString.length; i++) {
    const c = operationString[i]
  
    

    if (operations.includes(c)) {
     
      if (operationString[i-1] === ")") {
        newOperationString += c;
        continue
      }
      if (operationString[i+1] === "(") {
        newOperationString += "a" + c;
        continue
      }



      // if another operation exists, dont put a. if it is end or parenthesis, dont put
      if (operations.includes(operationString[i+1])) {
      newOperationString += "a" + c;

      } else {

      newOperationString += "a" + c + "a";
      }

    } else {
      newOperationString += c;
    }
  
  
  }

  return newOperationString
  }