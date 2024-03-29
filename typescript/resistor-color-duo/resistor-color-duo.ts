export function decodedValue(colors: string[]): number {
  let code = `${getColorCode(colors[0])}` + `${getColorCode(colors[1])}`
  return Number(code)
}

function getColorCode(color: string): number {
  switch (color) {
    case "black":   return 0;
    case "brown":   return 1;
    case "red":     return 2;
    case "orange":  return 3;
    case "yellow":  return 4;
    case "green":   return 5;
    case "blue":    return 6;
    case "violet":  return 7;
    case "grey":    return 8;
    case "white":   return 9;
    default:  throw new TypeError("Unknown color");
  }
}
