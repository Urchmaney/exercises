const numberWord = {
  1: 'One',
  2: 'Two',
  3: 'Three',
  4: 'Four',
  5: 'Five',
  6: 'Six',
  7: 'Seven',
  8: 'Eight',
  9: 'Nine',
  10: 'Ten',
  11: 'Eleven',
  12: 'Twelve',
  13: 'Thirteen',
  14: 'Fourteen',
  15: 'Fifteen',
  16: 'Sixteen',
  17: 'Seventeen',
  18: 'Eighteen',
  19: 'Nineteen',
  20: 'Twenty',
  30: 'Thirty',
  40: 'Forty',
  50: 'Fifty',
  60: 'Sixty',
  70: 'Seventy',
  80: 'Eighty',
  90: 'Ninety',
}

const nths = [
  { value: 100, word: 'Hundred' },
  { value: 1000, word: 'Thousand' },
  { value: 1000000, word: 'Million' },
  { value: 1000000000, word: 'Trillion' }
]

const convertToWord = (number, index = nths.length - 1, appendAnd = false) => {
  if (number <= 0) return '';
  let word = '';
  if (index < 0) {
    if(numberWord[number]) return appendAnd ? ` and ${numberWord[number]}` : numberWord[number]
    word = numberWord[Math.floor(number / 10) * 10] || ''
    word += ' ';
    word += numberWord[number % 10] || ''
    appendAnd ? word = ` and ${word}` : ''
    return word;
  }
  const nth = nths[index];
  word = convertToWord(Math.floor(number / nth.value), index - 1);
  if (Math.floor(number / nth.value) > 0) word += ` ${nth.word} `;
  word += convertToWord(number % nth.value, index - 1, appendAnd || Math.floor(number / nth.value) > 0);
  return word;
}

console.log(convertToWord(6));
console.log(convertToWord(1006));
console.log(convertToWord(10000016));
console.log(convertToWord(16));
console.log(convertToWord(153694663));
console.log(convertToWord(8))

//The limit for the code is from 1 to trillion. Zero exclusive.