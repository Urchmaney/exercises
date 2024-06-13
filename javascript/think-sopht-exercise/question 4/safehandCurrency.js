const usdRates = {
  'USD': 1,
  'GBP': 0.73,
  'CAD': 1.26,
  'CNY': 6.46,
  'GHS': 6.05
};

// currency arguments should be USD, GBP, CAD, CNY and GHS
const ghsCurrencyConverter = (amount, currency) => {
  if (!usdRates[currency]) return 0.00

  const currencyUSDRate = usdRates[currency];
  const ghsUSDRate = usdRates['GHS'];
  return ((amount * ghsUSDRate) / currencyUSDRate).toFixed(2);
}

console.log(ghsCurrencyConverter(40, 'CAD'));
console.log(ghsCurrencyConverter(40, 'USD'));
console.log(ghsCurrencyConverter(40, 'GHS'));
console.log(ghsCurrencyConverter(39, 'CNY'));