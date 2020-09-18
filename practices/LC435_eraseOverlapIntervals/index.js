/**
 * @param {number} n
 * @return {string[]}
 */
var fizzBuzz = function(n) {
    var result = [];
    for(count = 1; count <= n; count++){
      if( count % 15 === 0){
        result.push("FizzBuzz")
      }
      else if( count % 5 === 0){
        result.push("Buzz")
      }
      else if( count % 3 === 0){
        result.push("Fizz")
      }
      else {
        result.push(count.toString())
      }
    }
    return result;
};