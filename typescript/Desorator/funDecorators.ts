function get(purl: any, ptype: any) {

  // 1、对于静态成员来说是类的构造函数，对于实例成员是类的原型对象。
  // 2、成员的名字。
  // 3、成员的属性描述符。

  return function (target: any, methodName: any, desc: any) {
    // console.log(target)
    // console.log(methodName)
    // console.log(desc)

    //注意：不要重写HttpClient里面的方法
    // 作用：扩展功能，不是用来覆盖原来的功能的
    target.url = '3yteam';

    //这样可以覆盖原功能
    // target[methodName] = function(){
    //     console.log('方法被改写了');
    // }

    //desc.value 是原方法（数据）
    //这里this变了导致后面拿不到url
    let fun = desc.value; //把原来的方法存起来
    desc.value = function () {
      try {
        //这样就可以拿到url了
        // fun.call(target, purl, ptype)
        //拿不到url
        fun(purl, ptype)
      } catch {
        console.log('msg');
      }
    }

  }
}

class HttpClient {
  public url: any | undefined;
  constructor() {
  }
  @get('http://www.baidu.com', 'post')

  //方式一：
  getData() {
      console.log(arguments);
      // this 拿不到url
      console.log('我的链接是：' + this.url);
  }
  //方式二：
  // getData(...arg:any[]){
  //     console.log('arg----------');
  //     console.log(arg);
  //     console.log('我的链接是2：'+this.url);
  // }
}
var http: any = new HttpClient();
http.getData();
console.log(http.url)