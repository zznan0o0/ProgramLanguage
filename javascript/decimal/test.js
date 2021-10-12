// const Decimal = require('./Decimal');

// const de = new Decimal('123456789.987654321', 4);

// console.log(de)


//
var Decimal = require('decimal');

var d = Decimal('15641614561561561.154154456456456').add('24651651456156564568.25641651456145614561');
// console.log(d.toString())


const BigNumber = require('bignumber.js');

//这个支持大数据
var b = new BigNumber('4524224542452425542542.57')
// console.log(b.multipliedBy("154156561514546.311312312313361561356"))
// console.log(b.multipliedBy("154156561514546.311312312313361561356").toString())



//这个和Decimal一样只支持正常数据
const Decimaljs = require('decimal.js');

let dj = new Decimaljs("541456141546151565614561356.65145614561561356156");
console.log(dj.plus("26215611561516556165156156156.1561561561665156161561561").toFixed())