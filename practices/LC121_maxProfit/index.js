const maxDepth = require("../LC104_maxDepth");

var maxProfit = function(prices) {
let maxprofit = 0;
let cheapestPrice = prices[0];

for(let i =0; i <prices.length; i++){
  const price = prices[i];
  if ( price < cheapestPrice) cheapestPrice = price;

  const currentProfit = price - cheapestPrice
  maxprofit = Math.max(currentProfit,maxprofit);
}
return maxprofit
};

module.exports = maxProfit;
