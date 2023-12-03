import { data } from "./fileReader.ts";

interface Part {
  id: string;
  x: number;
  y: number;
}

const processData = (input: string): [number, number] => {
  const parts: Part[] = [];
  const symbols: Part[] = [];

  // Extract parts and symbols (including asterisks) from the input
  input.split("\n").forEach((line, y) => {
    let idBuffer = "";
    for (let x = 0; x < line.length; x++) {
      const char = line[x];
      if (/\d/.test(char)) {
        idBuffer += char;
      } else {
        if (idBuffer) {
          parts.push({ id: idBuffer, x: x - idBuffer.length, y });
          idBuffer = "";
        }
        if (char !== ".") {
          symbols.push({ id: char, x, y });
        }
      }
    }
    if (idBuffer) {
      parts.push({ id: idBuffer, x: line.length - idBuffer.length, y });
    }
  });

  const partOne = parts.reduce((total, part) => {
    let isPartAdjacentToSymbol = false;
    for (let i = 0; i < part.id.length; i++) {
      const digitX = part.x + i;
      for (let yOffset = -1; yOffset <= 1; yOffset++) {
        for (let xOffset = -1; xOffset <= 1; xOffset++) {
          if (yOffset === 0 && xOffset === 0) continue;
          const adjacentSymbolExists = symbols.some(
            (symbol) =>
              symbol.x === digitX + xOffset && symbol.y === part.y + yOffset
          );
          if (adjacentSymbolExists) {
            isPartAdjacentToSymbol = true;
            break;
          }
        }
        if (isPartAdjacentToSymbol) break;
      }
      if (isPartAdjacentToSymbol) {
        console.log(
          `Adding part ${part.id} at (${digitX}, ${part.y}) to total.`
        );
        total += parseInt(part.id, 10);
        break;
      }
    }

    if (!isPartAdjacentToSymbol) {
      console.log(
        `Part ${part.id} at (${part.x}, ${part.y}) has no adjacent symbol.`
      );
    }
    return total;
  }, 0);

  const gearRatios = symbols
    .filter((symbol) => symbol.id === "*")
    .map((asterisk) => {
      const adjacentParts = parts.filter((part) => {
        for (let xOffset = -1; xOffset <= part.id.length; xOffset++) {
          for (let yOffset = -1; yOffset <= 1; yOffset++) {
            if (
              part.x + xOffset === asterisk.x &&
              part.y + yOffset === asterisk.y
            ) {
              return true;
            }
          }
        }
        return false;
      });

      if (adjacentParts.length === 2) {
        return parseInt(adjacentParts[0].id) * parseInt(adjacentParts[1].id);
      }
      return 0;
    });

  const totalGearRatio = gearRatios.reduce((acc, ratio) => acc + ratio, 0);
  return [partOne, totalGearRatio];
};

// Example usage
const input = data;
const [partOneResult, partTwoResult] = processData(input);
console.log("Part One Total:", partOneResult);
console.log("Part Two Total Gear Ratio:", partTwoResult);
