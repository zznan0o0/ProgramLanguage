// 注意  TypeScript不允许同时装饰一个成员的get和set访问器。取而代之的是，一个成员的所有装饰的必须应用在文档顺序的第一个访问器上。这是因为，在装饰器应用于一个属性描述符时，它联合了get和set访问器，而不是分开声明的。

function configable(value: boolean) {
  // 1. 对于静态成员来说是类的构造函数，对于实例成员是类的原型对象。
  // 2. 成员的名字。
  // 3. 成员的属性描述符。

  return function (target: any, propertyKey: string, desc: PropertyDescriptor) {
    console.log(desc)
    desc.configurable = value;
    desc.enumerable = value;
    // desc.writable = value;
    //访问器属性：可以起到保护的作用
  }
}
class Point {
  private _x: number;
  private _y: number;
  constructor(x: number, y: number) {
    this._x = x;
    this._y = y;
  }

    //TypeScript支持通过getters/setters来截取对对象成员的访问。 它能帮助你有效的控制对对象成员的访问。
  //实例化后用XX访问_x
  @configable(false)
  get XX() {
    return this._x;
  }
  @configable(false)
  get YY() {
    return this._y;
  }

  set XX(_x: number){
    this._x = _x;
  }
}


var p: Point = new Point(1, 2);

console.log(p.XX)
// p.XX = 0;