class Decimal {
  constructor(val, numSize=15) {
    this.initVal = val;
    this.numSize = numSize;
    const splitArr = this.splitNum(val)
    this.intArr = splitArr.intArr;
    this.floatArr = splitArr.floatArr;
    this.intVal = splitArr.intVal;
    this.floatVal = splitArr.floatVal;
  }


  addInt(val){
    const {intArr, floatArr, intVal, floatVal} = this.splitNum;
    const maxIntLength = Math.max(intArr.length, this.intArr.length);

    //int
    for(let i = maxIntLength-1; i >= 0; i--){
      this.intArr[i] = this.addStrNum(this.intArr[i], intArr[i]*1);

      const length = (this.intArr[i]).length;
      if(length > this.numSize){
        const num = this.intArr[i][0];
        this.intArr[i] = this.intArr[i].slice(1);
        if(i == 0){
          this.intArr = [num, ...this.intArr];
        }
        else{
          this.intArr[i-1] = this.addStrNum(this.intArr[i-1], num);
        }
      }
    }

    //float
    const maxFloatLength = Math.max(floatArr.length, this.floatArr.length);
    if(floatArr.length)

    for(let i = maxFloatLength - 1; i >= 0; i--){

    }
  }


  addStrNum(strVal1, strVal2){
    strVal1 = !strVal1 ? 0 : strVal1*1;
    strVal2 = !strVal2 ? 0 : strVal2*1;
    return (strVal1 + strVal2) + '';
  }

  splitNum(val) {
    val = val + "";
    const spVal = val.split(".");
    const intVal = spVal[0];
    const floatVal = spVal[1] || '';
    const intArr = this.splitIntToArr(intVal);
    const floatArr = this.splitFloatToArr(floatVal);
    return {intArr, floatArr, intVal, floatVal}
  }

  //整数部分右对齐
  splitIntToArr(strVal) {
    const arr = [];
    let count = parseInt(strVal.length / this.numSize);
    const remVal = strVal.length % this.numSize;
    remVal > 0 && count++;

    for(let i = 0; i < count; i++){
      let startIndex = strVal.length - this.numSize*(i+1);
      // 会被减到负数所以防止一下
      startIndex = startIndex < 0 ? 0 : startIndex;
      let length = i == count - 1 && remVal > 0 ? remVal : this.numSize;
      arr[count - 1 - i] = strVal.slice(startIndex, startIndex + length);
    }

    return arr;
  }

  //小数部分左对齐
  splitFloatToArr(strVal){
    const arr = [];
    let count = parseInt(strVal.length / this.numSize);
    const remVal = strVal.length % this.numSize;
    remVal > 0 && count++;

    for(let i = 0; i < count; i++){
      const startIndex = i*this.numSize;
      const endIndex = startIndex + (i == count - 1 && remVal > 0 ? remVal : this.numSize);
      arr.push(strVal.slice(startIndex, endIndex));
    }

    return arr;
  }
}


module.exports =  Decimal;