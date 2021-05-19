function logParams(params: any) {
  // 1、对于静态成员来说是类的构造函数，对于实例成员是类的原型对象。
  // 2、方法的名字。
  // 3、参数在函数参数列表中的索引。
  return function (target: any, methodName: any, paramsIndex: any) {
    console.log(params);
    console.log(target);
    console.log(methodName);
    console.log(paramsIndex);
    debugger
    target.apiUrl = params;
  };
}

class HttpClient {
  public url: any | undefined;
  constructor() {}
  getData(@logParams("3yteam.com") uuid: any) {
    console.log(uuid);
  }
}

var http: any = new HttpClient();
http.getData(123456);
